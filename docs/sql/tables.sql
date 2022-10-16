CREATE TABLE `resource`  (
 `id` char(64) NOT NULL AUTO_INCREMENT,
 `vendor` tinyint(1) NOT NULL,
 `region` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 `create_at` bigint(20) NOT NULL,
 `expire_at` bigint(20) NOT NULL,
 `type` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 `update_at` bigint(20) NOT NULL,
 `sync_at` bigint(20) NOT NULL,
 `accout` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 `public_ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 `private_ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 PRIMARY KEY (`id`) USING BTREE,
 INDEX `status`(`status`) USING BTREE,
 INDEX `private_ip`(`public_ip`) USING BTREE,
 INDEX `public_ip`(`public_ip`) USING BTREE,
 INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

CREATE TABLE `host`  (
 `resource_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 `cpu` tinyint(5) NOT NULL,
 `memory` int(64) NOT NULL,
 `gpu_amount` tinyint(255) NOT NULL,
 `gpu_spec` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 `os_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 `os_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 `serial_number` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
 PRIMARY KEY (`resource_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;