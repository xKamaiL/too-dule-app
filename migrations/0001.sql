create table members
(
    id         uuid                 default gen_random_uuid(),
    username   varchar     not null unique,
    password   varchar     not null,
    email      varchar     not null,
    created_at timestamptz not null default now(),
    primary key (id)
);
