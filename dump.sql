BEGIN;

CREATE TABLE account
(
    account_id VARCHAR(100),
    name VARCHAR(100),
    email VARCHAR(100),
    balance INTEGER,
    PRIMARY KEY (account_id)
);

END;