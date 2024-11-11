-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_accounts(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_name VARCHAR(100) UNIQUE,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    phone_number VARCHAR(25),
    email VARCHAR(100) UNIQUE
);
-- default value when database initializes
INSERT INTO user_accounts (user_name, first_name, last_name, phone_number, email)
VALUES ('johndoe', 'John', 'Doe', '123-456-7890', 'johndoe@example.com');

INSERT INTO user_accounts (user_name, first_name, last_name, phone_number, email)
VALUES ('janedoe', 'Jane', 'Doe', '987-654-3210', 'janedoe@example.com');

INSERT INTO user_accounts (user_name, first_name, last_name, phone_number, email)
VALUES ('alice_smith', 'Alice', 'Smith', '555-123-4567', 'alice.smith@example.com');

INSERT INTO user_accounts (user_name, first_name, last_name, phone_number, email)
VALUES ('bob_jones', 'Bob', 'Jones', '444-777-8888', 'bob.jones@example.com');

INSERT INTO user_accounts (user_name, first_name, last_name, phone_number, email)
VALUES ('charlie_brown', 'Charlie', 'Brown', '333-444-5555', 'charlie.brown@example.com');

-- INSERT INTO user_accounts (user_name, first_name, last_name, phone_number, email)
-- VALUES ('emily_davis', 'Emily', 'Davis', '666-555-4444', 'emily.davis@example.com');

-- INSERT INTO user_accounts (user_name, first_name, last_name, phone_number, email)
-- VALUES ('michael_lee', 'Michael', 'Lee', '222-333-4444', 'michael.lee@example.com');

-- INSERT INTO user_accounts (user_name, first_name, last_name, phone_number, email)
-- VALUES ('susan_clark', 'Susan', 'Clark', '111-222-3333', 'susan.clark@example.com');

-- INSERT INTO user_accounts (user_name, first_name, last_name, phone_number, email)
-- VALUES ('david_wilson', 'David', 'Wilson', '999-888-7777', 'david.wilson@example.com');

-- INSERT INTO user_accounts (user_name, first_name, last_name, phone_number, email)
-- VALUES ('lisa_white', 'Lisa', 'White', '555-666-7777', 'lisa.white@example.com');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
