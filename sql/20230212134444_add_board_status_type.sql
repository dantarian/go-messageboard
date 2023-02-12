-- +goose Up
-- +goose StatementBegin
CREATE TYPE board_status AS ENUM('open', 'closed');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE board_status;
-- +goose StatementEnd
