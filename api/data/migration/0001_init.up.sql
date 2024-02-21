
CREATE SEQUENCE user_id_seq;
CREATE TABLE user_accounts
(
    "user_id" INT DEFAULT nextval('user_id_seq') PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL UNIQUE
);

CREATE TYPE update_type AS ENUM ('updated', 'default');
ALTER TABLE user_accounts
    ADD COLUMN "topic" update_type NOT NULL default 'default',
    ADD COLUMN "topic_body" VARCHAR;

CREATE SEQUENCE friendship_id_seq;
CREATE TABLE friendships
(
    "friendship_id" INT DEFAULT nextval('friendship_id_seq') PRIMARY KEY,
    "user_email_1"  VARCHAR NOT NULL,
    "user_email_2"  VARCHAR NOT NULL,
    CONSTRAINT unique_friends UNIQUE (user_email_1, user_email_2)
);

CREATE SEQUENCE subscription_id_seq;
CREATE TYPE status_type AS ENUM ('followed', 'blocked', 'default');

CREATE TABLE subscription
(
    "subscription_id" INT DEFAULT nextval('subscription_id_seq') PRIMARY KEY,
    "requester" VARCHAR NOT NULL,
    "target" VARCHAR NOT NULL,
    "status" status_type NOT NULL,
    CONSTRAINT subscribed UNIQUE (requester, target)
);