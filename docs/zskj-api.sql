/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1-3306
 Source Server Type    : MySQL
 Source Server Version : 80020
 Source Host           : localhost:3306
 Source Schema         : zskj_api

 Target Server Type    : MySQL
 Target Server Version : 80020
 File Encoding         : 65001

 Date: 08/12/2020 12:27:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `config_id` int NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) DEFAULT '' COMMENT '参数键值',
  `config_type` char(1) DEFAULT 'N' COMMENT '系统内置（Y是 N否）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`config_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='参数配置表';

-- ----------------------------
-- Records of sys_config
-- ----------------------------
BEGIN;
INSERT INTO `sys_config` VALUES (2, '用户管理-账号初始密码', 'sys.user.initPassword', '123456', '1', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '初始化密码 123456');
INSERT INTO `sys_config` VALUES (4, '静态资源网盘存储', 'sys.resource.url', '/static', '1', 'admin', '2020-02-18 20:10:33', '', '2020-03-23 20:51:39', 'public目录下的静态资源存储到OSS/COS等网盘，将键值设为/public表示本地');
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` bigint NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint DEFAULT '0' COMMENT '父部门id',
  `ancestors` varchar(50) DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) DEFAULT '' COMMENT '部门名称',
  `order_num` int DEFAULT '0' COMMENT '显示顺序',
  `leader` varchar(20) DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `status` char(1) DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`dept_id`)
) ENGINE=InnoDB AUTO_INCREMENT=125 DEFAULT CHARSET=utf8 COMMENT='部门表';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` VALUES (100, 0, '0', '蓉易购', 0, 'admin', '13982673816', '1005650533@qq.com', '0', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2020-09-25 17:37:20');
INSERT INTO `sys_dept` VALUES (110, 100, '0,100', '技术部', 2, '', '', '', '0', '0', 'admin', '2019-12-02 17:07:02', 'admin', '2020-09-25 17:49:48');
INSERT INTO `sys_dept` VALUES (112, 100, '0,100', '监管部', 1, '曾尚兵', '18788996255', 'ddd@163.com', '0', '0', 'admin', '2020-03-21 16:30:26', 'admin', '2020-11-17 00:29:09');
INSERT INTO `sys_dept` VALUES (113, 110, '0,100,110', '应用运维', 2, 'もう一度', '', '', '0', '0', 'admin', '2020-09-20 11:29:21', 'admin', '2020-09-25 17:47:51');
INSERT INTO `sys_dept` VALUES (116, 110, '0,100,110', '系统运维', 2, 'もう一度', '', '', '0', '0', 'admin', '2020-09-25 17:40:56', 'admin', '2020-09-25 17:48:02');
INSERT INTO `sys_dept` VALUES (117, 110, '0,100,110', '数据库运维', 3, 'もう一度', '', '', '0', '0', 'admin', '2020-09-25 17:41:29', 'admin', '2020-09-25 17:48:11');
INSERT INTO `sys_dept` VALUES (118, 110, '0,100,110', '运维安全', 4, 'もう一度', '', '', '0', '0', 'admin', '2020-09-25 17:42:02', 'admin', '2020-09-25 17:48:18');
INSERT INTO `sys_dept` VALUES (119, 100, '0,100', '财务部', 3, 'もう一度', '', '', '0', '0', 'admin', '2020-09-25 17:44:37', 'admin', '2020-09-25 17:49:57');
INSERT INTO `sys_dept` VALUES (120, 100, '0,100', '人事部', 2, 'もう一度', '13881887716', 'walkerr@163.com', '0', '0', 'admin', '2020-09-25 17:45:22', '', '2020-11-16 15:51:15');
INSERT INTO `sys_dept` VALUES (121, 100, '0,100', '市场部', 1, 'もう一度', '13881887710', '320553500@qq.com', '0', '0', 'admin', '2020-09-25 17:45:57', 'admin', '2020-11-14 22:17:29');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict` (
  `dict_id` int NOT NULL AUTO_INCREMENT,
  `dict_type` varchar(32) DEFAULT '',
  `dict_sort` int DEFAULT '0',
  `dict_label` varchar(64) DEFAULT '',
  `dict_value` varchar(32) DEFAULT '',
  `parent_id` int DEFAULT '0',
  `list_class` varchar(64) DEFAULT '',
  `is_default` int DEFAULT '0',
  `status` int DEFAULT '0',
  `create_by` varchar(64) DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_by` varchar(64) DEFAULT '',
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `remark` varchar(255) DEFAULT '',
  PRIMARY KEY (`dict_id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_dict
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict` VALUES (1, '001', 1, '市州地区', '0', 0, 'default', 0, 0, 'admin', '2020-10-23 14:26:43', '', '2020-10-27 06:16:09', '地区');
INSERT INTO `sys_dict` VALUES (2, '00101', 1, '重庆市', '1', 1, 'primary', 0, 0, 'admin', '2020-10-23 14:26:43', '', '2020-10-27 06:56:15', '');
INSERT INTO `sys_dict` VALUES (3, '002', 2, '系统开关', '0', 0, 'default', 0, 0, 'admin', '2020-10-23 14:26:43', '', '2020-10-27 06:24:51', '');
INSERT INTO `sys_dict` VALUES (4, '00201', 1, '正常', '0', 3, 'primary', 1, 0, 'admin', '2020-10-23 14:26:43', '', '2020-10-27 06:23:24', '');
INSERT INTO `sys_dict` VALUES (5, '00202', 2, '停用', '1', 3, 'danger', 0, 0, 'admin', '2020-10-23 14:26:43', '', '2020-10-27 07:10:36', '');
INSERT INTO `sys_dict` VALUES (6, '003', 3, '性别', '0', 0, 'default', 0, 0, 'admin', '2020-10-23 14:26:43', '', '2020-10-27 07:37:09', '');
INSERT INTO `sys_dict` VALUES (7, '00301', 1, '男', '1', 6, 'default', 1, 0, 'admin', '2020-10-23 14:26:43', '', '2020-10-27 07:38:46', '');
INSERT INTO `sys_dict` VALUES (8, '00302', 2, '女', '0', 6, 'primary', 0, 0, 'admin', '2020-10-23 14:26:43', '', '2020-10-27 07:38:49', '');
INSERT INTO `sys_dict` VALUES (9, '004', 4, '菜单状态', '0', 0, '', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (10, '00401', 1, '显示', '0', 9, 'primary', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (11, '00402', 2, '隐藏', '1', 9, 'warning', 0, 0, '', '2020-10-23 14:26:43', '', '2020-11-16 15:44:49', '11');
INSERT INTO `sys_dict` VALUES (12, '00102', 2, '泸州市', '2', 1, '', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (13, '00103', 3, '广安市', '3', 1, '', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (14, '00104', 4, '成都市', '4', 1, 'default', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (15, '00105', 5, '西昌市', '5', 1, '', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (16, '005', 5, '系统是否', '0', 0, 'default', 0, 0, 'admin', '2020-10-23 14:26:43', '', '2020-11-14 05:22:27', '');
INSERT INTO `sys_dict` VALUES (17, '00501', 1, '是', '1', 16, 'primary', 0, 0, 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (18, '00502', 2, '否', '0', 16, 'warning', 0, 0, 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (19, '006', 6, '任务分组', '0', 0, '', 0, 0, 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (20, '00601', 1, '默认', '0', 19, 'default', 1, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (21, '00602', 2, '系统', '1', 19, 'primary', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (22, '007', 0, '操作类型', '0', 0, '', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (23, '00701', 1, '新增', '1', 22, 'primary', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (24, '00702', 2, '修改', '2', 22, 'primary', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (25, '00703', 3, '删除', '3', 22, 'primary', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (26, '00704', 4, '授权', '4', 22, 'primary', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (27, '00705', 5, '导出', '5', 22, 'primary', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (28, '00706', 6, '导入', '6', 22, 'primary', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (29, '00707', 7, '强退', '7', 22, 'primary', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (30, '00708', 8, '清空数据', '8', 22, 'primary', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (31, '008', 0, '操作状态', '0', 0, '', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (32, '00801', 1, '成功', '0', 31, 'primary', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (33, '00802', 2, '失败', '1', 31, 'danger', 0, 0, '', '2020-11-16 17:19:58', '', '2020-11-16 17:26:47', '');
INSERT INTO `sys_dict` VALUES (39, '00709', 9, '其他', '0', 22, 'primary', 0, 0, '', '2020-11-15 11:30:29', '', '2020-11-15 11:30:49', '');
INSERT INTO `sys_dict` VALUES (40, '009', 9, '菜单类型', '0', 0, 'default', 1, 0, 'admin', '2020-11-17 16:01:07', '', '2020-11-17 08:03:01', '菜单类型');
INSERT INTO `sys_dict` VALUES (41, '00901', 2, '系统菜单', '2', 40, 'default', 1, 0, 'admin', '2020-11-17 16:01:32', '', '2020-11-17 08:03:15', '');
INSERT INTO `sys_dict` VALUES (42, '00901', 1, '业务菜单', '1', 40, 'default', 1, 0, 'admin', '2020-11-17 16:03:43', '', '2020-11-17 16:03:43', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job` (
  `job_id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` varchar(200) DEFAULT '' COMMENT '参数',
  `job_group` varchar(64) NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` varchar(20) DEFAULT '1' COMMENT '计划执行策略（1多次执行 2执行一次）',
  `concurrent` char(1) DEFAULT '1' COMMENT '是否并发执行（0允许 1禁止）',
  `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1暂停）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`job_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 COMMENT='定时任务调度表';

-- ----------------------------
-- Records of sys_job
-- ----------------------------
BEGIN;
INSERT INTO `sys_job` VALUES (10, 'test1', '', '0', 'test1', '0 9,18 * * *', '2', '1', '1', 'admin', '2020-02-26 15:30:27', '', '2020-03-24 16:12:46', 'qqq');
INSERT INTO `sys_job` VALUES (12, 'test2', 'helloworld|yjgo1', '0', 'test2', '@every 1s', '1', '1', '1', 'admin', '2020-02-27 10:20:26', 'admin', '2020-11-15 18:17:03', '11');
COMMIT;

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log` (
  `info_id` bigint NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) DEFAULT '' COMMENT '操作系统',
  `status` char(1) DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) DEFAULT '' COMMENT '提示消息',
  `login_time` datetime DEFAULT NULL COMMENT '访问时间',
  PRIMARY KEY (`info_id`)
) ENGINE=InnoDB AUTO_INCREMENT=985 DEFAULT CHARSET=utf8 COMMENT='系统访问记录';

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
BEGIN;
INSERT INTO `sys_login_log` VALUES (978, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2020-11-28 12:15:27');
INSERT INTO `sys_login_log` VALUES (979, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2020-11-29 12:12:26');
INSERT INTO `sys_login_log` VALUES (980, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2020-12-02 14:12:20');
INSERT INTO `sys_login_log` VALUES (981, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2020-12-04 09:36:22');
INSERT INTO `sys_login_log` VALUES (982, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '1', '账号或密码不正确', '2020-12-04 12:32:38');
INSERT INTO `sys_login_log` VALUES (983, 'admin', '::1', '内网IP', 'Chrome', 'Windows 10', '0', '登陆成功', '2020-12-04 12:32:56');
INSERT INTO `sys_login_log` VALUES (984, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_7', '0', '登陆成功', '2020-12-05 14:25:13');
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menu_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `url` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `parent_id` int DEFAULT '0',
  `sort_id` int DEFAULT '0',
  `perms` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `sys_type` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `menu_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `icon` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `visible` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=295 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` VALUES (1, '系统管理', '#', 0, 4, '', '2', 'M', 'menu-icon fa fa-home blue', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (2, '系统监控', '#', 0, 5, '', '2', 'M', 'menu-icon fa fa-leaf brown', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (3, '日志管理', '#', 2, 6, '', '2', 'M', 'menu-icon fa fa-book blue', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (4, '系统用户', '/system/user', 1, 1, 'system:user:view', '2', 'C', 'menu-icon fa fa-users green', '0', 'admin', '2020-11-16 17:19:58', 'admin', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (5, '用户查询', '#', 4, 1, 'system:user:list', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (6, '用户增加', '#', 4, 2, 'system:user:add', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (7, '用户修改', '#', 4, 3, 'system:user:edit', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (8, '用户删除', '#', 4, 4, 'system:user:remove', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (9, '用户导出', '#', 4, 5, 'system:user:export', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (10, '重置密码', '#', 4, 6, 'system:user:resetPwd', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (11, '操作日志', '/monitor/operlog', 3, 2, 'monitor:operlog:view', '2', 'C', 'menu-icon fa fa-leaf brown', '0', 'admin', '2020-11-16 17:19:58', 'admin', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (12, '登陆日志', '/monitor/loginlog', 3, 3, 'monitor:loginlog:view', '2', 'C', 'menu-icon fa fa-leaf brown', '0', 'admin', '2020-11-16 17:19:58', 'admin', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (13, '岗位管理', '/system/post', 1, 5, 'system:post:view', '2', 'C', 'menu-icon fa fa-leaf brown', '0', 'admin', '2020-11-16 17:19:58', 'admin', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (14, '部门管理', '/system/dept', 1, 4, 'system:dept:view', '2', 'C', 'menu-icon fa fa-leaf brown', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (15, '参数设置', '/system/config', 1, 7, 'system:config:view', '2', 'C', 'menu-icon fa fa-leaf brown', '0', 'admin', '2020-11-16 17:19:58', 'admin', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (16, '定时任务', '/monitor/job', 2, 2, 'monitor:job:view', '2', 'C', 'menu-icon fa fa-leaf brown', '0', 'admin', '2020-11-16 17:19:58', 'admin', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (17, '服务监控', '/monitor/server', 2, 3, 'monitor:server:view', '2', 'C', 'menu-icon fa fa-leaf brown', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (38, '菜单管理', '/system/menu', 1, 3, 'system:menu:view', '2', 'C', 'menu-icon fa fa-folder-open-o brown', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (39, '角色管理', '/system/role', 1, 2, 'system:role:view', '2', 'C', 'menu-icon fa fa-download orange', '0', 'admin', '2020-11-16 17:19:58', 'admin', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (43, '数据字典', '/system/dict', 1, 6, 'system:dict:view', '2', 'C', 'menu-icon fa fa-book red', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (180, '字典查询', '#', 43, 1, 'system:dict:list', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (181, '字典增加', '#', 43, 2, 'system:dict:add', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (182, '字典修改', '#', 43, 3, 'system:dict:edit', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (183, '字典删除', '#', 43, 4, 'system:dict:remove', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (184, '菜单查询', '#', 38, 1, 'system:menu:list', '2', 'F', ' #', '0', 'admin', '2020-11-16 17:19:58', 'admin', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (185, '菜单增加', '#', 38, 2, 'system:menu:add', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (186, '菜单修改', '#', 38, 3, 'system:menu:edit', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (187, '菜单删除', '#', 38, 4, 'system:menu:remove', '2', 'F', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (199, '在线用户', '/monitor/online', 2, 1, 'monitor:online:view', '2', 'C', 'menu-icon fa fa-laptop green', '0', 'admin', '2020-11-16 17:19:58', 'admin', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (200, '部门查询', '/system/dept/list', 14, 1, 'system:dept:list', '2', 'F', 'menu-icon fa fa-users orange', '0', 'admin', '2020-11-14 11:53:10', 'admin', '2020-11-14 03:53:43', '');
INSERT INTO `sys_menu` VALUES (201, '部门增加', '/system/dept/add', 14, 2, 'system:dept:add', '2', 'F', '#', '0', 'admin', '2020-11-14 12:26:02', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (202, '部门修改', '/system/dept/edit', 14, 3, 'system:dept:edit', '2', 'F', '#', '0', 'admin', '2020-11-14 12:26:36', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (203, '部门删除', '/system/dept/remove', 14, 4, 'system:dept:remove', '2', 'F', '#', '0', 'admin', '2020-11-14 12:27:22', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (204, '角色增加', '/system/role/add', 39, 2, 'system:role:add', '2', 'F', '#', '0', 'admin', '2020-11-14 12:38:32', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (205, '角色查询', '/system/role/list', 39, 1, 'system:role:list', '2', 'F', '#', '0', 'admin', '2020-11-14 12:39:04', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (206, '角色修改', '/system/role/edit', 39, 3, 'system:role:edit', '2', 'F', '#', '0', 'admin', '2020-11-14 12:39:33', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (207, '角色删除', '/system/role/remove', 39, 4, 'system:role:remove', '2', 'F', '#', '0', 'admin', '2020-11-14 12:40:06', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (208, '岗位查询', '/system/post/list', 13, 1, 'system:post:list', '2', 'F', '#', '0', 'admin', '2020-11-14 13:05:02', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (209, '岗位增加', '/system/post/add', 13, 2, 'system:post:add', '2', 'F', '#', '0', 'admin', '2020-11-14 13:05:32', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (210, '岗位修改', '/system/post/edit', 13, 3, 'system:post:edit', '2', 'F', '#', '0', 'admin', '2020-11-14 13:06:07', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (211, '岗位删除', '/system/post/remove', 13, 4, 'system:post:remove', '2', 'F', '#', '0', 'admin', '2020-11-14 13:06:39', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (212, '参数查询', '/system/config/list', 15, 1, 'system:config:list', '2', 'F', '#', '0', 'admin', '2020-11-14 13:25:58', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (213, '参数增加', '/system/config/add', 15, 2, 'system:config:add', '2', 'F', '#', '0', 'admin', '2020-11-14 13:26:24', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (214, '参数修改', '/system/config/edit', 15, 3, 'system:config:edit', '2', 'F', '#', '0', 'admin', '2020-11-14 13:26:55', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (215, '参数删除', '/system/config/remove', 15, 4, 'system:config:remove', '2', 'F', '#', '0', 'admin', '2020-11-14 13:27:22', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (216, '在线查询', '/monitor/online/list', 199, 1, 'monitor:online:list', '2', 'F', '#', '0', 'admin', '2020-11-14 13:51:45', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (217, '任务查询', '/monitor/job/list', 16, 1, 'monitor:job:list', '2', 'F', '#', '0', 'admin', '2020-11-14 14:07:48', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (218, '任务增加', '/monitor/job/add', 16, 2, 'monitor:job:add', '2', 'F', '#', '0', 'admin', '2020-11-14 14:20:12', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (219, '任务修改', '/monitor/job/edit', 16, 3, 'monitor:job:edit', '2', 'F', '#', '0', 'admin', '2020-11-14 14:20:41', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (220, '任务删除', '/monitor/job/remove', 16, 4, 'monitor:job:remove', '2', 'F', '#', '0', 'admin', '2020-11-14 14:21:10', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (221, '操作日志查询', '/monitor/operlog/list', 11, 1, 'monitor:operlog:list', '2', 'F', '#', '0', 'admin', '2020-11-14 14:37:17', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (222, '操作日志删除', '/monitor/operlog/remove', 11, 2, 'monitor:operlog:remove', '2', 'F', '#', '0', 'admin', '2020-11-14 14:38:07', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (223, '登陆日志查询', '/monitor/loginlog/list', 12, 1, 'monitor:loginlog:list', '2', 'F', '#', '0', 'admin', '2020-11-14 14:51:38', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (224, '登陆日志删除', '/monitor/loginlog/remove', 12, 2, 'monitor:loginlog:remove', '2', 'F', '#', '0', 'admin', '2020-11-14 14:52:15', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (225, '详细信息', '/monitor/operlog/detail', 11, 3, 'monitor:operlog:detail', '2', 'F', '#', '0', 'admin', '2020-11-15 16:48:17', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (226, '账户解锁', '/monitor/loginlog/unlock', 12, 3, 'monitor:loginlog:unlock', '2', 'F', '#', '0', 'admin', '2020-11-15 17:12:31', '', '2020-11-16 17:19:58', '');
INSERT INTO `sys_menu` VALUES (227, '状态修改', '/monitor/job/changeStatus', 16, 5, 'monitor:job:changeStatus', '2', 'F', ' #', '0', 'admin', '2020-11-15 17:59:49', 'admin', '2020-11-15 10:00:20', '');
INSERT INTO `sys_menu` VALUES (228, '任务详情', '/monitor/job/detail', 16, 6, 'monitor:job:detail', '2', 'F', ' #', '0', 'admin', '2020-11-15 18:00:59', 'admin', '2020-11-15 10:01:26', '');
INSERT INTO `sys_menu` VALUES (229, '任务导出', '/monitor/job/export', 16, 7, 'monitor:job:export', '2', 'F', ' #', '0', 'admin', '2020-11-15 18:01:50', 'admin', '2020-11-15 10:02:04', '');
INSERT INTO `sys_menu` VALUES (230, '批量强退', '/monitor/online/batchForceLogout', 199, 2, 'monitor:online:batchForceLogout', '2', 'F', ' #', '0', 'admin', '2020-11-16 19:13:12', 'admin', '2020-11-16 11:17:33', '');
INSERT INTO `sys_menu` VALUES (231, '单条强退出', '/monitor/online/forceLogout', 199, 3, 'monitor:online:forceLogout', '2', 'F', ' #', '0', 'admin', '2020-11-16 19:13:43', 'admin', '2020-11-16 11:14:15', '');
INSERT INTO `sys_menu` VALUES (232, '实例演示', '#', 0, 5, '', '2', 'M', 'menu-icon fa fa-leaf green', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:30:01', '');
INSERT INTO `sys_menu` VALUES (233, '表单演示', '#', 232, 1, '', '2', 'M', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (234, '表格演示', '#', 232, 2, '', '2', 'M', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (235, '弹框演示', '#', 232, 3, '', '2', 'M', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (236, '操作演示', '#', 232, 4, '', '2', 'M', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (237, '报表演示', '#', 232, 5, '', '2', 'M', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (238, '图标演示', '#', 232, 6, '', '2', 'M', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:20:43', '');
INSERT INTO `sys_menu` VALUES (239, '栅格演示', '/demo/form/grid', 233, 1, 'demo:grid:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:22:12', '');
INSERT INTO `sys_menu` VALUES (240, '下拉框', '/demo/form/select', 233, 2, 'demo:select:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:22:18', '');
INSERT INTO `sys_menu` VALUES (241, '时间轴', '/demo/form/timeline', 233, 3, 'demo:timeline:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:22:23', '');
INSERT INTO `sys_menu` VALUES (242, '基本表单', '/demo/form/basic', 233, 4, 'demo:basic:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:22:28', '');
INSERT INTO `sys_menu` VALUES (243, '卡片列表', '/demo/form/cards', 233, 5, 'demo:cards:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:22:33', '');
INSERT INTO `sys_menu` VALUES (244, '功能扩展', '/demo/form/jasny', 233, 6, 'demo:jasny:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:22:39', '');
INSERT INTO `sys_menu` VALUES (245, '拖动排序', '/demo/form/sortable', 233, 7, 'demo:sortable:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:22:43', '');
INSERT INTO `sys_menu` VALUES (246, '选项卡&面板', '/demo/form/tabs_panels', 233, 8, 'demo:tabs_panels:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:22:48', '');
INSERT INTO `sys_menu` VALUES (247, '表单校验', '/demo/form/validate', 233, 9, 'demo:validate:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:22:53', '');
INSERT INTO `sys_menu` VALUES (248, '表单向导', '/demo/form/wizard', 233, 10, 'demo:wizard:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:22:58', '');
INSERT INTO `sys_menu` VALUES (249, '文件上传', '/demo/form/upload', 233, 11, 'demo:upload:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:23:03', '');
INSERT INTO `sys_menu` VALUES (250, '日期和时间', '/demo/form/datetime', 233, 12, 'demo:datetime:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:23:08', '');
INSERT INTO `sys_menu` VALUES (251, '富文本编辑器', '/demo/form/summernote', 233, 13, 'demo:summernote:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:23:13', '');
INSERT INTO `sys_menu` VALUES (252, '左右互选', '/demo/form/duallistbox', 233, 14, 'demo:duallistbox:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:23:17', '');
INSERT INTO `sys_menu` VALUES (253, '按钮演示', '/demo/form/button', 233, 15, 'demo:button:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:23:19', '');
INSERT INTO `sys_menu` VALUES (254, '数据汇总', '/demo/table/footer', 234, 2, 'demo:footer:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:23:24', '');
INSERT INTO `sys_menu` VALUES (255, '组合表头', '/demo/table/groupHeader', 234, 3, 'demo:groupHeader:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:23:29', '');
INSERT INTO `sys_menu` VALUES (256, '记住翻页', '/demo/table/remember', 234, 4, 'demo:remember:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:23:38', '');
INSERT INTO `sys_menu` VALUES (257, '跳转指定页', '/demo/table/pageGo', 234, 5, 'demo:pageGo:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:23:42', '');
INSERT INTO `sys_menu` VALUES (258, '查询参数', '/demo/table/params', 234, 6, 'demo:params:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:23:44', '');
INSERT INTO `sys_menu` VALUES (259, '点击加载表格', '/demo/table/button', 234, 7, 'demo:button:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:23:53', '');
INSERT INTO `sys_menu` VALUES (260, '表格冻结列', '/demo/table/fixedColumns', 234, 8, 'demo:fixedColumns:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:23:57', '');
INSERT INTO `sys_menu` VALUES (261, '触发事件', '/demo/table/event', 234, 9, 'demo:event:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:24:02', '');
INSERT INTO `sys_menu` VALUES (262, '细节视图', '/demo/table/detail', 234, 10, 'demo:detail:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:24:03', '');
INSERT INTO `sys_menu` VALUES (263, '父子视图', '/demo/table/child', 234, 11, 'demo:child:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:24:10', '');
INSERT INTO `sys_menu` VALUES (264, '图片预览', '/demo/table/image', 234, 12, 'demo:image:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:24:15', '');
INSERT INTO `sys_menu` VALUES (265, '动态增删改查', '/demo/table/curd', 234, 13, 'demo:curd:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:24:16', '');
INSERT INTO `sys_menu` VALUES (266, '表格拖曳', '/demo/table/recorder', 234, 14, 'demo:recorder:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:24:25', '');
INSERT INTO `sys_menu` VALUES (267, '行内编辑', '/demo/table/editable', 234, 15, 'demo:editable:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:24:31', '');
INSERT INTO `sys_menu` VALUES (268, '其它操作', '/demo/table/other', 234, 16, 'demo:other:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:24:35', '');
INSERT INTO `sys_menu` VALUES (269, '查询条件', '/demo/table/search', 234, 1, 'demo:search:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:25:01', '');
INSERT INTO `sys_menu` VALUES (270, '弹层组件', '/demo/modal/layer', 235, 2, 'demo:layer:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:25:05', '');
INSERT INTO `sys_menu` VALUES (271, '弹层表格', '/demo/modal/table', 235, 3, 'demo:table:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:25:10', '');
INSERT INTO `sys_menu` VALUES (272, '模态窗口', '/demo/modal/dialog', 235, 1, 'demo:dialog:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:25:14', '');
INSERT INTO `sys_menu` VALUES (273, '其他操作', '/demo/operate/other', 236, 2, 'demo:other:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:25:20', '');
INSERT INTO `sys_menu` VALUES (274, '表格操作', '/demo/operate/table', 236, 1, 'demo:table:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:25:21', '');
INSERT INTO `sys_menu` VALUES (275, 'Peity', '/demo/report/metrics', 237, 2, 'demo:metrics:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:25:26', '');
INSERT INTO `sys_menu` VALUES (276, 'SparkLine', '/demo/report/peity', 237, 3, 'demo:peity:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:25:31', '');
INSERT INTO `sys_menu` VALUES (277, '图表组合', '/demo/report/sparkline', 237, 4, 'demo:sparkline:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:25:38', '');
INSERT INTO `sys_menu` VALUES (278, '百度Echarts', '/demo/report/echarts', 237, 1, 'demo:echarts:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:25:43', '');
INSERT INTO `sys_menu` VALUES (279, 'Glyphicons', '/demo/icon/glyphicons', 238, 2, 'demo:glyphicons:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:25:48', '');
INSERT INTO `sys_menu` VALUES (280, 'Font Awesome', '/demo/icon/fontawesome', 238, 1, 'demo:fontawesome:view', '2', 'C', '#', '0', 'admin', '2020-11-16 17:19:58', '', '2020-11-16 17:25:50', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice` (
  `notice_id` int NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `notice_title` varchar(50) NOT NULL COMMENT '公告标题',
  `notice_type` char(1) NOT NULL COMMENT '公告类型（1通知 2公告）',
  `notice_content` varchar(2000) DEFAULT NULL COMMENT '公告内容',
  `status` char(1) DEFAULT '0' COMMENT '公告状态（0正常 1关闭）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`notice_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通知公告表';

-- ----------------------------
-- Records of sys_notice
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log` (
  `oper_id` bigint NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) DEFAULT '' COMMENT '模块标题',
  `business_type` int DEFAULT '0' COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) DEFAULT '' COMMENT '请求方式',
  `operator_type` int DEFAULT '0' COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(255) DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) DEFAULT '' COMMENT '操作地点',
  `oper_param` varchar(2000) DEFAULT '' COMMENT '请求参数',
  `json_result` varchar(2000) DEFAULT '' COMMENT '返回参数',
  `status` int DEFAULT '0' COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1020 DEFAULT CHARSET=utf8 COMMENT='操作日志记录';

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
BEGIN;
INSERT INTO `sys_oper_log` VALUES (931, '操作日志管理', 3, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/operlog/clean', '127.0.0.1', '内网IP', '\"all\"', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}', 0, '', '2020-11-24 14:26:35');
INSERT INTO `sys_oper_log` VALUES (932, '登陆日志管理', 3, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/loginlog/remove', '127.0.0.1', '内网IP', '{\"Ids\":\"977,976,975,974,973,972,971,970,969,968,967\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":11,\"otype\":3}', 0, '', '2020-11-24 14:26:40');
INSERT INTO `sys_oper_log` VALUES (933, '角色管理', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/system/role/edit', '127.0.0.1', '内网IP', '{\"RoleId\":3,\"RoleName\":\"普通用户\",\"RoleKey\":\"putongyonghu\",\"RoleSort\":\"2\",\"Status\":\"0\",\"Remark\":\"普通用户\",\"MenuIds\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":3,\"otype\":2}', 0, '', '2020-11-24 14:30:38');
INSERT INTO `sys_oper_log` VALUES (934, '角色管理', 3, 'POST', 'POST', 1, 'admin', '系统运维', '/system/role/remove', '127.0.0.1', '内网IP', '{\"Ids\":\"12\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":3}', 0, '', '2020-11-24 14:30:51');
INSERT INTO `sys_oper_log` VALUES (935, '修改用户', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/system/user/edit', '127.0.0.1', '内网IP', '{\"UserId\":31,\"UserName\":\"五阿哥！\",\"Phonenumber\":\"13881887762\",\"Email\":\"WAG198321@qq.com\",\"DeptId\":121,\"Sex\":\"1\",\"Status\":\"0\",\"RoleIds\":\"3\",\"PostIds\":\"1\",\"Remark\":\"管理员，有扣分权限\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":31,\"otype\":2}', 0, '', '2020-11-24 14:39:35');
INSERT INTO `sys_oper_log` VALUES (936, '角色管理', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/system/role/edit', '127.0.0.1', '内网IP', '{\"RoleId\":1,\"RoleName\":\"管理员\",\"RoleKey\":\"admin\",\"RoleSort\":\"1\",\"Status\":\"0\",\"Remark\":\"管理员\",\"MenuIds\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":1,\"otype\":2}', 0, '', '2020-11-24 14:42:05');
INSERT INTO `sys_oper_log` VALUES (937, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"eeeeee\",\"JobParams\":\"eeee\",\"JobGroup\":\"0\",\"InvokeTarget\":\"eeeeeeeee\",\"CronExpression\":\"eeeeeee\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"eeeee\"}', '{\"code\":500,\"msg\":\"Key: \'AddReq.Status\' Error:Field validation for \'Status\' failed on the \'required\' tag\",\"otype\":1}', 1, '', '2020-11-28 20:10:04');
INSERT INTO `sys_oper_log` VALUES (938, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test3\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test3\",\"CronExpression\":\"* * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"111\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":13,\"otype\":1}', 0, '', '2020-11-28 20:33:12');
INSERT INTO `sys_oper_log` VALUES (939, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test4\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test4\",\"CronExpression\":\"30 * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"1\"}', '{\"code\":500,\"msg\":\"操作失败\",\"otype\":1}', 1, '', '2020-11-28 20:58:54');
INSERT INTO `sys_oper_log` VALUES (940, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test4\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test4\",\"CronExpression\":\"30 * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"1\"}', '{\"code\":500,\"msg\":\"当前task目录下没有绑定这个方法\",\"otype\":1}', 1, '', '2020-11-28 20:59:30');
INSERT INTO `sys_oper_log` VALUES (941, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test2\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test4\",\"CronExpression\":\"30 * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":14,\"otype\":1}', 0, '', '2020-11-28 20:59:44');
INSERT INTO `sys_oper_log` VALUES (942, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test2\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"30 * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"qq\"}', '{\"code\":500,\"msg\":\"任务名称已经存在\",\"otype\":1}', 1, '', '2020-11-28 21:00:43');
INSERT INTO `sys_oper_log` VALUES (943, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test2\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"30 * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"qq\"}', '{\"code\":500,\"msg\":\"任务名称已经存在\",\"otype\":1}', 1, '', '2020-11-28 21:00:47');
INSERT INTO `sys_oper_log` VALUES (944, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test2\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test21\",\"CronExpression\":\"30 * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"qq\"}', '{\"code\":500,\"msg\":\"任务名称已经存在\",\"otype\":1}', 1, '', '2020-11-28 21:00:56');
INSERT INTO `sys_oper_log` VALUES (945, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test21\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test21\",\"CronExpression\":\"30 * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"qq\"}', '{\"code\":500,\"msg\":\"当前task目录下没有绑定这个方法\",\"otype\":1}', 1, '', '2020-11-28 21:01:07');
INSERT INTO `sys_oper_log` VALUES (946, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test21\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"30 * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"qq\"}', '{\"code\":500,\"msg\":\"当前task目录下没有绑定这个方法\",\"otype\":1}', 1, '', '2020-11-28 21:01:14');
INSERT INTO `sys_oper_log` VALUES (947, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test21\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test1\",\"CronExpression\":\"30 * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"qq\"}', '{\"code\":500,\"msg\":\"当前task目录下没有绑定这个方法\",\"otype\":1}', 1, '', '2020-11-28 21:01:32');
INSERT INTO `sys_oper_log` VALUES (948, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test21\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test1\",\"CronExpression\":\"30 * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"qq\"}', '{\"code\":500,\"msg\":\"当前task目录下没有绑定这个方法\",\"otype\":1}', 1, '', '2020-11-28 21:02:31');
INSERT INTO `sys_oper_log` VALUES (949, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test21\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test1\",\"CronExpression\":\"30 * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"qq\"}', '{\"code\":500,\"msg\":\"当前task目录下没有绑定这个方法\",\"otype\":1}', 1, '', '2020-11-28 21:03:17');
INSERT INTO `sys_oper_log` VALUES (950, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test1\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test1\",\"CronExpression\":\"30 * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"qq\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":15,\"otype\":1}', 0, '', '2020-11-28 21:03:33');
INSERT INTO `sys_oper_log` VALUES (951, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test3\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test3\",\"CronExpression\":\"* * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"mm\"}', '{\"code\":500,\"msg\":\"当前task目录下没有绑定这个方法\",\"otype\":1}', 1, '', '2020-11-28 21:05:45');
INSERT INTO `sys_oper_log` VALUES (952, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test2\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"* * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"mm\"}', '{\"code\":500,\"msg\":\"任务名称已经存在\",\"otype\":1}', 1, '', '2020-11-28 21:05:59');
INSERT INTO `sys_oper_log` VALUES (953, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test21\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"* * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"mm\"}', '{\"code\":500,\"msg\":\"当前task目录下没有绑定这个方法\",\"otype\":1}', 1, '', '2020-11-28 21:06:10');
INSERT INTO `sys_oper_log` VALUES (954, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test2\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test21\",\"CronExpression\":\"* * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"mm\"}', '{\"code\":500,\"msg\":\"任务名称已经存在\",\"otype\":1}', 1, '', '2020-11-28 21:06:23');
INSERT INTO `sys_oper_log` VALUES (955, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test21\",\"JobParams\":\"\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"* * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"mm\"}', '{\"code\":500,\"msg\":\"当前task目录下没有绑定这个方法\",\"otype\":1}', 1, '', '2020-11-28 21:06:33');
INSERT INTO `sys_oper_log` VALUES (956, '定时任务管理', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/edit', '127.0.0.1', '内网IP', '{\"JobId\":12,\"JobName\":\"test2\",\"JobParams\":\"helloworld|yjgo1\",\"JobGroup\":\"1\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"@every 4s\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"11\"}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":2}', 0, '', '2020-11-29 09:54:34');
INSERT INTO `sys_oper_log` VALUES (957, '定时任务管理', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/edit', '127.0.0.1', '内网IP', '{\"JobId\":12,\"JobName\":\"test2\",\"JobParams\":\"helloworld|yjgo1\",\"JobGroup\":\"1\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"@every 4s\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"11\"}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":2}', 0, '', '2020-11-29 10:00:02');
INSERT INTO `sys_oper_log` VALUES (958, '定时任务管理', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/edit', '127.0.0.1', '内网IP', '{\"JobId\":12,\"JobName\":\"test2\",\"JobParams\":\"helloworld|yjgo1\",\"JobGroup\":\"1\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"@every 4s\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"11\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":12,\"otype\":2}', 0, '', '2020-11-29 10:10:44');
INSERT INTO `sys_oper_log` VALUES (959, '定时任务管理', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/edit', '127.0.0.1', '内网IP', '{\"JobId\":12,\"JobName\":\"test2\",\"JobParams\":\"helloworld|yjgo1\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"@every 4s\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"11\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":12,\"otype\":2}', 0, '', '2020-11-29 10:10:50');
INSERT INTO `sys_oper_log` VALUES (960, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 10:17:29');
INSERT INTO `sys_oper_log` VALUES (961, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 10:20:11');
INSERT INTO `sys_oper_log` VALUES (962, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 10:21:40');
INSERT INTO `sys_oper_log` VALUES (963, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 10:22:58');
INSERT INTO `sys_oper_log` VALUES (964, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 10:24:49');
INSERT INTO `sys_oper_log` VALUES (965, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 10:30:55');
INSERT INTO `sys_oper_log` VALUES (966, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 10:31:19');
INSERT INTO `sys_oper_log` VALUES (967, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 10:31:28');
INSERT INTO `sys_oper_log` VALUES (968, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 10:32:05');
INSERT INTO `sys_oper_log` VALUES (969, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 10:32:24');
INSERT INTO `sys_oper_log` VALUES (970, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 10:32:51');
INSERT INTO `sys_oper_log` VALUES (971, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 10:33:13');
INSERT INTO `sys_oper_log` VALUES (972, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":500,\"msg\":\"任务已存在\",\"otype\":0}', 1, '', '2020-11-29 10:33:19');
INSERT INTO `sys_oper_log` VALUES (973, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":500,\"msg\":\"任务已存在\",\"otype\":0}', 1, '', '2020-11-29 10:33:23');
INSERT INTO `sys_oper_log` VALUES (974, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 10:33:26');
INSERT INTO `sys_oper_log` VALUES (975, '定时任务管理', 1, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/add', '127.0.0.1', '内网IP', '{\"JobName\":\"test2\",\"JobParams\":\"test2\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"* * * * *\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"\",\"Status\":\"\",\"Remark\":\"看看\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":16,\"otype\":1}', 0, '', '2020-11-29 11:12:51');
INSERT INTO `sys_oper_log` VALUES (976, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 11:13:01');
INSERT INTO `sys_oper_log` VALUES (977, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 11:13:04');
INSERT INTO `sys_oper_log` VALUES (978, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 11:14:00');
INSERT INTO `sys_oper_log` VALUES (979, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 11:14:10');
INSERT INTO `sys_oper_log` VALUES (980, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 11:14:27');
INSERT INTO `sys_oper_log` VALUES (981, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 11:14:39');
INSERT INTO `sys_oper_log` VALUES (982, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 11:14:48');
INSERT INTO `sys_oper_log` VALUES (983, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":500,\"msg\":\"任务已存在\",\"otype\":0}', 1, '', '2020-11-29 11:14:52');
INSERT INTO `sys_oper_log` VALUES (984, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":500,\"msg\":\"任务已存在\",\"otype\":0}', 1, '', '2020-11-29 11:14:58');
INSERT INTO `sys_oper_log` VALUES (985, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 11:16:45');
INSERT INTO `sys_oper_log` VALUES (986, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 11:16:53');
INSERT INTO `sys_oper_log` VALUES (987, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":500,\"msg\":\"任务已存在\",\"otype\":0}', 1, '', '2020-11-29 11:17:01');
INSERT INTO `sys_oper_log` VALUES (988, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 11:20:45');
INSERT INTO `sys_oper_log` VALUES (989, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 11:21:01');
INSERT INTO `sys_oper_log` VALUES (990, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":500,\"msg\":\"任务已存在\",\"otype\":0}', 1, '', '2020-11-29 11:21:04');
INSERT INTO `sys_oper_log` VALUES (991, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":500,\"msg\":\"任务已存在\",\"otype\":0}', 1, '', '2020-11-29 11:21:07');
INSERT INTO `sys_oper_log` VALUES (992, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 11:21:12');
INSERT INTO `sys_oper_log` VALUES (993, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":500,\"msg\":\"任务名称已经存在\",\"otype\":0}', 1, '', '2020-11-29 11:23:07');
INSERT INTO `sys_oper_log` VALUES (994, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":500,\"msg\":\"任务名称已经存在\",\"otype\":0}', 1, '', '2020-11-29 11:23:11');
INSERT INTO `sys_oper_log` VALUES (995, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":500,\"msg\":\"其他任务正在运行\",\"otype\":0}', 1, '', '2020-11-29 11:24:05');
INSERT INTO `sys_oper_log` VALUES (996, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":500,\"msg\":\"任务已存在\",\"otype\":0}', 1, '', '2020-11-29 11:24:15');
INSERT INTO `sys_oper_log` VALUES (997, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":500,\"msg\":\"其他任务正在运行\",\"otype\":0}', 1, '', '2020-11-29 11:24:28');
INSERT INTO `sys_oper_log` VALUES (998, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":500,\"msg\":\"其他任务正在运行\",\"otype\":0}', 1, '', '2020-11-29 11:24:41');
INSERT INTO `sys_oper_log` VALUES (999, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 11:26:39');
INSERT INTO `sys_oper_log` VALUES (1000, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":500,\"msg\":\"Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near \'=? where job_id=?\' at line 1\",\"otype\":0}', 1, '', '2020-11-29 11:32:45');
INSERT INTO `sys_oper_log` VALUES (1001, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 11:34:05');
INSERT INTO `sys_oper_log` VALUES (1002, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":16}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 11:35:25');
INSERT INTO `sys_oper_log` VALUES (1003, '定时任务管理', 3, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/remove', '127.0.0.1', '内网IP', '{\"Ids\":\"16\"}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":3}', 0, '', '2020-11-29 11:48:02');
INSERT INTO `sys_oper_log` VALUES (1004, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 11:51:10');
INSERT INTO `sys_oper_log` VALUES (1005, '定时任务管理', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/edit', '127.0.0.1', '内网IP', '{\"JobId\":12,\"JobName\":\"test2\",\"JobParams\":\"helloworld|yjgo1\",\"JobGroup\":\"0\",\"InvokeTarget\":\"test2\",\"CronExpression\":\"@every 1s\",\"MisfirePolicy\":\"1\",\"Concurrent\":\"1\",\"Status\":\"1\",\"Remark\":\"11\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":12,\"otype\":2}', 0, '', '2020-11-29 11:51:28');
INSERT INTO `sys_oper_log` VALUES (1006, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 11:51:37');
INSERT INTO `sys_oper_log` VALUES (1007, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 11:51:49');
INSERT INTO `sys_oper_log` VALUES (1008, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":10}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 11:52:01');
INSERT INTO `sys_oper_log` VALUES (1009, '保存头像', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/system/user/profile/updateAvatar', '127.0.0.1', '内网IP', '{\"userid\":1}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":2}', 0, '', '2020-11-29 12:07:26');
INSERT INTO `sys_oper_log` VALUES (1010, '保存头像', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/system/user/profile/updateAvatar', '127.0.0.1', '内网IP', '{\"userid\":1}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":2}', 0, '', '2020-11-29 12:07:27');
INSERT INTO `sys_oper_log` VALUES (1011, '保存头像', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/system/user/profile/updateAvatar', '127.0.0.1', '内网IP', '{\"userid\":1}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":2}', 0, '', '2020-11-29 12:07:27');
INSERT INTO `sys_oper_log` VALUES (1012, '保存头像', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/system/user/profile/updateAvatar', '127.0.0.1', '内网IP', '{\"userid\":1}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":2}', 0, '', '2020-11-29 12:07:27');
INSERT INTO `sys_oper_log` VALUES (1013, '保存头像', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/system/user/profile/updateAvatar', '127.0.0.1', '内网IP', '{\"userid\":1}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":2}', 0, '', '2020-11-29 12:07:27');
INSERT INTO `sys_oper_log` VALUES (1014, '保存头像', 2, 'POST', 'POST', 1, 'admin', '系统运维', '/system/user/profile/updateAvatar', '127.0.0.1', '内网IP', '{\"userid\":1}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":2}', 0, '', '2020-11-29 12:08:56');
INSERT INTO `sys_oper_log` VALUES (1015, '字典管理', 3, 'POST', 'POST', 1, 'admin', '系统运维', '/system/dict/remove', '127.0.0.1', '内网IP', '{\"Ids\":\"40\"}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":3}', 0, '', '2020-11-29 12:09:38');
INSERT INTO `sys_oper_log` VALUES (1016, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-11-29 12:09:58');
INSERT INTO `sys_oper_log` VALUES (1017, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-11-29 12:10:41');
INSERT INTO `sys_oper_log` VALUES (1018, '定时任务管理启动', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/start', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"操作成功\",\"otype\":0}', 0, '', '2020-12-03 18:55:31');
INSERT INTO `sys_oper_log` VALUES (1019, '定时任务管理停止', 0, 'POST', 'POST', 1, 'admin', '系统运维', '/monitor/job/stop', '127.0.0.1', '内网IP', '{\"jobId\":12}', '{\"code\":0,\"msg\":\"停止成功\",\"otype\":0}', 0, '', '2020-12-03 18:55:39');
COMMIT;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `post_id` bigint NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) NOT NULL COMMENT '岗位名称',
  `post_sort` int NOT NULL COMMENT '显示顺序',
  `status` char(1) NOT NULL COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='岗位信息表';

-- ----------------------------
-- Records of sys_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_post` VALUES (1, '1001', '产品经理', 1, '0', 'admin', '2018-03-16 11:33:00', 'admin', '2020-11-14 22:35:56', '产品经理');
INSERT INTO `sys_post` VALUES (5, '2001', '项目经理', 2, '0', 'admin', '2020-09-25 17:53:40', '', '2020-11-16 17:19:58', '项目经理');
INSERT INTO `sys_post` VALUES (6, '3001', '交互设计', 3, '0', 'admin', '2020-09-25 17:54:14', '', '2020-11-16 17:19:58', '交互设计');
INSERT INTO `sys_post` VALUES (7, '4001', '市场推广', 4, '0', 'admin', '2020-09-25 17:54:38', '', '2020-11-16 17:19:58', '市场推广');
INSERT INTO `sys_post` VALUES (8, '5001', '运营总监', 5, '0', 'admin', '2020-09-25 17:55:32', '', '2020-11-16 17:19:58', '运营总监');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) NOT NULL COMMENT '角色名称',
  `role_key` varchar(100) NOT NULL COMMENT '角色权限字符串',
  `role_sort` int NOT NULL COMMENT '显示顺序',
  `data_scope` char(1) DEFAULT '1' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `status` char(1) NOT NULL COMMENT '角色状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COMMENT='角色信息表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` VALUES (1, '管理员', 'admin', 1, '1', '0', '0', 'admin', '2018-03-16 11:33:00', '', '2020-11-24 14:42:05', '管理员');
INSERT INTO `sys_role` VALUES (3, '普通用户', 'putongyonghu', 2, '5', '0', '0', 'admin', '2020-03-01 09:13:21', 'admin', '2020-11-24 14:30:38', '普通用户');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` int NOT NULL COMMENT '角色ID',
  `dept_id` int NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`,`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色和部门关联表';

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_dept` VALUES (3, 100);
INSERT INTO `sys_role_dept` VALUES (3, 119);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` int NOT NULL,
  `menu_id` int NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `dept_id` bigint DEFAULT NULL COMMENT '部门ID',
  `login_name` varchar(30) NOT NULL COMMENT '登录账号',
  `user_name` varchar(30) NOT NULL COMMENT '用户昵称',
  `user_type` varchar(2) DEFAULT '00' COMMENT '用户类型（00系统用户）',
  `email` varchar(50) DEFAULT '' COMMENT '用户邮箱',
  `phonenumber` varchar(11) DEFAULT '' COMMENT '手机号码',
  `sex` char(1) DEFAULT '0' COMMENT '用户性别（0男 1女 2未知）',
  `avatar` varchar(100) DEFAULT '' COMMENT '头像路径',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `salt` varchar(20) DEFAULT '' COMMENT '盐加密',
  `status` char(1) DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `login_ip` varchar(50) DEFAULT '' COMMENT '最后登陆IP',
  `login_date` datetime DEFAULT NULL COMMENT '最后登陆时间',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` VALUES (1, 116, 'admin', '超级管理员', '00', 'yunwei@163.com', '13881887781', '1', '/upload/admin1605425941269768000.png', '066bdc9ebfa2bd6b7b9c708970bab386', 'UJyWpT', '0', '0', '127.0.0.1', '2020-12-05 14:25:13', 'admin', '2018-03-16 11:33:00', 'admin', '2020-12-05 06:25:13', '超级管理员');
INSERT INTO `sys_user` VALUES (26, 121, 'zhouye', '客服～公瑾', '00', '1250181129@qq.com', '15888888761', '0', '', '6406d20c0bd3b531fb4afa03a2fa6059', '4Buf2K', '0', '0', '127.0.0.1', '2020-10-19 15:42:23', 'admin', '2020-09-25 18:01:55', 'admin', '2020-11-22 12:44:54', '客服～公瑾');
INSERT INTO `sys_user` VALUES (31, 121, 'WAG198321', '五阿哥！', '00', 'WAG198321@qq.com', '13881887762', '1', '', 'ce0585ff9d5a3382441fa349b1c27021', 'eB8v7L', '0', '0', '', NULL, 'admin', '2020-11-22 20:47:34', 'admin', '2020-11-24 14:39:35', '管理员，有扣分权限');
COMMIT;

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post` (
  `user_id` int NOT NULL,
  `post_id` int NOT NULL,
  PRIMARY KEY (`user_id`,`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_post` VALUES (26, 7);
INSERT INTO `sys_user_post` VALUES (31, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `user_id` int NOT NULL,
  `role_id` int NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` VALUES (26, 3);
INSERT INTO `sys_user_role` VALUES (31, 3);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
