-- CREATE users table

BEGIN;
create table if not exists users
(
    id uuid default uuid_generate_v1() not null
        constraint users_pkey
            primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    first_name text not null,
    last_name text not null,
    email text not null
);

alter table users owner to "user";

create unique index if not exists uix_users_email
    on users (email);

create unique index if not exists idx_user_first_last_name_unique
    on users (first_name, last_name);

COMMIT;

-- CREATE calendars table
BEGIN;
create table if not exists calendars
(
    id uuid default uuid_generate_v1() not null
        constraint calendars_pkey
            primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name text not null,
    user_id uuid not null
        constraint calendars_user_id_users_id_foreign
            references users
            on update cascade on delete cascade
);

alter table calendars owner to "user";

create unique index if not exists uix_calendars_name
    on calendars (name);
COMMIT;

-- CREATE appointments table
BEGIN;
create table if not exists appointments
(
    id uuid default uuid_generate_v1() not null
        constraint appointments_pkey
            primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    subject text not null,
    description text,
    calendar_id uuid not null
        constraint appointments_calendar_id_calendars_id_foreign
            references calendars
            on update cascade on delete cascade,
    start timestamp with time zone,
    "end" timestamp with time zone,
    whole_day boolean
);

alter table appointments owner to "user";

create index if not exists idx_appointments_subject
    on appointments (subject);

create unique index if not exists idx_calendar_id_subject_unique
    on appointments (calendar_id, subject);
COMMIT;

-- CREATE users_appointments table
BEGIN;
create table if not exists users_appointments
(
    user_id uuid default uuid_generate_v1() not null,
    appointment_id uuid default uuid_generate_v1() not null,
    constraint users_appointments_pkey
        primary key (user_id, appointment_id)
);

alter table users_appointments owner to "user";

COMMIT;
