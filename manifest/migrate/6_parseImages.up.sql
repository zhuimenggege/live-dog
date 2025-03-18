SET FOREIGN_KEY_CHECKS = 0;

ALTER TABLE `media_parse`
ADD COLUMN `images_url` text NULL COMMENT '图集 url' AFTER `music_cover_url`,
ADD COLUMN `images_cover_url` varchar(1000) NULL COMMENT '图集封面 url' AFTER `images_url`;

SET FOREIGN_KEY_CHECKS = 1;