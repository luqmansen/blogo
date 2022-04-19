
create table if not exists users
(
    id         bigserial primary key,
    username   varchar(30) not null unique,
    password   text        not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table if not exists posts
(
    id         bigserial primary key,
    author_id  bigint       not null,
    title      varchar(100) not null,
    content    text         not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,

    foreign key (author_id) references users (id)

);

create table if not exists comments
(
    id             bigserial primary key,
    parent_post_id bigint not null references posts (id),
    parent_id      bigint references comments(id),
    author_id      bigint not null references users (id),
    content        text   not null,
    created_at     timestamp default current_timestamp,
    updated_at     timestamp default current_timestamp

);

