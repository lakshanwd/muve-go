create table tbl_passenger (
    passenger_id int auto_increment,
    name varchar(100) not null,
    phone varchar(20) not null,
    constraint primary key(passenger_id)
) engine=innodb;

create table tbl_passenger_payment_option (
    payment_option_id int auto_increment,
    passenger_id int not null,
    payment_type enum('cash','visa','mastercard') default 'cash',
    card_no varchar(12) default null,
    csv varchar(3) default null,
    constraint foreign key(passenger_id) references tbl_passenger(passenger_id) on update cascade on delete restrict,
    constraint primary key(payment_option_id)
) engine=innodb;

create table tbl_driver (
    driver_id int auto_increment,
    driver_name varchar(100) not null,
    constraint primary key(driver_id)
) engine=innodb;

create table tbl_driver_vehicle (
    vehicle_id int auto_increment,
    driver_id int not null,
    vehicle_no varchar(15) not null,
    constraint foreign key(driver_id) references tbl_driver(driver_id) on update cascade on delete restrict,
    constraint primary key(vehicle_id)
) engine=innodb;

create table tbl_booking (
    booking_id int auto_increment,
    vehicle_id int not null,
    pick_up_address varchar(255) not null,
    drop_address varchar(255) not null,
    total_distance decimal(6,3),
    created_on datetime not null,
    constraint foreign key(vehicle_id) references tbl_driver_vehicle(vehicle_id) on update cascade on delete restrict,
    constraint primary key(booking_id)
) engine=innodb;

create table tbl_fare (
    fare_id int auto_increment,
    booking_id int not null,
    payment_option_id int not null,
    bill_amount decimal(10,3) not null,
    constraint foreign key(booking_id) references tbl_booking(booking_id) on update cascade on delete restrict,
    constraint foreign key(payment_option_id) references tbl_passenger_payment_option(payment_option_id) on update cascade on delete restrict,
    constraint primary key(fare_id)
) engine=innodb;