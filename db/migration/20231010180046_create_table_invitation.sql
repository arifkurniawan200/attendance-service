-- +goose Up
-- +goose StatementBegin
CREATE TABLE invitations (
     id INT AUTO_INCREMENT PRIMARY KEY,
     member_id INT NOT NULL,
     gathering_id INT NOT NULL,
     status ENUM('Sent', 'Approve', 'Reject') NOT NULL,
     FOREIGN KEY (member_id) REFERENCES members(id),
     FOREIGN KEY (gathering_id) REFERENCES gatherings(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE invitations;
-- +goose StatementEnd
