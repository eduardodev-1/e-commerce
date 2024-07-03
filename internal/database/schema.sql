BEGIN;
CREATE TABLE IF NOT EXISTS tb_category (id bigserial not null,name varchar(255),primary key (id));
CREATE TABLE IF NOT EXISTS tb_order (status smallint check (status between 0 and 4), client_id bigint, id bigserial not null, moment TIMESTAMP WITHOUT TIME ZONE, primary key (id));
CREATE TABLE IF NOT EXISTS tb_order_item (price float(53), quantity integer, order_id bigint not null, product_id bigint not null, primary key (order_id, product_id));
CREATE TABLE IF NOT EXISTS tb_payment (moment TIMESTAMP WITHOUT TIME ZONE, order_id bigint not null, primary key (order_id));
CREATE TABLE IF NOT EXISTS tb_product (price float(53), id bigserial not null, description TEXT, img_url varchar(255), name varchar(255), primary key (id));
CREATE TABLE IF NOT EXISTS tb_product_category (category_id bigint not null, product_id bigint not null, primary key (category_id, product_id));
CREATE TABLE IF NOT EXISTS tb_role (id bigserial not null, authority varchar(255), primary key (id));
CREATE TABLE IF NOT EXISTS tb_user (birth_date date, id bigserial not null, email varchar(255) unique, name varchar(255), password varchar(255), phone varchar(255), primary key (id));
CREATE TABLE IF NOT EXISTS tb_user_role (role_id bigint not null, user_id bigint not null, primary key (role_id, user_id));
alter table if exists tb_order add constraint FKi0x0rv7d65vsceuy33km9567n foreign key (client_id) references tb_user;
alter table if exists tb_order_item add constraint FK4h5xid5qehset7qwe5l9c997x foreign key (product_id) references tb_product;
alter table if exists tb_order_item add constraint FKgeobgl2xu916he8vhljktwxnx foreign key (order_id) references tb_order;
alter table if exists tb_payment add constraint FKokaf4il2cwit4h780c25dv04r foreign key (order_id) references tb_order;
alter table if exists tb_product_category add constraint FK5r4sbavb4nkd9xpl0f095qs2a foreign key (category_id) references tb_category;
alter table if exists tb_product_category add constraint FKgbof0jclmaf8wn2alsoexxq3u foreign key (product_id) references tb_product;
alter table if exists tb_user_role add constraint FKea2ootw6b6bb0xt3ptl28bymv foreign key (role_id) references tb_role;
alter table if exists tb_user_role add constraint FK7vn3h53d0tqdimm8cp45gc0kl foreign key (user_id) references tb_user;
COMMIT;