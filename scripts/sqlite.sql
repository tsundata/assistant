CREATE TABLE "todos"
(
    "id"                INTEGER       NOT NULL,
    "content"           VARCHAR(1024) NOT NULL,
    "priority"          TINYINT       NOT NULL DEFAULT '0',
    "is_remind_at_time" TINYINT       NOT NULL DEFAULT '0',
    "remind_at"         DATETIME      NOT NULL,
    "repeat_method"     VARCHAR(50)   NOT NULL DEFAULT '',
    "repeat_rule"       VARCHAR(256)  NOT NULL DEFAULT '',
    "repeat_end_at"     DATETIME      NOT NULL,
    "category"          VARCHAR(50)   NOT NULL DEFAULT '',
    "remark"            VARCHAR(1024) NOT NULL DEFAULT '',
    "complete"          TINYINT       NOT NULL,
    "created_at"        DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"        DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE TABLE `triggers`
(
    `id`         INTEGER      NOT NULL,
    `type`       varchar(16)  NOT NULL DEFAULT '',
    `kind`       varchar(16)  NOT NULL DEFAULT '',
    `flag`       varchar(128) NOT NULL DEFAULT '',
    `secret`     varchar(128) NOT NULL DEFAULT '',
    `when`       varchar(128) NOT NULL DEFAULT '',
    `message_id` INTEGER      NOT NULL,
    `created_at` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `users`
(
    `id`         INTEGER     NOT NULL,
    `name`       VARCHAR(50) NOT NULL,
    `mobile`     VARCHAR(50) NOT NULL DEFAULT '',
    `remark`     VARCHAR(50) NOT NULL DEFAULT '',
    `created_at` DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (`id`, `name`, `mobile`, `remark`, `created_at`, `updated_at`)
VALUES (1, 'me', '', '', '2021-07-01 00:00:00', '2021-07-01 00:00:00');

CREATE TABLE `roles`
(
    `id`          INTEGER     NOT NULL,
    `user_id`     INTEGER     NOT NULL,
    `profession`  VARCHAR(50) NOT NULL,
    `exp`         INTEGER     NOT NULL DEFAULT '0',
    `level`       INTEGER     NOT NULL DEFAULT '1',
    `strength`    INTEGER     NOT NULL DEFAULT '0',
    `culture`     INTEGER     NOT NULL DEFAULT '0',
    `environment` INTEGER     NOT NULL DEFAULT '0',
    `charisma`    INTEGER     NOT NULL DEFAULT '0',
    `talent`      INTEGER     NOT NULL DEFAULT '0',
    `intellect`   INTEGER     NOT NULL DEFAULT '0',
    `created_at`  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `role_records`
(
    `id`          INTEGER     NOT NULL,
    `user_id`     INTEGER     NOT NULL,
    `profession`  VARCHAR(50) NOT NULL,
    `exp`         INTEGER     NOT NULL DEFAULT '0',
    `level`       INTEGER     NOT NULL DEFAULT '0',
    `strength`    INTEGER     NOT NULL DEFAULT '0',
    `culture`     INTEGER     NOT NULL DEFAULT '0',
    `environment` INTEGER     NOT NULL DEFAULT '0',
    `charisma`    INTEGER     NOT NULL DEFAULT '0',
    `talent`      INTEGER     NOT NULL DEFAULT '0',
    `intellect`   INTEGER     NOT NULL DEFAULT '0',
    `created_at`  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

INSERT INTO `role_records` (`id`, `user_id`, `profession`, `exp`, `level`, `strength`, `culture`, `environment`,
                            `charisma`,
                            `talent`, `intellect`)
VALUES (1, 1, 'super', 0, 1, 0, 0, 0, 0, 0, 0);


CREATE TABLE `messages`
(
    `id`         INTEGER       NOT NULL,
    `uuid`       varchar(36)   NOT NULL DEFAULT '',
    `type`       varchar(12)   NOT NULL DEFAULT '',
    `channel`    varchar(20)   NOT NULL DEFAULT '',
    `text`       varchar(2048) NOT NULL DEFAULT '',
    `created_at` DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE (`uuid`)
);