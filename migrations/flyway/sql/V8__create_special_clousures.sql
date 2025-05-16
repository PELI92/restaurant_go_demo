CREATE TABLE special_closures (
    id BINARY(16) PRIMARY KEY,
    restaurant_id BINARY(16) NOT NULL,
    closed_date DATE NOT NULL,
    reason VARCHAR(255),
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id) ON DELETE CASCADE
);
