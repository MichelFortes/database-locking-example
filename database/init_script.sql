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


insert into events (id, payload) values (gen_random_uuid(), 'json0');
insert into events (id, payload) values (gen_random_uuid(), 'json1');
insert into events (id, payload) values (gen_random_uuid(), 'json2');
insert into events (id, payload) values (gen_random_uuid(), 'json3');
insert into events (id, payload) values (gen_random_uuid(), 'json4');
insert into events (id, payload) values (gen_random_uuid(), 'json5');
insert into events (id, payload) values (gen_random_uuid(), 'json6');
insert into events (id, payload) values (gen_random_uuid(), 'json7');
insert into events (id, payload) values (gen_random_uuid(), 'json8');
insert into events (id, payload) values (gen_random_uuid(), 'json9');