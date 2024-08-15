INSERT INTO tb_location (street_number,street_name,city,state,country,postcode,coordinates_latitude,coordinates_longitude,timezone_offset,timezone_description)
VALUES (7170,'Lønneveien','Kårvåg','Bergen','Norway','4325','45.2549','-140.0653','+3:30','Tehran');
INSERT INTO tb_user (gender,name_title,name_first,name_last,email,login_uuid,login_username,login_password,login_salt,login_md5,login_sha1,login_sha256,dob_date,dob_age,registered_date,registered_age,phone,cell,id_name,id_value,picture_large,picture_medium,picture_thumbnail,nat,location_id)
VALUES ('male','Mr','Administrator','Torstensen','admin@gmail.com','3e9b8621-8e4e-4870-a695-7e36de8e6631','admin','$2a$10$N7SkKCa3r17ga.i.dF9iy.BFUBL2n3b6Z1CWSZWi/qy7ABq/E6VpO','I6r6GBa3','6133f62f8b15e40a99b803650850b315','bbd30114162d66ff28b3549ee86fd9117772e695','f2731308c173cf6aafda4cd2e16dc1c4efa3ff6a381b33c9c740a0d41afcaa3a','1994-07-01T13:54:14.593Z'::TIMESTAMP,30,'2006-10-15T10:00:49.494Z'::TIMESTAMP,17,'89593420','45700142','FN','01079421592','https://randomuser.me/api/portraits/men/39.jpg','https://randomuser.me/api/portraits/med/men/39.jpg','https://randomuser.me/api/portraits/thumb/men/39.jpg','NO',1);
INSERT INTO tb_user (gender,name_title,name_first,name_last,email,login_uuid,login_username,login_password,login_salt,login_md5,login_sha1,login_sha256,dob_date,dob_age,registered_date,registered_age,phone,cell,id_name,id_value,picture_large,picture_medium,picture_thumbnail,nat,location_id)
VALUES ('male','Mr','Seller','Torstensen','seller@gmail.com','3e9b8621-8e4e-4870-a695-7e36de8e6631','seller','$2a$10$N7SkKCa3r17ga.i.dF9iy.BFUBL2n3b6Z1CWSZWi/qy7ABq/E6VpO','I6r6GBa3','6133f62f8b15e40a99b803650850b315','bbd30114162d66ff28b3549ee86fd9117772e695','f2731308c173cf6aafda4cd2e16dc1c4efa3ff6a381b33c9c740a0d41afcaa3a','1994-07-01T13:54:14.593Z'::TIMESTAMP,30,'2006-10-15T10:00:49.494Z'::TIMESTAMP,17,'89593420','45700142','FN','01079421592','https://randomuser.me/api/portraits/men/40.jpg','https://randomuser.me/api/portraits/med/men/40.jpg','https://randomuser.me/api/portraits/thumb/men/40.jpg','NO',1);
INSERT INTO tb_user (gender,name_title,name_first,name_last,email,login_uuid,login_username,login_password,login_salt,login_md5,login_sha1,login_sha256,dob_date,dob_age,registered_date,registered_age,phone,cell,id_name,id_value,picture_large,picture_medium,picture_thumbnail,nat,location_id)
VALUES ('male','Mr','Client','Torstensen','client@gmail.com','3e9b8621-8e4e-4870-a695-7e36de8e6631','client','$2a$10$N7SkKCa3r17ga.i.dF9iy.BFUBL2n3b6Z1CWSZWi/qy7ABq/E6VpO','I6r6GBa3','6133f62f8b15e40a99b803650850b315','bbd30114162d66ff28b3549ee86fd9117772e695','f2731308c173cf6aafda4cd2e16dc1c4efa3ff6a381b33c9c740a0d41afcaa3a','1994-07-01T13:54:14.593Z'::TIMESTAMP,30,'2006-10-15T10:00:49.494Z'::TIMESTAMP,17,'89593420','45700142','FN','01079421592','https://randomuser.me/api/portraits/men/41.jpg','https://randomuser.me/api/portraits/med/men/41.jpg','https://randomuser.me/api/portraits/thumb/men/41.jpg','NO',1);
INSERT INTO tb_role (authority)
VALUES ('ROLE_ADMIN');
INSERT INTO tb_role (authority)
VALUES ('ROLE_SELLER');
INSERT INTO tb_role (authority)
VALUES ('ROLE_CLIENT');
INSERT INTO tb_user_role (user_id, role_id)
VALUES (1, 1);
INSERT INTO tb_user_role (user_id, role_id)
VALUES (2, 2);
INSERT INTO tb_user_role (user_id, role_id)
VALUES (3, 3);
INSERT INTO tb_category(name)
VALUES ('Livros');
INSERT INTO tb_category(name)
VALUES ('Eletrônicos');
INSERT INTO tb_category(name)
VALUES ('Computadores');
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('The Lord of the Rings', 90.5,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/1-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('Smart TV', 2190.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/2-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('Macbook Pro', 1250.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/3-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer', 1200.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/4-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('Rails for Dummies', 100.99,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/5-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Ex', 1350.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/6-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer X', 1350.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/7-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Alfa', 1850.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/8-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Tera', 1950.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/9-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Y', 1700.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/10-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Nitro', 1450.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/11-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Card', 1850.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/12-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Plus', 1350.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/13-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Hera', 2250.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/14-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Weed', 2200.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/15-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Max', 2340.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/16-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Turbo', 1280.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/17-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Hot', 1450.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/18-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Ez', 1750.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/19-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Tr', 1650.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/20-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Tx', 1680.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/21-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Er', 1850.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/22-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Min', 2250.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/23-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Boo', 2350.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/24-big.jpg', 2, 99);
INSERT INTO tb_product (name, price, description, img_url, seller, quantity)
VALUES ('PC Gamer Foo', 4170.0,
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
        'https://raw.githubusercontent.com/devsuperior/dscatalog-resources/master/backend/img/25-big.jpg', 2, 99);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (1, 1);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (2, 2);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (2, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (3, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (4, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (5, 1);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (6, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (7, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (8, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (9, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (10, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (11, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (12, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (13, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (14, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (15, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (16, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (17, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (18, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (19, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (20, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (21, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (22, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (23, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (24, 3);
INSERT INTO tb_product_category (product_id, category_id)
VALUES (25, 3);
INSERT INTO tb_order (moment, status, client_id)
VALUES (TIMESTAMP WITH TIME ZONE '2022-07-25T13:00:00Z', 1, 1);
INSERT INTO tb_order (moment, status, client_id)
VALUES (TIMESTAMP WITH TIME ZONE '2022-07-29T15:50:00Z', 3, 2);
INSERT INTO tb_order (moment, status, client_id)
VALUES (TIMESTAMP WITH TIME ZONE '2022-08-03T14:20:00Z', 0, 1);
INSERT INTO tb_order_item (order_id, product_id, quantity, price)
VALUES (1, 1, 2, 90.5);
INSERT INTO tb_order_item (order_id, product_id, quantity, price)
VALUES (1, 3, 1, 1250.0);
INSERT INTO tb_order_item (order_id, product_id, quantity, price)
VALUES (2, 3, 1, 1250.0);
INSERT INTO tb_order_item (order_id, product_id, quantity, price)
VALUES (3, 1, 1, 90.5);
INSERT INTO tb_payment (order_id, moment)
VALUES (1, TIMESTAMP WITH TIME ZONE '2022-07-25T15:00:00Z');
INSERT INTO tb_payment (order_id, moment)
VALUES (2, TIMESTAMP WITH TIME ZONE '2022-07-30T11:00:00Z');
