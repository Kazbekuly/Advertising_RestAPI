CREATE TABLE users
(
        id          serial          not null unique,
        username    varchar(255)   not null,
        login       varchar(255)    not null unique,
        password    varchar(255)    not null

);

CREATE TABLE todo_lists
(
    id serial                        not null unique ,
    name            varchar(255)     not null,
    description     varchar(255)     not null,
    image_link      varchar(255),
    price           varchar (255)    not null
);