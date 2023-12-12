create table users
(
    user_id     serial primary key,
    username    varchar(64)   not null unique,
    password    varchar(64)   not null,
    token       varchar(36)   not null unique,
    avatar      varchar(255)  not null default '',
    email       varchar(255)  not null default '',
    url         varchar(255)  not null default '',
    signature   varchar(1000) not null default '',
    create_time timestamp     not null default current_timestamp,
    latest_time timestamp     not null default current_timestamp
);
