CREATE TABLE "user"
(
    id         bigserial                   PRIMARY KEY
);

CREATE TABLE segment
(
    id         bigserial                   PRIMARY KEY,
    slug       text                        NOT NULL UNIQUE
);