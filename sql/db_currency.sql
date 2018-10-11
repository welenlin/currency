/*
Navicat MySQL Data Transfer

Source Server         : LOCAL
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : db_currency2

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2018-10-10 23:01:40
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `exchange`
-- ----------------------------
DROP TABLE IF EXISTS `exchange`;
CREATE TABLE `exchange` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `from_currency` varchar(255) NOT NULL,
  `to_currency` varchar(255) NOT NULL,
  `rate` decimal(10,4) NOT NULL,
  `rate_date` date NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_exchange` (`from_currency`,`to_currency`,`rate_date`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of exchange
-- ----------------------------

-- ----------------------------
-- Table structure for `exchange_list`
-- ----------------------------
DROP TABLE IF EXISTS `exchange_list`;
CREATE TABLE `exchange_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `from_currency` varchar(10) NOT NULL,
  `to_currency` varchar(10) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_exchangelist` (`from_currency`,`to_currency`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of exchange_list
-- ----------------------------
