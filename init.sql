CREATE DATABASE IF NOT EXISTS Algorithm_learning_assistant;
USE Algorithm_learning_assistant;

-- ----------------------------
-- Table structure for systems
-- ----------------------------
DROP TABLE IF EXISTS `problems`;
DROP TABLE IF EXISTS `user_websites`;
DROP TABLE IF EXISTS `websites`;
DROP TABLE IF EXISTS `systems`;
DROP TABLE IF EXISTS `users`;

CREATE TABLE `systems`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `minimum_solved` bigint NOT NULL,
  `semester_start` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ding_header` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `ding_bottom` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `cur_term` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_systems_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of systems
-- ----------------------------
INSERT INTO `systems` VALUES (1, NULL, NULL, NULL, 3, '2025-02-17', '各位算法小能手注意查收本周战报~ \n ------✨本周战绩速览✨------\n', '💡温馨提醒：\n 保持解题节奏就像打游戏签到领金币，连续登录会有惊喜加成哦~ 暂时落后的同学别着急，下周「补题buff」已生效！\n 代码不息，刷题不止 \n 我们下周同一时间，继续见证成长！( •̀ ω •́ )✧ \n  （有任何建议欢迎随时滴滴~）\n 详细解题列表请点击:\n http://114.55.128.130:8080/api/user/list?department=familySix', '大三下');

-- ----------------------------
-- Table structure for user_websites
-- ----------------------------
CREATE TABLE `user_websites`  (
  `user_id` bigint UNSIGNED NOT NULL,
  `website_id` bigint UNSIGNED NOT NULL,
  `username` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`user_id`, `website_id`) USING BTREE,
  INDEX `fk_user_websites_website`(`website_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_websites
-- ----------------------------
INSERT INTO `user_websites` VALUES (1, 1, 'kan-fan-xing', 'https://leetcode.cn/u/kan-fan-xing/');
INSERT INTO `user_websites` VALUES (2, 1, 'festive-i2ubinwnk', 'https://leetcode.cn/u/festive-i2ubinwnk/');
INSERT INTO `user_websites` VALUES (3, 1, 'ding-mao-s', 'https://leetcode.cn/u/ding-mao-s/');
INSERT INTO `user_websites` VALUES (4, 1, 'gui-tu-960', 'https://leetcode.cn/u/gui-tu-960/');
INSERT INTO `user_websites` VALUES (5, 1, 'xun_xun', 'https://leetcode.cn/u/xun_xun/');
INSERT INTO `user_websites` VALUES (6, 1, 'hardcore-swirlesrz0', 'https://leetcode.cn/u/hardcore-swirlesrz0/');
INSERT INTO `user_websites` VALUES (7, 1, 'practical-snyderqvy', 'https://leetcode.cn/u/practical-snyderqvy/');
INSERT INTO `user_websites` VALUES (8, 1, 'trusting-6rothendieckqgx', 'https://leetcode.cn/u/trusting-6rothendieckqgx/');
INSERT INTO `user_websites` VALUES (9, 1, '6oofy-gangulyxsi', 'https://leetcode.cn/u/6oofy-gangulyxsi/');
INSERT INTO `user_websites` VALUES (10, 1, 'zao-an-e', 'https://leetcode.cn/u/zao-an-e/');

-- ----------------------------
-- Table structure for users
-- ----------------------------
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `department` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, NULL, NULL, NULL, '宁赵星', 'familySix');
INSERT INTO `users` VALUES (2, NULL, NULL, NULL, '李壮', 'familySix');
INSERT INTO `users` VALUES (3, NULL, NULL, NULL, '田家杰', 'familySix');
INSERT INTO `users` VALUES (4, NULL, NULL, NULL, '方腾飞', 'familySix');
INSERT INTO `users` VALUES (5, NULL, NULL, NULL, '蒋睿勋', 'familySix');
INSERT INTO `users` VALUES (6, NULL, NULL, NULL, '王玉龙', 'familySix');
INSERT INTO `users` VALUES (7, NULL, NULL, NULL, '王怡晗', 'familySix');
INSERT INTO `users` VALUES (8, NULL, NULL, NULL, '贺丽帆', 'familySix');
INSERT INTO `users` VALUES (9, NULL, NULL, NULL, '韩硕博', 'familySix');
INSERT INTO `users` VALUES (10, NULL, NULL, NULL, '雪怡琦', 'familySix');

-- ----------------------------
-- Table structure for problems
-- ----------------------------
DROP TABLE IF EXISTS `problems`;
CREATE TABLE `problems`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `number` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `translated_title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `title_slug` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `question_id` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `submit_time` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `term` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `week` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_problems_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_problems_user_id`(`user_id`) USING BTREE,
  CONSTRAINT `fk_users_problems` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for websites
-- ----------------------------
CREATE TABLE `websites`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_websites_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of websites
-- ----------------------------
INSERT INTO `websites` VALUES (1, NULL, NULL, NULL, '力扣', 'https://leetcode.cn/');

SET FOREIGN_KEY_CHECKS = 1;