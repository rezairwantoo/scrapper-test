create table products (
	id serial CONSTRAINT rmr_remisier_pk PRIMARY KEY,
	name VARCHAR(255) null,
	description text null,
	price VARCHAR(255) null,
	rating VARCHAR(15) null,
	merchant_name VARCHAR(255) null,
	image_link text null
);