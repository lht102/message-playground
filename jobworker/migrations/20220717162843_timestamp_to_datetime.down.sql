-- reverse: modify "jobs" table
ALTER TABLE `jobs` MODIFY COLUMN `updated_at` timestamp NOT NULL, MODIFY COLUMN `created_at` timestamp NOT NULL, MODIFY COLUMN `completed_at` timestamp NULL;
