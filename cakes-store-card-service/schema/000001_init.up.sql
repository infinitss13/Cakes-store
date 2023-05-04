CREATE TABLE usersCart (
    id serial not null unique,
    userID integer not null unique,
    cartItems JSON[] NOT NULL
);