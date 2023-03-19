# Data Definition Language (DDL)

## 1. Create database alta_online_shop

create database alta_online_shop

## 2. Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.

## a. create table user

create table users(
user_id int(11) not null auto_increment,
name varchar(255) not null,
address text,
status smallint,
dob date,
gender char(1),
create_at timestamp,
updated_at timestamp,
primary key(user_id)
)

## b. Create table product, product type, operators, product description, payment_method.

create table products(
product_id int(11) not null auto_increment,
name varchar(100) not null,
code varchar(50) not null,
status smallint not null,
created_at timestamp,
updated_at timestamp,
product_type_id int(11),
operator_id int(11),
primary key(product_id),
constraint FK_ProductType foreign key(product_type_id) references product_types(product_type_id),
constraint FK_ProductOperator foreign key(operator_id) references operators(operator_id)
)

create table product_types (
product_type_id int(11) not null auto_increment,
name varchar(255) not null,
create_at timestamp,
update_at timestamp,
primary key(product_type_id)
)

create table product_descriptions (
product_description_id int(11) not null auto_increment,
description text,
created_at timestamp,
updated_at timestamp,
primary key(product_description_id)
)

create table operators (
operator_id int(11) not null auto_increment,
name varchar(255) not null,
created_at timestamp,
updated_at timestamp,
primary key(operator_id)
)

create table payment_methods (
payment_method_id int(11) not null auto_increment,
name varchar(255) not null,
status smallint,
created_at timestamp,
updated_at timestamp,
primary key(payment_method_id)
)

## c. Create table transaction, transaction detail

create table transactions (
transaction_id int(11) not null auto_increment,
status smallint,
total_qty int(11),
total_price numeric(25,2),
created_at timestamp,
updated_at timestamp,
user_id int(11),
payment_method_id int(11),
primary key(transaction_id),
foreign key(user_id) references users(user_id),
foreign key(payment_method_id) references payment_methods(payment_method_id)
)

create table transaction_details (
transaction_id int(11) not null,
product_id int(11) not null,
status varchar(10) not null,
qty int(11) not null,
price numeric(25,2) not null,
created_at timestamp,
updated_at timestamp,
foreign key(transaction_id) references transactions(transaction_id),
foreign key(product_id) references products(product_id)
)

## 3. Create tabel kurir dengan field id, name, created_at, updated_at.

create table kurir (
kurir_id int(11) not null auto_increment,
name varchar(255) not null,
create_at timestamp,
updated_at timestamp,
primary key(kurir_id)
)

## 4. Tambahkan ongkos_dasar column di tabel kurir.

alter table kurir
add ongkos_dasar numeric(25,2) not null

## 5. Rename tabel kurir menjadi shipping

alter table kurir
rename to shipping

## 6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.

drop table shipping
