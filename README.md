# sample-db-migration-tool
Sample db migration tool written in golang.

DATABASE SCHEMA:

  - customerNumber int(11) NOT NULL,
  - customerName varchar(50) NOT NULL,
  - contactLastName varchar(50) NOT NULL,
  - contactFirstName varchar(50) NOT NULL,
  - phone varchar(50) NOT NULL,
  - addressLine1 varchar(50) NOT NULL,
  - addressLine2 varchar(50) DEFAULT NULL,
  - city varchar(50) NOT NULL,
  - state varchar(50) DEFAULT NULL,
  - postalCode varchar(15) DEFAULT NULL,
  - country varchar(50) NOT NULL,
  - salesRepEmployeeNumber int(11) DEFAULT NULL,
  - creditLimit decimal(10,2) DEFAULT NULL,

Engine: InnoDB