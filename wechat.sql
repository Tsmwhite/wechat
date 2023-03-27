/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 100508 (10.5.8-MariaDB)
 Source Host           : localhost:3306
 Source Schema         : thewhite

 Target Server Type    : MySQL
 Target Server Version : 100508 (10.5.8-MariaDB)
 File Encoding         : 65001

 Date: 27/03/2023 18:58:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for contacts
-- ----------------------------
DROP TABLE IF EXISTS `contacts`;
CREATE TABLE `contacts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `room` char(32) NOT NULL DEFAULT '',
  `user` char(32) NOT NULL DEFAULT '',
  `friend` char(32) NOT NULL DEFAULT '',
  `remark` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `avatar` varchar(255) NOT NULL DEFAULT '',
  `type` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `sort` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `last_time` int(10) unsigned NOT NULL DEFAULT 0,
  `last_msg` varchar(500) NOT NULL DEFAULT '',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `update_time` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `friend` (`friend`),
  KEY `user` (`user`)
) ENGINE=InnoDB AUTO_INCREMENT=107 DEFAULT CHARSET=utf8mb4 COMMENT='联系人关系';

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `path` varchar(255) NOT NULL DEFAULT '',
  `is_use` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `create_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `is_use` (`is_use`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE,
  KEY `uuid` (`uuid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for friends
-- ----------------------------
DROP TABLE IF EXISTS `friends`;
CREATE TABLE `friends` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` char(32) NOT NULL DEFAULT '',
  `friend` char(32) NOT NULL DEFAULT '',
  `remark` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `avatar` varchar(255) NOT NULL DEFAULT '',
  `create_time` int(10) unsigned NOT NULL DEFAULT 0,
  `type` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `hot` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `friend` (`friend`),
  KEY `user` (`user`)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COMMENT='好友关系';

-- ----------------------------
-- Table structure for group_members
-- ----------------------------
DROP TABLE IF EXISTS `group_members`;
CREATE TABLE `group_members` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `room` char(32) NOT NULL DEFAULT '',
  `user` char(32) NOT NULL DEFAULT '',
  `create_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `room` (`room`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',
  `type` int(11) NOT NULL DEFAULT 0,
  `second_type` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `parent` char(32) NOT NULL DEFAULT '',
  `reads` varchar(2500) NOT NULL DEFAULT '',
  `send_time` int(11) NOT NULL DEFAULT 0,
  `receive_time` int(11) NOT NULL DEFAULT 0,
  `forward_time` int(11) NOT NULL DEFAULT 0,
  `log_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `remark` varchar(50) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`uuid`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent` (`parent`) USING BTREE,
  KEY `send_time` (`send_time`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE,
  KEY `recipient_sender` (`recipient`,`sender`),
  KEY `recipient_type` (`recipient`,`type`) USING BTREE,
  KEY `recipient_second_type` (`recipient`,`second_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1233 DEFAULT CHARSET=utf8mb4 COMMENT='发送消息表';

-- ----------------------------
-- Table structure for messages_01
-- ----------------------------
DROP TABLE IF EXISTS `messages_01`;
CREATE TABLE `messages_01` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',
  `type` int(11) NOT NULL DEFAULT 0,
  `second_type` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `parent` char(32) NOT NULL DEFAULT '',
  `reads` varchar(2500) NOT NULL DEFAULT '',
  `send_time` int(11) NOT NULL DEFAULT 0,
  `receive_time` int(11) NOT NULL DEFAULT 0,
  `forward_time` int(11) NOT NULL DEFAULT 0,
  `log_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `remark` varchar(50) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`uuid`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent` (`parent`) USING BTREE,
  KEY `send_time` (`send_time`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE,
  KEY `recipient_sender` (`recipient`,`sender`),
  KEY `recipient_type` (`recipient`,`type`) USING BTREE,
  KEY `recipient_second_type` (`recipient`,`second_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=417 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for messages_02
-- ----------------------------
DROP TABLE IF EXISTS `messages_02`;
CREATE TABLE `messages_02` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',
  `type` int(11) NOT NULL DEFAULT 0,
  `second_type` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `parent` char(32) NOT NULL DEFAULT '',
  `reads` varchar(2500) NOT NULL DEFAULT '',
  `send_time` int(11) NOT NULL DEFAULT 0,
  `receive_time` int(11) NOT NULL DEFAULT 0,
  `forward_time` int(11) NOT NULL DEFAULT 0,
  `log_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `remark` varchar(50) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`uuid`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent` (`parent`) USING BTREE,
  KEY `send_time` (`send_time`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE,
  KEY `recipient_sender` (`recipient`,`sender`),
  KEY `recipient_type` (`recipient`,`type`) USING BTREE,
  KEY `recipient_second_type` (`recipient`,`second_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2081 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for messages_03
-- ----------------------------
DROP TABLE IF EXISTS `messages_03`;
CREATE TABLE `messages_03` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',
  `type` int(11) NOT NULL DEFAULT 0,
  `second_type` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `parent` char(32) NOT NULL DEFAULT '',
  `reads` varchar(2500) NOT NULL DEFAULT '',
  `send_time` int(11) NOT NULL DEFAULT 0,
  `receive_time` int(11) NOT NULL DEFAULT 0,
  `forward_time` int(11) NOT NULL DEFAULT 0,
  `log_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `remark` varchar(50) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`uuid`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent` (`parent`) USING BTREE,
  KEY `send_time` (`send_time`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE,
  KEY `recipient_sender` (`recipient`,`sender`),
  KEY `recipient_type` (`recipient`,`type`) USING BTREE,
  KEY `recipient_second_type` (`recipient`,`second_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1256 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for messages_04
-- ----------------------------
DROP TABLE IF EXISTS `messages_04`;
CREATE TABLE `messages_04` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',
  `type` int(11) NOT NULL DEFAULT 0,
  `second_type` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `parent` char(32) NOT NULL DEFAULT '',
  `reads` varchar(2500) NOT NULL DEFAULT '',
  `send_time` int(11) NOT NULL DEFAULT 0,
  `receive_time` int(11) NOT NULL DEFAULT 0,
  `forward_time` int(11) NOT NULL DEFAULT 0,
  `log_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `remark` varchar(50) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`uuid`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent` (`parent`) USING BTREE,
  KEY `send_time` (`send_time`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE,
  KEY `recipient_sender` (`recipient`,`sender`),
  KEY `recipient_type` (`recipient`,`type`) USING BTREE,
  KEY `recipient_second_type` (`recipient`,`second_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=834 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for messages_05
-- ----------------------------
DROP TABLE IF EXISTS `messages_05`;
CREATE TABLE `messages_05` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',
  `type` int(11) NOT NULL DEFAULT 0,
  `second_type` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `parent` char(32) NOT NULL DEFAULT '',
  `reads` varchar(2500) NOT NULL DEFAULT '',
  `send_time` int(11) NOT NULL DEFAULT 0,
  `receive_time` int(11) NOT NULL DEFAULT 0,
  `forward_time` int(11) NOT NULL DEFAULT 0,
  `log_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `remark` varchar(50) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`uuid`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent` (`parent`) USING BTREE,
  KEY `send_time` (`send_time`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE,
  KEY `recipient_sender` (`recipient`,`sender`),
  KEY `recipient_type` (`recipient`,`type`) USING BTREE,
  KEY `recipient_second_type` (`recipient`,`second_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1249 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for messages_06
-- ----------------------------
DROP TABLE IF EXISTS `messages_06`;
CREATE TABLE `messages_06` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',
  `type` int(11) NOT NULL DEFAULT 0,
  `second_type` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `parent` char(32) NOT NULL DEFAULT '',
  `reads` varchar(2500) NOT NULL DEFAULT '',
  `send_time` int(11) NOT NULL DEFAULT 0,
  `receive_time` int(11) NOT NULL DEFAULT 0,
  `forward_time` int(11) NOT NULL DEFAULT 0,
  `log_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `remark` varchar(50) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`uuid`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent` (`parent`) USING BTREE,
  KEY `send_time` (`send_time`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE,
  KEY `recipient_sender` (`recipient`,`sender`),
  KEY `recipient_type` (`recipient`,`type`) USING BTREE,
  KEY `recipient_second_type` (`recipient`,`second_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1667 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for messages_07
-- ----------------------------
DROP TABLE IF EXISTS `messages_07`;
CREATE TABLE `messages_07` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',
  `type` int(11) NOT NULL DEFAULT 0,
  `second_type` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `parent` char(32) NOT NULL DEFAULT '',
  `reads` varchar(2500) NOT NULL DEFAULT '',
  `send_time` int(11) NOT NULL DEFAULT 0,
  `receive_time` int(11) NOT NULL DEFAULT 0,
  `forward_time` int(11) NOT NULL DEFAULT 0,
  `log_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `remark` varchar(50) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`uuid`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent` (`parent`) USING BTREE,
  KEY `send_time` (`send_time`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE,
  KEY `recipient_sender` (`recipient`,`sender`),
  KEY `recipient_type` (`recipient`,`type`) USING BTREE,
  KEY `recipient_second_type` (`recipient`,`second_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2507 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for messages_08
-- ----------------------------
DROP TABLE IF EXISTS `messages_08`;
CREATE TABLE `messages_08` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',
  `type` int(11) NOT NULL DEFAULT 0,
  `second_type` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `parent` char(32) NOT NULL DEFAULT '',
  `reads` varchar(2500) NOT NULL DEFAULT '',
  `send_time` int(11) NOT NULL DEFAULT 0,
  `receive_time` int(11) NOT NULL DEFAULT 0,
  `forward_time` int(11) NOT NULL DEFAULT 0,
  `log_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `remark` varchar(50) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`uuid`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent` (`parent`) USING BTREE,
  KEY `send_time` (`send_time`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE,
  KEY `recipient_sender` (`recipient`,`sender`),
  KEY `recipient_type` (`recipient`,`type`) USING BTREE,
  KEY `recipient_second_type` (`recipient`,`second_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1665 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for messages_09
-- ----------------------------
DROP TABLE IF EXISTS `messages_09`;
CREATE TABLE `messages_09` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',
  `type` int(11) NOT NULL DEFAULT 0,
  `second_type` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `parent` char(32) NOT NULL DEFAULT '',
  `reads` varchar(2500) NOT NULL DEFAULT '',
  `send_time` int(11) NOT NULL DEFAULT 0,
  `receive_time` int(11) NOT NULL DEFAULT 0,
  `forward_time` int(11) NOT NULL DEFAULT 0,
  `log_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `remark` varchar(50) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`uuid`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent` (`parent`) USING BTREE,
  KEY `send_time` (`send_time`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE,
  KEY `recipient_sender` (`recipient`,`sender`),
  KEY `recipient_type` (`recipient`,`type`) USING BTREE,
  KEY `recipient_second_type` (`recipient`,`second_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1249 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for messages_handle
-- ----------------------------
DROP TABLE IF EXISTS `messages_handle`;
CREATE TABLE `messages_handle` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',
  `type` int(11) NOT NULL DEFAULT 0,
  `second_type` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `parent` char(32) NOT NULL DEFAULT '',
  `send_time` int(11) NOT NULL DEFAULT 0,
  `log_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `remark` varchar(50) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  `is_read` tinyint(1) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`uuid`) USING BTREE,
  KEY `sender` (`sender`) USING BTREE,
  KEY `recipient` (`recipient`) USING BTREE,
  KEY `type` (`type`) USING BTREE,
  KEY `second_type` (`second_type`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent` (`parent`) USING BTREE,
  KEY `send_time` (`send_time`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COMMENT='好友消息处理';

-- ----------------------------
-- Table structure for receive_messages
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages`;
CREATE TABLE `receive_messages` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='接受消息表';

-- ----------------------------
-- Table structure for receive_messages_01
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_01`;
CREATE TABLE `receive_messages_01` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_02
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_02`;
CREATE TABLE `receive_messages_02` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_03
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_03`;
CREATE TABLE `receive_messages_03` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_04
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_04`;
CREATE TABLE `receive_messages_04` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_05
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_05`;
CREATE TABLE `receive_messages_05` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_06
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_06`;
CREATE TABLE `receive_messages_06` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_07
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_07`;
CREATE TABLE `receive_messages_07` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_08
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_08`;
CREATE TABLE `receive_messages_08` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_09
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_09`;
CREATE TABLE `receive_messages_09` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_10
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_10`;
CREATE TABLE `receive_messages_10` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_11
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_11`;
CREATE TABLE `receive_messages_11` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_12
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_12`;
CREATE TABLE `receive_messages_12` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_13
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_13`;
CREATE TABLE `receive_messages_13` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_14
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_14`;
CREATE TABLE `receive_messages_14` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_15
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_15`;
CREATE TABLE `receive_messages_15` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_16
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_16`;
CREATE TABLE `receive_messages_16` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_17
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_17`;
CREATE TABLE `receive_messages_17` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_18
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_18`;
CREATE TABLE `receive_messages_18` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for receive_messages_19
-- ----------------------------
DROP TABLE IF EXISTS `receive_messages_19`;
CREATE TABLE `receive_messages_19` (
  `id` int(10) unsigned NOT NULL DEFAULT 0,
  `msg_uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接收者uuid',
  `room` char(32) NOT NULL DEFAULT '' COMMENT '接收房间uuid',
  `second_type` int(11) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `is_read` tinyint(1) NOT NULL DEFAULT 0,
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`msg_uuid`,`recipient`) USING BTREE,
  KEY `recipient_type` (`room`,`recipient`,`second_type`),
  KEY `recipient_sender` (`room`,`recipient`,`sender`,`is_read`) USING BTREE,
  KEY `recipient_read_status` (`room`,`recipient`,`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for rooms
-- ----------------------------
DROP TABLE IF EXISTS `rooms`;
CREATE TABLE `rooms` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `title` varchar(255) NOT NULL DEFAULT '',
  `type` tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT '0 单聊；1 群聊',
  `members` varchar(2500) NOT NULL DEFAULT '',
  `creator` char(32) NOT NULL DEFAULT '',
  `description` varchar(255) NOT NULL DEFAULT '',
  `create_time` int(11) unsigned NOT NULL DEFAULT 0,
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `member_num` int(5) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COMMENT='聊天房间';

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `name` varchar(20) NOT NULL DEFAULT '',
  `mobile` char(11) NOT NULL DEFAULT '',
  `mail` varchar(50) NOT NULL DEFAULT '',
  `avatar` varchar(500) NOT NULL DEFAULT '',
  `role_id` int(11) NOT NULL DEFAULT 0,
  `salt` char(6) NOT NULL DEFAULT '',
  `password` varchar(50) NOT NULL DEFAULT '',
  `pass_look` varchar(50) NOT NULL DEFAULT '',
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `register_time` int(11) NOT NULL DEFAULT 0,
  `login_time` int(11) NOT NULL DEFAULT 0,
  `update_time` int(11) NOT NULL DEFAULT 0,
  `register_ip` varchar(255) NOT NULL DEFAULT '',
  `login_ip` varchar(255) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `mail` (`mail`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE,
  UNIQUE KEY `uuid` (`uuid`) USING BTREE,
  KEY `mobile` (`mobile`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1036844 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for verify_codes
-- ----------------------------
DROP TABLE IF EXISTS `verify_codes`;
CREATE TABLE `verify_codes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(50) NOT NULL DEFAULT '' COMMENT '账户',
  `code` char(6) NOT NULL DEFAULT '',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '1 账户验证',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0 待验证；1  已验证',
  `create_time` int(10) unsigned NOT NULL DEFAULT 0,
  `update_time` int(10) unsigned NOT NULL DEFAULT 0,
  `expire_time` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `account` (`account`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
