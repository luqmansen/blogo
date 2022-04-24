alter table react_users drop constraint react_users_author_id_fkey;
alter table react_users drop constraint react_users_react_id_fkey;
drop table if exists react_users;
drop table if exists reacts;
