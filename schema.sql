CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS cryptos CASCADE;
DROP TABLE IF EXISTS wallets CASCADE;
DROP TABLE IF EXISTS portfolioEntry CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;
DROP TABLE IF EXISTS portfolio CASCADE;
DROP TABLE IF EXISTS friends CASCADE;
DROP TABLE IF EXISTS privacy CASCADE;
DROP TABLE IF EXISTS portfolioType CASCADE;
DROP TABLE IF EXISTS alerts CASCADE;
DROP TABLE IF EXISTS phoneNumber CASCADE;
DROP TABLE IF EXISTS portfolioSort CASCADE; 

-- CREATE TABLE phoneNumber(
-- 	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     countryCode VARCHAR(3),
--     phoneNumber VARCHAR(12)
-- );

CREATE TABLE privacy(
    id SERIAL PRIMARY KEY,
    level VARCHAR(25)
);

CREATE TABLE portfolioType(
    id SERIAL PRIMARY KEY,
    type VARCHAR(25)
);

CREATE TABLE users(
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	username VARCHAR(50) UNIQUE NOT NULL,
	firstname VARCHAR(50),
	lastname VARCHAR(50),
    avatar VARCHAR(2083),
	email VARCHAR(255) UNIQUE NOT NULL,
    -- phone uuid REFERENCES phoneNumber,
	password BYTEA NOT NULL,
    provider VARCHAR(50),
    email_confirmed BOOLEAN DEFAULT false,
    email_confirmed_on TIMESTAMP,
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

CREATE TABLE portfolio(
	id SERIAL PRIMARY KEY,
	userID UUID REFERENCES users NOT NULL,
    title VARCHAR(255) NOT NULL,
    portfolioType smallint REFERENCES portfolioType DEFAULT 1,
    startingBalance NUMERIC(10, 2),
    privacy smallint REFERENCES privacy DEFAULT 1,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE portfolioSort(
    userID UUID REFERENCES users NOT NULL,
    portfolioID int REFERENCES portfolio NOT NULL,
    index SERIAL
);

CREATE TABLE transactions(
	id SERIAL PRIMARY KEY,
	userID UUID REFERENCES users NOT NULL,
    portfolio int REFERENCES portfolio NOT NULL,
    price NUMERIC(10, 5),
    transactionDate TIMESTAMP NOT NULL,
    qty INT NOT NULL,
    crypto int REFERENCES cryptos NOT NULL,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE alerts(
    id SERIAL PRIMARY KEY,
    userID UUID REFERENCES users NOT NULL,
    crypto int REFERENCES cryptos NOT NULL,
    emailAlert BOOLEAN DEFAULT false,
    -- smsAlert BOOLEAN DEFAULT false,
    pushAlert BOOLEAN DEFAULT false,
    targetPrice NUMERIC(10, 5),
    -- "==", ">=", "<=", "<", ">"
    comparisonOperator VARCHAR(2)
);

CREATE TABLE friends(
    id SERIAL PRIMARY KEY,
	userID UUID REFERENCES users NOT NULL,
    friendID UUID REFERENCES users NOT NULL
);

INSERT INTO privacy (level) VALUES ('private'), ('friends_only'), ('public');
INSERT INTO portfolioType (type) VALUES ('normal'), ('paper_trade'), ('competitive');