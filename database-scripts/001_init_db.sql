CREATE DATABASE IF NOT EXISTS `cloudrover`;

USE `cloudrover`;

DROP TABLE IF EXISTS `alerts`;

CREATE TABLE `alerts` (
  `id` bigint(8) NOT NULL,
  `date_happened` int(11) DEFAULT NULL,
  `device_name` varchar(50) DEFAULT NULL,
  `alert_type` varchar(20) DEFAULT NULL,
  `title` varchar(1000) DEFAULT NULL,
  `url` varchar(500) DEFAULT NULL,
  `tags` json DEFAULT NULL,
  `host` varchar(45) DEFAULT NULL,
  `payload` TEXT DEFAULT NULL,
  `priority` varchar(10) DEFAULT NULL,
  `text` TEXT DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `tag_config`;

CREATE TABLE `tag_config` (
  `id` int NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(45) NOT NULL,
  `alias` varchar(45) NOT NULL,
  `no_of_tag_parts` int NOT NULL DEFAULT '1',
  `description` varchar(45) DEFAULT NULL,
  `required` tinyint DEFAULT '0',
  `aggregate` tinyint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `cloudrover`.`tag_config` (`tag_name`, `alias`, `no_of_tag_parts`, `description`, `required`, `aggregate`)
VALUES ('organization', 'organization', '2', 'Name of the Organization', true, true);

INSERT INTO `cloudrover`.`tag_config` (`tag_name`, `alias`, `no_of_tag_parts`, `description`, `required`, `aggregate`)
VALUES ('department', 'department', '2', 'Name of the Department', true, true);

INSERT INTO `cloudrover`.`tag_config` (`tag_name`, `alias`, `no_of_tag_parts`, `description`, `required`, `aggregate`)
VALUES ('product', 'product', '2', 'Name of the Product', true, true);

INSERT INTO `cloudrover`.`tag_config` (`tag_name`, `alias`, `no_of_tag_parts`, `description`, `required`, `aggregate`)
VALUES ('team', 'team', '2', 'Name of the team', true, true);

INSERT INTO `cloudrover`.`tag_config` (`tag_name`, `alias`, `no_of_tag_parts`, `description`, `required`, `aggregate`)
VALUES ('service', 'servuce', '2', 'Name of the Service', true, true);

INSERT INTO `cloudrover`.`tag_config` (`tag_name`, `alias`, `no_of_tag_parts`, `description`, `required`, `aggregate`)
VALUES ('cloudprovider', 'cloudprovider', '1', 'Cloud Provider', true, true);

INSERT INTO `cloudrover`.`tag_config` (`tag_name`, `alias`, `no_of_tag_parts`, `description`, `required`, `aggregate`)
VALUES ('cloudresource', 'cloudresource', '1', 'Name of the Cloud Resource', true, true);

DROP TABLE IF EXISTS `past_reports`;

CREATE TABLE `past_reports` (
  `id` int NOT NULL AUTO_INCREMENT,
  `date_happened` int(11) DEFAULT NULL,
  `name` varchar(45) NOT NULL,
  `report` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;