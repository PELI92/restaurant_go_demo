CREATE TABLE opening_hours (
    id BINARY(16) NOT NULL PRIMARY KEY,
    restaurant_id BINARY(16) NOT NULL,
    day_of_week TINYINT NOT NULL,
    open_time TIME NOT NULL,
    close_time TIME NOT NULL,
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id) ON DELETE CASCADE
);              