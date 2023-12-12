create table user_words
(
    user_id    integer references users (user_id) not null,
    word_id    integer references vocabulary (word_id) not null,
    count_mark integer default 0,
    last_mark  timestamp,

    primary key (user_id, word_id)
);
