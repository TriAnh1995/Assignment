
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