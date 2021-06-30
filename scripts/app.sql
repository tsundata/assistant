#
# Dump of table apps
# ------------------------------------------------------------

CREATE TABLE `apps`
(
    `id`    int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`  varchar(16)      NOT NULL DEFAULT '',
    `type`  varchar(12)      NOT NULL DEFAULT '',
    `token` varchar(256)     NOT NULL DEFAULT '',
    `extra` varchar(2048)    NOT NULL DEFAULT '',
    `time`  timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;



#
# Dump of table credentials
# ------------------------------------------------------------

CREATE TABLE `credentials`
(
    `id`      int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`    varchar(16)      NOT NULL DEFAULT '',
    `type`    varchar(12)      NOT NULL DEFAULT '',
    `content` varchar(2048)    NOT NULL DEFAULT '',
    `time`    timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;



#
# # Dump of table messages
# ------------------------------------------------------------

CREATE TABLE `messages`
(
    `id`   int(11) unsigned NOT NULL AUTO_INCREMENT,
    `uuid` varchar(36)      NOT NULL DEFAULT '',
    `type` varchar(12)      NOT NULL DEFAULT '',
    `text` varchar(2048)    NOT NULL DEFAULT '',
    `time` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;



#
# Dump of table pages
# ------------------------------------------------------------

CREATE TABLE `pages`
(
    `id`      int(11) unsigned                       NOT NULL AUTO_INCREMENT,
    `uuid`    varchar(36) CHARACTER SET utf8mb4      NOT NULL DEFAULT '',
    `type`    varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `title`   varchar(256) CHARACTER SET utf8mb4     NOT NULL DEFAULT '',
    `content` text CHARACTER SET utf8mb4             NOT NULL,
    `time`    timestamp                              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;



#
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
    `time`       timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


#
# Dump of table todos
# ------------------------------------------------------------

CREATE TABLE `todos`
(
    `id`                INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `content`           VARCHAR(1024)    NOT NULL DEFAULT '',
    `priority`          TINYINT(4)       NOT NULL DEFAULT '0',
    `is_remind_at_time` TINYINT(4)       NOT NULL DEFAULT '0',
    `remind_at`         TIMESTAMP        NULL     DEFAULT NULL,
    `repeat_method`     VARCHAR(50)      NOT NULL,
    `repeat_rule`       VARCHAR(256)     NOT NULL DEFAULT '',
    `repeat_end_at`     TIMESTAMP        NULL     DEFAULT NULL,
    `category`          VARCHAR(50)      NOT NULL,
    `remark`            VARCHAR(1024)    NOT NULL,
    `complete`          TINYINT(4)       NOT NULL DEFAULT '0',
    `created_at`        TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


#
# Dump of table roles
# ------------------------------------------------------------

CREATE TABLE `roles`
(
    `id`          INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     INT(10)          NOT NULL,
    `profession`  VARCHAR(50)      NOT NULL,
    `exp`         INT(11)          NOT NULL DEFAULT '0',
    `level`       INT(11)          NOT NULL DEFAULT '1',
    `strength`    INT(11)          NOT NULL DEFAULT '0',
    `culture`     INT(11)          NOT NULL DEFAULT '0',
    `environment` INT(11)          NOT NULL DEFAULT '0',
    `charisma`    INT(11)          NOT NULL DEFAULT '0',
    `talent`      INT(11)          NOT NULL DEFAULT '0',
    `intellect`   INT(11)          NOT NULL DEFAULT '0',
    `time`        TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

INSERT INTO `roles` (`id`, `user_id`, `profession`, `exp`, `level`, `strength`, `culture`, `environment`, `charisma`,
                     `talent`, `intellect`, `time`)
VALUES (1, 1, 'super', 0, 1, 0, 0, 0, 0, 0, 0, '2021-06-01 00:00:00');



#
# Dump of table role_records
# ------------------------------------------------------------


CREATE TABLE `role_records`
(
    `id`          INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     INT(10)          NOT NULL,
    `profession`  VARCHAR(50)      NOT NULL,
    `exp`         INT(11)          NOT NULL DEFAULT '0',
    `level`       INT(11)          NOT NULL DEFAULT '0',
    `strength`    INT(11)          NOT NULL DEFAULT '0',
    `culture`     INT(11)          NOT NULL DEFAULT '0',
    `environment` INT(11)          NOT NULL DEFAULT '0',
    `charisma`    INT(11)          NOT NULL DEFAULT '0',
    `talent`      INT(11)          NOT NULL DEFAULT '0',
    `intellect`   INT(11)          NOT NULL DEFAULT '0',
    `time`        TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

INSERT INTO `role_records` (`id`, `user_id`, `profession`, `exp`, `level`, `strength`, `culture`, `environment`,
                            `charisma`,
                            `talent`, `intellect`, `time`)
VALUES (1, 1, 'super', 0, 1, 0, 0, 0, 0, 0, 0, '2021-06-01 00:00:00');
