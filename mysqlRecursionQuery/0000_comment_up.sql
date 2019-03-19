drop table if exists  `comment`;
create table `comment` (
 `id` int(11) not null auto_increment,
 `level` int(11) default '0',
 `nick_name` varchar(255) default '0',
 `parent_Id` int(11) default '0',
 `status` int(11) default '0',
 primary key  (`id`)
) engine =InnoDB auto_increment=65 default charset = utf8mb4;