CREATE TABLE IF NOT EXISTS `settings` (
	`id` int(5) NOT NULL AUTO_INCREMENT,
	`key` varchar(150) NOT NULL,
	`value` varchar(20) NOT NULL,
	`type` varchar(150) NOT NULL DEFAULT 'core',
	`created_at` datetime NOT NULL,
	`created_by` int(10) NOT NULL,
	`updated_at` datetime,
	`updated_by` int(10),
	PRIMARY KEY (`id`),
	UNIQUE KEY `ix_settings_key` (`key`),
	KEY `type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS `users` (
	`id` int(10) NOT NULL AUTO_INCREMENT,
	`name` varchar(150) NOT NULL,
	`slug` varchar(150) NOT NULL,
	`password` varchar(60) NOT NULL,
	`email` varchar(254) NOT NULL,
	`image` text,
	`cover` text,
	`bio` varchar(200),
	`website` text,
	`location` text,
	`accessibility` text,
	`status` varchar(150) NOT NULL DEFAULT 'active',
	`language` varchar(6) NOT NULL DEFAULT 'en_US',
	`last_login` datetime,
	`created_at` datetime NOT NULL,
	`created_by` int(10) NOT NULL,
	`updated_at` datetime,
	`updated_by` int(10),
	PRIMARY KEY (`id`),
	UNIQUE KEY `ix_users_email` (`email`),
	KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS `tokens` (
	`id` INT(10) NOT NULL AUTO_INCREMENT,
	`value` varchar(40) NOT NULL,
	`user_id` int(10) NOT NULL,
	`created_at` datetime,
	`expired_at` datetime,
	PRIMARY KEY (`id`),
	UNIQUE KEY `ix_token_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;
