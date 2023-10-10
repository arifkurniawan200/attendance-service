-- +goose Up
-- +goose StatementBegin
CREATE TABLE gatherings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    creator INT NOT NULL,
    type ENUM('online', 'offline') NOT NULL,
    schedule_at DATETIME NOT NULL,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (creator) REFERENCES members(id),
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE gatherings;
-- +goose StatementEnd
