CREATE DATABASE IF NOT EXISTS base
CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE base;


CREATE TABLE IF NOT EXISTS `link` (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` VARCHAR (100) NOT NULL COMMENT "连接唯一id",
  `context` TEXT COMMENT "内容",
  PRIMARY KEY (id) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8
COMMENT="链接分享表";

begin;
-- just for test case
insert into link (uid, content) values ('11', 'ge'),('12', 'all'),('33', 'great');

commit;
