create table products (
    id serial primary key,
    name varchar,
    description varchar,
    price decimal,
    quantity integer
);

insert into products(name, description, price, quantity) values
    ('Camiseta','Camiseta Azul', 39, 5),
	('Tenis', 'Confortável', 139, 3),
	('Fone', 'Muito bom', 59, 9),
    ('Camiseta Polo','Camiseta Polo', 69, 5),
	('Tenis Preto', 'Confortável com estilo', 139, 3),
	('Fone sem fio', 'Muito bom', 59, 9);