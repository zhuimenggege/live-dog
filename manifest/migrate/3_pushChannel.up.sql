SET FOREIGN_KEY_CHECKS = 0;

ALTER TABLE `push_channel` DROP COLUMN `url`;
ALTER TABLE `push_channel_email` 
ADD UNIQUE INDEX `idx_channel`(`channel_id`) USING HASH COMMENT '渠道 ID 唯一索引';

-- ----------------------------
-- Table structure for push_channel_web
-- ----------------------------
DROP TABLE IF EXISTS `push_channel_web`;
CREATE TABLE `push_channel_web` (
  `id` int(6) NOT NULL AUTO_INCREMENT COMMENT '记录 ID',
  `channel_id` int(6) NOT NULL COMMENT '渠道 ID',
  `url` varchar(255) DEFAULT NULL COMMENT '推送 URL',
  `http_method` varchar(10) DEFAULT NULL COMMENT '请求方式',
  `secret` varchar(100) DEFAULT NULL COMMENT '密钥/token/key',
  `app_id` varchar(50) DEFAULT NULL COMMENT '应用 ID',
  `corp_id` varchar(50) DEFAULT NULL COMMENT '企业 ID',
  `receiver_id` varchar(255) DEFAULT NULL COMMENT '接收人 ID',
  `receiver_type` varchar(20) DEFAULT NULL COMMENT '接收人类型',
  `extra_params` varchar(255) DEFAULT NULL COMMENT '额外参数',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `action_time` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_channel` (`channel_id`) USING HASH COMMENT '渠道 ID 唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='消息渠道推送表（web）';

SET FOREIGN_KEY_CHECKS = 1;
