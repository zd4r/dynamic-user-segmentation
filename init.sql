CREATE TABLE "user"
(
    id         bigint                      PRIMARY KEY
);

CREATE TABLE segment
(
    id         bigserial                   PRIMARY KEY,
    slug       text                        NOT NULL UNIQUE
);