# Database Schema

## Table: users

- `id` (Serial Primary Key): A unique identifier for each user.
- `username` (VARCHAR): The username of the user, which should be unique.
- `password_hash` (VARCHAR): The hashed password of the user.
- `date_of_birth` (DATE): The date of birth of the user.
- `created_at` (TIMESTAMPTZ): The timestamp with time zone indicating when the user record was created. It defaults to the current timestamp.
- `UNIQUE (username)`: Ensures that each username is unique within the table.

## Table: records

- `id` (Serial Primary Key): A unique identifier for each record.
- `user_id` (INT): A foreign key referencing the `id` of a user who performed the search.
- `city_name` (VARCHAR): The name of the city for which weather data was searched.
- `weather_data` (TEXT): The textual representation of the weather data.
- `search_time` (TIMESTAMPTZ): The timestamp with time zone indicating when the search was performed. It defaults to the current timestamp.
- `CONSTRAINT fk_user FOREIGN KEY (user_id)`: Enforces a foreign key relationship with the `users` table, referencing the `id` of the user, with cascading deletes.
- `UNIQUE (user_id, city_name, search_time)`: Ensures that there are no duplicate records for the same user, city, and search time.

- `INDEX idx_records_user_id`: Index on the `user_id` column for optimization.
- `INDEX idx_records_city_name`: Index on the `city_name` column for optimization.
