use actix_web::{HttpResponse, web};
use crate::{elastic::search_products, cache::redis_cache::get_or_store};

pub async fn search_handler(query: web::Query<std::collections::HashMap<String, String>>) -> HttpResponse {
    let key = format!("search:{:?}", query);

    match get_or_store(&key, || search_products(query.clone())) .await {
        Ok(data) => HttpResponse::Ok().json(data),
        Err(_) => HttpResponse::InternalServerError().body("Error al buscar productos"),
    }
}
