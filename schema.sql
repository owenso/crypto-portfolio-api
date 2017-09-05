CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS cryptos CASCADE;
DROP TABLE IF EXISTS wallets CASCADE;
DROP TABLE IF EXISTS portfolioEntry CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;
DROP TABLE IF EXISTS portfolio CASCADE;
DROP TABLE IF EXISTS friends CASCADE;

CREATE TABLE users(
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	username VARCHAR(50) UNIQUE NOT NULL,
	firstname VARCHAR(50),
	lastname VARCHAR(50),
	email VARCHAR(255) UNIQUE NOT NULL,
	password BYTEA NOT NULL,
    provider VARCHAR(50),
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp,
    lastseen TIMESTAMP DEFAULT current_timestamp
);
CREATE UNIQUE INDEX "lcase_username_idx" ON users (LOWER(username));

CREATE TABLE cryptos(
    id SERIAL PRIMARY KEY,
    symbol VARCHAR(10) UNIQUE,
    name VARCHAR(255),
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE wallets(
	id SERIAL PRIMARY KEY,
	userID UUID REFERENCES users NOT NULL,
    address VARCHAR(255) NOT NULL,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE transactions(
	id SERIAL PRIMARY KEY,
	userID UUID REFERENCES users NOT NULL,
    portfolio SERIAL REFERENCES portfolio NOT NULL,
    price NUMERIC(10, 5),
    transactionDate TIMESTAMP NOT NULL,
    qty INT NOT NULL,
    crypto INTEGER REFERENCES cryptos NOT NULL,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE portfolio(
	id SERIAL PRIMARY KEY,
	userID UUID REFERENCES users NOT NULL,
    title VARCHAR(255) NOT NULL,
    -- PORTFOLIO TYPE = 0: NORMAL 1: PAPERTRADE 2: COMPETITIVE
    type smallint NOT NULL,
    startingBalance NUMERIC(10, 2),
    -- PRIVACY TYPE = 0: PRIVATE 1: FRIENDS ONLY 2: PUBLIC
    privacy smallint DEFAULT 0,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE friends(
    id SERIAL PRIMARY KEY,
	userID UUID REFERENCES users NOT NULL,
    friendID UUID REFERENCES users NOT NULL
)