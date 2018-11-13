/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.5.53-log : Database - auditintegral
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`auditintegral` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `auditintegral`;

/*Table structure for table `clause` */

DROP TABLE IF EXISTS `clause`;

CREATE TABLE `clause` (
  `id` int(11) NOT NULL COMMENT '管理办法及条款ID',
  `department_id` int(11) DEFAULT NULL COMMENT '所属部门ID',
  `title` char(200) DEFAULT NULL COMMENT '管理办法及条款标题',
  `author_id` int(11) DEFAULT NULL COMMENT '发布人ID',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `state` enum('draft','publish') DEFAULT 'draft' COMMENT '状态（draft草稿publish发布）',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `clause` */

/*Table structure for table `clause_content` */

DROP TABLE IF EXISTS `clause_content`;

CREATE TABLE `clause_content` (
  `id` int(11) NOT NULL COMMENT '管理办法内容ID',
  `is_title` tinyint(1) DEFAULT '0' COMMENT '是否为标题（0内容1标题）',
  `title_level` char(20) DEFAULT NULL COMMENT '标题级别',
  `content` varchar(1000) DEFAULT NULL COMMENT '内容',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `clause_content` */

/*Table structure for table `clause_file` */

DROP TABLE IF EXISTS `clause_file`;

CREATE TABLE `clause_file` (
  `clause_id` int(11) DEFAULT NULL COMMENT '管理办法ID',
  `file_id` int(11) DEFAULT NULL COMMENT '附件ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `clause_file` */

/*Table structure for table `department` */

DROP TABLE IF EXISTS `department`;

CREATE TABLE `department` (
  `id` int(9) NOT NULL COMMENT '部门ID',
  `parent_id` int(9) DEFAULT NULL COMMENT '上级部门ID',
  `code` char(50) DEFAULT NULL COMMENT '部门编码',
  `level` char(20) DEFAULT NULL COMMENT '部门级别',
  `grade` int(3) DEFAULT NULL COMMENT '所在部门树的层级',
  `address` char(250) DEFAULT NULL COMMENT '地址',
  `phone` char(11) DEFAULT NULL COMMENT '手机号或者电话',
  `delete` tinyint(1) DEFAULT '0' COMMENT '是否删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `department` */

/*Table structure for table `dictionary` */

DROP TABLE IF EXISTS `dictionary`;

CREATE TABLE `dictionary` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '字典ID',
  `type_id` int(11) NOT NULL COMMENT '字典类型ID',
  `value` char(50) DEFAULT NULL COMMENT '字典值',
  `title` char(50) DEFAULT NULL COMMENT '字典名称',
  `order` int(4) DEFAULT '0' COMMENT '排序顺序',
  `describe` char(250) DEFAULT NULL COMMENT '字典描述',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`,`type_id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

/*Data for the table `dictionary` */

insert  into `dictionary`(`id`,`type_id`,`value`,`title`,`order`,`describe`,`delete`) values (-1,-1,'other','其他',0,'系统分类（不允许删除）',0),(1,-1,'system','系统',0,NULL,0);

/*Table structure for table `dictionary_type` */

DROP TABLE IF EXISTS `dictionary_type`;

CREATE TABLE `dictionary_type` (
  `type_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '字典类型ID',
  `title` char(50) NOT NULL COMMENT '字典类型名称',
  `type` char(50) NOT NULL COMMENT '字典分类',
  `key` char(20) NOT NULL COMMENT '字典分类键',
  `is_use` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `user_id` int(11) DEFAULT NULL COMMENT '创建人ID',
  `describe` char(250) DEFAULT NULL COMMENT '字典类型描述',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '软删除',
  PRIMARY KEY (`type_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `dictionary_type` */

insert  into `dictionary_type`(`type_id`,`title`,`type`,`key`,`is_use`,`create_time`,`user_id`,`describe`,`delete`) values (-1,'字典分类','dictionaryType','dictionaryType',1,'2018-11-12 15:37:43',NULL,'系统字典分类',0),(-2,'人员岗位','jobType','jobType',1,'2018-11-12 15:58:30',NULL,'人员岗位分类',0);

/*Table structure for table `file` */

DROP TABLE IF EXISTS `file`;

CREATE TABLE `file` (
  `id` int(11) NOT NULL COMMENT '附件ID',
  `name` char(100) DEFAULT NULL COMMENT '附件名称',
  `suffix` char(10) DEFAULT NULL COMMENT '附件后缀',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上传时间',
  `path` char(250) DEFAULT NULL COMMENT '附件存放位置',
  `filename` char(100) DEFAULT NULL COMMENT '存放名称',
  `delete` tinyint(1) DEFAULT NULL COMMENT '是否删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `file` */

/*Table structure for table `login` */

DROP TABLE IF EXISTS `login`;

CREATE TABLE `login` (
  `user_code` int(11) NOT NULL COMMENT '员工号',
  `password` char(50) DEFAULT NULL COMMENT '登录密码',
  `is_use` enum('1','0') DEFAULT '1' COMMENT '是否禁用（1启用0禁用）',
  `change_pd_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改密码时间',
  `login_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登录时间',
  `author_id` int(11) DEFAULT NULL COMMENT '授权人',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`user_code`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `login` */

/*Table structure for table `logs` */

DROP TABLE IF EXISTS `logs`;

CREATE TABLE `logs` (
  `id` int(11) DEFAULT NULL COMMENT '日志ID',
  `type` char(50) DEFAULT NULL COMMENT '日志类型',
  `user_id` int(11) NOT NULL COMMENT '操作人ID',
  `msg` varchar(5000) DEFAULT NULL COMMENT '操作说明',
  `data` varchar(5000) DEFAULT NULL COMMENT '参数',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `logs` */

/*Table structure for table `notice` */

DROP TABLE IF EXISTS `notice`;

CREATE TABLE `notice` (
  `id` int(11) DEFAULT NULL COMMENT '公告ID',
  `department_id` int(11) NOT NULL COMMENT '机构ID',
  `title` char(100) DEFAULT NULL COMMENT '标题',
  `content` longtext COMMENT '公告内容',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '发布时间',
  `range` enum('1','2') DEFAULT '1' COMMENT '通知范围（1所有部门，2指定部门）',
  `inform_id` char(1) DEFAULT NULL COMMENT '指定部门时的部门ids',
  `state` enum('draft','publish') DEFAULT 'draft' COMMENT '状态（draft草稿publish发布）',
  `delete` int(11) DEFAULT '0' COMMENT '删除（0未删1已删）',
  PRIMARY KEY (`department_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `notice` */

/*Table structure for table `notice_file` */

DROP TABLE IF EXISTS `notice_file`;

CREATE TABLE `notice_file` (
  `file_id` int(11) NOT NULL COMMENT '通知附件ID',
  `notice_id` int(11) NOT NULL COMMENT '通知ID',
  PRIMARY KEY (`file_id`,`notice_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `notice_file` */

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '人员ID',
  `department_id` int(11) NOT NULL COMMENT '部门ID',
  `name` char(50) DEFAULT NULL COMMENT '姓名',
  `user_code` char(20) NOT NULL COMMENT '员工号',
  `class` char(20) DEFAULT NULL COMMENT '民族',
  `sex` enum('0','1','2') NOT NULL DEFAULT '0' COMMENT '性别（0保密1女2男）',
  `id_card` char(18) DEFAULT NULL COMMENT '身份证',
  `add_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  `update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `delete` tinyint(1) DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`,`department_id`,`user_code`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `user` */

/*Table structure for table `user_job` */

DROP TABLE IF EXISTS `user_job`;

CREATE TABLE `user_job` (
  `user_id` int(11) NOT NULL COMMENT '人员ID',
  `job` char(20) NOT NULL COMMENT '岗位',
  PRIMARY KEY (`user_id`,`job`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `user_job` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
