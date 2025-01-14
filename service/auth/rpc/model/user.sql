CREATE TABLE `t_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL,
  `password` varchar(32) NOT NULL,
  `email` varchar(50) DEFAULT NULL,
  `phone` varchar(32) DEFAULT NULL,
  `createtime` date NOT NULL,
  `updatetime` date NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1为激活0为未激活',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;