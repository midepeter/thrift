--
CREATE TABLE transactions (
  transaction_id VARCHAR(50) NOT NULL,
  user_id INT NOT NULL UNIQUE,
  currency_id INT  NOT NULL,
  transaction_amount DECIMAL(18,4) NOT NULL,
  transaction_date DATE
);

--
CREATE VIEW balances AS
SELECT user_id, currency_id, SUM(transaction_amount) AS balance_amount,
COUNT(*) AS transaction_count FROM transactions G
GROUP BY user_id, currency_id;

--
CREATE INDEX UQ_balances_user_id_currency_id ON balances (
  user_id,
  currency_id
);



