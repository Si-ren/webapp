create database if not exists  cmdb default  charset utf8mb4;
create table if not exists  user (
    id bigint primary  key  auto_increment,
    name varchar(32) not null  default '',
    password varchar(32) not null  default '',
)engine=innodb default charset utf8mb4;

insert into user(name,password) values ("siri",md5("lsl"));