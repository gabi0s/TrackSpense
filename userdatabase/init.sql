-- users table creation
CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nom VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    solde DECIMAL(10,2) NOT NULL
);

-- user inserts
INSERT INTO users (nom, password, solde) VALUES
('Alice', '123', 1500.00),
('Bob', '456', 2300.50),
('Charlie', '789', 500.75);
