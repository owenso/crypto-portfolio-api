-- psql cryptowallet < schema.sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- DROP TABLE IF EXISTS users CASCADE;
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
DROP TABLE IF EXISTS competitionWinners CASCADE; 

-- CREATE TABLE phoneNumber(
-- 	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     countryCode VARCHAR(3),
--     phoneNumber VARCHAR(12)
-- );

CREATE TABLE privacy(
    id SERIAL PRIMARY KEY,
    name VARCHAR(25),
    enabled BOOLEAN,
    description TEXT
);

CREATE TABLE portfolioType(
    id SERIAL PRIMARY KEY,
    name VARCHAR(25),
    enabled BOOLEAN,
    description TEXT
);

CREATE TABLE users(
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	username VARCHAR(50) UNIQUE NOT NULL,
	firstname VARCHAR(50),
	lastname VARCHAR(50),
    avatar VARCHAR(2083) DEFAULT 'http://placecage.com/400/400',
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
    active BOOLEAN DEFAULT true,
    added_by UUID REFERENCES users ,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP
);

CREATE TABLE wallets(
	id SERIAL PRIMARY KEY,
	userID UUID REFERENCES users NOT NULL,
    address VARCHAR(255) NOT NULL,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE portfolio(
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
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
    portfolioID UUID REFERENCES portfolio NOT NULL,
    index SERIAL
);

CREATE TABLE exchanges(
    int SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE transactions(
	id SERIAL PRIMARY KEY,
	userID UUID REFERENCES users NOT NULL,
    portfolio UUID REFERENCES portfolio NOT NULL,
    exchange int REFERENCES exchanges,
    price NUMERIC(10, 5),
    fee NUMERIC(10, 5),
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

CREATE TABLE competitionWinners(
    id SERIAL PRIMARY KEY,
    userID UUID REFERENCES users NOT NULL,
    competition_period DATE NOT NULL
);

INSERT INTO privacy (name, description, enabled) VALUES ('private', 'Only you will see your portfolio details. Your portfolio will not be listed anywhere.', true);
INSERT INTO privacy (name, description, enabled) VALUES ('friends_only', 'Only those you accept as a friend will see your portfolio details. However, they will not be able to see exact figures, only percentages.', true);
INSERT INTO privacy (name, description, enabled) VALUES ('public', 'Anyone can see your portfolio details. However, they will not be able to see exact figures, only percentages.' , true);


INSERT INTO portfolioType (name, description, enabled) VALUES ('normal', 'A portfolio to keep track of your trades.', true);
INSERT INTO portfolioType (name, description, enabled) VALUES ('competition', 'Competition portfolios all start off with 10 BTC. At the end of each month, the portfolio with the largest gain wins and the competion is reset.  Competition portfolios must be public. ', false);
