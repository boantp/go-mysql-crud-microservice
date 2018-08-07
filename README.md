# Go MYSQL CRUD

Build Simple CRUD with go + mysql

## Built With

* [Golang](https://godoc.org/) - open source programming language 
* [MySQL](https://godoc.org/github.com/go-sql-driver/mysql) - Golang Package mysql

### Prerequisites

First you need to prepare the database at you server and change the connection in folder config/db.go

#  create database
```
CREATE DATABASE /*!32312 IF NOT EXISTS*/`rent_car` /*!40100 DEFAULT CHARACTER SET latin1 */;
```

# create table

* customer
* car
* order

```
USE `rent_car`;

/*Table structure for table `car` */

DROP TABLE IF EXISTS `car`;

CREATE TABLE `car` (
  `car_id` int(11) NOT NULL AUTO_INCREMENT,
  `car_name` varchar(100) NOT NULL,
  `car_year` int(4) NOT NULL,
  `default_price` decimal(15,2) DEFAULT NULL,
  `car_status` int(1) DEFAULT NULL,
  PRIMARY KEY (`car_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

/*Data for the table `car` */

insert  into `car`(`car_id`,`car_name`,`car_year`,`default_price`,`car_status`) values (1,'Yarris',2012,100000.00,1),(2,'Jazz',2012,200000.00,1),(3,'Datsun',2012,200000.00,1),(4,'Serena',2011,100000.00,1);

/*Table structure for table `customer` */

DROP TABLE IF EXISTS `customer`;

CREATE TABLE `customer` (
  `customer_id` bigint(10) NOT NULL AUTO_INCREMENT,
  `customer_name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `phone_number` varchar(30) NOT NULL,
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `customer` */

insert  into `customer`(`customer_id`,`customer_name`,`email`,`phone_number`) values (1,'Boan','boantuapasaribu@gmail.com','25234535345345');

/*Table structure for table `order` */

DROP TABLE IF EXISTS `order`;

CREATE TABLE `order` (
  `order_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_code` varchar(50) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `passenger_name` varchar(100) NOT NULL,
  `passenger_phone_number` varchar(30) NOT NULL,
  `car_id` int(11) DEFAULT NULL,
  `pickup_address` text NOT NULL,
  `destination_address` text NOT NULL,
  `order_quantity` int(11) NOT NULL DEFAULT '1',
  `total_fare_amount` decimal(15,2) NOT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `order` */

insert  into `order`(`order_id`,`order_code`,`customer_id`,`passenger_name`,`passenger_phone_number`,`car_id`,`pickup_address`,`destination_address`,`order_quantity`,`total_fare_amount`) values (1,'123456789',1,'Boan','02856944555',1,'testeatasdasdasdasdasdasdasd','testeatasdasdasdasdasdasdasd',1,200000.00);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

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

```
http://localhost:8080/book
```

```
http://localhost:8080/order
```
