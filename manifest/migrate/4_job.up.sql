SET FOREIGN_KEY_CHECKS = 0;

ALTER TABLE `sys_job` DROP COLUMN `job_group`,
MODIFY COLUMN `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'cron执行表达式' AFTER `invoke_target`,
ADD COLUMN `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '任务类型：0 系统' AFTER `status`,
ADD COLUMN `job_params` json NULL COMMENT '自定义任务参数' AFTER `type`,
DROP PRIMARY KEY,
ADD PRIMARY KEY (`job_id`) USING BTREE,
ADD UNIQUE INDEX `idx_name`(`job_name`) USING BTREE COMMENT '任务名称唯一索引';

ALTER TABLE `sys_job_log` DROP COLUMN `job_group`,
ADD COLUMN `job_id` bigint(20) NOT NULL COMMENT '对应任务 ID' AFTER `job_log_id`;

UPDATE sys_menu SET api_path="put/monitor/job/changeStatus" WHERE menu_id = 20105;
INSERT INTO `sys_menu`  VALUES (20107, '任务执行一次', 201, 7, '#', 'put/monitor/job/run', '', '', 1, 0, 'F', '0', '0', 'monitor:job:run', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (20108, '任务日志列表', 201, 8, '#', 'get/monitor/jobLog/list', '', '', 1, 0, 'F', '0', '0', 'monitor:jobLog:list', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (20109, '删除任务日志', 201, 9, '#', 'delete/monitor/job/{jobLogId}', '', '', 1, 0, 'F', '0', '0', 'monitor:jobLog:remove', '#', 'admin', sysdate(), '', NULL, '');

SET FOREIGN_KEY_CHECKS = 1;