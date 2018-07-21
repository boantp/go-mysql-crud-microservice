# Go API E-comm

Build API with golang for e-comm process

## Built With

* [Golang](https://godoc.org/) - open source programming language 
* [MySQL](https://godoc.org/github.com/go-sql-driver/mysql) - Golang Package mysql

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

First you need to prepare the database at you server

#  create database
```
CREATE DATABASE /*!32312 IF NOT EXISTS*/`rent_car` /*!40100 DEFAULT CHARACTER SET latin1 */;
```

# create table

* customer
* car
* car_detail
* car_type
* order
* order_detail

```
USE `rent_car`;

/*Table structure for table `car` */

DROP TABLE IF EXISTS `car`;

CREATE TABLE `car` (
  `car_id` int(11) NOT NULL AUTO_INCREMENT,
  `car_name` varchar(100) NOT NULL,
  `car_type_id` int(11) NOT NULL,
  `car_year` int(4) NOT NULL,
  `default_price` decimal(15,2) DEFAULT NULL,
  `flag_with_bbm` tinyint(1) DEFAULT NULL,
  `flag_with_driver` tinyint(1) DEFAULT NULL,
  `seat` int(11) DEFAULT NULL,
  `is_have_ac` tinyint(1) NOT NULL DEFAULT '1',
  `is_have_radio` tinyint(1) NOT NULL DEFAULT '1',
  `is_have_baggage` tinyint(1) NOT NULL DEFAULT '1',
  `car_image` varchar(150) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`car_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

/*Data for the table `car` */

insert  into `car`(`car_id`,`car_name`,`car_type_id`,`car_year`,`default_price`,`flag_with_bbm`,`flag_with_driver`,`seat`,`is_have_ac`,`is_have_radio`,`is_have_baggage`,`car_image`,`created_at`,`updated_at`) values (1,'Yarris',1,2012,100000.00,1,1,5,1,1,1,NULL,'0000-00-00 00:00:00',NULL),(2,'Jazz',2,2012,200000.00,1,1,5,1,1,1,NULL,'0000-00-00 00:00:00',NULL),(3,'Datsun',3,2012,200000.00,1,1,5,1,1,1,NULL,'0000-00-00 00:00:00',NULL),(4,'Serena',4,2011,100000.00,1,1,6,1,1,1,NULL,'0000-00-00 00:00:00',NULL);

/*Table structure for table `car_detail` */

DROP TABLE IF EXISTS `car_detail`;

CREATE TABLE `car_detail` (
  `car_detail_id` int(11) NOT NULL AUTO_INCREMENT,
  `car_id` int(11) NOT NULL,
  `nopol` varchar(10) NOT NULL,
  `nopol_color` varchar(30) NOT NULL,
  `car_color` varchar(30) NOT NULL,
  `chasis_number` varchar(100) DEFAULT NULL,
  `engine_number` varchar(100) DEFAULT NULL,
  `car_status` tinyint(1) DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`car_detail_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

/*Data for the table `car_detail` */

insert  into `car_detail`(`car_detail_id`,`car_id`,`nopol`,`nopol_color`,`car_color`,`chasis_number`,`engine_number`,`car_status`,`created_at`,`updated_at`) values (1,1,'1234','test','test','1234','1234',1,'0000-00-00 00:00:00',NULL),(2,2,'4567','test','test','4567','4567',1,'0000-00-00 00:00:00',NULL),(3,3,'1478','test','test','1478','1478',1,'0000-00-00 00:00:00',NULL),(4,4,'5555','test','test','5555','5555',1,'0000-00-00 00:00:00',NULL);

/*Table structure for table `car_type` */

DROP TABLE IF EXISTS `car_type`;

CREATE TABLE `car_type` (
  `car_type_id` int(11) NOT NULL AUTO_INCREMENT,
  `car_type_name` varchar(50) NOT NULL,
  PRIMARY KEY (`car_type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

/*Data for the table `car_type` */

insert  into `car_type`(`car_type_id`,`car_type_name`) values (1,'Toyota'),(2,'Honda'),(3,'Daihatsu'),(4,'Nissan');

/*Table structure for table `customer` */

DROP TABLE IF EXISTS `customer`;

CREATE TABLE `customer` (
  `customer_id` bigint(10) NOT NULL AUTO_INCREMENT,
  `customer_code` varchar(50) NOT NULL,
  `customer_name` varchar(100) NOT NULL,
  `gender` char(1) NOT NULL,
  `email` varchar(100) NOT NULL,
  `phone_number` varchar(30) NOT NULL,
  `date_of_birth` varchar(10) NOT NULL,
  `customer_photo` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `customer` */

insert  into `customer`(`customer_id`,`customer_code`,`customer_name`,`gender`,`email`,`phone_number`,`date_of_birth`,`customer_photo`,`created_at`,`updated_at`) values (1,'qwewdas234312432423','Boan','1','boantuapasaribu@gmail.com','25234535345345','12-12-2012',NULL,'0000-00-00 00:00:00',NULL);

/*Table structure for table `order` */

DROP TABLE IF EXISTS `order`;

CREATE TABLE `order` (
  `order_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_code` varchar(50) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `passenger_name` varchar(100) NOT NULL,
  `passenger_phone_number` varchar(30) NOT NULL,
  `pickup_address` text NOT NULL,
  `destination_address` text NOT NULL,
  `order_quantity` int(11) NOT NULL DEFAULT '1',
  `total_fare_amount` decimal(15,2) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `order` */

insert  into `order`(`order_id`,`order_code`,`customer_id`,`passenger_name`,`passenger_phone_number`,`pickup_address`,`destination_address`,`order_quantity`,`total_fare_amount`,`created_at`,`updated_at`) values (1,'123456789',1,'Boan','02856944555','testeatasdasdasdasdasdasdasd','testeatasdasdasdasdasdasdasd',1,200000.00,'0000-00-00 00:00:00',NULL);

/*Table structure for table `order_detail` */

DROP TABLE IF EXISTS `order_detail`;

CREATE TABLE `order_detail` (
  `order_detail_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) NOT NULL,
  `estimated_fare` decimal(15,2) DEFAULT NULL,
  `car_id` int(11) NOT NULL,
  `order_status` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`order_detail_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `order_detail` */

insert  into `order_detail`(`order_detail_id`,`order_id`,`estimated_fare`,`car_id`,`order_status`,`created_at`,`updated_at`) values (1,1,200000.00,2,1,'0000-00-00 00:00:00',NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

```

# insert a record
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (4, 'Jasmine', 5, '983 Star Ave., Brooklyn, NY, 00912 ', 55700.00, '1997-12-13' ), (5, 'Orranda', 9, '745 Hammer Lane, Hammerfield, Texas, 75839', 65350.00 , '1992-12-13');
```

### Installing

After we already set the database now we prepare to running the server and try following routes to test the code

running the server : go to the project folder from your cmd and running

```
go run main.go
```

And now you go to your localhost with your browser

```
http://localhost:8080
```

End with an example of getting some data out of the system or using it for a little demo

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```
