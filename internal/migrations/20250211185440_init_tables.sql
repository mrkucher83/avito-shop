-- +goose Up
-- +goose StatementBegin
CREATE TABLE employee (
                          id BIGSERIAL PRIMARY KEY,
                          username VARCHAR(255) NOT NULL,
                          password VARCHAR(255) NOT NULL,
                          coins INT NOT NULL DEFAULT 1000
);

CREATE INDEX IF NOT EXISTS idx_employee_id ON employee(id);

CREATE TABLE IF NOT EXISTS merch (
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(255) NOT NULL UNIQUE,
    price   INT NOT NULL CHECK (price > 0)
);

CREATE INDEX IF NOT EXISTS idx_merch_id ON merch(id);

INSERT INTO merch (name, price) VALUES
    ('t-shirt', 80),
    ('cup', 20),
    ('book', 50),
    ('pen', 10),
    ('powerbank', 200),
    ('hoody', 300),
    ('umbrella', 200),
    ('socks', 10),
    ('wallet', 50),
    ('pink-hoody', 500)
ON CONFLICT (name) DO NOTHING;

CREATE TABLE IF NOT EXISTS transactions (
    id            BIGSERIAL PRIMARY KEY,
    sender_id     BIGINT REFERENCES employee(id) ON DELETE SET NULL,
    receiver_id   BIGINT REFERENCES employee(id) ON DELETE SET NULL,
    amount        INT NOT NULL CHECK (amount > 0),
    created_at    TIMESTAMPTZ  DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_transactions_sender_id ON transactions(sender_id);
CREATE INDEX IF NOT EXISTS idx_transactions_receiver_id ON transactions(receiver_id);
CREATE INDEX IF NOT EXISTS idx_transactions_created_at ON transactions(created_at);

CREATE TABLE IF NOT EXISTS purchases (
    id              BIGSERIAL PRIMARY KEY,
    employee_id     BIGINT REFERENCES employee(id) ON DELETE CASCADE,
    merch_id        INT REFERENCES merch(id) ON DELETE CASCADE,
    quantity        INT NOT NULL CHECK (quantity > 0),
    purchased_at    TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_purchases_employee_id ON purchases(employee_id);
CREATE INDEX IF NOT EXISTS idx_purchases_purchased_at ON purchases(purchased_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS purchases;
DROP TABLE IF EXISTS transactions;
DELETE FROM merch WHERE name IN ('t-shirt', 'cup', 'book', 'pen', 'powerbank', 'hoody', 'umbrella', 'socks', 'wallet', 'pink-hoody');
DROP TABLE IF EXISTS merch;
DROP TABLE IF EXISTS employee;
-- +goose StatementEnd
