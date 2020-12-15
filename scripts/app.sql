/*!40101 SET @OLD_CHARACTER_SET_CLIENT = @@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS = @@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION = @@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0 */;
/*!40101 SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES = @@SQL_NOTES, SQL_NOTES = 0 */;


# Dump of table events
# ------------------------------------------------------------

DROP TABLE IF EXISTS `events`;

CREATE TABLE `events`
(
    `id`    int(11) unsigned                  NOT NULL AUTO_INCREMENT,
    `uuid`  varchar(64) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
    `type`  varchar(20) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
    `time`  timestamp                         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `event` text CHARACTER SET utf8mb4        NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


# Dump of table pages
# ------------------------------------------------------------

DROP TABLE IF EXISTS `pages`;

CREATE TABLE `pages`
(
    `id`      int(11) unsigned                   NOT NULL AUTO_INCREMENT,
    `uuid`    varchar(64) CHARACTER SET utf8mb4  NOT NULL DEFAULT '',
    `title`   varchar(256) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
    `content` text CHARACTER SET utf8mb4         NOT NULL,
    `time`    timestamp                          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


# Dump of table subscribes
# ------------------------------------------------------------

DROP TABLE IF EXISTS `subscribes`;

CREATE TABLE `pages`
(
    `id`      int(11) unsigned                   NOT NULL AUTO_INCREMENT,
    `uuid`    varchar(64) CHARACTER SET utf8mb4  NOT NULL DEFAULT '',
    `title`   varchar(256) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
    `content` text CHARACTER SET utf8mb4         NOT NULL,
    `time`    timestamp                          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


/*!40111 SET SQL_NOTES = @OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE = @OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS = @OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT = @OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS = @OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION = @OLD_COLLATION_CONNECTION */;
