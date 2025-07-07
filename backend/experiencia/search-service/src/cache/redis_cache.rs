use redis::AsyncCommands;
use serde_json::Value;
use std::future::Future;

pub async fn get_or_store<F, Fut>(key: &str, fetch_fn: F) -> Result<Value, redis::RedisError>
where
    F: FnOnce() -> Fut,
    Fut: Future<Output = Result<Value, Box<dyn std::error::Error>>>,
{
    let client = redis::Client::open("redis://redis:6379/")?;
    let mut conn = client.get_async_connection().await?;

    // Intentar obtener del caché
    if let Ok(cached) = conn.get::<_, String>(key).await {
        return Ok(serde_json::from_str(&cached).unwrap());
    }

    // Si no hay en caché, buscar y guardar
    let data = fetch_fn().await.unwrap();
    let data_str = serde_json::to_string(&data).unwrap();
    let _: () = conn.set_ex(key, data_str, 60).await?; // TTL: 60 segundos

    Ok(data)
}
