CREATE TABLE records (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    city_name VARCHAR(255) NOT NULL,
    weather_data TEXT NOT NULL,
    search_time TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE (user_id, city_name, search_time)
);

CREATE INDEX idx_records_user_id ON records (user_id);  
CREATE INDEX idx_records_city_name ON records (city_name); 