/* *****************************************************************************
// Setup the preferences
// ****************************************************************************/
SET NAMES utf8 COLLATE 'utf8_unicode_ci';
SET foreign_key_checks = 1;
SET time_zone = '+00:00';
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';
SET storage_engine = InnoDB;
SET CHARACTER SET utf8;

/* *****************************************************************************
// Remove old database
// ****************************************************************************/
DROP DATABASE IF EXISTS todo_dev;

/* *****************************************************************************
// Create new database
// ****************************************************************************/
CREATE DATABASE todo_dev DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;
USE todo_dev;

/* *****************************************************************************
// Create the tables
// ****************************************************************************/
CREATE TABLE tasks (
    id TINYINT(1) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    completed TINYINT(1) UNSIGNED NOT NULL DEFAULT 0,
	due TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    
    PRIMARY KEY (id)
);

/* *****************************************************************************
// Seed Data
// ****************************************************************************/

INSERT INTO tasks (name, completed, due, created_at, updated_at) VALUES
("This is task 1",  0, DATE_ADD(NOW(), INTERVAL 1 DAY),  CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);