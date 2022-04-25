create table if not exists reacts
(
    id            serial primary key,
    name          varchar(30)  not null,
    thumbnail_url varchar(250) not null,
    created_at    timestamp default current_timestamp,
    updated_at    timestamp default current_timestamp
);


create table if not exists react_users
(
    id         bigserial primary key,
    author_id  bigint not null references users (id),
    post_id    bigint references posts (id),
    comment_id bigint references comments (id),
    react_id   int    not null references reacts (id),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,

    unique (author_id, post_id),
    unique (author_id, comment_id)
);
