/*
 Navicat Premium Data Transfer

 Source Server         : aliyun_trial
 Source Server Type    : MySQL
 Source Server Version : 50743
 Source Host           : 127.0.0.1:3306
 Source Schema         : dousheng_db

 Target Server Type    : MySQL
 Target Server Version : 50743
 File Encoding         : 65001

 Date: 12/09/2023 15:57:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for auth
-- ----------------------------
DROP TABLE IF EXISTS `auth`;
CREATE TABLE `auth`  (
  `user_id` bigint(20) NOT NULL,
  `passwd_crypt` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `salt` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  PRIMARY KEY (`user_id`) USING BTREE,
  CONSTRAINT `auth_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auth
-- ----------------------------
INSERT INTO `auth` VALUES (1460162896, '4e38d641f250285325cc1a2d2b7c1e9b', '4TFQX3vFO7mkjJvW89VLQw==');
INSERT INTO `auth` VALUES (2883017410, '10f777e31e7b87590691b06ab13654a4', 'lvjmQ/mS2ww9NLaWcu3uGQ==');
INSERT INTO `auth` VALUES (3384586993, 'dd07d1131ab08a9ec79ee1370ffec2ad', '/PPv9WocAjysSzNkXV1eDQ==');
INSERT INTO `auth` VALUES (5529799995, 'cc78efbf915783d139600a03c94d5ae3', 'vO7iiFW4mZ6EqyX2FIqJnA==');
INSERT INTO `auth` VALUES (6732519909, '2597d652c21cfd8c843a69ba88a3467a', 'IkaK3bpfWl5HbSP1DVBjww==');
INSERT INTO `auth` VALUES (8258216771, '7235171aa03a23c24b5a6ebf419c45bc', 'bu8TZMMmUfdCfviEv9FKXQ==');
INSERT INTO `auth` VALUES (8398219489, 'f74cf062156a4fbffbc3a130a883aa77', 'TVbDqzsRlc6iXE3kQWWyZA==');
INSERT INTO `auth` VALUES (9568277701, 'feb31cc571f188f994fe57deb0a8bd38', '2aAU1tjFHi+YTuauIxM6gw==');

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `video_id` bigint(20) NOT NULL,
  `content` varchar(900) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `video_id`(`video_id`) USING BTREE,
  CONSTRAINT `comments_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `comments_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (3322241640, 3384586993, 100044, 'test_comment No.2', '2023-09-06 14:16:32', NULL);
INSERT INTO `comments` VALUES (5278995311, 8258216771, 100044, 'nice', '2023-09-12 05:10:47', NULL);
INSERT INTO `comments` VALUES (6427205422, 3384586993, 100044, 'test_comment No.1 ', '2023-09-06 14:15:36', '2023-09-06 14:15:58');
INSERT INTO `comments` VALUES (6557640619, 5529799995, 100045, 'test2222', '2023-09-06 14:26:39', NULL);
INSERT INTO `comments` VALUES (7627340682, 5529799995, 100044, 'test33333', '2023-09-06 14:26:58', NULL);
INSERT INTO `comments` VALUES (8883110565, 5529799995, 100045, 'test11111', '2023-09-06 14:26:34', NULL);

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` bigint(20) NOT NULL,
  `from_user_id` bigint(20) NOT NULL,
  `to_user_id` bigint(20) NOT NULL,
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `deleated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `from_user_id`(`from_user_id`) USING BTREE,
  INDEX `to_user_id`(`to_user_id`) USING BTREE,
  CONSTRAINT `message_ibfk_1` FOREIGN KEY (`from_user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `message_ibfk_2` FOREIGN KEY (`to_user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------

-- ----------------------------
-- Table structure for relation
-- ----------------------------
DROP TABLE IF EXISTS `relation`;
CREATE TABLE `relation`  (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  `to_user_id` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `to_user_id`(`to_user_id`) USING BTREE,
  CONSTRAINT `relation_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `relation_ibfk_2` FOREIGN KEY (`to_user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of relation
-- ----------------------------

-- ----------------------------
-- Table structure for user_favorite_videos
-- ----------------------------
DROP TABLE IF EXISTS `user_favorite_videos`;
CREATE TABLE `user_favorite_videos`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `video_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `video_id`(`video_id`) USING BTREE,
  CONSTRAINT `user_favorite_videos_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `user_favorite_videos_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 100000029 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_favorite_videos
-- ----------------------------
INSERT INTO `user_favorite_videos` VALUES (100000026, 3384586993, 100044);
INSERT INTO `user_favorite_videos` VALUES (100000027, 5529799995, 100003);
INSERT INTO `user_favorite_videos` VALUES (100000028, 5529799995, 100000);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(20) NOT NULL,
  `username` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `follow_count` bigint(20) NULL DEFAULT NULL,
  `follower_count` bigint(20) NULL DEFAULT 0,
  `is_follow` tinyint(4) NOT NULL COMMENT '0未关注，1已关注',
  `avatar` varchar(2083) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `background_image` varchar(2083) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `signature` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `total_favorited` bigint(20) NULL DEFAULT NULL,
  `work_count` bigint(20) NULL DEFAULT NULL,
  `favorite_count` bigint(20) NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL COMMENT 'cre',
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1460162896, 'wangty002', 0, 0, 0, '', '', '', 6, 3, 3, '2023-08-22 05:09:49', '2023-09-06 14:25:50', NULL);
INSERT INTO `users` VALUES (2883017410, 'wangty0015', 0, 0, 0, '', '', '', 1, 2, 1, '2023-09-05 20:05:05', '2023-09-05 20:44:19', NULL);
INSERT INTO `users` VALUES (3384586993, 'wangty0016', 0, 0, 0, '', '', '', 1, 1, 1, '2023-09-06 14:09:59', '2023-09-06 14:14:39', NULL);
INSERT INTO `users` VALUES (5529799995, 'wangty0018', 0, 0, 0, '', '', '', 0, 1, 2, '2023-09-06 14:23:05', '2023-09-06 14:25:54', NULL);
INSERT INTO `users` VALUES (6732519909, 'wangty001', 0, 0, 0, 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.duitang.com%2Fuploads%2Fitem%2F202003%2F06%2F20200306103939_JuZhW.jpeg&refer=http%3A%2F%2Fc-ssl.duitang.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1696066387&t=b4d6358ec4ec4e92bd2898e8bc04c1ef', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.duitang.com%2Fuploads%2Fitem%2F202003%2F06%2F20200306103939_JuZhW.jpeg&refer=http%3A%2F%2Fc-ssl.duitang.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1696066387&t=b4d6358ec4ec4e92bd2898e8bc04c1ef', '这是我的个人简介！😀', 10006, 5, 228, '2023-08-22 05:09:42', '2023-09-06 14:25:54', NULL);
INSERT INTO `users` VALUES (8258216771, 'test@163.com', 0, 0, 0, '', '', '', 0, 0, 0, '2023-09-12 05:09:07', '2023-09-12 05:09:07', NULL);
INSERT INTO `users` VALUES (8398219489, 'wangty003', 0, 0, 0, '', '', '', 0, 0, 0, '2023-08-22 05:09:54', '2023-08-22 05:09:54', NULL);
INSERT INTO `users` VALUES (9568277701, 'wangty004', 0, 0, 0, '', '', '', 0, 0, 0, '2023-08-22 05:09:58', '2023-08-22 05:09:58', NULL);

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `play_url` varchar(2083) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `cover_url` varchar(2083) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `favorite_count` bigint(20) NOT NULL,
  `comment_count` bigint(20) NOT NULL,
  `is_favorite` tinyint(4) NOT NULL DEFAULT 0 COMMENT '0表示未点赞，1表示已点赞。（自己给视频点赞）',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `videos_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 100046 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (100000, 6732519909, 'https://oss.agilestudio.cn/1596432356843366.mp4', 'http://oss.agilestudio.cn/snaphost-demo.jpg', 1, 0, 0, 'NO.3  what do you want from me', '2023-08-25 18:10:11', '2023-09-06 14:25:54', NULL);
INSERT INTO `videos` VALUES (100001, 6732519909, 'https://v-cdn.zjol.com.cn/276982.mp4', 'https://oss-dousheng.oss-cn-beijing.aliyuncs.com/video/6732519909/1596432356843388.mp4?x-oss-process=video/snapshot,t_17000,f_jpg,w_800,h_600', 0, 0, 0, 'NO.2 山水暴涨了。。', '2023-08-24 19:10:11', '2023-08-25 08:47:03', NULL);
INSERT INTO `videos` VALUES (100002, 1460162896, 'http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4', 'http://oss.agilestudio.cn/snaphost-demo.jpg', 12000, 345, 1, 'No.1', '2023-08-23 19:10:11', '2023-08-23 19:10:11', NULL);
INSERT INTO `videos` VALUES (100003, 1460162896, 'https://prod-streaming-video-msn-com.akamaized.net/a8c412fa-f696-4ff2-9c76-e8ed9cdffe0f/604a87fc-e7bc-463e-8d56-cde7e661d690.mp4', 'https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcdM.img', 1, 0, 0, 'No.4', '2023-08-26 21:29:55', '2023-09-06 14:25:50', NULL);
INSERT INTO `videos` VALUES (100044, 3384586993, 'http://123.57.251.188:22441/oss-dousheng/video/3384586993/1694009570858.mp4', 'http://123.57.251.188:22441/oss-dousheng/video/3384586993/1694009570858.mp4?x-oss-process=video/snapshot,t_2000,f_jpg,w_0,h_0', 1, 3, 0, '芜湖海鸥~No.5', '2023-09-06 14:12:51', '2023-09-12 05:10:47', NULL);
INSERT INTO `videos` VALUES (100045, 5529799995, 'http://123.57.251.188:22441/oss-dousheng/video/5529799995/1694010304827.mp4', 'http://123.57.251.188:22441/oss-dousheng/video/5529799995/1694010304827.mp4?x-oss-process=video/snapshot,t_2000,f_jpg,w_0,h_0', 0, 2, 0, 'no.6 喜羊羊', '2023-09-06 14:25:05', '2023-09-06 14:26:39', NULL);

SET FOREIGN_KEY_CHECKS = 1;
