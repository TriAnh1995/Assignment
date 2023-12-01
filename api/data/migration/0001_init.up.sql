
CREATE SEQUENCE user_id_seq;
CREATE TABLE user_accounts (
     "user_id" INT DEFAULT nextval('user_id_seq') PRIMARY KEY,
     "name" VARCHAR NOT NULL,
     "email" VARCHAR NOT NULL UNIQUE
);

CREATE SEQUENCE friendship_id_seq;
CREATE TABLE friendships
(
    "friendship_id" INT DEFAULT nextval('friendship_id_seq') PRIMARY KEY,
    "user_email_1"  VARCHAR NOT NULL,
    "user_email_2"  VARCHAR NOT NULL,
    CONSTRAINT unique_friends UNIQUE (user_email_1, user_email_2)
);

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    title VARCHAR(10) CHECK (title IN ('mr', 'ms', 'mrs', 'miss', 'dr', '')),
    first_name VARCHAR(50) CHECK (LENGTH(first_name) BETWEEN 2 AND 50),
    last_name VARCHAR(50) CHECK (LENGTH(last_name) BETWEEN 2 AND 50),
    gender VARCHAR(10) CHECK (gender IN ('male', 'female', 'other', '')),
    email VARCHAR(255) UNIQUE,
    date_of_birth DATE CHECK (date_of_birth >= '1900-01-01' AND date_of_birth <= CURRENT_DATE),
    register_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    phone VARCHAR(20),
    password_hash VARCHAR(255)
    );
alter table users add column password_hash VARCHAR(255)