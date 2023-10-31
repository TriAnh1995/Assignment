
CREATE SEQUENCE user_id_seq;
CREATE TABLE user_accounts (
                               "user_id" INT DEFAULT nextval('user_id_seq') PRIMARY KEY,
                               "name" VARCHAR NOT NULL,
                               "email" VARCHAR NOT NULL UNIQUE
);