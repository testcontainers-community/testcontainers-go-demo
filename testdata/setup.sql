CREATE TABLE students (
	id bigint NOT NULL AUTO_INCREMENT,
    fname varchar(50) not null,
    lname varchar(50) not null,
    date_of_birth datetime not null,
    email varchar(50) not null,
    address varchar(50) not null,
    gender varchar(50) not null,
    PRIMARY KEY (id)
);

insert into students (id, fname, lname, date_of_birth, email, gender, address) values (1, 'Caddric', 'Likely', '2000-07-06 02:43:37', 'clikely0@mail.com', 'Male', '9173 Boyd Street');
insert into students (id, fname, lname, date_of_birth, email, gender, address) values (2, 'Jerad', 'Ciccotti', '1993-02-11 15:59:56', 'jciccotti1@mail.com', 'Male', '34 Declaration Drive');
insert into students (id, fname, lname, date_of_birth, email, gender, address) values (3, 'Hillier', 'Caslett', '1992-09-04 13:38:46', 'hcaslett2@mail.com', 'Male', '36 Duke Trail');
insert into students (id, fname, lname, date_of_birth, email, gender, address) values (4, 'Bertine', 'Roddan', '1991-02-18 09:10:05', 'broddan3@mail.com', 'Female', '2896 Kropf Road');
insert into students (id, fname, lname, date_of_birth, email, gender, address) values (5, 'Theda', 'Brockton', '1991-10-29 09:08:48', 'tbrockton4@mail.com', 'Female', '93 Hermina Plaza');
insert into students (id, fname, lname, date_of_birth, email, gender, address) values (6, 'Leon', 'Ashling', '1994-08-14 23:51:42', 'lashling5@mail.com', 'Male', '39 Kipling Pass');
insert into students (id, fname, lname, date_of_birth, email, gender, address) values (7, 'Aldo', 'Pettitt', '1994-08-14 22:03:40', 'apettitt6@mail.com', 'Male', '38 Dryden Road');
insert into students (id, fname, lname, date_of_birth, email, gender, address) values (8, 'Filmore', 'Cordingly', '1999-11-20 02:35:48', 'fcordingly7@mail.com', 'Male', '34 Pawling Park');
insert into students (id, fname, lname, date_of_birth, email, gender, address) values (9, 'Katalin', 'MacCroary', '1994-11-08 11:59:19', 'kmaccroary8@mail.com', 'Female', '2540 Maryland Parkway');
insert into students (id, fname, lname, date_of_birth, email, gender, address) values (10, 'Franky', 'Puddan', '1995-04-23 17:07:29', 'fpuddan9@mail.com', 'Female', '3214 Washington Road');