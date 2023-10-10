-- +goose Up
-- +goose StatementBegin
CREATE TABLE attendees (
     id INT AUTO_INCREMENT PRIMARY KEY,
     member_id INT NOT NULL,
     gathering_id INT NOT NULL,
     FOREIGN KEY (member_id) REFERENCES members(id),
     FOREIGN KEY (gathering_id) REFERENCES gatherings(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table attendees;
-- +goose StatementEnd
