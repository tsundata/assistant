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
    `user_id`    BIGINT(19)          NOT NULL DEFAULT '0',
    `message_id` BIGINT(19)          NOT NULL DEFAULT '0',
    `status`     TINYINT(3)          NOT NULL DEFAULT '0',
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    `updated_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    INDEX `user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


create table if not exists `todos`
(
    `id`                BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`           BIGINT(10)          NOT NULL,
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
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `objectives`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`       bigint unsigned NOT NULL,
    `sequence`      int             NOT NULL,
    `title`         varchar(50)     NOT NULL,
    `memo`          varchar(300)    NOT NULL,
    `motive`        varchar(300)    NOT NULL,
    `feasibility`   varchar(300)    NOT NULL,
    `is_plan`       tinyint         NOT NULL,
    `plan_start`    int             NOT NULL,
    `plan_end`      int             NOT NULL,
    `total_value`   int             NOT NULL,
    `current_value` int             NOT NULL,
    `status`        tinyint         NOT NULL DEFAULT '0',
    `created_at`    int             NOT NULL DEFAULT '0',
    `updated_at`    int             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`, `sequence`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `key_results`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`       bigint unsigned NOT NULL,
    `sequence`      int             NOT NULL,
    `objective_id`  bigint          NOT NULL,
    `title`         varchar(50)     NOT NULL,
    `memo`          varchar(300)    NOT NULL,
    `initial_value` int             NOT NULL DEFAULT '0',
    `target_value`  int             NOT NULL,
    `current_value` int             NOT NULL,
    `value_mode`    varchar(20)     NOT NULL,
    `status`        tinyint         NOT NULL DEFAULT '0',
    `created_at`    int             NOT NULL DEFAULT '0',
    `updated_at`    int             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`, `sequence`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `key_result_values`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `key_result_id` bigint          NOT NULL,
    `value`         int             NOT NULL,
    `created_at`    int             NOT NULL,
    PRIMARY KEY (`id`),
    KEY `key_result_id` (`key_result_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `inboxes`
(
    `id`          BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     BIGINT(19)          NOT NULL,
    `sequence`    INT(10)             NOT NULL,
    `sender`      BIGINT(19)          NOT NULL,
    `sender_type` VARCHAR(20)         NOT NULL,
    `title`       VARCHAR(100)        NOT NULL DEFAULT '',
    `content`     VARCHAR(2048)       NOT NULL DEFAULT '',
    `type`        VARCHAR(50)         NOT NULL DEFAULT '',
    `payload`     VARCHAR(2048)       NOT NULL DEFAULT '',
    `status`      TINYINT(3)          NOT NULL,
    `created_at`  INT(10)             NOT NULL DEFAULT '0',
    `updated_at`  INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `user_sequence_id` (`user_id`, `sequence`) USING BTREE,
    INDEX `sender` (`sender`) USING BTREE,
    INDEX `status` (`status`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `counters`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    bigint          NOT NULL DEFAULT '0',
    `flag`       varchar(128)    NOT NULL DEFAULT '',
    `digit`      bigint          NOT NULL DEFAULT '0',
    `status`     tinyint         NOT NULL DEFAULT '0',
    `created_at` int             NOT NULL DEFAULT '0',
    `updated_at` int             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `counter_records`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `counter_id` bigint          NOT NULL DEFAULT '0',
    `digit`      bigint          NOT NULL DEFAULT '0',
    `created_at` int             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
