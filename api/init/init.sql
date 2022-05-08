create table if not exists todos (
	id serial primary key,
	content text not null,
	priority  int,
	done bool,
	created_at timestamp with time zone,
	updated_at timestamp with time zone
);
