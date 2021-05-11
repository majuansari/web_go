-- +goose Up
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(256) DEFAULT NULL,
  `name` varchar(256) DEFAULT NULL,
  `phone` varchar(256) DEFAULT NULL,
  `city` varchar(256) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

-- INSERT INTO users VALUES
-- (0, 'root', 'Root', ''),
-- (1, 'maju', 'Maju', '7259797591');

-- +goose Down
DROP TABLE users;
