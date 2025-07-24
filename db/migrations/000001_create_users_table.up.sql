create type user_role as enum('admin', 'manager', 'client');

create table users(
    id serial primary key,
    role user_role not null,
    email varchar(255) not null,
    password text not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    deleted_at timestamp
);