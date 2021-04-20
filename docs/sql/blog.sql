create database
if not exists blog_service default character set utf8mb4 default collate utf8mb4_general_ci;

create table `blog_tag` (
    `id` int(10) unsigned not null auto_increment,
    `name` varchar(100) default '' comment 'Tag Name',
    `created_on` int(10) unsigned default '0' comment 'Created On',
    `created_by` varchar(100) default '' comment 'Creator',
    `modified_on` int(10) unsigned default '0' comment 'Modified On',
    `modified_by` varchar(100) default '' comment 'Modifier',
    `deleted_on` int(10) unsigned default '0' comment 'Deleted On',
    `is_del` tinyint(3) unsigned default '0' comment 'Is Deleted',
    `state` tinyint(3) unsigned default '1' comment 'State 0 as forbidden, 1 as allowed',
    primary key (`id`)
) engine=InnoDB default charset=utf8mb4 comment 'Tag Management';

create table `blog_article` (
    `id` int(10) unsigned not null auto_increment,
    `title` varchar(100) default '' comment 'Article Title',
    `desc` varchar(255) default '' comment 'Article Description',
    `cover_image_url` varchar(255) default '' comment 'Cover Image URL',
    `content` longtext comment 'Article Content',
    `created_on` int(10) unsigned default '0' comment 'Created On',
    `created_by` varchar(100) default '' comment 'Creator',
    `modified_on` int(10) unsigned default '0' comment 'Modified On',
    `modified_by` varchar(100) default '' comment 'Modifier',
    `deleted_on` int(10) unsigned default '0' comment 'Deleted On',
    `is_del` tinyint(3) unsigned default '0' comment 'Is Deleted',
    `state` tinyint(3) unsigned default '1' comment 'State 0 as forbidden, 1 as allowed',
    primary key (`id`)
) engine=InnoDB default charset=utf8mb4 comment 'Article Management';

create table `blog_article_tag` (
    `id` int(10) unsigned not null auto_increment,
    `article_id` int(11) not null comment 'Article ID',
    `tag_id` int(10) unsigned not null default '0' comment 'Tag ID',
    `created_on` int(10) unsigned default '0' comment 'Created On',
    `created_by` varchar(100) default '' comment 'Creator',
    `modified_on` int(10) unsigned default '0' comment 'Modified On',
    `modified_by` varchar(100) default '' comment 'Modifier',
    `deleted_on` int(10) unsigned default '0' comment 'Deleted On',
    `is_del` tinyint(3) unsigned default '0' comment 'Is Deleted',
    primary key (`id`)
) engine=InnoDB default charset=utf8mb4 comment 'Article-Tag Association';

create table `blog_auth` (
    `id` int(10) unsigned not null auto_increment,
    `app_key` varchar(20) default '' comment 'Key',
    `app_secret` varchar(50) default '' comment 'Secret',
    `created_on` int(10) unsigned default '0' comment 'Created On',
    `created_by` varchar(100) default '' comment 'Creator',
    `modified_on` int(10) unsigned default '0' comment 'Modified On',
    `modified_by` varchar(100) default '' comment 'Modifier',
    `deleted_on` int(10) unsigned default '0' comment 'Deleted On',
    `is_del` tinyint(3) unsigned default '0' comment 'Is Deleted',
    primary key (`id`) using BTREE
) engine=InnoDB default charset=utf8mb4 comment 'authentication management';

insert into `blog_service`.`blog_auth` (
    id, app_key, app_secret, created_on, created_by, modified_on, modified_by, deleted_on, is_del
) VALUES (1, 'lachimere', 'blogservice', 0, 'lachimere', 0, '', 0, 0);
