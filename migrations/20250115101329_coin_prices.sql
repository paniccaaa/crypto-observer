-- +goose Up
-- +goose StatementBegin
CREATE TABLE coin_prices (
    id SERIAL PRIMARY KEY,
    coin_id INT NOT NULL REFERENCES coins(id) ON DELETE CASCADE,
    price DECIMAL(20, 8) NOT NULL,
    timestamp INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS coin_prices;
-- +goose StatementEnd
