use actix_web::web;

pub fn configure(cfg: &mut web::ServiceConfig) {
    cfg.service(
        web::scope("/tasks")
            .route("", web::post().to(crate::handlers::create_task))
            .route("", web::get().to(crate::handlers::get_tasks))
            .route("/{id}", web::put().to(crate::handlers::update_task))
            .route("/{id}", web::delete().to(crate::handlers::delete_task)),
    );
}
