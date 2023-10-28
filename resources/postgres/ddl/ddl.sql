-- public.products definition

-- Drop table

-- DROP TABLE public.products;

CREATE TABLE public.products (
	id serial4 NOT NULL,
	"name" varchar(255) NULL,
	description text NULL,
	price varchar(255) NULL,
	rating varchar(15) NULL,
	merchant_name varchar(255) NULL,
	image_link text NULL,
	CONSTRAINT rmr_remisier_pk PRIMARY KEY (id)
);