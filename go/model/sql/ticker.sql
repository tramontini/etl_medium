CREATE TABLE tickers (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    symbol VARCHAR NOT NULL,
    rank SMALLINT NOT NULL,
    first_appeared_at TIMESTAMP NOT NULL
);