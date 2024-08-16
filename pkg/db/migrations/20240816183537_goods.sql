-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS goodsDB (
    id VARCHAR(36) PRIMARY KEY, 
    name VARCHAR(100) NOT NULL, 
    provider INT NOT NULL,
    price INT NOT NULL,
    quantity INT NOT NULL
);

INSERT INTO goodsDB (id, name, provider, price, quantity) 
VALUES ('8617bf49-39a9-4268-b113-7b6bcd189ba2', 'Goods 3', '1', '152','64'),
    ('5217bf49-39a9-4268-b113-7b6bcd189ba3', 'Goods 4', '2', '263','35');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE goodsDB;
-- +goose StatementEnd
