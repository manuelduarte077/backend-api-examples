use actix_web::{web, App, HttpResponse, HttpServer, Responder};
use serde::{Deserialize, Serialize};
use sqlx::{sqlite::SqlitePool, Executor, FromRow};
use std::fs;
use std::path::Path;

// Estructura de usuario
#[derive(Serialize, Deserialize, FromRow)]
struct User {
    id: Option<i64>,
    name: String,
    age: i32,
}

// Obtener todos los usuarios
async fn get_users(pool: web::Data<SqlitePool>) -> impl Responder {
    let result = sqlx::query_as::<_, User>("SELECT * FROM users")
        .fetch_all(pool.get_ref())
        .await;

    match result {
        Ok(users) => HttpResponse::Ok().json(users),
        Err(e) => {
            eprintln!("Error al obtener usuarios: {:?}", e);
            HttpResponse::InternalServerError().body(format!("Error al obtener usuarios: {:?}", e))
        }
    }
}

// Crear un nuevo usuario
async fn create_user(pool: web::Data<SqlitePool>, new_user: web::Json<User>) -> impl Responder {
    let result = sqlx::query("INSERT INTO users (name, age) VALUES (?, ?)")
        .bind(&new_user.name)
        .bind(new_user.age)
        .execute(pool.get_ref())
        .await;

    match result {
        Ok(_) => HttpResponse::Created().body("Usuario creado exitosamente"),
        Err(e) => {
            eprintln!("Error al crear usuario: {:?}", e);
            HttpResponse::InternalServerError().body(format!("Error al crear usuario: {:?}", e))
        }
    }
}

// Actualizar un usuario existente
async fn update_user(
    pool: web::Data<SqlitePool>,
    path: web::Path<i64>, // ID del usuario en la ruta
    updated_user: web::Json<User>,
) -> impl Responder {
    let id = path.into_inner();
    let result = sqlx::query("UPDATE users SET name = ?, age = ? WHERE id = ?")
        .bind(&updated_user.name)
        .bind(updated_user.age)
        .bind(id)
        .execute(pool.get_ref())
        .await;

    match result {
        Ok(res) if res.rows_affected() > 0 => {
            HttpResponse::Ok().body("Usuario actualizado exitosamente")
        }
        Ok(_) => HttpResponse::NotFound().body("Usuario no encontrado"),
        Err(e) => {
            eprintln!("Error al actualizar usuario: {:?}", e);
            HttpResponse::InternalServerError()
                .body(format!("Error al actualizar usuario: {:?}", e))
        }
    }
}

// Eliminar un usuario
async fn delete_user(pool: web::Data<SqlitePool>, path: web::Path<i64>) -> impl Responder {
    let id = path.into_inner();
    let result = sqlx::query("DELETE FROM users WHERE id = ?")
        .bind(id)
        .execute(pool.get_ref())
        .await;

    match result {
        Ok(res) if res.rows_affected() > 0 => {
            HttpResponse::Ok().body("Usuario eliminado exitosamente")
        }
        Ok(_) => HttpResponse::NotFound().body("Usuario no encontrado"),
        Err(e) => {
            eprintln!("Error al eliminar usuario: {:?}", e);
            HttpResponse::InternalServerError().body(format!("Error al eliminar usuario: {:?}", e))
        }
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    // Crear el directorio de la base de datos si no existe
    let db_path = "data/users.db";
    let db_dir = Path::new("data");
    if !db_dir.exists() {
        fs::create_dir_all(db_dir).expect("No se pudo crear el directorio de la base de datos");
    }

    // Crear pool de conexión a la base de datos
    let pool = SqlitePool::connect(&format!("sqlite:{}", db_path))
        .await
        .expect("No se pudo conectar a la base de datos");

    // Inicializar la base de datos
    pool.execute(
        "CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            age INTEGER NOT NULL
        );",
    )
    .await
    .expect("No se pudo inicializar la base de datos");

    // Configuración del servidor
    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(pool.clone())) // Pool compartido
            .route("/users", web::get().to(get_users)) // Obtener usuarios
            .route("/users", web::post().to(create_user)) // Crear usuario
            .route("/users/{id}", web::put().to(update_user)) // Actualizar usuario
            .route("/users/{id}", web::delete().to(delete_user)) // Eliminar usuario
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}
