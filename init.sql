CREATE TABLE "user"
(
    id         bigint                      PRIMARY KEY
);

CREATE TABLE segment
(
    id         bigserial                   PRIMARY KEY,
    slug       text                        NOT NULL UNIQUE
);

CREATE TABLE experiment
(
    user_id    bigserial,
    segment_id bigserial,
    PRIMARY KEY (user_id, segment_id),
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE,
    FOREIGN KEY (segment_id) REFERENCES segment(id) ON DELETE CASCADE
);

CREATE TABLE report
(
    id              bigserial                   PRIMARY KEY,
    user_id         bigserial,
    segment_slug    text                        NOT NULL,
    action          text                        NOT NULL,
    date            timestamp(0) with time zone NOT NULL DEFAULT NOW()
)