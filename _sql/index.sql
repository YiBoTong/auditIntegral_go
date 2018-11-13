# 创建数据库
CREATE DATABASE auditIntegral DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
# 创建字典类型表
CREATE TABLE `auditintegral`.`dictionary_type`( `type_id` INT NOT NULL AUTO_INCREMENT COMMENT '字典类型ID', `title` CHAR(50) NOT NULL COMMENT '字典类型名称', `key` CHAR(20) NOT NULL COMMENT '字典分类', `is_use` BOOL NOT NULL DEFAULT 1 COMMENT '是否启用', `describe` VARCHAR(250) COMMENT '字典类型描述', `delete` BOOL NOT NULL DEFAULT 0 COMMENT '软删除', PRIMARY KEY (`type_id`) );
ALTER TABLE `auditintegral`.`dictionary_type` ADD COLUMN `type` CHAR(50) NOT NULL COMMENT '字典分类' AFTER `title`, CHANGE `key` `key` CHAR(20) CHARSET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字典分类键';
# 创建字典表
CREATE TABLE `auditintegral`.`dictionary`( `id` INT NOT NULL AUTO_INCREMENT COMMENT '字典ID', `type_id` INT NOT NULL COMMENT '字典类型ID', `value` VARCHAR(50) COMMENT '字典值', `title` VARCHAR(50) COMMENT '字典名称', `order` INT(4) DEFAULT 0 COMMENT '排序顺序', `describe` VARCHAR(250) COMMENT '字典描述', PRIMARY KEY (`id`, `type_id`) );