create table if not exists `apps`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(16)      NOT NULL DEFAULT '',
    `type`       varchar(12)      NOT NULL DEFAULT '',
    `token`      varchar(256)     NOT NULL DEFAULT '',
    `extra`      varchar(2048)    NOT NULL DEFAULT '',
    `created_at` int(11)          NOT NULL DEFAULT '0',
    `updated_at` int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `credentials`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(16)      NOT NULL DEFAULT '',
    `type`       varchar(12)      NOT NULL DEFAULT '',
    `content`    varchar(2048)    NOT NULL DEFAULT '',
    `created_at` int(11)          NOT NULL DEFAULT '0',
    `updated_at` int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `messages`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `uuid`       varchar(36)      NOT NULL DEFAULT '',
    `type`       varchar(12)      NOT NULL DEFAULT '',
    `channel`    varchar(20)      NOT NULL DEFAULT '',
    `text`       varchar(2048)    NOT NULL DEFAULT '',
    `created_at` int(11)          NOT NULL DEFAULT '0',
    `updated_at` int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `pages`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `uuid`       varchar(36)      NOT NULL DEFAULT '',
    `type`       varchar(10)      NOT NULL DEFAULT '',
    `title`      varchar(256)     NOT NULL DEFAULT '',
    `content`    text             NOT NULL,
    `created_at` int(11)          NOT NULL DEFAULT '0',
    `updated_at` int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `triggers`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `type`       varchar(16)      NOT NULL DEFAULT '',
    `kind`       varchar(16)      NOT NULL DEFAULT '',
    `flag`       varchar(128)     NOT NULL DEFAULT '',
    `secret`     varchar(128)     NOT NULL DEFAULT '',
    `when`       varchar(128)     NOT NULL DEFAULT '',
    `message_id` int(11)          NOT NULL,
    `created_at` int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `todos`
(
    `id`                INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `content`           VARCHAR(1024)    NOT NULL DEFAULT '',
    `priority`          TINYINT(4)       NOT NULL DEFAULT '0',
    `is_remind_at_time` TINYINT(4)       NOT NULL DEFAULT '0',
    `remind_at`         int(11)          NOT NULL DEFAULT '0',
    `repeat_method`     VARCHAR(50)      NOT NULL,
    `repeat_rule`       VARCHAR(256)     NOT NULL DEFAULT '',
    `repeat_end_at`     int(11)          NOT NULL DEFAULT '0',
    `category`          VARCHAR(50)      NOT NULL,
    `remark`            VARCHAR(1024)    NOT NULL,
    `complete`          TINYINT(4)       NOT NULL DEFAULT '0',
    `created_at`        int(11)          NOT NULL DEFAULT '0',
    `updated_at`        int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `roles`
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
    `created_at`  int(11)          NOT NULL DEFAULT '0',
    `updated_at`  int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

INSERT INTO `roles` (`user_id`, `profession`, `exp`, `level`, `strength`, `culture`, `environment`, `charisma`,
                     `talent`, `intellect`)
VALUES (1, 'super', 0, 1, 0, 0, 0, 0, 0, 0);


create table if not exists `role_records`
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
    `created_at`  int(11)          NOT NULL DEFAULT '0',
    `updated_at`  int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;


INSERT INTO `role_records` (`user_id`, `profession`, `exp`, `level`, `strength`, `culture`, `environment`,
                            `charisma`,
                            `talent`, `intellect`)
VALUES (1, 'super', 0, 1, 0, 0, 0, 0, 0, 0);


create table if not exists `users`
(
    `id`         INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `username`   VARCHAR(50)      NOT NULL,
    `password`   VARCHAR(50)      NOT NULL DEFAULT '',
    `nickname`   VARCHAR(50)      NOT NULL DEFAULT '',
    `mobile`     VARCHAR(50)      NOT NULL DEFAULT '',
    `remark`     VARCHAR(50)      NOT NULL DEFAULT '',
    `created_at` int(11)          NOT NULL DEFAULT '0',
    `updated_at` int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;


INSERT INTO `users` (`username`, `nickname`, `mobile`, `remark`, `created_at`, `updated_at`)
VALUES ('admin', 'me', '', '', '1625068800', '1625068800');


create table if not exists `objectives`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(50)      NOT NULL DEFAULT '',
    `tag_id`     int(11)          NOT NULL,
    `created_at` int(11)          NOT NULL DEFAULT '0',
    `updated_at` int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `key_results`
(
    `id`           int(11) unsigned NOT NULL AUTO_INCREMENT,
    `objective_id` int(11)          NOT NULL,
    `name`         varchar(50)      NOT NULL DEFAULT '',
    `tag_id`       int(11)          NOT NULL,
    `complete`     TINYINT(4)       NOT NULL DEFAULT '0',
    `created_at`   int(11)          NOT NULL DEFAULT '0',
    `updated_at`   int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `tags`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(50)      NOT NULL DEFAULT '',
    `created_at` int(11)          NOT NULL DEFAULT '0',
    `updated_at` int(11)          NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;