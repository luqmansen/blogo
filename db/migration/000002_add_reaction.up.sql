create table if not exists reacts
(
    id            int primary key,
    name          varchar(30)  not null,
    thumbnail_url varchar(250) not null,
    created_at    timestamp default current_timestamp,
    updated_at    timestamp default current_timestamp
);


create table if not exists react_users
(
    id            bigserial primary key,
    author_id     bigint not null references users (id),
    resource_kind int    not null,
    react_id      int    not null references reacts (id),
    created_at    timestamp default current_timestamp,
    updated_at    timestamp default current_timestamp
);
