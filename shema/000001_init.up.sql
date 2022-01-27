CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    phone         serial not null unique
);

CREATE TABLE users_auth
(
    id            serial       not null unique,
    phone         serial not null unique,
    code          serial not null
);