create table todos (
    id serial not null,
    todo varchar(255),
    done boolean,
    primary key(id)
);

insert into todos(todo, done) values("Task one", false);