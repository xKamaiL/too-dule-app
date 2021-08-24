create table todos
(
    id         uuid                 default gen_random_uuid(),
    owner_id   uuid        not null,
    title      varchar     not null,
    content    varchar     not null,

    is_active  boolean     not null default false,

    due_date   timestamptz not null,
    created_at timestamptz not null default now(),

    primary key (id)
);
