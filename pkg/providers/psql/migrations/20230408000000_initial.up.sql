create extension if not exists "uuid-ossp";

-- Time table
create table if not exists time_table
(
    id         uuid         default uuid_generate_v4(),
    created_at timestamptz  not null default now(),

    unique (id)
);
