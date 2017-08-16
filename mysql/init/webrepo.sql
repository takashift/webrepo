CREATE TABLE userinfo(id int unique not null auto_increment primary key, OAuth utf8mb4, email varchar(255) unique, password varchar(255), name utf8mb4, safe_search int not null);
