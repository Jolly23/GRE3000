create table vocabulary
(
    word_id serial primary key,
    word    varchar(64)  not null,
    mean    varchar(512) not null
);
