begin;

drop index if exists reply_bot.admins_telegram_id_uindex;
drop table if exists reply_bot.admins;

drop index if exists reply_bot.blacklists_telegram_id_uindex;
drop table if exists reply_bot.blacklists;

drop table if exists reply_bot.msg_sessions;

commit;