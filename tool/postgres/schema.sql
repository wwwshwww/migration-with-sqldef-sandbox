
CREATE TABLE users (
    id bigint NOT NULL PRIMARY KEY,
    name text,
    age int
);
CREATE INDEX index_name ON users(name);

CREATE TABLE bigdata (
    data bigint
);