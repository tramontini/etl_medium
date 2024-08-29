CREATE TABLE prices (
    ticker_id VARCHAR NOT NULL,
    fiat_id VARCHAR NOT NULL,
    price NUMERIC(20, 8) NOT NULL,
    last_updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (ticker_id, fiat_id),
    FOREIGN KEY (ticker_id) REFERENCES tickers(id),
    FOREIGN KEY (fiat_id) REFERENCES fiats(fiat_id)
);
