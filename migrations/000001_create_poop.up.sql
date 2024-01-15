create table if not exists poop(
    id varchar(36) not null primary key,
    time timestamp,
    lat float,
    long float,
    user_id varchar(36)
);
