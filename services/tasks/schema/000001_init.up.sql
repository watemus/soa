CREATE TABLE tasks
(
    id        serial       not null unique,
    task_name varchar(255) not null,
    body      varchar(255),
    author    varchar(255) not null
);
