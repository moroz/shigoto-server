create extension if not exists citext with schema public;

create table users (
  id integer primary key generated by default as identity,
  email citext unique not null,
  password_hash text,
  inserted_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc')
);