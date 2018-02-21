DROP TABLE IF EXISTS accounts;

CREATE TABLE accounts (
	-- カラム
	id                   uuid not NULL primary key,
	created_at           timestamp with time zone not NULL,
	updated_at           timestamp with time zone not NULL,
	deleted_at           timestamp with time zone default null,
    email                text not NULL unique,
    password             text not NULL
);
