CREATE DATABASE appdb;

CREATE extension IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS user_account
(
    id uuid PRIMARY KEY default uuid_generate_v4(),
    created_at BIGINT DEFAULT extract(epoch from now()),
    email TEXT UNIQUE,
    first_name TEXT DEFAULT '',
    last_name TEXT DEFAULT ''
);