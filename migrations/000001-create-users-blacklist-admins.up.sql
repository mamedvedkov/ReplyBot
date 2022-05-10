begin;

create table if not exists reply_bot.blacklists
(
    id          serial
        constraint blacklist_pk
            primary key,
    telegram_id integer not null,
    created_at  timestamp(6) with time zone default now(),
    updated_at  timestamp(6) with time zone default now()
);

create unique index if not exists blacklists_telegram_id_uindex
    on reply_bot.blacklists (telegram_id);

create table if not exists reply_bot.admins
(
    id          serial,
    telegram_id int  not null,
    is_active   bool not null               default true,
    created_at  timestamp(6) with time zone default now(),
    updated_at  timestamp(6) with time zone default now()
);

create unique index if not exists admins_telegram_id_uindex
    on reply_bot.admins (telegram_id);

create table if not exists reply_bot.msg_sessions
(
    msg_id      int
        constraint msg_sessions_pk primary key,
    telegram_id int not null,
    created_at  timestamp(6) with time zone default now()
);

commit;