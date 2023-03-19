# NOTE : saat disave otomatis ada backslashnya '\', ex : \*

# Data Manipulation Language (DML)

## insert

### 1. Insert 5 operators pada table operators.

insert into operators (name, created_at, updated_at)
values ('baba','2023-03-18 08:20:00','2023-03-18 08:20:00'),
('bibi','2023-03-18 08:20:00','2023-03-18 08:20:00'),
('bubu','2023-03-18 08:20:00','2023-03-18 08:20:00'),
('bebe','2023-03-18 08:20:00','2023-03-18 08:20:00'),
('bobo','2023-03-18 08:20:00','2023-03-18 08:20:00')

### 2. Insert 3 product type.

insert into product_types (name, create_at, update_at)
values ('baju','2023-03-18 08:20:00','2023-03-18 08:20:00'),
('sandal','2023-03-18 08:20:00','2023-03-18 08:20:00'),
('sepatu','2023-03-18 08:20:00','2023-03-18 08:20:00')

### 3. Insert 2 product dengan product type id = 1, dan operators id = 3.

insert into products (name,code,status,created_at, updated_at, product_type_id,operator_id)
values ('t-shirt', 'shirt123',1,'2023-03-18 08:20:00','2023-03-18 08:20:00',1,3),
('hnm', 'hnm123',1,'2023-03-18 08:20:00','2023-03-18 08:20:00',1,3)

### 4. Insert 3 product dengan product type id = 2, dan operators id = 1.

insert into products (name,code,status,created_at, updated_at, product_type_id,operator_id)
values ('eiger', 'eiger',1,'2023-03-18 08:20:00','2023-03-18 08:20:00',2,1),
('homyped', 'homy',1,'2023-03-18 08:20:00','2023-03-18 08:20:00',2,1),
('swallow', 'sllw',1,'2023-03-18 08:20:00','2023-03-18 08:20:00',2,1)

### 5. Insert 3 product dengan product type id = 3, dan operators id = 4.

insert into products (name,code,status,created_at, updated_at, product_type_id,operator_id)
values ('ventella', 'ventella',1,'2023-03-18 08:20:00','2023-03-18 08:20:00',3,4),
('nah', 'nah',1,'2023-03-18 08:20:00','2023-03-18 08:20:00',3,4),
('compass', 'compss',1,'2023-03-18 08:20:00','2023-03-18 08:20:00',3,4)

### 6. Insert product description pada setiap product.

insert into product_descriptions (description,created_at, updated_at)
values ('baju t-shirt', '2023-03-18 08:40:00', '2023-03-18 08:40:00'),
('baju hnm', '2023-03-18 08:40:00', '2023-03-18 08:40:00'),
('sandal eiger', '2023-03-18 08:40:00', '2023-03-18 08:40:00'),
('sandal homyped', '2023-03-18 08:40:00', '2023-03-18 08:40:00'),
('sandal swallow', '2023-03-18 08:40:00', '2023-03-18 08:40:00'),
('sepatu ventella', '2023-03-18 08:40:00', '2023-03-18 08:40:00'),
('sepatu nah', '2023-03-18 08:40:00', '2023-03-18 08:40:00'),
('sepatu compass', '2023-03-18 08:40:00', '2023-03-18 08:40:00')

### 7. Insert 3 payment methods.

insert into payment_methods (name, status, created_at, updated_at)
values ('cash on delivery', 1, '2023-03-18 08:40:00','2023-03-18 08:40:00' ),
('transfer', 1, '2023-03-18 08:40:00','2023-03-18 08:40:00'),
('virtual account',1,'2023-03-18 08:40:00','2023-03-18 08:40:00')

### 8. Insert 5 user pada tabel user.

insert into users (name,address, status, dob,gender,create_at, updated_at)
values ('nana', 'jakarta',1,'2001-01-01','perempuan','2023-03-18 08:40:00','2023-03-18 08:40:00'),
('nini', 'bogor',1,'2001-01-01','perempuan','2023-03-18 08:40:00','2023-03-18 08:40:00'),
('nunu', 'depok',1,'2001-01-01','laki - laki','2023-03-18 08:40:00','2023-03-18 08:40:00'),
('nene', 'tangerang',1,'2001-01-01','perempuan','2023-03-18 08:40:00','2023-03-18 08:40:00'),
('nono', 'bekasi',1,'2001-01-01','perempuan','2023-03-18 08:40:00','2023-03-18 08:40:00')

### 9. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)

insert into transactions (status, total_qty,total_price, created_at,updated_at,user_id,payment_method_id)
values (1,2,'15000','2023-03-18 09:00:00','2023-03-18 09:00:00',1,1 ),
(1,3,'27000','2023-03-18 09:00:00','2023-03-18 09:00:00',1,2 ),
(1,5,'30000','2023-03-18 09:00:00','2023-03-18 09:00:00',1,3 ),
(1,2,'15000','2023-03-18 09:00:00','2023-03-18 09:00:00',2,1 ),
(1,3,'27000','2023-03-18 09:00:00','2023-03-18 09:00:00',2,2 ),
(1,5,'30000','2023-03-18 09:00:00','2023-03-18 09:00:00',2,3 ),
(1,2,'15000','2023-03-18 09:00:00','2023-03-18 09:00:00',3,1 ),
(1,3,'27000','2023-03-18 09:00:00','2023-03-18 09:00:00',3,2 ),
(1,5,'30000','2023-03-18 09:00:00','2023-03-18 09:00:00',3,3 ),
(1,2,'15000','2023-03-18 09:00:00','2023-03-18 09:00:00',4,1 ),
(1,3,'27000','2023-03-18 09:00:00','2023-03-18 09:00:00',4,2 ),
(1,5,'30000','2023-03-18 09:00:00','2023-03-18 09:00:00',4,3 ),
(1,2,'15000','2023-03-18 09:00:00','2023-03-18 09:00:00',5,1 ),
(1,3,'27000','2023-03-18 09:00:00','2023-03-18 09:00:00',5,2 ),
(1,5,'30000','2023-03-18 09:00:00','2023-03-18 09:00:00',5,3 )

### 10. Insert 3 product di masing-masing transaksi.

insert into transaction_details (transaction_id, product_id,status,qty,price,created_at,updated_at)
values (1,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(1,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(1,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(2,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(2,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(2,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(3,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(3,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(3,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(4,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(4,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(4,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(5,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(5,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(5,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(6,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(6,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(6,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(7,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(7,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(7,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(8,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(8,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(8,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(9,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(9,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(9,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(10,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(10,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(10,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(11,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(11,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(11,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(12,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(12,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(12,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(13,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(13,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(13,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(14,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(14,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(14,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(15,1,'lunas',2,15000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(15,2,'lunas',2,27000,'2023-03-18 09:00:00','2023-03-18 09:00:00'),
(15,3,'lunas',2,30000,'2023-03-18 09:00:00','2023-03-18 09:00:00')

## update

### a. Ubah data product id 1 dengan nama ‘product dummy’.

update products set name = 'product dummy' where product_id = 1

### b. Update qty = 3 pada transaction detail dengan product id = 1.

update transaction_details set qty = 3 where product_id = 1

## delete

### a. Delete data pada tabel product dengan id = 1.

delete from products where product_id = 1

### b. Delete pada pada tabel product dengan product type id 1.

delete from products where product_type_id = 1

# Join, Union, Sub Query, Function

## 1. Gabungkan data transaksi dari user id 1 dan user id 2

select _ from transactions where user_id = 1
union
select _ from transactions where user_id = 2

## 2. Tampilkan jumlah harga transaksi user id 1.

select sum(total_price) from transactions t where user_id = 1

## 3. Tampilkan total transaksi dengan product type 2.

select sum(total_price) from transactions t join transaction_details td on t.transaction_id = td.transaction_id join products p on td.product_id = p.product_id join product_types pt on p.product_type_id = pt.product_type_id where pt.product_type_id = 2

## 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan.

select \* from products p left join product_types pt on p.product_type_id = pt.product_type_id

## 5. Tampilkan semua field table transaction, field name table product dan field name table user.

select t.\*, p.name, u.name from transactions t join users u on t.user_id = u.user_id join transaction_details td on t.transaction_id = td.transaction_id join products p on td.product_id = p.product_id

## 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.

delimiter $$
create trigger delete_data_transaction_user
	before delete 
	on transactions
	for each row 
begin 
	delete from transaction_details td 
	where td.transaction_id = old.transaction_id;
end$$
delimiter ;

## 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.

delimiter $$
create trigger update_totalqty_transaction
	after delete
	on transaction_details
	for each row
begin
	update transactions
	set total_qty = old.qty;
	where trasaction_id = old.transaction_id
end$$

## 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.

select\*from products p where p.product_id not in (select td.product_id from transaction_details td)
