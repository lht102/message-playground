-- modify "jobs" table
ALTER TABLE `jobs` MODIFY COLUMN `completed_at` datetime(6) NULL, MODIFY COLUMN `created_at` datetime(6) NOT NULL, MODIFY COLUMN `updated_at` datetime(6) NOT NULL;
