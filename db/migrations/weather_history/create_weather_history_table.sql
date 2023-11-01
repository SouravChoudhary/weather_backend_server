CREATE TABLE records (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    city_name VARCHAR(255) NOT NULL,
    weather_data TEXT NOT NULL,
    search_time TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_records_user_id ON records (UserID);

CREATE INDEX idx_records_city_name ON records (CityName);
