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
