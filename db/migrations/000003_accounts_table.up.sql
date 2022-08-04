CREATE TABLE accounts (   
    account_id INT NOT NULL,
    account_balance VARCHAR(255),
    account_status VARCHAR(255),
    currency VARCHAR(255),
    created_at date NOT NULL,
    updated_at date NOT NULL,
    deleted_at date NOT NULL
);
