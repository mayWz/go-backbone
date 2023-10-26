CREATE TABLE users (
    id BINARY(16) PRIMARY KEY,
    name varchar(255) NOT NULL,
    gender varchar(255) NOT NULL,
    email varchar(50) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);