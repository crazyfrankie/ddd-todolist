CREATE DATABASE todolist charset=utf8mb4;

USE todolist;

CREATE TABLE `user`(
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'Primary Key ID',
    `name` VARCHAR(255) NOT NULL COMMENT 'User Nickname',
    `email` VARCHAR(255) NOT NULL COMMENT 'Email',
    `password` VARCHAR(255) NOT NULL COMMENT 'Password (Encrypted)',
    `created_at` BIGINT NOT NULL COMMENT 'Creation Time (Milliseconds)',
    `updated_at` BIGINT NOT NULL COMMENT 'Update Time (Milliseconds)',
    `deleted_at` BIGINT NULL DEFAULT NULL COMMENT 'Deletion Time (Milliseconds)',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='User Table';

CREATE TABLE `task` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'Primary Key ID',
    `content` TEXT COMMENT 'Task Content',
    `user_id` BIGINT NOT NULL COMMENT 'Associated User ID',
    `due_time` BIGINT NULL DEFAULT NULL COMMENT '',
    `priority` ENUM('important and urgent', 'important but not urgent', 'not important but urgent', 'neither important or urgent') NOT NULL DEFAULT 'neither important or urgent' COMMENT 'Task priority level',
    `is_completed` BOOLEAN DEFAULT FALSE COMMENT 'is completed',
    `created_at` BIGINT NOT NULL COMMENT 'Creation Time (Milliseconds)',
    `updated_at` BIGINT NOT NULL COMMENT 'Update Time (Milliseconds)',
    `deleted_at` BIGINT NULL DEFAULT NULL COMMENT 'Deletion Time (Milliseconds)',
    PRIMARY KEY (`id`),
    INDEX `idx_task_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Task Table';