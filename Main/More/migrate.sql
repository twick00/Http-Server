CREATE DATABASE IF NOT EXISTS test;
USE test;
CREATE TABLE IF NOT EXISTS games ( 
	title varchar(45),
    year int(4),
    genre varchar(45),
    barcode int(18),
	id int(20)
    );
