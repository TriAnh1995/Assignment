CREATE SEQUENCE user_id_seq;
CREATE TABLE user_accounts
(
    "user_id" INT DEFAULT nextval('user_id_seq') PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL UNIQUE
);

CREATE SEQUENCE friendship_id_seq;
CREATE TYPE subscription_type AS ENUM ('followed', 'blocked', 'default');
CREATE TYPE friendship_type AS ENUM ('friend', 'stranger', 'default');
CREATE TABLE relationship
(
    "friendship_id" INT DEFAULT nextval('friendship_id_seq') PRIMARY KEY,
    "user_email_1"  VARCHAR NOT NULL,
    "user_email_2"  VARCHAR NOT NULL,
    "subscription" subscription_type NOT NULL,
    "friendship" friendship_type NOT NULL,
    CONSTRAINT unique_relationship UNIQUE (user_email_1, user_email_2)
);