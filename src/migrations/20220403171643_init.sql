-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `quotes` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `content` LONGTEXT NOT NULL,
  `author` TEXT NOT NULL,
  `like` BIGINT(20) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `quotes_of_the_days` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `quotes_id` BIGINT(20) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
