use actix_web::{App, HttpServer};
use dotenv::dotenv;
use std::env;
mod routes;
mod handlers;
mod elastic;
mod cache;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();
    env_logger::init();

    let port = env::var("PORT").unwrap_or_else(|_| "3030".to_string());

    println!("🚀 search-service ejecutándose en http://localhost:{}", port);

    HttpServer::new(|| {
        App::new()
            .configure(routes::config)
    })
    .bind(("0.0.0.0", port.parse::<u16>().unwrap()))?
    .run()
    .await
}
