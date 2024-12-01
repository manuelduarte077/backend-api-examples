
use crate::models::{NewTask, Task};
use crate::schema::tasks::dsl::*;
use actix_web::{web, HttpResponse, Responder};
use diesel::prelude::*;
use diesel::RunQueryDsl;

type DbPool = web::Data<crate::db::DbPool>;

pub async fn create_task(pool: DbPool, new_task: web::Json<NewTask>) -> impl Responder {
    let mut conn = pool.get().expect("Could not get connection from pool");

    let new_task = NewTask {
        title: new_task.title.clone(),
        description: new_task.description.clone(),
        completed: new_task.completed,
    };

    match diesel::insert_into(tasks)
        .values(&new_task)
        .execute(&mut conn) {
        Ok(_) => HttpResponse::Created().finish(),
        Err(err) => {
            eprintln!("Error saving new task: {}", err);
            HttpResponse::InternalServerError().finish()
        }
    }
}

pub async fn get_tasks(pool: DbPool) -> impl Responder {
    let mut conn = pool.get().expect("Could not get connection from pool");

    match tasks.load::<Task>(&mut conn) {
        Ok(results) => HttpResponse::Ok().json(results),
        Err(err) => {
            eprintln!("Error loading tasks: {}", err);
            HttpResponse::InternalServerError().finish()
        }
    }
}

pub async fn update_task(
    pool: DbPool,
    task_id: web::Path<i32>,
    task_update: web::Json<NewTask>,
) -> impl Responder {
    let mut conn = pool.get().expect("Could not get connection from pool");

    match diesel::update(tasks.find(*task_id))
        .set((
            title.eq(&task_update.title),
            description.eq(&task_update.description),
            completed.eq(task_update.completed),
        ))
        .execute(&mut conn) {
        Ok(_) => HttpResponse::Ok().finish(),
        Err(err) => {
            eprintln!("Error updating task: {}", err);
            HttpResponse::InternalServerError().finish()
        }
    }
}

pub async fn delete_task(pool: DbPool, task_id: web::Path<i32>) -> impl Responder {
    let mut conn = pool.get().expect("Could not get connection from pool");

    match diesel::delete(tasks.find(*task_id)).execute(&mut conn) {
        Ok(_) => HttpResponse::Ok().finish(),
        Err(err) => {
            eprintln!("Error deleting task: {}", err);
            HttpResponse::InternalServerError().finish()
        }
    }
}