CREATE TABLE restaurants (
    id BINARY(16) NOT NULL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(100),
    website VARCHAR(100),
    address_id BINARY(16),
    opening_hours VARCHAR(255),
    status TINYINT DEFAULT 1 CHECK (status IN (0, 1, 2, 3)),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (address_id) REFERENCES addresses(id) ON DELETE CASCADE
);