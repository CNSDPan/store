CREATE TABLE `store_users` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`store_user_id` bigint(20) NOT NULL COMMENT '店主ID',
`user_id` bigint(20) NOT NULL COMMENT '用户IID',
`store_id` bigint(20) NOT NULL COMMENT '店铺ID',
`created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
PRIMARY KEY (`id`) USING BTREE,
UNIQUE KEY `unx_suid` (`store_user_id`) USING BTREE,
UNIQUE KEY `unx_u_s_id` (`user_id`,`store_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='店主表';