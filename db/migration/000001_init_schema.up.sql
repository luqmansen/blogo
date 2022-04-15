create table if not exists users
(
    id         bigint unsigned primary key auto_increment,
    username   varchar(30) not null unique,
    password   text        not null,
    created_at timestamp default current_timestamp,
    updated_at datetime  default current_timestamp on update current_timestamp
);

create table if not exists posts
(
    id         bigint unsigned primary key auto_increment,
    author_id  bigint unsigned not null,
    title      varchar(100)    not null,
    content    text            not null,
    created_at timestamp default current_timestamp,
    updated_at datetime  default current_timestamp on update current_timestamp,

    foreign key (author_id) references users (id)

);

create table if not exists comments
(
    id          bigint unsigned primary key auto_increment,
    parent_post bigint unsigned,
    parent_id   bigint unsigned,
    author_id   bigint unsigned not null,
    content     text            not null,
    created_at  timestamp default current_timestamp,
    updated_at  datetime  default current_timestamp on update current_timestamp,

    foreign key (author_id) references users (id),
    foreign key (parent_id) references comments (id)
);

