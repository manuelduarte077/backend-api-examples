use diesel::prelude::*;
use serde::{Deserialize, Serialize};
use chrono::NaiveDateTime;



#[derive(Queryable, Serialize)]
pub struct Task {
    pub id: Option<i32>,               // Nullable(Integer)
    pub title: String,                 // Text
    pub description: Option<String>,   // Nullable(Text)
    pub completed: bool,               // Bool
    pub created_at: Option<NaiveDateTime>, // Nullable(Timestamp)
    pub updated_at: Option<NaiveDateTime>, // Nullable(Timestamp)
}


#[derive(Insertable, Deserialize)]
#[diesel(table_name = crate::schema::tasks)]
pub struct NewTask {
    pub title: String,
    pub description: Option<String>,
    pub completed: bool,
}