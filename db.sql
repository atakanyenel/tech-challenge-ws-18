-- Adminer 4.7.0 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP DATABASE IF EXISTS `test`;
CREATE DATABASE `test` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */;
USE `test`;

DROP TABLE IF EXISTS `measurements`;
CREATE TABLE `measurements` (
  `socket_id` int(11) NOT NULL,
  `time` timestamp NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `measurements` (`socket_id`, `time`) VALUES
(1,	'2019-01-07 00:00:00'),
(1,	'2019-01-07 00:00:00'),
(2,	'2019-01-07 16:13:13'),
(2,	'2019-01-07 16:16:20'),
(2,	'2019-01-07 16:24:35'),
(1,	'2019-01-07 16:28:36'),
(1,	'2019-01-07 16:28:37'),
(1,	'2019-01-07 16:28:38'),
(1,	'2019-01-07 16:28:39'),
(1,	'2019-01-07 16:28:40'),
(1,	'2019-01-08 01:16:49'),
(2,	'2019-01-08 01:16:53'),
(1,	'2019-01-08 01:16:59'),
(1,	'2019-01-08 16:38:31'),
(1,	'2019-01-08 16:38:41'),
(1,	'2019-01-08 16:38:51'),
(1,	'2019-01-08 16:42:28'),
(1,	'2019-01-08 16:42:38'),
(1,	'2019-01-08 16:42:48'),
(1,	'2019-01-08 16:43:06'),
(1,	'2019-01-10 16:49:12'),
(1,	'2019-01-10 16:49:22'),
(1,	'2019-01-10 16:49:32'),
(1,	'2019-01-10 16:49:42'),
(1,	'2019-01-10 16:49:52'),
(1,	'2019-01-11 14:42:09'),
(1,	'2019-01-11 14:42:19'),
(1,	'2019-01-11 14:42:29'),
(1,	'2019-01-11 14:42:39'),
(1,	'2019-01-11 14:42:49');

DROP TABLE IF EXISTS `sockets`;
CREATE TABLE `sockets` (
  `socket_id` int(11) NOT NULL,
  `type` varchar(80) NOT NULL,
  `statu` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `sockets` (`socket_id`, `type`, `statu`) VALUES
(1,	'charging',	'ACTIVE'),
(2,	'lighting',	'FAIL'),
(3,	'cooking',	'STOPPED');


DROP TABLE IF EXISTS `ads`;
CREATE TABLE `ads` (
  `image_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `reason` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `title` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `product_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `ads` (`image_url`, `text`, `reason`, `title`, `product_url`) VALUES
('https://images-na.ssl-images-amazon.com/images/I/51B5x6fi4HL._SL1500_.jpg',	'This is the new PS4.',	'Entertainment > 2 hours every day',	'Playstation 4',	'https://www.amazon.de/dp/B07KMV94JF/ref=asc_df_B07KMV94JF57979247/?tag=googshopde-21&creative=22434&creativeASIN=B07KMV94JF&linkCode=df0&hvadid=308847264890&hvpos=1o1&hvnetw=g&hvrand=16041305715093971334&hvpone=&hvptwo=&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=9042512&hvtargid=pla-647432777687&th=1&psc=1&tag=&ref=&adgrpid=64245981791&hvpone=&hvptwo=&hvadid=308847264890&hvpos=1o1&hvnetw=g&hvrand=16041305715093971334&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=9042512&hvtargid=pla-647432777687');
-- 2019-01-11 14:46:13

