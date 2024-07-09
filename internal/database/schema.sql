CREATE TABLE IF NOT EXISTS tb_category (id bigserial not null,name varchar(255),primary key (id));
CREATE TABLE IF NOT EXISTS tb_order (id bigserial not null,status smallint check (status between 0 and 4), client_id bigint, moment TIMESTAMP WITHOUT TIME ZONE, primary key (id));
CREATE TABLE IF NOT EXISTS tb_order_item (order_id bigint not null, product_id bigint not null, primary key (order_id, product_id), price float(53), quantity integer);
CREATE TABLE IF NOT EXISTS tb_payment (order_id bigint not null, moment TIMESTAMP WITHOUT TIME ZONE, primary key (order_id));
CREATE TABLE IF NOT EXISTS tb_product (id bigserial not null, price float(53), description TEXT, img_url varchar(255), name varchar(255), primary key (id), seller bigint);
CREATE TABLE IF NOT EXISTS tb_product_category (category_id bigint not null, product_id bigint not null, primary key (category_id, product_id));
CREATE TABLE IF NOT EXISTS tb_role (id bigserial not null, authority varchar(255), primary key (id));
CREATE TABLE tb_location (
                             id SERIAL PRIMARY KEY,
                             user_id INT REFERENCES tb_user(id),
                             street_number INT,
                             street_name VARCHAR(100),
                             city VARCHAR(100),
                             state VARCHAR(100),
                             country VARCHAR(100),
                             postcode VARCHAR(20),
                             coordinates_latitude NUMERIC,
                             coordinates_longitude NUMERIC,
                             timezone_offset VARCHAR(20),
                             timezone_description VARCHAR(100)
);
CREATE TABLE tb_user (
                         id SERIAL PRIMARY KEY,
                         gender VARCHAR(10),
                         name_title VARCHAR(20),
                         name_first VARCHAR(50),
                         name_last VARCHAR(50),
                         email VARCHAR(100),
                         login_uuid UUID,
                         login_username VARCHAR(50),
                         login_password VARCHAR(50),
                         login_salt VARCHAR(50),
                         login_md5 VARCHAR(50),
                         login_sha1 VARCHAR(50),
                         login_sha256 VARCHAR(50),
                         dob_date TIMESTAMP,
                         dob_age INTEGER,
                         registered_date TIMESTAMP,
                         registered_age INTEGER,
                         phone VARCHAR(20),
                         cell VARCHAR(20),
                         id_name VARCHAR(50),
                         id_value VARCHAR(50),
                         picture_large TEXT,
                         picture_medium TEXT,
                         picture_thumbnail TEXT,
                         nat VARCHAR(10),
                         location_id INT REFERENCES tb_location(id)
);
CREATE TABLE IF NOT EXISTS tb_user_role (role_id bigint not null, user_id bigint not null, primary key (role_id, user_id));
alter table if exists tb_order add constraint fk_order_user foreign key (client_id) references tb_user;
alter table if exists tb_order_item add constraint fk_order_item_product foreign key (product_id) references tb_product;
alter table if exists tb_order_item add constraint fk_order_item_order foreign key (order_id) references tb_order;
alter table if exists tb_payment add constraint fk_payment_order foreign key (order_id) references tb_order;
alter table if exists tb_product add constraint fk_product_user foreign key (seller) references tb_user;
alter table if exists tb_product_category add constraint fk_product_category_category foreign key (category_id) references tb_category;
alter table if exists tb_product_category add constraint fk_product_category_product foreign key (product_id) references tb_product;
alter table if exists tb_user_role add constraint fk_user_role_role foreign key (role_id) references tb_role;
alter table if exists tb_user_role add constraint fk_user_role_user foreign key (user_id) references tb_user;