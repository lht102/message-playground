-- create "jobs" table
CREATE TABLE `jobs` (`id` char(36) NOT NULL, `request_uuid` char(36) NOT NULL, `state` enum('QUEUED','PROCESSING','COMPLETED') NOT NULL, `description` varchar(255) NOT NULL, `completed_at` timestamp NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `request_uuid` (`request_uuid`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
