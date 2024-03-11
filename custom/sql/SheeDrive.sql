/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80034
 Source Host           : localhost:3306
 Source Schema         : sheedrive

 Target Server Type    : MySQL
 Target Server Version : 80034
 File Encoding         : 65001

 Date: 11/03/2024 10:00:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address`  (
  `id` bigint(0) NOT NULL COMMENT '主键ID',
  `belong_id` bigint(0) NOT NULL COMMENT '所属ID',
  `belong_category` int(0) NOT NULL COMMENT '所属分类 1:经销商,2:用户',
  `lng_lat` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '经纬度',
  `province` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT '省',
  `city` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT '市',
  `district` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT '区',
  `detail` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '详细地址',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `delete_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '地址簿' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of address
-- ----------------------------
INSERT INTO `address` VALUES (1, 1, 1, '121.505531,31.233544', '上海市', '市辖区', '浦东新区', '银城中路501号', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `address` VALUES (2, 2, 1, '121.494740,31.184338', '上海市', '市辖区', '浦东新区', '上南路205号', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `address` VALUES (3, 3, 1, '121.473701,31.230416', '上海市', '市辖区', '黄浦区', '人民大道200号', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `address` VALUES (4, 4, 1, '121.204669,31.057037', '上海市', '市辖区', '松江区', '龙腾路333号', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `address` VALUES (5, 5, 1, '121.385756,31.238682', '上海市', '市辖区', '普陀区', '真北路1108号', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `address` VALUES (6, 2, 2, '121.518589,31.142639', '上海市', '市辖区', '浦东新区', '永泰路775弄1号', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `address` VALUES (7, 2, 2, '121.518589,31.142639', '上海市', '市辖区', '浦东新区', '永泰路775弄2号', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `address` VALUES (8, 3, 2, '121.505667,31.068941', '上海市', '市辖区', '闵行区', '浦锦街道浦驰南路238弄1号', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `address` VALUES (9, 4, 2, '121.497363,30.959489', '上海市', '市辖区', '奉贤区', '广丰路100弄1号', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '姓名',
  `username` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '用户名',
  `password` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '密码',
  `avatar` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT '头像',
  `phone` varchar(11) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '手机号',
  `status` int(0) NOT NULL DEFAULT 1 COMMENT '状态 0:禁用, 1:正常',
  `isRoot` int(0) NOT NULL DEFAULT 0 COMMENT '是否是超级管理员 0:否, 1:是',
  `token` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT 'token',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `delete_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username`) USING BTREE,
  UNIQUE INDEX `idx_phone`(`phone`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '管理员信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES (1, '超级管理员', 'admin', '14e1b600b1fd579f47433b88e8d85291', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/avatar_chaojiguanliyuan.jpg', '13001801111', 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhZG1pbiIsImV4cCI6MTcxMDIwNTE2NX0.pIvkKgVzpj4XJ1vVYLuO7tRmcan2EKfRavehTbvmFlk', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `admin` VALUES (2, '汤日成', 'TangRiCheng', '14e1b600b1fd579f47433b88e8d85291', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/avatar_tangricheng_r.jpg', '13001802222', 1, 1, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `admin` VALUES (3, '姜宁山', 'JiangNingShan', '14e1b600b1fd579f47433b88e8d85291', NULL, '13001803333', 1, 0, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `admin` VALUES (4, '葛龙', 'GeLong', '14e1b600b1fd579f47433b88e8d85291', NULL, '13001804444', 1, 0, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `admin` VALUES (5, '姚刚会', 'YaoGangHui', '14e1b600b1fd579f47433b88e8d85291', NULL, '13001805555', 1, 0, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `admin` VALUES (6, '柴榕', 'ChaiRong', '14e1b600b1fd579f47433b88e8d85291', NULL, '13001806666', 1, 0, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);

-- ----------------------------
-- Table structure for car_detail
-- ----------------------------
DROP TABLE IF EXISTS `car_detail`;
CREATE TABLE `car_detail`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `year` year NOT NULL COMMENT '年份',
  `brand` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '品牌',
  `model` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '型号',
  `version` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '版本',
  `image` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT '图片',
  `category` int(0) NOT NULL COMMENT '类型 0:其他, 1:轿车, 2:SUV, 3:MPV, 4:卡车, 5:跑车',
  `price` bigint(0) NOT NULL COMMENT '指导价',
  `type` int(0) NOT NULL COMMENT '类型 0:其他, 1:纯电动, 2:插电混动, 3:增程, 4:汽油, 5:汽油+48V轻混系统, 6:油电混动, 7:柴油',
  `seats` int(0) NOT NULL COMMENT '座位数 0:7座以上 1:1座, 2:2座, 4:4座, 5:5座, 6:6座, 7:7座',
  `describe_info` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT '描述信息',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `delete_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '汽车细节表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of car_detail
-- ----------------------------
INSERT INTO `car_detail` VALUES (1, 2024, '大众', '帕萨特', '商务版', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/images/pasate.jpg', 1, 181900, 4, 5, '2023款 大众帕萨特 280TSI 商务版', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (2, 2024, '大众', '帕萨特', '豪华版', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/images/pasate.jpg', 1, 227300, 4, 5, '2023款 大众帕萨特 380TSI 豪华版', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (3, 2024, '奥迪', 'A6L', '45 TFSI 甄选致雅版', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/images/a6.jpg', 1, 454900, 4, 5, '2023款 奥迪A6L 45 TFSI 甄选致雅版', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (4, 2024, '奥迪', 'A6L', '55 TFSI quattro 旗舰致雅型', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/images/a6.jpg', 1, 656800, 4, 5, '2023款 奥迪A6L 55 TFSI quattro 旗舰致雅型', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (5, 2024, '保时捷', 'Cayenne', '3.0T', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/images/kayan.jpg', 2, 948000, 4, 5, '2024款 保时捷 Cayenne 3.0T', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (6, 2023, '本田', '雅阁', '卓越版', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/images/yage.jpg', 1, 213800, 4, 5, '2023款 本田雅阁 锐T动 260TURBO 卓越版', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (7, 2023, '本田', '雅阁', '卓越版', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/images/yage.jpg', 1, 213800, 4, 5, '2023款 本田雅阁 锐T动 260TURBO 卓越版', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (8, 2023, '本田', 'CR-V', 'CVT两驱锋尚7座版', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/images/crv.jpg', 2, 208900, 4, 7, '2023款 本田CR-V 240TURBO CVT 两驱锋尚7座版', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (9, 2023, '理想汽车', 'L9', 'Pro', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/images/l9.jpg', 2, 429800, 3, 7, '2023款 理想L9 Pro', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (10, 2022, '理想汽车', 'L9', 'Max', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/images/l9.jpg', 2, 459800, 3, 7, '2022款 理想L9 Max', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (11, 2024, '沃尔沃', 'S90', 'B5 智逸豪华版', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/images/s90.jpg', 1, 306900, 4, 5, '2024款 沃尔沃S90 B5 智逸豪华版', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (12, 2024, '宝马', '5系', '525Li 豪华套装', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/images/525.jpg', 1, 439900, 5, 5, '2024款 宝马5系 525Li 豪华套装', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (13, 2024, '奥迪', 'A4L', '40 TFSI 时尚动感型', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/car_aodiA4L.jpg', 1, 225300, 4, 5, '2024款 奥迪A4L 40 TFSI 时尚动感型', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (14, 2024, '奥迪', 'Q3', '35 TFSI 进取致雅型', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/car_aodiQ3.jpg', 2, 279800, 4, 5, '2024款 奥迪Q3 35 TFSI 进取致雅型', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `car_detail` VALUES (15, 2024, '奥迪', 'Q3', '35 TFSI 时尚致雅型', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/car_aodiQ3.jpg', 2, 295800, 4, 5, '2024款 奥迪Q3 35 TFSI 时尚致雅型', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint(0) NOT NULL COMMENT '主键ID',
  `order_id` bigint(0) NOT NULL COMMENT '订单ID',
  `content` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '评价内容',
  `total_score` int(0) NOT NULL COMMENT '总评分 1~5星',
  `dealer_score` int(0) NOT NULL COMMENT '经销商评分 1~5星',
  `car_score` int(0) NOT NULL COMMENT '汽车评分 1~5星',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `delete_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_order_id`(`order_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '评价表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1, 1, '服务好，员工热情', 5, 5, 5, '2024-03-11 09:59:42', NULL);

-- ----------------------------
-- Table structure for dealer
-- ----------------------------
DROP TABLE IF EXISTS `dealer`;
CREATE TABLE `dealer`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '名称',
  `username` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '用户名',
  `password` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '密码',
  `avatar` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT '头像',
  `phone` varchar(11) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '手机号',
  `describe_info` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT '描述信息',
  `status` int(0) NOT NULL DEFAULT 1 COMMENT '状态 0:禁用, 1:正常',
  `token` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT 'token',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `delete_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username`) USING BTREE,
  UNIQUE INDEX `idx_phone`(`phone`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '经销商信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dealer
-- ----------------------------
INSERT INTO `dealer` VALUES (1, '上海子鼠汽车', 'zishu', '14e1b600b1fd579f47433b88e8d85291', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/dealer_zishu.jpg', '15001801111', '五星经销商，好评不断！', 1, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `dealer` VALUES (2, '上海丑牛汽车', 'chouniu', '14e1b600b1fd579f47433b88e8d85291', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/dealer_chouniu.jpg', '15001802222', '百万客户，好评连连！', 1, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `dealer` VALUES (3, '上海寅虎汽车', 'yinhu', '14e1b600b1fd579f47433b88e8d85291', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/dealer_yinhu.jpeg', '15001803333', '您身边的汽车管家！', 1, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `dealer` VALUES (4, '上海卯兔汽车', 'maotu', '14e1b600b1fd579f47433b88e8d85291', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/dealer_maotu.jpg', '15001804444', '尊享服务，就在你家门口！', 1, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `dealer` VALUES (5, '上海辰龙汽车', 'chenlong', '14e1b600b1fd579f47433b88e8d85291', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/dealer_chenlong.jpg', '15001805555', '优质好车，选择辰龙！', 1, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` bigint(0) NOT NULL COMMENT '主键ID',
  `user_id` bigint(0) NOT NULL COMMENT '用户ID',
  `dealer_id` bigint(0) NOT NULL COMMENT '经销商ID',
  `car_id` bigint(0) NOT NULL COMMENT '车辆ID',
  `addr_id` bigint(0) NOT NULL COMMENT '用户地址ID',
  `status` int(0) NOT NULL DEFAULT 1 COMMENT '订单状态 -1:异常,0:取消,1:未确认,2:已确认,3:签署协议,4:试驾中,5:试驾结束,6:待评价,7:已评价',
  `order_time` date NOT NULL COMMENT '预定时间',
  `confirm_time` datetime(0) NULL DEFAULT NULL COMMENT '确认时间',
  `sign_time` datetime(0) NULL DEFAULT NULL COMMENT '签署协议时间',
  `start_time` datetime(0) NULL DEFAULT NULL COMMENT '试驾开始时间',
  `end_time` datetime(0) NULL DEFAULT NULL COMMENT '试驾结束时间',
  `precomment_time` datetime(0) NULL DEFAULT NULL COMMENT '试驾待评价时间',
  `comment_time` datetime(0) NULL DEFAULT NULL COMMENT '评价时间',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `delete_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '订单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (1, 2, 1, 3, 6, 7, '2024-03-11', '2024-03-11 09:58:37', '2024-03-11 09:58:40', '2024-03-11 09:58:42', '2024-03-11 09:58:45', '2024-03-11 09:58:47', '2024-03-11 09:58:49', '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `order` VALUES (2, 3, 1, 4, 8, 1, '2024-03-11', NULL, NULL, NULL, NULL, NULL, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `order` VALUES (3, 4, 1, 13, 9, 1, '2024-03-11', NULL, NULL, NULL, NULL, NULL, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `order` VALUES (4, 2, 2, 1, 7, 1, '2024-03-11', NULL, NULL, NULL, NULL, NULL, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);

-- ----------------------------
-- Table structure for stock
-- ----------------------------
DROP TABLE IF EXISTS `stock`;
CREATE TABLE `stock`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `dealer_id` bigint(0) NOT NULL COMMENT '经销商ID',
  `car_id` bigint(0) NOT NULL COMMENT '车辆ID',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uc_dealer_car`(`dealer_id`, `car_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '库存表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of stock
-- ----------------------------
INSERT INTO `stock` VALUES (1, 1, 3, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (2, 1, 4, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (3, 1, 13, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (4, 1, 14, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (5, 1, 15, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (6, 2, 1, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (7, 2, 2, '2024-03-11 09:35:14');
INSERT INTO `stock` VALUES (8, 3, 6, '2024-03-11 09:35:31');
INSERT INTO `stock` VALUES (9, 3, 7, '2024-03-11 09:35:42');
INSERT INTO `stock` VALUES (10, 3, 8, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (11, 4, 3, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (12, 4, 4, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (13, 4, 13, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (14, 5, 3, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (15, 5, 4, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (16, 5, 13, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (17, 5, 14, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (18, 5, 15, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (19, 1, 1, '2024-03-11 08:52:12');
INSERT INTO `stock` VALUES (20, 1, 2, '2024-03-11 08:52:12');

-- ----------------------------
-- Table structure for swiper
-- ----------------------------
DROP TABLE IF EXISTS `swiper`;
CREATE TABLE `swiper`  (
  `id` bigint(0) NOT NULL COMMENT '主键ID',
  `car_id` bigint(0) NOT NULL COMMENT '车辆ID',
  `image_url` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '图片地址',
  `describe_info` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '描述信息',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `delete_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '轮播图表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of swiper
-- ----------------------------
INSERT INTO `swiper` VALUES (1, 3, 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/lunbo-audi.jpg', '交织热血 飞驰人生', '2024-03-11 08:52:12', NULL);
INSERT INTO `swiper` VALUES (2, 12, 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/lunbo-bmw.jpg', '豪华越级，更悦心', '2024-03-11 08:52:12', NULL);
INSERT INTO `swiper` VALUES (3, 11, 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/lunbo-volvo.jpg', '过万境，见心境。', '2024-03-11 08:52:12', NULL);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(0) NOT NULL COMMENT '主键ID',
  `last_name` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '姓',
  `first_name` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '名',
  `username` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '用户名',
  `password` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '密码',
  `avatar` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT '头像',
  `phone` varchar(11) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '手机号',
  `id_number` varchar(18) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '身份证号',
  `sex` varchar(2) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT '性别',
  `birthday` datetime(0) NULL DEFAULT NULL COMMENT '生日',
  `status` int(0) NOT NULL DEFAULT 1 COMMENT '状态 0:禁用, 1:正常',
  `token` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NULL DEFAULT NULL COMMENT 'token',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `delete_time` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username`) USING BTREE,
  UNIQUE INDEX `idx_phone`(`phone`) USING BTREE,
  UNIQUE INDEX `uc_idNumber`(`id_number`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '用户信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (2, '张', '三', '15001801111', '14e1b600b1fd579f47433b88e8d85291', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/user_zhangsan.jpg', '15001801111', '310115200101011111', '男', '2001-01-01 00:00:00', 1, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `user` VALUES (3, '李', '四', '15001802222', '14e1b600b1fd579f47433b88e8d85291', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/user_lisi.jpg', '15001802222', '310115200101022222', '女', '2002-01-02 00:00:00', 1, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `user` VALUES (4, '王', '五', '15001803333', '14e1b600b1fd579f47433b88e8d85291', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/user_wangwu.jpeg', '15001803333', '310115200101031111', '男', '2002-01-03 00:00:00', 1, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);
INSERT INTO `user` VALUES (5, '赵', '六', '15001804444', '14e1b600b1fd579f47433b88e8d85291', 'https://sheedrive.oss-cn-shanghai.aliyuncs.com/sys/user_zhaoliu.jpeg', '15001804444', '310115200101042222', '女', '2002-01-04 00:00:00', 1, NULL, '2024-03-11 08:52:12', '2024-03-11 08:52:12', NULL);

SET FOREIGN_KEY_CHECKS = 1;
