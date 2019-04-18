/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : lottery

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-04-18 20:02:39
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `lt_blackip`
-- ----------------------------
DROP TABLE IF EXISTS `lt_blackip`;
CREATE TABLE `lt_blackip` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'IP地址',
  `black_time` int(11) NOT NULL DEFAULT '0' COMMENT '黑名单限制到期时间',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ip` (`ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='黑名单表';

-- ----------------------------
-- Records of lt_blackip
-- ----------------------------

-- ----------------------------
-- Table structure for `lt_code`
-- ----------------------------
DROP TABLE IF EXISTS `lt_code`;
CREATE TABLE `lt_code` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `gift_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '奖品ID',
  `code` varchar(255) NOT NULL DEFAULT '' COMMENT '虚拟券编码',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `sys_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0：正常 1：作废 2：已发送',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`),
  KEY `gift_id` (`gift_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='虚拟券编码表';

-- ----------------------------
-- Records of lt_code
-- ----------------------------

-- ----------------------------
-- Table structure for `lt_gift`
-- ----------------------------
DROP TABLE IF EXISTS `lt_gift`;
CREATE TABLE `lt_gift` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品名称',
  `prize_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '奖品数量',
  `left_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '剩余奖品数量',
  `prize_code` varchar(50) NOT NULL DEFAULT '0' COMMENT '0-9999标识100%中奖概率，0-0标识万分之一的中奖概率',
  `prize_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '发放周期，以天为单位：D',
  `img` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品图片',
  `prize_order` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '奖品排序序号，小的排在前面',
  `gtype` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '奖品类型，0：虚拟币 1：实物小奖 2：实物大奖',
  `gdata` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品扩展数据，比如:虚拟币数量',
  `time_begin` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '抽奖活动开始时间',
  `time_end` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '抽奖活动结束时间',
  `prize_data` mediumtext COMMENT '发奖计划， [[时间1,数量1],[时间2,数量2]...]',
  `prize_begin` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '发奖计划周期的开始',
  `prize_end` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '发奖计划周期的结束',
  `sys_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0：正常 1：删除',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '操作人IP',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='奖品表';

-- ----------------------------
-- Records of lt_gift
-- ----------------------------

-- ----------------------------
-- Table structure for `lt_result`
-- ----------------------------
DROP TABLE IF EXISTS `lt_result`;
CREATE TABLE `lt_result` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `gift_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '奖品ID',
  `gift_name` varchar(255) NOT NULL DEFAULT '0' COMMENT '奖品名称',
  `gift_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '奖品类型，0：虚拟币 1：实物小奖 2：实物大奖',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `prize_code` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '抽奖编号（4位的随机数）',
  `gift_data` varchar(255) NOT NULL DEFAULT '' COMMENT '获奖信息',
  `sys_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0：正常 1：删除 2：作弊',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '操作人IP',
  PRIMARY KEY (`id`),
  KEY `gift_id` (`gift_id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='奖品结果表';

-- ----------------------------
-- Records of lt_result
-- ----------------------------

-- ----------------------------
-- Table structure for `lt_user`
-- ----------------------------
DROP TABLE IF EXISTS `lt_user`;
CREATE TABLE `lt_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `realname` varchar(50) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号码',
  `black_time` int(11) NOT NULL DEFAULT '0' COMMENT '黑名单限制到期时间',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '联系地址',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'IP地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Records of lt_user
-- ----------------------------

-- ----------------------------
-- Table structure for `lt_userday`
-- ----------------------------
DROP TABLE IF EXISTS `lt_userday`;
CREATE TABLE `lt_userday` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `day` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '日期到天',
  `num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '次数',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid` (`uid`,`day`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户每日抽奖记录表';

-- ----------------------------
-- Records of lt_userday
-- ----------------------------
