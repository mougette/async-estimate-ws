create table users (
    id integer primary key autoincrement,
    name text not null unique,
    password text not null
);

create table company (
    id integer primary key autoincrement,
    name text not null
);

create table tickets (
    id integer primary key autoincrement,
    company_id integer not null,
    jira_url text
);


