# 小羊试驾SheeDrive_开发文档

<img src="..\程序截图\title_SheeDrive.png" width="80%" height="auto">

> 本文档建立于 2024-01-25 ，用于记录本人毕业设计的开发过程

## 数据库设计
1. 使用 MySQL 创建名为`sheedrive`的数据库，字符集采用`utf8mb4_general_ci`

<img src="..\程序截图\PixPin_2024-01-23_08-52-48.png" width="40%" height="auto">

2. 运行如下 SQL 代码，建表并插入测试数据
```sql
/*
SheeDrive小羊试驾 数据库

Author: 汤晟
Date: 2024-01-23

Note: 
1. 用户表、订单表、地址表使用雪花算法，其他表的主键都自增
2. 用户手机号将作为用户名
3. 因为现实场景需要驾驶证等信息，用户姓名和身份证号为必填项，并且可以通过切字符串的形式获取性别和生日
4. 经销商主营可以通过limit库存表获得
*/
USE `sheedrive`;

SET FOREIGN_KEY_CHECKS=0;

/*Table structure for table `admin` */
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(32) COLLATE utf8_bin NOT NULL COMMENT '姓名',
  `username` varchar(32) COLLATE utf8_bin NOT NULL COMMENT '用户名',
  `password` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '密码',
  `avatar` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '头像',
  `phone` varchar(11) COLLATE utf8_bin NOT NULL COMMENT '手机号',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态 0:禁用, 1:正常',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_user` bigint(20) NOT NULL COMMENT '创建人',
  `update_user` bigint(20) NOT NULL COMMENT '修改人',
  `delete_user` bigint(20) DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='管理员信息';

INSERT INTO `admin` VALUES (1, '超级管理员', 'admin', '123456', null, '13001801111', 1, NOW(), NOW(), null, 1, 1, null);
INSERT INTO `admin` VALUES (2, '超级管理员2', 'admin2', '123456', null, '13001802222', 1, NOW(), NOW(), null, 1, 1, null);

/*Table structure for table `user` */
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `last_name` varchar(32) COLLATE utf8_bin NOT NULL COMMENT '姓',
  `first_name` varchar(32) COLLATE utf8_bin NOT NULL COMMENT '名',
  `username` varchar(32) COLLATE utf8_bin NOT NULL COMMENT '用户名',
  `password` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '密码',
  `avatar` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '头像',
  `phone` varchar(11) COLLATE utf8_bin NOT NULL COMMENT '手机号',
  `id_number` varchar(18) COLLATE utf8_bin NOT NULL COMMENT '身份证号',
  `sex` varchar(2) COLLATE utf8_bin DEFAULT NULL COMMENT '性别',
  `birthday` datetime DEFAULT NULL COMMENT '生日',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态 0:禁用, 1:正常',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_user` bigint(20) NOT NULL COMMENT '创建人',
  `update_user` bigint(20) NOT NULL COMMENT '修改人',
  `delete_user` bigint(20) DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='用户信息';

INSERT INTO `user` VALUES (1, '汤', '日成', '15001807369', '123456', null, '15001807369', '310115200207179212', '男', '2002-07-17', '1', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `user` VALUES (2, '张', '三', '15001801111', '123456', null, '15001801111', '310115200101011111', '男', '2001-01-01', '1', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `user` VALUES (3, '李', '四', '15001802222', '123456', null, '15001802222', '310115200101022222', '女', '2002-01-02', '1', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `user` VALUES (4, '王', '五', '15001803333', '123456', null, '15001803333', '310115200101031111', '男', '2002-01-03', '1', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `user` VALUES (5, '赵', '六', '15001804444', '123456', null, '15001804444', '310115200101042222', '女', '2002-01-04', '1', NOW(), NOW(), null, 1, 1, null);

/*Table structure for table `dealer` */
DROP TABLE IF EXISTS `dealer`;
CREATE TABLE `dealer`(
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(32) COLLATE utf8_bin NOT NULL COMMENT '名称',
  `username` varchar(32) COLLATE utf8_bin NOT NULL COMMENT '用户名',
  `password` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '密码',
  `avatar` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '头像',
  `phone` varchar(11) COLLATE utf8_bin NOT NULL COMMENT '手机号',
  `describe_info` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '描述信息',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态 0:禁用, 1:正常',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_user` bigint(20) NOT NULL COMMENT '创建人',
  `update_user` bigint(20) NOT NULL COMMENT '修改人',
  `delete_user` bigint(20) DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='经销商信息';

INSERT INTO `dealer` VALUES (1, '上海子鼠汽车', 'zishu', '123456', null, '15001801111', '五星经销商，好评不断！', 1, NOW(), NOW(), null, 1, 1, null);
INSERT INTO `dealer` VALUES (2, '上海丑牛汽车', 'chouniu', '123456', null, '15001802222', '百万客户，好评连连！', 1, NOW(), NOW(), null, 1, 1, null);
INSERT INTO `dealer` VALUES (3, '上海寅虎汽车', 'yinhu', '123456', null, '15001803333', '您身边的汽车管家！', 1, NOW(), NOW(), null, 1, 1, null);
INSERT INTO `dealer` VALUES (4, '上海卯兔汽车', 'maotu', '123456', null, '15001804444', '尊享服务，就在你家门口！', 1, NOW(), NOW(), null, 1, 1, null);
INSERT INTO `dealer` VALUES (5, '上海辰龙汽车', 'chenlong', '123456', null, '15001805555', '优质好车，选择辰龙！', 1, NOW(), NOW(), null, 1, 1, null);

/*Table structure for table `address` */
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `belong_id` bigint(20) NOT NULL COMMENT '所属ID',
  `belong_category` int(11) NOT NULL COMMENT '所属分类 1:经销商,2:用户',
  `latitude` double NOT NULL COMMENT '纬度',
  `longitude` double NOT NULL COMMENT '经度',
  `accuracy` double NOT NULL COMMENT '精度',
  `country` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '国家',
  `province` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '省',
  `city` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '市',
  `district` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '区',
  `street` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '街道',
  `street_number` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '门牌号',
  `poi_name` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT 'POI信息',
  `postal_code` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '邮政编码',
  `city_code` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '城市代码',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_user` bigint(20) NOT NULL COMMENT '创建人',
  `update_user` bigint(20) NOT NULL COMMENT '修改人',
  `delete_user` bigint(20) DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='地址簿';

/*Table structure for table `car_detail` */
DROP TABLE IF EXISTS `car_detail`;
CREATE TABLE `car_detail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `year` year NOT NULL COMMENT '年份',
  `brand` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '品牌',
  `model` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '型号',
  `version` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '版本',
  `image` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '图片',
  `category` int(11) NOT NULL COMMENT '类型 0:其他, 1:轿车, 2:SUV, 3:MPV, 4:卡车, 5:跑车',
  `color` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '颜色',
  `price` bigint(20) NOT NULL COMMENT '指导价',
  `type` int(11) NOT NULL COMMENT '类型 0:其他, 1:纯电动, 2:插电混动, 3:增程, 4:汽油, 5:汽油+48V轻混系统, 6:油电混动, 7:柴油',
  `seats` int(11) NOT NULL COMMENT '座位数 0:7座以上 1:1座, 2:2座, 4:4座, 5:5座, 6:6座, 7:7座',
  `describe_info` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '描述信息',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_user` bigint(20) NOT NULL COMMENT '创建人',
  `update_user` bigint(20) NOT NULL COMMENT '修改人',
  `delete_user` bigint(20) DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='汽车细节表';

INSERT INTO `car_detail` VALUES (1, 2024, '大众', '帕萨特', '商务版', null, '1', '黑', 181900, 4, 5, '2023款 大众帕萨特 280TSI 商务版', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `car_detail` VALUES (2, 2024, '大众', '帕萨特', '豪华版', null, '1', '黑', 227300, 4, 5, '2023款 大众帕萨特 380TSI 豪华版', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `car_detail` VALUES (3, 2024, '奥迪', 'A6L', '45 TFSI 甄选致雅版', null, '1', '黑', 454900, 4, 5, '2023款 奥迪A6L 45 TFSI 甄选致雅版', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `car_detail` VALUES (4, 2024, '奥迪', 'A6L', '55 TFSI quattro 旗舰致雅型', null, '1', '黑', 656800, 4, 5, '2023款 奥迪A6L 55 TFSI quattro 旗舰致雅型', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `car_detail` VALUES (5, 2024, '保时捷', 'Cayenne', '3.0T', null, '2', '白', 948000, 4, 5, '2024款 保时捷 Cayenne 3.0T', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `car_detail` VALUES (6, 2023, '本田', '雅阁', '卓越版', null, '1', '黑', 213800, 4, 5, '2023款 本田雅阁 锐T动 260TURBO 卓越版', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `car_detail` VALUES (7, 2023, '本田', '雅阁', '卓越版', null, '1', '灰', 213800, 4, 5, '2023款 本田雅阁 锐T动 260TURBO 卓越版', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `car_detail` VALUES (8, 2023, '本田', 'CR-V', 'CVT两驱锋尚7座版', null, '2', '红', 208900, 4, 7, '2023款 本田CR-V 240TURBO CVT 两驱锋尚7座版', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `car_detail` VALUES (9, 2023, '理想', 'L9', 'Pro', null, '2', '银', 429800, 3, 7, '2023款 理想L9 Pro', NOW(), NOW(), null, 1, 1, null);
INSERT INTO `car_detail` VALUES (10, 2022, '理想', 'L9', 'Max', null, '2', '银', 459800, 3, 7, '2022款 理想L9 Max', NOW(), NOW(), null, 1, 1, null);

/*Table structure for table `stock` */
DROP TABLE IF EXISTS `stock`;
CREATE TABLE `stock` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `dealer_id` bigint(20) NOT NULL COMMENT '经销商ID',
  `car_id` bigint(20) NOT NULL COMMENT '车辆ID',
  `stock` int(11) NOT NULL COMMENT '库存数',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_user` bigint(20) NOT NULL COMMENT '创建人',
  `update_user` bigint(20) NOT NULL COMMENT '修改人',
  `delete_user` bigint(20) DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='库存表';

INSERT INTO `stock` VALUES (1, 1, 1, 10, NOW(), NOW(), null, 1, 1, null);
INSERT INTO `stock` VALUES (2, 1, 2, 10, NOW(), NOW(), null, 1, 1, null);
INSERT INTO `stock` VALUES (3, 1, 3, 10, NOW(), NOW(), null, 1, 1, null);
INSERT INTO `stock` VALUES (4, 2, 4, 10, NOW(), NOW(), null, 1, 1, null);
INSERT INTO `stock` VALUES (5, 3, 5, 10, NOW(), NOW(), null, 1, 1, null);
INSERT INTO `stock` VALUES (6, 3, 6, 10, NOW(), NOW(), null, 1, 1, null);

/*Table structure for table `order` */
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `dealer_id` bigint(20) NOT NULL COMMENT '经销商ID',
  `car_id` bigint(20) NOT NULL COMMENT '车辆ID',
  `addr_id` bigint(20) NOT NULL COMMENT '用户地址ID',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '订单状态 -1:异常,0:取消,1:未确认,2:已确认,3:签署协议,4:试驾中,5:试驾结束,6:待评价,7:已评价',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_user` bigint(20) NOT NULL COMMENT '创建人',
  `update_user` bigint(20) NOT NULL COMMENT '修改人',
  `delete_user` bigint(20) DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='订单表';

/*Table structure for table `comment` */
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `order_id` bigint(20) NOT NULL COMMENT '订单ID',
  `content` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '评价内容',
  `total_score` int(11) NOT NULL COMMENT '总评分 1~5星',
  `dealer_score` int(11) NOT NULL COMMENT '经销商评分 1~5星',
  `car_score` int(11) NOT NULL COMMENT '汽车评分 1~5星',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_user` bigint(20) NOT NULL COMMENT '创建人',
  `delete_user` bigint(20) DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`) USING BTREE 
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='评价表';
```

## 后端接口开发
后端使用基于 Go 语言的 GoFrame 框架开发
### 准备工作
下面简单介绍准备工作，详细操作省略：
1. 安装 Go 语言开发环境，并配置环境变量
2. 安装 GoFrame 框架
3. 修改 Go 语言开发环境

### 创建项目
在指定目录下使用如下命令初始化项目，指定项目名称`sheedrive`
```bash
gf init sheedrive -u
```

1. 使用 VSCode 打开项目，删除`api`、`controller`中自动生成的文件，删除`cmd.go`中的路由绑定

2. 添加依赖的库：`mysql`驱动和`jwt`
```bash
go get github.com/gogf/gf/contrib/drivers/mysql/v2
```
```bash
go get github.com/golang-jwt/jwt/v5
```

3. 在`main.go`中初始化导入`mysql`驱动 
```go
import (
    _ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)
```

4. 在`manifest/config/config.yaml`中配置数据库连接
```yaml
database:
  link: "mysql:root:2002@tcp(127.0.0.1:3306)/sheedrive?loc=Local&parseTime=true"
```


### 代码生成
数据库相关的操作与数据结构放在`dao`与`model`中。在GoFrame中，`dao`与`model`的内容可以自动生成
1. 在`hack/config.yaml`中配置`dao`
```yaml
gfcli:
  gen:
    dao:
      link: "mysql:root:2002@tcp(localhost:3306)/sheedrive?loc=Local&parseTime=True"
      tables: "admin,user,dealer,address,car_detail,stock,order,comment"
```
其中，`link`表示数据库连接的URL；`tables`表示需要生成的`dao`和`model`的数据表，多个表用逗号隔开

2. 在命令行中执行该命令生成`dao`和`model`的代码
```bash
gf gen dao
```

### 管理员相关功能
管理员实体具有的功能如下：

|一级功能|二级子功能|功能描述|
|:--:|:--:|:--|
|管理员登录|登录|管理员在B端系统使用用户名和密码进行登录操作|
|管理员登录|登出|管理员登录系统后点击“退出”按钮实现登出操作|
|管理员管理|添加管理员|管理员在“管理员管理”菜单项中点击“添加”按钮后，填入表单数据后确定添加管理员|
|管理员管理|删除管理员|管理员在“管理员管理”菜单项中点击“删除”按钮后确定删除管理员|
|管理员管理|修改管理员|管理员在“管理员管理”菜单项中点击“修改”按钮后，填入表单数据后确定修改管理员|
|管理员管理|分页查询管理员|管理员进入“管理员管理”菜单项或执行增删改操作后，执行分页查询管理员列表|
|管理员管理|关键字查询管理员|管理员在“管理员管理”菜单项的查询输入框中填入搜索关键字，查询指定管理员列表|
|经销商管理|添加经销商|管理员在“经销商管理”菜单项中点击“添加”按钮后，填入表单数据后确定添加经销商|
|经销商管理|删除经销商|管理员在“经销商管理”菜单项中点击“删除”按钮后确定删除经销商|
|经销商管理|修改经销商|管理员在“经销商管理”菜单项中点击“修改”按钮后，填入表单数据后确定修改经销商|
|经销商管理|分页查询经销商|管理员进入“经销商管理”菜单项或执行增删改操作后，执行分页查询经销商列表|
|经销商管理|关键字查询经销商|管理员在“经销商管理”菜单项的查询输入框中填入搜索关键字，查询指定经销商列表|
|用户管理|添加用户|管理员在“用户管理”菜单项中点击“添加”按钮后，填入表单数据后确定添加用户|
|用户管理|删除用户|管理员在“用户管理”菜单项中点击“删除”按钮后确定删除用户|
|用户管理|修改用户|管理员在“用户管理”菜单项中点击“修改”按钮后，填入表单数据后确定修改用户|
|用户管理|分页查询用户|管理员进入“用户管理”菜单项或执行增删改操作后，执行分页查询用户列表|
|用户管理|关键字查询用户|管理员在“用户管理”菜单项的查询输入框中填入搜索关键字，查询指定用户列表|
|车辆信息管理|添加车辆信息|管理员在“车辆信息管理”菜单项中点击“添加”按钮后，填入表单数据后确定添加车辆信息|
|车辆信息管理|删除车辆信息|管理员在“车辆信息管理”菜单项中点击“删除”按钮后确定删除车辆信息|
|车辆信息管理|修改车辆信息|管理员在“车辆信息管理”菜单项中点击“修改”按钮后，填入表单数据后确定修改车辆信息|
|车辆信息管理|分页查询车辆信息|管理员进入“车辆信息管理”菜单项或执行增删改操作后，执行分页查询车辆信息列表|
|车辆信息管理|关键字查询车辆信息|管理员在“车辆信息管理”菜单项的查询输入框中填入搜索关键字，查询指定车辆信息列表|
|订单管理|添加订单|管理员在“订单管理”菜单项中点击“添加”按钮后，填入表单数据后确定添加订单|
|订单管理|删除订单|管理员在“订单管理”菜单项中点击“删除”按钮后确定删除订单|
|订单管理|修改订单|管理员在“订单管理”菜单项中点击“修改”按钮后，填入表单数据后确定修改订单|
|订单管理|分页查询订单|管理员进入“订单管理”菜单项或执行增删改操作后，执行分页查询订单列表|
|订单管理|关键字查询订单|管理员在“订单管理”菜单项的查询输入框中填入搜索关键字，查询指定订单列表|
|评论管理|添加评论|管理员在“评论管理”菜单项中点击“添加”按钮后，填入表单数据后确定添加评论|
|评论管理|删除评论|管理员在“评论管理”菜单项中点击“删除”按钮后确定删除评论|
|评论管理|修改评论|管理员在“评论管理”菜单项中点击“修改”按钮后，填入表单数据后确定修改评论|
|评论管理|分页查询评论|管理员进入“评论管理”菜单项或执行增删改操作后，执行分页查询评论列表|
|评论管理|关键字查询评论|管理员在“评论管理”菜单项的查询输入框中填入搜索关键字，查询指定评论列表|
|地址管理|添加地址|管理员在“地址管理”菜单项中点击“添加”按钮后，填入表单数据后确定添加地址|
|地址管理|删除地址|管理员在“地址管理”菜单项中点击“删除”按钮后确定删除地址|
|地址管理|修改地址|管理员在“地址管理”菜单项中点击“修改”按钮后，填入表单数据后确定修改地址|
|地址管理|分页查询地址|管理员进入“地址管理”菜单项或执行增删改操作后，执行分页查询地址列表|
|地址管理|关键字查询地址|管理员在“地址管理”菜单项的查询输入框中填入搜索关键字，查询指定地址列表|
|库存管理|添加库存|管理员在“库存管理”菜单项中点击“添加”按钮后，填入表单数据后确定添加库存|
|库存管理|删除库存|管理员在“库存管理”菜单项中点击“删除”按钮后确定删除库存|
|库存管理|修改库存|管理员在“库存管理”菜单项中点击“修改”按钮后，填入表单数据后确定修改库存|
|库存管理|分页查询库存|管理员进入“库存管理”菜单项或执行增删改操作后，执行分页查询库存列表|
|库存管理|关键字查询库存|管理员在“库存管理”菜单项的查询输入框中填入搜索关键字，查询指定库存列表|

#### API数据结构
在`api/admin/admin.go`中编写 API 数据结构






































































