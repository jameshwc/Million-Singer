CREATE DATABASE IF NOT EXISTS million_singer;
USE million_singer;

CREATE TABLE IF NOT EXISTS `tours` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tours_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `collects` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(200),
  PRIMARY KEY (`id`),
  KEY `idx_collects_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `songs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `video_id` varchar(200),
  `start_time` varchar(200),
  `end_time` varchar(200),
  `language` varchar(200),
  `name` varchar(200),
  `singer` varchar(200),
  `genre` varchar(200),
  `miss_lyrics` varchar(200),
  PRIMARY KEY (`id`),
  KEY `idx_songs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `lyrics` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `index` int DEFAULT NULL,
  `line` longtext,
  `start_at` bigint DEFAULT NULL,
  `end_at` bigint DEFAULT NULL,
  `song_id` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_lyrics_deleted_at` (`deleted_at`),
  KEY `fk_songs_lyrics` (`song_id`),
  CONSTRAINT `fk_songs_lyrics` FOREIGN KEY (`song_id`) REFERENCES `songs` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `email` longtext,
  `password` longtext,
  `active` tinyint(1) DEFAULT NULL,
  `last_login` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `collect_songs` (
  `song_id` int unsigned NOT NULL,
  `collect_id` int unsigned NOT NULL,
  PRIMARY KEY (`song_id`,`collect_id`),
  KEY `fk_collect_songs_collect` (`collect_id`),
  CONSTRAINT `fk_collect_songs_collect` FOREIGN KEY (`collect_id`) REFERENCES `collects` (`id`),
  CONSTRAINT `fk_collect_songs_song` FOREIGN KEY (`song_id`) REFERENCES `songs` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `tour_collects` (
  `tour_id` int unsigned NOT NULL,
  `collect_id` int unsigned NOT NULL,
  PRIMARY KEY (`tour_id`,`collect_id`),
  KEY `fk_tour_collects_collect` (`collect_id`),
  CONSTRAINT `fk_tour_collects_collect` FOREIGN KEY (`collect_id`) REFERENCES `collects` (`id`),
  CONSTRAINT `fk_tour_collects_tour` FOREIGN KEY (`tour_id`) REFERENCES `tours` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;