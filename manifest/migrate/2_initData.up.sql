-- ----------------------------
-- 初始化-角色信息表数据
-- ----------------------------
insert into sys_role values(1, '超级管理员',  'admin',  1, 1, 1, '0', '0', 'admin', sysdate(), '', null, '超级管理员');
insert into sys_role values(2, '普通角色',    'common', 2, 2, 1, '0', '0', 'admin', sysdate(), '', null, '普通角色');

-- ----------------------------
-- 初始化-用户和角色关联表数据
-- ----------------------------
insert into sys_user_role values (1, 1);

-- ----------------------------
-- 初始化-字典类型表
-- ----------------------------
insert into sys_dict_type values(1,  '用户性别', 'sys_user_sex',        '0', 'admin', sysdate(), '', null, '用户性别列表');
insert into sys_dict_type values(2,  '菜单状态', 'sys_show_hide',       '0', 'admin', sysdate(), '', null, '菜单状态列表');
insert into sys_dict_type values(3,  '系统开关', 'sys_normal_disable',  '0', 'admin', sysdate(), '', null, '系统开关列表');
insert into sys_dict_type values(4,  '任务状态', 'sys_job_status',      '0', 'admin', sysdate(), '', null, '任务状态列表');
insert into sys_dict_type values(5,  '任务分组', 'sys_job_group',       '0', 'admin', sysdate(), '', null, '任务分组列表');
insert into sys_dict_type values(6,  '系统是否', 'sys_yes_no',          '0', 'admin', sysdate(), '', null, '系统是否列表');
insert into sys_dict_type values(7,  '通知类型', 'sys_notice_type',     '0', 'admin', sysdate(), '', null, '通知类型列表');
insert into sys_dict_type values(8,  '通知状态', 'sys_notice_status',   '0', 'admin', sysdate(), '', null, '通知状态列表');
insert into sys_dict_type values(9,  '操作类型', 'sys_oper_type',       '0', 'admin', sysdate(), '', null, '操作类型列表');
insert into sys_dict_type values(10, '系统状态', 'sys_common_status',   '0', 'admin', sysdate(), '', null, '登录状态列表');
insert into sys_dict_type values(11, '节目类型', 'anchor_show_type',    '0', 'admin', sysdate(), '', null, '主播节目类型');

-- ----------------------------
-- 初始化-字典数据表
-- ----------------------------
insert into sys_dict_data values(1,  1,  '男',       '0',       'sys_user_sex',        '',   '',        'Y', '0', 'admin', sysdate(), '', null, '性别男');
insert into sys_dict_data values(2,  2,  '女',       '1',       'sys_user_sex',        '',   '',        'N', '0', 'admin', sysdate(), '', null, '性别女');
insert into sys_dict_data values(3,  3,  '未知',     '2',       'sys_user_sex',        '',   '',        'N', '0', 'admin', sysdate(), '', null, '性别未知');
insert into sys_dict_data values(4,  1,  '显示',     '0',       'sys_show_hide',       '',   'primary', 'Y', '0', 'admin', sysdate(), '', null, '显示菜单');
insert into sys_dict_data values(5,  2,  '隐藏',     '1',       'sys_show_hide',       '',   'danger',  'N', '0', 'admin', sysdate(), '', null, '隐藏菜单');
insert into sys_dict_data values(6,  1,  '正常',     '0',       'sys_normal_disable',  '',   'primary', 'Y', '0', 'admin', sysdate(), '', null, '正常状态');
insert into sys_dict_data values(7,  2,  '停用',     '1',       'sys_normal_disable',  '',   'danger',  'N', '0', 'admin', sysdate(), '', null, '停用状态');
insert into sys_dict_data values(8,  1,  '正常',     '0',       'sys_job_status',      '',   'primary', 'Y', '0', 'admin', sysdate(), '', null, '正常状态');
insert into sys_dict_data values(9,  2,  '暂停',     '1',       'sys_job_status',      '',   'danger',  'N', '0', 'admin', sysdate(), '', null, '停用状态');
insert into sys_dict_data values(10, 1,  '默认',     'DEFAULT', 'sys_job_group',       '',   '',        'Y', '0', 'admin', sysdate(), '', null, '默认分组');
insert into sys_dict_data values(11, 2,  '系统',     'SYSTEM',  'sys_job_group',       '',   '',        'N', '0', 'admin', sysdate(), '', null, '系统分组');
insert into sys_dict_data values(12, 1,  '是',       'Y',       'sys_yes_no',          '',   'primary', 'Y', '0', 'admin', sysdate(), '', null, '系统默认是');
insert into sys_dict_data values(13, 2,  '否',       'N',       'sys_yes_no',          '',   'danger',  'N', '0', 'admin', sysdate(), '', null, '系统默认否');
insert into sys_dict_data values(14, 1,  '通知',     '1',       'sys_notice_type',     '',   'warning', 'Y', '0', 'admin', sysdate(), '', null, '通知');
insert into sys_dict_data values(15, 2,  '公告',     '2',       'sys_notice_type',     '',   'success', 'N', '0', 'admin', sysdate(), '', null, '公告');
insert into sys_dict_data values(16, 1,  '正常',     '0',       'sys_notice_status',   '',   'primary', 'Y', '0', 'admin', sysdate(), '', null, '正常状态');
insert into sys_dict_data values(17, 2,  '关闭',     '1',       'sys_notice_status',   '',   'danger',  'N', '0', 'admin', sysdate(), '', null, '关闭状态');
insert into sys_dict_data values(18, 99, '其他',     '0',       'sys_oper_type',       '',   'info',    'N', '0', 'admin', sysdate(), '', null, '其他操作');
insert into sys_dict_data values(19, 1,  '新增',     '1',       'sys_oper_type',       '',   'info',    'N', '0', 'admin', sysdate(), '', null, '新增操作');
insert into sys_dict_data values(20, 2,  '修改',     '2',       'sys_oper_type',       '',   'info',    'N', '0', 'admin', sysdate(), '', null, '修改操作');
insert into sys_dict_data values(21, 3,  '删除',     '3',       'sys_oper_type',       '',   'danger',  'N', '0', 'admin', sysdate(), '', null, '删除操作');
insert into sys_dict_data values(22, 4,  '授权',     '4',       'sys_oper_type',       '',   'primary', 'N', '0', 'admin', sysdate(), '', null, '授权操作');
insert into sys_dict_data values(23, 5,  '导出',     '5',       'sys_oper_type',       '',   'warning', 'N', '0', 'admin', sysdate(), '', null, '导出操作');
insert into sys_dict_data values(24, 6,  '导入',     '6',       'sys_oper_type',       '',   'warning', 'N', '0', 'admin', sysdate(), '', null, '导入操作');
insert into sys_dict_data values(25, 7,  '强退',     '7',       'sys_oper_type',       '',   'danger',  'N', '0', 'admin', sysdate(), '', null, '强退操作');
insert into sys_dict_data values(26, 8,  '生成代码', '8',       'sys_oper_type',       '',   'warning', 'N', '0', 'admin', sysdate(), '', null, '生成操作');
insert into sys_dict_data values(27, 9,  '清空数据', '9',       'sys_oper_type',       '',   'danger',  'N', '0', 'admin', sysdate(), '', null, '清空操作');
insert into sys_dict_data values(28, 1,  '成功',     '0',       'sys_common_status',   '',   'primary', 'N', '0', 'admin', sysdate(), '', null, '正常状态');
insert into sys_dict_data values(29, 2,  '失败',     '1',       'sys_common_status',   '',   'danger',  'N', '0', 'admin', sysdate(), '', null, '停用状态');
insert into sys_dict_data values(30, 0,  '歌曲',    '1',        'anchor_show_type',    '',   'default', 'Y', '0', 'admin', sysdate(), '', null, '');
insert into sys_dict_data values(31, 1,  '弹唱',    '2',        'anchor_show_type',    '',   'default', 'N', '0', 'admin', sysdate(), '', null, '');

-- ----------------------------
-- 初始化-参数配置表
-- ----------------------------
insert into sys_config values(1, '主框架页-默认皮肤样式名称',     'sys.index.skinName',            'skin-blue',     'Y', 'admin', sysdate(), '', null, '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow' );
insert into sys_config values(2, '用户管理-账号初始密码',         'sys.user.initPassword',         '123456',        'Y', 'admin', sysdate(), '', null, '初始化密码 123456' );
insert into sys_config values(3, '主框架页-侧边栏主题',           'sys.index.sideTheme',           'theme-dark',    'Y', 'admin', sysdate(), '', null, '深色主题theme-dark，浅色主题theme-light' );
insert into sys_config values(4, '账号自助-验证码开关',           'sys.account.captchaEnabled',    'true',          'Y', 'admin', sysdate(), '', null, '是否开启验证码功能（true开启，false关闭）');
insert into sys_config values(5, '账号自助-是否开启用户注册功能', 'sys.account.registerUser',      'false',         'Y', 'admin', sysdate(), '', null, '是否开启注册用户功能（true开启，false关闭）');
insert into sys_config values(6, '用户登录-黑名单列表',           'sys.login.blackIPList',         '',              'Y', 'admin', sysdate(), '', null, '设置登录IP黑名单限制，多个匹配项以;分隔，支持匹配（*通配、网段）');

-- ----------------------------
-- 初始化-菜单权限表
-- ----------------------------
INSERT INTO `sys_menu`  VALUES (1, '系统管理', 0, 1, 'system', '', NULL, '', 1, 0, 'M', '0', '0', '', 'system', 'admin', sysdate(), '', NULL, '系统管理目录');
INSERT INTO `sys_menu`  VALUES (2, '系统监控', 0, 2, 'monitor', '', NULL, '', 1, 0, 'M', '0', '0', '', 'monitor', 'admin', sysdate(), '', NULL, '系统监控目录');
INSERT INTO `sys_menu`  VALUES (5, '直播管理', 0, 0, 'live', '', '', '', 1, 0, 'M', '0', '0', '', 'druid', 'admin', sysdate(), 'admin', '2024-11-29 15:34:52', '');
INSERT INTO `sys_menu`  VALUES (100, '用户管理', 1, 1, 'user', 'get/system/user/list', 'system/user/index', '', 1, 0, 'C', '0', '0', 'system:user:list', 'user', 'admin', sysdate(), 'admin', null, '用户管理菜单');
INSERT INTO `sys_menu`  VALUES (101, '角色管理', 1, 2, 'role', 'get/system/role/list', 'system/role/index', '', 1, 0, 'C', '0', '0', 'system:role:list', 'peoples', 'admin', sysdate(), '', NULL, '角色管理菜单');
INSERT INTO `sys_menu`  VALUES (102, '菜单管理', 1, 3, 'menu', 'get/system/menu/list', 'system/menu/index', '', 1, 0, 'C', '0', '0', 'system:menu:list', 'tree-table', 'admin', sysdate(), '', NULL, '菜单管理菜单');
INSERT INTO `sys_menu`  VALUES (105, '字典管理', 1, 6, 'dict_type', 'get/system/dict/type/list', 'system/dict/type', '', 1, 0, 'C', '0', '0', 'system:dict:type:list', 'dict', 'admin', sysdate(), 'admin', null, '字典管理菜单');
INSERT INTO `sys_menu`  VALUES (106, '参数设置', 1, 7, 'config', 'get/system/config/list', 'system/config/index', '', 1, 0, 'C', '0', '0', 'system:config:list', 'edit', 'admin', sysdate(), '', NULL, '参数设置菜单');
INSERT INTO `sys_menu`  VALUES (107, '通知公告', 1, 8, 'notice', 'get/system/notice/list', 'system/notice/index', '', 1, 0, 'C', '0', '0', 'system:notice:list', 'message', 'admin', sysdate(), '', NULL, '通知公告菜单');
INSERT INTO `sys_menu`  VALUES (110, '定时任务', 2, 3, 'job', 'get/monitor/job/list', 'monitor/job/index', '', 1, 0, 'C', '0', '0', 'monitor:job:list', 'job', 'admin', sysdate(), 'admin', null, '定时任务菜单');
INSERT INTO `sys_menu`  VALUES (112, '服务监控', 2, 5, 'server', 'get/monitor/server/list', 'monitor/server/index', '', 1, 0, 'C', '0', '0', 'monitor:server:list', 'server', 'admin', sysdate(), 'admin', null, '服务监控菜单');
INSERT INTO `sys_menu`  VALUES (500, '操作日志', 2, 2, 'operlog', 'get/monitor/operlog/list', 'monitor/operlog/index', '', 1, 0, 'C', '0', '0', 'monitor:operlog:list', 'form', 'admin', sysdate(), 'admin', null, '操作日志菜单');
INSERT INTO `sys_menu`  VALUES (1000, '用户查询', 100, 1, '', 'get/system/user/{userId}', '', '', 1, 0, 'F', '0', '0', 'system:user:query', '#', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (1001, '用户新增', 100, 2, '', 'post/system/user', '', '', 1, 0, 'F', '0', '0', 'system:user:add', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1002, '用户修改', 100, 3, '', 'put/system/user', '', '', 1, 0, 'F', '0', '0', 'system:user:edit', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1003, '用户删除', 100, 4, '', 'delete/system/user/{userId}', '', '', 1, 0, 'F', '0', '0', 'system:user:remove', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1004, '用户导出', 100, 5, '', '', '', '', 1, 0, 'F', '0', '0', 'system:user:export', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1005, '用户导入', 100, 6, '', '', '', '', 1, 0, 'F', '0', '0', 'system:user:import', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1006, '重置密码', 100, 7, '', 'put/system/user/resetPwd', '', '', 1, 0, 'F', '0', '0', 'system:user:resetPwd', '#', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (1007, '角色查询', 101, 1, '', 'get/system/role/{roleId}', '', '', 1, 0, 'F', '0', '0', 'system:role:query', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1008, '角色新增', 101, 2, '', 'post/system/role', '', '', 1, 0, 'F', '0', '0', 'system:role:add', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1009, '角色修改', 101, 3, '', 'put/system/role', '', '', 1, 0, 'F', '0', '0', 'system:role:edit', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1010, '角色删除', 101, 4, '', 'delete/system/role/{roleId}', '', '', 1, 0, 'F', '0', '0', 'system:role:remove', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1011, '角色导出', 101, 5, '', '', '', '', 1, 0, 'F', '0', '0', 'system:role:export', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1012, '菜单查询', 102, 1, '', 'get/system/menu/{menuId}', '', '', 1, 0, 'F', '0', '0', 'system:menu:query', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1013, '菜单新增', 102, 2, '', 'post/system/menu', '', '', 1, 0, 'F', '0', '0', 'system:menu:add', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1014, '菜单修改', 102, 3, '', 'put/system/menu', '', '', 1, 0, 'F', '0', '0', 'system:menu:edit', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1015, '菜单删除', 102, 4, '', 'delete/system/menu/{menuId}', '', '', 1, 0, 'F', '0', '0', 'system:menu:remove', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1025, '字典查询', 105, 1, '#', 'get/system/dict/type/{dictId}', '', '', 1, 0, 'F', '0', '0', 'system:dict:query', '#', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (1026, '字典新增', 105, 2, '#', 'post/system/dict/type', '', '', 1, 0, 'F', '0', '0', 'system:dict:add', '#', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (1027, '字典修改', 105, 3, '#', 'put/system/dict/type', '', '', 1, 0, 'F', '0', '0', 'system:dict:edit', '#', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (1028, '字典删除', 105, 4, '#', 'delete/system/dict/type/{dictId}', '', '', 1, 0, 'F', '0', '0', 'system:dict:remove', '#', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (1029, '字典导出', 105, 5, '#', '', '', '', 1, 0, 'F', '0', '0', 'system:dict:export', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1030, '参数查询', 106, 1, '#', 'get/system/config/{configId}', '', '', 1, 0, 'F', '0', '0', 'system:config:query', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1031, '参数新增', 106, 2, '#', 'post/system/config', '', '', 1, 0, 'F', '0', '0', 'system:config:add', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1032, '参数修改', 106, 3, '#', 'put/system/config', '', '', 1, 0, 'F', '0', '0', 'system:config:edit', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1033, '参数删除', 106, 4, '#', 'delete/system/config/{configId}', '', '', 1, 0, 'F', '0', '0', 'system:config:remove', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1034, '参数导出', 106, 5, '#', '', '', '', 1, 0, 'F', '0', '0', 'system:config:export', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1035, '公告查询', 107, 1, '#', 'get/system/notice/{noticeId}', '', '', 1, 0, 'F', '0', '0', 'system:notice:query', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1036, '公告新增', 107, 2, '#', 'post/system/notice', '', '', 1, 0, 'F', '0', '0', 'system:notice:add', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1037, '公告修改', 107, 3, '#', 'put/system/notice', '', '', 1, 0, 'F', '0', '0', 'system:notice:edit', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1038, '公告删除', 107, 4, '#', 'delete/system/notice/{noticeId}', '', '', 1, 0, 'F', '0', '0', 'system:notice:remove', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1039, '操作查询', 500, 1, '#', 'get/monitor/operlog/{operlogId}', '', '', 1, 0, 'F', '0', '0', 'monitor:operlog:query', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1040, '操作删除', 500, 2, '#', 'delete/monitor/operlog/{operlogId}', '', '', 1, 0, 'F', '0', '0', 'monitor:operlog:remove', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1041, '日志导出', 500, 3, '#', '', '', '', 1, 0, 'F', '0', '0', 'monitor:operlog:export', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1049, '任务查询', 110, 1, '#', 'get/monitor/job/{jobId}', '', '', 1, 0, 'F', '0', '0', 'monitor:job:query', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1050, '任务新增', 110, 2, '#', 'post/monitor/job', '', '', 1, 0, 'F', '0', '0', 'monitor:job:add', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1051, '任务修改', 110, 3, '#', 'put/monitor/job', '', '', 1, 0, 'F', '0', '0', 'monitor:job:edit', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1052, '任务删除', 110, 4, '#', 'delete/monitor/job/{jobId}', '', '', 1, 0, 'F', '0', '0', 'monitor:job:remove', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1053, '状态修改', 110, 5, '#', '', '', '', 1, 0, 'F', '0', '0', 'monitor:job:changeStatus', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (1054, '任务导出', 110, 6, '#', '', '', '', 1, 0, 'F', '0', '0', 'monitor:job:export', '#', 'admin', sysdate(), '', NULL, '');
INSERT INTO `sys_menu`  VALUES (2000, '错误', 4, 9, 'index', '', '', '', 1, 0, 'M', '0', '0', '', '404', 'admin', '2023-08-09 14:17:46', 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2001, '字典数据', 1, 5, 'dict_data', 'get/system/dict/data/list', 'system/dict/data', '', 1, 0, 'C', '0', '0', 'system:dict:data:list', 'table', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2003, '分配数据', 101, 6, '', 'put/system/role/dataScope', '', '', 1, 0, 'F', '0', '0', 'system:role:edit', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2004, '菜单更改状态', 101, 9, '', 'put/system/role/changeStatus', '', '', 1, 0, 'F', '0', '0', 'system:role:changestatus', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2005, '分配用户', 101, 9, '', 'get/system/role/authUser/allocatedList', '', '', 1, 0, 'F', '0', '0', 'system:role:list', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2006, '取消授权', 101, 10, '', 'put/system/role/authUser/cancel', '', '', 1, 0, 'F', '0', '0', 'system:role:eidt', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2007, '批量取消授权', 101, 11, '', 'put/system/role/authUser/cancelAll', '', '', 1, 0, 'F', '0', '0', 'system:role:eidt', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2008, '用户列表', 101, 12, '', 'get/system/role/authUser/unallocatedList', '', '', 1, 0, 'F', '0', '0', 'system:role:edit', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2009, '添加用户', 101, 13, '', 'put/system/role/authUser/selectAll', '', '', 1, 0, 'F', '0', '0', 'system:role:edit', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2011, '每日统计', 5, 2, 'daily', 'get/live/daily/list', 'live/daily/index', '', 1, 0, 'C', '0', '0', 'live:daily:list', 'number', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2012, '新增', 2011, 1, '', 'post/live/daily', '', '', 1, 0, 'F', '0', '0', 'live:daily:add', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2013, '删除', 2011, 2, '', 'delete/live/daily', '', '', 1, 0, 'F', '0', '0', 'live:daily:delete', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2014, '房间管理', 5, 1, 'manage', 'get/live/manage/list', 'live/manage/index', '', 1, 0, 'C', '0', '0', 'live:manage:list', 'icon', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2015, '新增', 2014, 1, '', 'post/live/manage', '', '', 1, 0, 'F', '0', '0', 'live:manage:add', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2016, '修改', 2014, 2, '', 'put/live/manage', '', '', 1, 0, 'F', '0', '0', 'live:manage:put', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2017, '删除', 2014, 3, '', 'delete/live/manage', '', '', 1, 0, 'F', '0', '0', 'live:manage:delete', '', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2018, '推送渠道', 1, 4, 'push', 'get/system/push/channel/list', 'system/push/index', '', 1, 0, 'C', '0', '0', 'system:push:list', 'message', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2019, '文件管理', 5, 3, 'file', 'get/file/manage/list', 'live/file/index', '', 1, 0, 'C', '0', '0', 'file:manage:list', 'documentation', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2020, 'Cookie管理', 5, 9, 'cookie', 'get/live/cookie/list', 'live/cookie/index', '', 1, 0, 'C', '0', '0', 'live:cookie:list', 'example', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2021, '个人信息查询', 100, 8, '', 'get/system/user/profile', '', '', 1, 0, 'F', '0', '0', 'system:user:profile', '#', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2022, '个人信息修改', 100, 9, '', 'put/system/user/profile', '', '', 1, 0, 'F', '0', '0', 'system:user:profile', '#', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2023, '个人密码修改', 100, 10, '', 'put/system/user/profile/updatePwd', '', '', 1, 0, 'F', '0', '0', 'system:user:updatePwd', '#', 'admin', sysdate(), 'admin', null, '');
INSERT INTO `sys_menu`  VALUES (2024, '头像修改', 100, 11, '', 'post/system/user/profile/avatar', '', '', 1, 0, 'F', '0', '0', 'system:user:avatar', '#', 'admin', sysdate(), 'admin', null, '');

-- ----------------------------
-- 初始化-角色和菜单关联表数据
-- ----------------------------
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 2021);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 2022);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 2023);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 2024);