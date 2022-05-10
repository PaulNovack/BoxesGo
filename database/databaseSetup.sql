CREATE DATABASE  IF NOT EXISTS `boxes` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `boxes`;
-- MySQL dump 10.13  Distrib 8.0.25, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: boxes
-- ------------------------------------------------------
-- Server version	8.0.29-0ubuntu0.20.04.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `boxes`
--

DROP TABLE IF EXISTS `boxes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `boxes` (
  `id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `name` varchar(255) NOT NULL,
  `weight` decimal(10,2) DEFAULT NULL,
  `picture` varchar(255) DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `boxes`
--

LOCK TABLES `boxes` WRITE;
/*!40000 ALTER TABLE `boxes` DISABLE KEYS */;
INSERT INTO `boxes` VALUES (3,1,'New Box Name 67',34.67,'','2022-05-10 12:29:47'),(4,1,'test2',34.65,'','2022-05-10 12:29:47'),(5,1,'Yan yah',0.00,'','2022-05-10 13:13:43'),(6,1,'Yan yah',0.00,'','2022-05-10 13:14:10'),(7,1,'Yan yah',0.00,'','2022-05-10 13:17:11'),(8,1,'Yan yah',0.00,'','2022-05-10 13:17:20'),(9,1,'New Box Name 67',34.67,'HERES MY PICRUTE.DFDE','2022-05-10 13:17:22'),(10,1,'Yan yah',0.00,'','2022-05-10 13:17:23'),(11,1,'Yan yah',0.00,'','2022-05-10 13:17:23'),(12,1,'Yan yah',0.00,'','2022-05-10 13:17:24'),(13,1,'Yan yah',0.00,'','2022-05-10 13:17:25'),(14,1,'Yan yah',0.00,'','2022-05-10 13:17:25'),(15,1,'Yan yah',0.00,'','2022-05-10 13:17:26'),(16,1,'Yan yah',0.00,'','2022-05-10 13:17:26'),(17,1,'Yan yah',0.00,'','2022-05-10 14:00:11');
/*!40000 ALTER TABLE `boxes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `items`
--

DROP TABLE IF EXISTS `items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `items` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `box_id` int NOT NULL,
  `quantity` bigint NOT NULL DEFAULT '1',
  `name` varchar(255) NOT NULL,
  `picture` varchar(255) DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `items`
--

LOCK TABLES `items` WRITE;
/*!40000 ALTER TABLE `items` DISABLE KEYS */;
INSERT INTO `items` VALUES (3,1,1,2,'New Item Name 67','HERES MY PICRUTE.DFDE','2022-05-10 12:40:18'),(4,1,5,10000000000,'New Item Name 4','My Picture Here','2022-05-10 12:40:18'),(5,1,1,2,'Yan yah','','2022-05-10 12:40:18'),(20,1,5,10000000000,'New Item Name 4','My Picture Here','2022-05-10 15:15:13'),(21,1,5,10,'New Item Name 4ddd','My Picture Here','2022-05-10 15:15:25'),(22,1,5,10,'New Item Name 4ddd','My Picture Here','2022-05-10 15:15:26'),(23,1,5,10,'New Item Name 4ddd','My Picture Here','2022-05-10 15:15:26'),(24,1,5,10,'New Item Name 4ddd','My Picture Here','2022-05-10 15:15:27'),(25,1,5,10,'New Item Name 4ddd','My Picture Here','2022-05-10 15:15:27'),(26,1,5,10,'New Item Name 4ddd','My Picture Here','2022-05-10 15:15:27'),(27,1,5,10,'New Item Name 4ddd','My Picture Here','2022-05-10 15:15:28');
/*!40000 ALTER TABLE `items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `authkey` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'paulnovack','paulnovack','4aad700b0645fbf90fa9d95a82494b3129f1a5b0cb8183867727ea5d21fc0821','2022-05-09 07:59:13'),(11,'paulnovac','paulnovack','25c983ce2863b4d8806da25d08ba4b0fc386926a6fbda8b4c1d6e7cb17e5d986','2022-05-10 08:15:36'),(12,'paulnovack3','paulnovack','fb271394052f6676139c7f47537153bc48eec675cfa01ee73efb949b5dc2ca85','2022-05-10 08:19:20'),(13,'paulnovack34','paulnovack','a7eae5b196cc756280f394764e16963e16c207366106ee5a5c5a7fa5c5a9d306','2022-05-10 08:19:33');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-10 15:21:32
