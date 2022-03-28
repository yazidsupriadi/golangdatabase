CREATE TABLE customer
(
    id varchar(100) not null,
    name varchar(100) not null,
    email varchar(100) not null,
    balance integer not null DEFAULT 0,
    rating float not null DEFAULT 0.0,
    birth_date DATE,
    married boolean DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY(id)
)ENGINE = InnoDb;

insert into customer(id,name,email,balance,rating,birth_date,married) 
values("1","Oncom","Oncom@gmail.com",100000,5,"1999-06-14",false);