CREATE TABLE usersData (
                           id serial not null unique,
                           firstName varchar(255) not null,
                           lastName varchar(255) not null,
                           email varchar(255) not null unique,
                           phoneNumber varchar(15) not null unique,
                           hashPassword varchar(255) not null,
                           dateOfBirth varchar(30),
                           role integer not null ,
                           imgUrl varchar(255)
);