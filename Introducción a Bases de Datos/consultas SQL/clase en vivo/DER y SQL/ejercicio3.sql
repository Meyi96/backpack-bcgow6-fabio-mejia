-- MySQL dump 10.13  Distrib 8.0.27, for macos11 (x86_64)
--
-- Host: localhost    Database: empresa_internet
-- ------------------------------------------------------
-- Server version	8.0.31

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
-- Table structure for table `client`
--

DROP TABLE IF EXISTS `client`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `client` (
  `idclient` int NOT NULL AUTO_INCREMENT,
  `dni` varchar(45) NOT NULL,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `city` varchar(200) NOT NULL,
  `province` varchar(200) NOT NULL,
  `birth_date` timestamp NOT NULL,
  PRIMARY KEY (`idclient`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `client`
--

LOCK TABLES `client` WRITE;
/*!40000 ALTER TABLE `client` DISABLE KEYS */;
INSERT INTO `client` VALUES (14,'1234','Fabio','Mejia','Palmira','Valle','1999-03-17 05:00:00'),(15,'3421','Juan','Parra','Cali','Valle','1992-05-10 04:00:00'),(16,'9543','Nicolas','Velez','Cali','Valle','1994-07-21 05:00:00'),(17,'5421','Franco','Perez','Medellin','Antioquia','2001-10-14 05:00:00'),(18,'6421','Sebastian','Rozo','Bello','Antioquia','1996-03-15 05:00:00'),(19,'3253','Gloria','Garcia','Bogota','Bogota','1999-01-17 05:00:00'),(20,'5772','Isabella','Rodriguez','Buga','Valle','1989-03-26 05:00:00'),(21,'4345','Maria','Muriel','Bogota','Bogota','1997-10-31 05:00:00'),(22,'5365','Valeria','Zamora','Medellin','Antioquia','1998-08-27 05:00:00'),(23,'7142','Sofia','Ossorio','Palmira','Valle','2000-06-05 05:00:00');
/*!40000 ALTER TABLE `client` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `plan`
--

DROP TABLE IF EXISTS `plan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `plan` (
  `idplan` int NOT NULL AUTO_INCREMENT,
  `price` decimal(9,2) NOT NULL,
  `speed` varchar(100) NOT NULL,
  PRIMARY KEY (`idplan`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `plan`
--

LOCK TABLES `plan` WRITE;
/*!40000 ALTER TABLE `plan` DISABLE KEYS */;
INSERT INTO `plan` VALUES (1,49.90,'50 Megas'),(2,99.90,'100 Megas'),(3,134.90,'200 Megas'),(4,169.90,'400 Megas'),(5,199.90,'1 GB');
/*!40000 ALTER TABLE `plan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service`
--

DROP TABLE IF EXISTS `service`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service` (
  `idservice` int NOT NULL AUTO_INCREMENT,
  `reference` varchar(45) NOT NULL,
  `discount` decimal(2,2) NOT NULL,
  `status` varchar(50) NOT NULL,
  `idclient` int NOT NULL,
  `idplan` int NOT NULL,
  PRIMARY KEY (`idservice`),
  UNIQUE KEY `reference_UNIQUE` (`reference`),
  KEY `client_service_idx` (`idclient`),
  KEY `plan_service_idx` (`idplan`),
  CONSTRAINT `client_service` FOREIGN KEY (`idclient`) REFERENCES `client` (`idclient`),
  CONSTRAINT `plan_service` FOREIGN KEY (`idplan`) REFERENCES `plan` (`idplan`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service`
--

LOCK TABLES `service` WRITE;
/*!40000 ALTER TABLE `service` DISABLE KEYS */;
INSERT INTO `service` VALUES (2,'12345',0.20,'activo',20,1),(3,'34213',0.15,'activo',14,5),(4,'49202',0.10,'activo',16,4),(5,'31342',0.25,'activo',16,3),(6,'23512',0.12,'activo',21,4);
/*!40000 ALTER TABLE `service` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-11-02 13:40:52
