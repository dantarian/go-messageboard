-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION pgcrypto;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP EXTENSION pgcrypto;
-- +goose StatementEnd
