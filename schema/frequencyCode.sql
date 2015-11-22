CREATE TABLE `frequencyCode` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `unique` varchar(255) NOT NULL,
  `object` varchar(255) NOT NULL,
  `timestamp` int(10) unsigned NOT NULL,
  `code` varchar(255) NOT NULL,
  `message` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniqueIdx` (`unique`),
  KEY `timeIdx` (`timestamp`),
  KEY `codeIdx` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;