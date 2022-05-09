CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id          uuid DEFAULT uuid_generate_v4 (),
    first_name  varchar(200),
    last_name   varchar(200),
    national_id varchar(200),
    cancelled   boolean,
    created_at  timestamp without time zone,
    updated_at  timestamp without time zone,
    PRIMARY KEY (id)
);