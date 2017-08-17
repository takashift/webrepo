CREATE TABLE userinfo(id int unique not null auto_increment primary key, OAuth varchar(255), email varchar(255) unique, password varchar(255), name varchar(255), safe_search int not null);
CREATE TABLE page_status(id int unique not null auto_increment primary key, title varchar(255));
