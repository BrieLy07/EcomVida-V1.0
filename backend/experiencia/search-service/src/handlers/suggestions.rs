use actix_web::{HttpResponse, web};
use crate::{elastic::suggest_terms, cache::redis_cache::get_or_store};

pub async fn suggestions_handler(query: web::Query<std::collections::HashMap<String, String>>) -> HttpResponse {
    let key = format!("suggestions:{:?}", query);

    match get_or_store(&key, || suggest_terms(query.clone())).await {
        Ok(data) => HttpResponse::Ok().json(data),
        Err(_) => HttpResponse::InternalServerError().body("Error al obtener sugerencias"),
    }
}
