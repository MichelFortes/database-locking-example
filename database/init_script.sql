create type event_status as enum ('idle', 'publishing', 'published');

create table public.events (
	id uuid not null,
	payload varchar not null,
	registered_at timestamp default now() not null,
	attempts int2 default 1 not null,
	status public.event_status default 'idle'::event_status not null,
	constraint events_pk primary key (id)
);

create index if not exists events_registered_at on events (registered_at);
create index if not exists events_attempts on events (attempts);
create index if not exists events_status on events (status);


insert into events (id, payload) values (gen_random_uuid(), 'json 0');
insert into events (id, payload) values (gen_random_uuid(), 'json 1');
insert into events (id, payload) values (gen_random_uuid(), 'json 2');
insert into events (id, payload) values (gen_random_uuid(), 'json 3');
insert into events (id, payload) values (gen_random_uuid(), 'json 4');
insert into events (id, payload) values (gen_random_uuid(), 'json 5');
insert into events (id, payload) values (gen_random_uuid(), 'json 6');
insert into events (id, payload) values (gen_random_uuid(), 'json 7');
insert into events (id, payload) values (gen_random_uuid(), 'json 8');
insert into events (id, payload) values (gen_random_uuid(), 'json 9');