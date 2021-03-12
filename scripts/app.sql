# Dump of table apps
# ------------------------------------------------------------

CREATE TABLE `apps`
(
    `id`    int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`  varchar(16)      NOT NULL DEFAULT '',
    `type`  varchar(12)      NOT NULL DEFAULT '',
    `token` varchar(256)     NOT NULL DEFAULT '',
    `extra` varchar(2048)    NOT NULL DEFAULT '',
    `time`  timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;



# Dump of table credentials
# ------------------------------------------------------------

CREATE TABLE `credentials`
(
    `id`      int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`    varchar(16)      NOT NULL DEFAULT '',
    `type`    varchar(12)      NOT NULL DEFAULT '',
    `content` varchar(2048)    NOT NULL DEFAULT '',
    `time`    timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;



# Dump of table messages
# ------------------------------------------------------------

CREATE TABLE `messages`
(
    `id`   int(11) unsigned NOT NULL AUTO_INCREMENT,
    `uuid` varchar(36)      NOT NULL DEFAULT '',
    `type` varchar(12)      NOT NULL DEFAULT '',
    `text` varchar(2048)    NOT NULL DEFAULT '',
    `time` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;



# Dump of table pages
# ------------------------------------------------------------

CREATE TABLE `pages`
(
    `id`      int(11) unsigned                   NOT NULL AUTO_INCREMENT,
    `uuid`    varchar(36) CHARACTER SET utf8mb4  NOT NULL DEFAULT '',
    `title`   varchar(256) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
    `content` text CHARACTER SET utf8mb4         NOT NULL,
    `time`    timestamp                          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;



# Dump of table triggers
# ------------------------------------------------------------

CREATE TABLE `triggers`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `type`       varchar(16)      NOT NULL DEFAULT '',
    `kind`       varchar(16)      NOT NULL DEFAULT '',
    `flag`       varchar(128)     NOT NULL DEFAULT '',
    `secret`     varchar(128)     NOT NULL DEFAULT '',
    `when`       varchar(128)     NOT NULL DEFAULT '',
    `message_id` int(11)          NOT NULL,
    `time`       timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `flag` (`flag`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
