create table lists(
    id serial primary key,
    title varchar,
    description varchar,
    assignee varchar,
    status varchar,
    deadline timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);