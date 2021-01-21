/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50732
Source Host           : localhost:3306
Source Database       : lottery

Target Server Type    : MYSQL
Target Server Version : 50732
File Encoding         : 65001

Date: 2021-01-16 16:34:14
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for lottery
-- ----------------------------
DROP TABLE IF EXISTS `lottery`;
CREATE TABLE `lottery` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '彩种id',
  `class_id` bigint(20) NOT NULL COMMENT '彩票分类id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '彩票名称',
  `code` varchar(50) NOT NULL DEFAULT '' COMMENT '彩票编码',
  `period_type` tinyint(1) DEFAULT NULL COMMENT '周期类型;1-1分钟',
  `disable` tinyint(1) DEFAULT '0' COMMENT '禁用,1-禁用;0-启用',
  `is_self` tinyint(1) NOT NULL COMMENT '官彩或私彩;1-私彩;0-管彩',
  `remarks` varchar(255) DEFAULT '' COMMENT '彩票备注',
  `display_order` int(11) NOT NULL COMMENT '排序',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`,`code`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='彩票种类';

-- ----------------------------
-- Table structure for lottery_class
-- ----------------------------
DROP TABLE IF EXISTS `lottery_class`;
CREATE TABLE `lottery_class` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '彩票分类id',
  `name` varchar(50) NOT NULL COMMENT '彩种分类名称',
  `code` varchar(50) NOT NULL COMMENT '彩种分类编码',
  `remarks` varchar(255) DEFAULT NULL COMMENT '备注',
  `display_order` int(11) NOT NULL COMMENT '排序',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='彩种分类表';

-- ----------------------------
-- Table structure for lottery_play_group
-- ----------------------------
DROP TABLE IF EXISTS `lottery_play_group`;
CREATE TABLE `lottery_play_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `class_id` bigint(20) NOT NULL COMMENT '彩票分类id',
  `name` varchar(50) NOT NULL COMMENT '玩法名称',
  `code` varchar(50) DEFAULT NULL COMMENT '玩法编码',
  `model` tinyint(1) NOT NULL DEFAULT '1' COMMENT '盘口类型;1-标准盘,2-两面盘',
  `display_order` int(11) NOT NULL COMMENT '排序',
  `remarks` varchar(255) DEFAULT NULL COMMENT '备注',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='彩票玩法种类表';

-- ----------------------------
-- Table structure for lottery_play_type
-- ----------------------------
DROP TABLE IF EXISTS `lottery_play_type`;
CREATE TABLE `lottery_play_type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `class_id` bigint(20) NOT NULL COMMENT '彩票分类id',
  `group_id` bigint(20) NOT NULL COMMENT '玩法总类id',
  `name` varchar(50) NOT NULL COMMENT '玩法名称',
  `code` varchar(50) NOT NULL COMMENT '玩法编码',
  `model` tinyint(1) NOT NULL DEFAULT '1' COMMENT '盘口类型;1-标准盘,2-两面盘',
  `display_order` int(11) NOT NULL DEFAULT '1' COMMENT '排序',
  `remarks` varchar(255) DEFAULT NULL COMMENT '备注',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COMMENT='彩票玩法子类表';

-- ----------------------------
-- Table structure for lottery_plays
-- ----------------------------
DROP TABLE IF EXISTS `lottery_plays`;
CREATE TABLE `lottery_plays` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `class_id` bigint(20) NOT NULL COMMENT '彩票分类id',
  `group_id` bigint(20) NOT NULL COMMENT '玩法总类id',
  `type_id` bigint(20) DEFAULT NULL COMMENT '玩法子类id',
  `name` varchar(50) NOT NULL COMMENT '玩法名称',
  `code` varchar(50) NOT NULL COMMENT '玩法编码',
  `odds` decimal(13,2) DEFAULT NULL COMMENT '赔率',
  `display_order` int(11) NOT NULL DEFAULT '1' COMMENT '排序',
  `remarks` varchar(255) DEFAULT NULL COMMENT '备注',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COMMENT='彩票玩法表';
