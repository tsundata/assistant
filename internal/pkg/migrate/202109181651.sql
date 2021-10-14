create table if not exists `apps`
(
    `id`         BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       varchar(16)         NOT NULL DEFAULT '',
    `type`       varchar(12)         NOT NULL DEFAULT '',
    `token`      varchar(256)        NOT NULL DEFAULT '',
    `extra`      varchar(2048)       NOT NULL DEFAULT '',
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    `updated_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `credentials`
(
    `id`         BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       varchar(16)         NOT NULL DEFAULT '',
    `type`       varchar(12)         NOT NULL DEFAULT '',
    `content`    varchar(2048)       NOT NULL DEFAULT '',
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    `updated_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `messages`
(
    `id`         BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `uuid`       varchar(36)         NOT NULL DEFAULT '',
    `type`       varchar(12)         NOT NULL DEFAULT '',
    `channel`    varchar(20)         NOT NULL DEFAULT '',
    `text`       varchar(2048)       NOT NULL DEFAULT '',
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    `updated_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `pages`
(
    `id`         BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `uuid`       varchar(36)         NOT NULL DEFAULT '',
    `type`       varchar(10)         NOT NULL DEFAULT '',
    `title`      varchar(256)        NOT NULL DEFAULT '',
    `content`    text                NOT NULL,
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    `updated_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `triggers`
(
    `id`         BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `type`       varchar(16)         NOT NULL DEFAULT '',
    `kind`       varchar(16)         NOT NULL DEFAULT '',
    `flag`       varchar(128)        NOT NULL DEFAULT '',
    `secret`     varchar(128)        NOT NULL DEFAULT '',
    `when`       varchar(128)        NOT NULL DEFAULT '',
    `message_id` INT(10)             NOT NULL,
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `todos`
(
    `id`                BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `content`           VARCHAR(1024)       NOT NULL DEFAULT '',
    `priority`          TINYINT(4)          NOT NULL DEFAULT '0',
    `is_remind_at_time` TINYINT(4)          NOT NULL DEFAULT '0',
    `remind_at`         INT(10)             NOT NULL DEFAULT '0',
    `repeat_method`     VARCHAR(50)         NOT NULL,
    `repeat_rule`       VARCHAR(256)        NOT NULL DEFAULT '',
    `repeat_end_at`     INT(10)             NOT NULL DEFAULT '0',
    `category`          VARCHAR(50)         NOT NULL,
    `remark`            VARCHAR(1024)       NOT NULL,
    `complete`          TINYINT(4)          NOT NULL DEFAULT '0',
    `created_at`        INT(10)             NOT NULL DEFAULT '0',
    `updated_at`        INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `roles`
(
    `id`          BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     INT(10)             NOT NULL,
    `profession`  VARCHAR(50)         NOT NULL,
    `exp`         INT(10)             NOT NULL DEFAULT '0',
    `level`       INT(10)             NOT NULL DEFAULT '1',
    `strength`    INT(10)             NOT NULL DEFAULT '0',
    `culture`     INT(10)             NOT NULL DEFAULT '0',
    `environment` INT(10)             NOT NULL DEFAULT '0',
    `charisma`    INT(10)             NOT NULL DEFAULT '0',
    `talent`      INT(10)             NOT NULL DEFAULT '0',
    `intellect`   INT(10)             NOT NULL DEFAULT '0',
    `created_at`  INT(10)             NOT NULL DEFAULT '0',
    `updated_at`  INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

INSERT INTO `roles` (`user_id`, `profession`, `exp`, `level`, `strength`, `culture`, `environment`, `charisma`,
                     `talent`, `intellect`)
VALUES (1, 'super', 0, 1, 0, 0, 0, 0, 0, 0);


create table if not exists `role_records`
(
    `id`          BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     INT(10)             NOT NULL,
    `profession`  VARCHAR(50)         NOT NULL,
    `exp`         INT(10)             NOT NULL DEFAULT '0',
    `level`       INT(10)             NOT NULL DEFAULT '0',
    `strength`    INT(10)             NOT NULL DEFAULT '0',
    `culture`     INT(10)             NOT NULL DEFAULT '0',
    `environment` INT(10)             NOT NULL DEFAULT '0',
    `charisma`    INT(10)             NOT NULL DEFAULT '0',
    `talent`      INT(10)             NOT NULL DEFAULT '0',
    `intellect`   INT(10)             NOT NULL DEFAULT '0',
    `created_at`  INT(10)             NOT NULL DEFAULT '0',
    `updated_at`  INT(10)             NOT NULL DEFAULT '0',
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
    `id`         BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `username`   VARCHAR(50)         NOT NULL,
    `password`   VARCHAR(256)        NOT NULL DEFAULT '',
    `nickname`   VARCHAR(50)         NOT NULL DEFAULT '',
    `mobile`     VARCHAR(50)         NOT NULL DEFAULT '',
    `remark`     VARCHAR(50)         NOT NULL DEFAULT '',
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    `updated_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;


INSERT INTO `users` (`username`, `password`, `nickname`, `mobile`, `remark`, `created_at`, `updated_at`)
VALUES ('admin', '$2a$10$UbySCK7RHJwyD7DYMjIyTOIfvL8t2KEmz.3jVFIwGlOvzV2P373uu', 'me', '', '', '1625068800',
        '1625068800');


create table if not exists `objectives`
(
    `id`         BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       varchar(50)         NOT NULL DEFAULT '',
    `tag_id`     INT(10)             NOT NULL,
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    `updated_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `key_results`
(
    `id`           BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `objective_id` INT(10)             NOT NULL,
    `name`         varchar(50)         NOT NULL DEFAULT '',
    `tag_id`       INT(10)             NOT NULL,
    `complete`     TINYINT(4)          NOT NULL DEFAULT '0',
    `created_at`   INT(10)             NOT NULL DEFAULT '0',
    `updated_at`   INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `tags`
(
    `id`         BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       varchar(50)         NOT NULL DEFAULT '',
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    `updated_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;