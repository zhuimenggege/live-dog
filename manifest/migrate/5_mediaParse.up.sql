SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for media_parse
-- ----------------------------
DROP TABLE IF EXISTS `media_parse`;
CREATE TABLE `media_parse` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '媒体解析主键 ID',
  `platform` varchar(20) NOT NULL COMMENT '平台',
  `author` varchar(100) NOT NULL COMMENT '作者名称',
  `author_uid` varchar(100) NOT NULL COMMENT '作者 UID',
  `desc` text COMMENT '媒体描述',
  `media_id` varchar(100) NOT NULL COMMENT '媒体 ID',
  `type` varchar(10) NOT NULL COMMENT '媒体类型',
  `video_url` varchar(1000) DEFAULT NULL COMMENT '视频 url',
  `video_cover_url` varchar(1000) DEFAULT NULL COMMENT '视频封面 url',
  `music_url` varchar(1000) DEFAULT NULL COMMENT '音乐 url',
  `music_cover_url` varchar(1000) DEFAULT NULL COMMENT '音乐封面 url',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='媒体解析表';

INSERT INTO `sys_menu` VALUES (306, '媒体解析', 3, 6, 'parse', 'get/media/parse/list', 'live/parse/index', '', 1, 0, 'C', '0', '0', 'live:parse:list', 'online', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu` VALUES (30601, '解析', 306, 1, '', 'post/media/parse', '', '', 1, 0, 'F', '0', '0', 'live:parse:add', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu` VALUES (30602, '详情', 306, 2, '', 'get/media/parse/{id}', '', '', 1, 0, 'F', '0', '0', 'live:parse:get', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu` VALUES (30603, '删除', 306, 3, '', 'delete/media/parse/{id}', '', '', 1, 0, 'F', '0', '0', 'live:parse:delete', '', 'admin', sysdate(), 'admin', null, '');

SET FOREIGN_KEY_CHECKS = 1;