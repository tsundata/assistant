CREATE TABLE `groups`
(
    `id`         BIGINT(19)   NOT NULL AUTO_INCREMENT,
    `sequence`   INT(10)      NOT NULL,
    `type`       TINYINT(1)   NOT NULL,
    `user_id`    BIGINT(19)   NOT NULL,
    `name`       VARCHAR(20)  NOT NULL,
    `avatar`     VARCHAR(256) NOT NULL,
    `created_at` INT(10)      NOT NULL,
    `updated_at` INT(10)      NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `user_sequence_id` (`user_id`, `sequence`) USING BTREE
)
    ENGINE = InnoDB;


CREATE TABLE `group_bots`
(
    `id`         BIGINT(19) NOT NULL AUTO_INCREMENT,
    `group_id`   BIGINT(19) NOT NULL,
    `bot_id`     BIGINT(19) NOT NULL,
    `created_at` INT(10)    NOT NULL,
    `updated_at` INT(10)    NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `group_id` (`group_id`) USING BTREE
)
    ENGINE = InnoDB;


CREATE TABLE `group_bot_settings`
(
    `group_id`   BIGINT(19)    NOT NULL,
    `bot_id`     BIGINT(19)    NOT NULL,
    `key`        VARCHAR(50)   NOT NULL,
    `value`      VARCHAR(1000) NOT NULL,
    `created_at` INT(10)       NOT NULL,
    `updated_at` INT(10)       NOT NULL,
    INDEX `group_id` (`group_id`) USING BTREE,
    INDEX `bot_id` (`bot_id`) USING BTREE
)
    ENGINE = InnoDB;


CREATE TABLE `group_settings`
(
    `group_id`   BIGINT(19)    NOT NULL,
    `key`        VARCHAR(50)   NOT NULL,
    `value`      VARCHAR(1000) NOT NULL,
    `created_at` INT(10)       NOT NULL,
    `updated_at` INT(10)       NOT NULL,
    INDEX `group_id` (`group_id`) USING BTREE
)
    ENGINE = InnoDB;


CREATE TABLE `messages`
(
    `id`            BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `group_id`      BIGINT(19)          NOT NULL,
    `user_id`       BIGINT(19)          NOT NULL,
    `sequence`      INT(10)             NOT NULL,
    `sender`        BIGINT(19)          NOT NULL,
    `sender_type`   VARCHAR(20)         NOT NULL,
    `receiver`      BIGINT(19)          NOT NULL,
    `receiver_type` VARCHAR(20)         NOT NULL,
    `type`          varchar(12)         NOT NULL DEFAULT '',
    `text`          varchar(2048)       NOT NULL DEFAULT '',
    `payload`       json                NOT NULL,
    `action`        varchar(5120)       NOT NULL DEFAULT '',
    `status`        TINYINT(2)          NOT NULL,
    `created_at`    INT(10)             NOT NULL DEFAULT '0',
    `updated_at`    INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    INDEX `group_id` (`group_id`) USING BTREE,
    INDEX `sender` (`sender`) USING BTREE,
    INDEX `receiver` (`receiver`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE `bots`
(
    `id`         BIGINT(19)    NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(30)   NOT NULL,
    `identifier` VARCHAR(20)   NOT NULL,
    `detail`     VARCHAR(250)  NOT NULL,
    `avatar`     VARCHAR(256)  NOT NULL,
    `extend`     VARCHAR(2000) NOT NULL,
    `status`     TINYINT(2)    NOT NULL,
    `created_at` INT(10)       NOT NULL,
    `updated_at` INT(10)       NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `identifier` (`identifier`) USING BTREE
)
    ENGINE = InnoDB;


CREATE TABLE `nodes`
(
    `id`         BIGINT(19)  NOT NULL AUTO_INCREMENT,
    `ip`         VARCHAR(45) NOT NULL,
    `port`       SMALLINT(5) NOT NULL,
    `created_at` INT(10)     NOT NULL,
    `updated_at` INT(10)     NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `ip_port` (`ip`, `port`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE `devices`
(
    `id`         BIGINT(19)  NOT NULL AUTO_INCREMENT,
    `user_id`    BIGINT(19)  NOT NULL,
    `name`       VARCHAR(50) NOT NULL,
    `created_at` INT(10)     NOT NULL,
    `updated_at` INT(10)     NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE `apps`
(
    `id`         BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`    BIGINT(19) UNSIGNED NOT NULL,
    `name`       varchar(16)         NOT NULL DEFAULT '',
    `type`       varchar(12)         NOT NULL DEFAULT '',
    `token`      varchar(256)        NOT NULL DEFAULT '',
    `extra`      varchar(2048)       NOT NULL DEFAULT '',
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    `updated_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    INDEX `user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE `credentials`
(
    `id`         BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`    BIGINT(20) UNSIGNED NOT NULL DEFAULT '0',
    `name`       varchar(16)         NOT NULL DEFAULT '',
    `type`       varchar(12)         NOT NULL DEFAULT '',
    `content`    varchar(2048)       NOT NULL DEFAULT '',
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    `updated_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `user_id_name` (`user_id`, `name`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;



CREATE TABLE `roles`
(
    `id`          BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     BIGINT(10)          NOT NULL,
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
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `user_id` (`user_id`) USING BTREE
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

INSERT INTO `roles` (`user_id`, `profession`, `exp`, `level`, `strength`, `culture`, `environment`, `charisma`,
                     `talent`, `intellect`)
VALUES (1, 'super', 0, 1, 0, 0, 0, 0, 0, 0);


CREATE TABLE `role_records`
(
    `id`          BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     BIGINT(10)          NOT NULL,
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


CREATE TABLE `users`
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


CREATE TABLE `tags`
(
    `id`         BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`    BIGINT              NOT NULL,
    `name`       varchar(50)         NOT NULL DEFAULT '',
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    `updated_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE `model_tags`
(
    `id`         BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`    BIGINT              NOT NULL,
    `service`    VARCHAR(50)         NOT NULL DEFAULT '',
    `model`      VARCHAR(50)         NOT NULL DEFAULT '',
    `model_id`   BIGINT(19)          NOT NULL,
    `tag_id`     BIGINT(19)          NOT NULL,
    `created_at` INT(10)             NOT NULL DEFAULT '0',
    `updated_at` INT(10)             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `service` (`service`) USING BTREE,
    INDEX `model` (`model`) USING BTREE,
    INDEX `model_id` (`model_id`) USING BTREE,
    INDEX `tag_id` (`tag_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `subscribes`
(
    `id`         BIGINT(19)  NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(20) NOT NULL,
    `status`     TINYINT(1)  NOT NULL DEFAULT '1',
    `created_at` INT(10)     NOT NULL,
    `updated_at` INT(10)     NOT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE `user_subscribes`
(
    `id`           BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`      BIGINT(20) UNSIGNED NOT NULL,
    `subscribe_id` BIGINT(20) UNSIGNED NOT NULL,
    `status`       TINYINT(3) UNSIGNED NOT NULL DEFAULT '1',
    `created_at`   INT(10)             NOT NULL,
    `updated_at`   INT(10)             NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `user_id` (`user_id`) USING BTREE,
    INDEX `subscribe_id` (`subscribe_id`) USING BTREE,
    INDEX `status` (`status`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `metadata`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    bigint          NOT NULL,
    `model`      varchar(30)     NOT NULL,
    `model_id`   bigint          NOT NULL,
    `sequence`   int             NOT NULL,
    `text`       varchar(100)    NOT NULL,
    `data`       json            NOT NULL,
    `extra`      json            NOT NULL,
    `created_at` int             NOT NULL,
    `updated_at` int             NOT NULL,
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`),
    KEY `model` (`model`),
    KEY `model_id` (`model_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
