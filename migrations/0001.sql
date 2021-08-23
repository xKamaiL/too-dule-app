create table members
(
    id         uuid                 default gen_random_uuid(),
    content    varchar     not null,
    done       boolean     not null default false,
    created_at timestamptz not null default now(),
    primary key (id)
);
