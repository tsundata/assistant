CREATE TABLE `groups`
(
    `id`         BIGINT(19)   NOT NULL AUTO_INCREMENT,
    `sequence`   INT(10)      NOT NULL,
    `uuid`       varchar(36)  NULL DEFAULT NULL,
    `user_id`    BIGINT(19)   NULL DEFAULT NULL,
    `name`       varchar(20)  NULL DEFAULT NULL,
    `avatar`     varchar(256) NULL DEFAULT NULL,
    `created_at` INT(10)      NULL DEFAULT NULL,
    `updated_at` INT(10)      NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `user_sequence_id` (`user_id`, `sequence`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE `nodes`
(
    `id`         BIGINT(19)  NOT NULL AUTO_INCREMENT,
    `ip`         VARCHAR(45) NULL DEFAULT NULL,
    `port`       SMALLINT(5) NULL DEFAULT NULL,
    `created_at` INT(10)     NULL DEFAULT NULL,
    `updated_at` INT(10)     NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `ip_port` (`ip`, `port`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
