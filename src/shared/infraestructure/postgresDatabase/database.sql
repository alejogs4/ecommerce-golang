CREATE TABLE users
(
  id varchar(150) NOT NULL,
  name varchar(100) NOT NULL,
  email varchar(150) NOT NULL,
  password varchar(250) NOT NULL,
  lastname varchar(150) NOT NULL,
  admin BOOLEAN DEFAULT '0',
  email_verified BOOLEAN DEFAULT '0',
  signup_at DATE DEFAULT now(),
  CONSTRAINT pk_users PRIMARY KEY (id),
  CONSTRAINT unique_key_users_email UNIQUE(email)
);

create type product_state as enum('ACTIVE', 'UNAVAILABLE', 'REMOVED');
CREATE TABLE products
(
  id varchar(150) NOT NULL,
  name varchar(150) NULL,
  description varchar(450) NULL,
  picture varchar(300) NULL,
  state product_state NOT NULL,
  quantity INTEGER NOT NULL,
  price float NOT NULL,
  created_at DATE DEFAULT now(),
  CONSTRAINT pk_products PRIMARY KEY(id),
  CONSTRAINT non_negative_quantity CHECK(quantity >= 0),
  CONSTRAINT non_negative_price CHECK(price >= 0)
);

create type cart_state as enum('InProgress', 'Ordered', 'Removed');
CREATE TABLE cart (
  id varchar(150) NOT NULL,
  user_id varchar(150),
  state cart_state NOT NULL,
  created_at DATE DEFAULT now(),
  CONSTRAINT pk_cart PRIMARY KEY(id),
  CONSTRAINT fk_cart_user FOREIGN KEY(user_id) REFERENCES users(id)
);

create type cart_item_state as enum('ADDED', 'REMOVED');
CREATE TABLE cart_item (
  id varchar(150) NOT NULL,
  cart_id varchar(150),
  product_id varchar(150),
  state cart_item_state NOT NULL,
  created_at DATE DEFAULT now(),
  CONSTRAINT pk_cart_item PRIMARY KEY(id),
  CONSTRAINT fk_cart_item FOREIGN KEY(cart_id) REFERENCES cart(id),
  CONSTRAINT fk_cart_product FOREIGN KEY(product_id) REFERENCES products(id)
);

