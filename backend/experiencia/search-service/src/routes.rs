use actix_web::web;
use crate::handlers::{search::search_handler, suggestions::suggestions_handler};

pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(
        web::scope("/search")
            .route("", web::get().to(search_handler))
            .route("/suggestions", web::get().to(suggestions_handler))
    );
}
