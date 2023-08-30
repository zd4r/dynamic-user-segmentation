CREATE TABLE "user"
(
    id         bigserial                   PRIMARY KEY,
    login      text                        NOT NULL UNIQUE
);

CREATE TABLE segment
(
    id         bigserial                   PRIMARY KEY,
    slug       text                        NOT NULL UNIQUE
);