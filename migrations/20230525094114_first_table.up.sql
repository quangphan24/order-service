CREATE TABLE IF NOT EXISTS orders
(
    `id`         VARCHAR(255)   NOT NULL  PRIMARY KEY,
    `total`      INTEGER NOT NULL,
    `wallet_id`  VARCHAR(255) NOT NULL,
    `status`     VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP    NULL     DEFAULT NULL
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS order_item
(
    `id`            VARCHAR(255)   NOT NULL PRIMARY KEY,
    `order_id`      VARCHAR(255) NOT NULL,
    `product_id`    VARCHAR(255) NOT NULL,
    `quantity`      INTEGER     NOT NULL,
    `price`         INTEGER     NOT NULL,
    `total`         INTEGER     NOT NULL,
    `created_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`    TIMESTAMP    NULL     DEFAULT NULL
    );