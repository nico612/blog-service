

# 创建数据库，
# 默认字符集为: utf8mb4 是一种支持更广泛的Unicode字符集的字符集，可以用于存储包括Emoji表情在内的各种字符。
# 默认排序规则utf8mb4_general_ci，这意味着该数据库中的所有表和列都将使用该字符集和排序规则。utf8mb4_general_ci是一种不区分大小写的排序规则。
CREATE DATABASE IF NOT EXISTS blog_service DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

# 公共字段
#   `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
#   `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
#   `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
#   `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
#   `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
#   `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',

# 创建标签表
# ENGINE：是一个MySQL数据库中的表选项，用于指定表所使用的存储引擎，存储引擎是负责管理数据存储和检索的核心组件，它决定了如何组织、存储和操作表中的数据。
#         InnoDB 是MySQL中最常用的存储引擎之一。它提供了许多高级功能，如事务支持、行级锁定、外键约束和崩溃恢复能力。这些功能使得InnoDB非常适合于处理事务性和并发性要求较高的应用程序。
#默认字符集： CHARSET=utf8mb4，  字符集定义了数据库中可以使用的字符集和编码方式。 utf8mb4 是一种支持更广泛的Unicode字符集的字符集，可以用于存储包括Emoji表情在内的各种字符
# 使用当前默认时间：`created_on` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
CREATE TABLE IF NOT EXISTS `blog_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) DEFAULT '' COMMENT '标签',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT  '' COMMENT  '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del`    tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 未删除 1 已删除',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 为禁用，1 为启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';


# 创建文章表
CREATE TABLE IF NOT EXISTS `blog_article`(
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(100) DEFAULT '' COMMENT '文章标题',
    `desc` varchar(255) DEFAULT '' COMMENT '文章简述',
    `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
    `content` longtext COMMENT '文章内容',
    `name` VARCHAR(100) DEFAULT '' COMMENT '标签',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT  '' COMMENT  '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del`    tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 未删除 1 已删除',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 为禁用，1 为启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

# 创建文章标签关联表
# 文章和标签之间的1对多的关联关系
CREATE TABLE IF NOT EXISTS `blog_article_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id` int(11) NOT NULL COMMENT '文章id',
    `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签ID, 有可能存在没有标签的情况',
    `content` longtext COMMENT '文章内容',
    `name` VARCHAR(100) DEFAULT '' COMMENT '标签',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT  '' COMMENT  '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';

# 创建JWT认证表

CREATE TABLE `blog_auth` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_key` varchar(20) DEFAULT '' COMMENT '用于签发的认证信息的Key',
    `app_secret` varchar(50) DEFAULT '' COMMENT '用与签发认证信息的Secret',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT  '' COMMENT  '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del`    tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 未删除 1 已删除',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='认证管理';

# 插入认证信息
# INSERT INTO `blog_service`.blog_auth
#     (`id`,  `app_key`, `app_secret`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`)
# VALUES
#     (1, 'eddycjy', 'go-programming-tour-book', 0, 'eddycjy', 0, 'eddycjy', 0);