CREATE TABLE `users`  (
                                        `id` int(11) NOT NULL AUTO_INCREMENT,
                                        `password` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
                                        `last_login` datetime(6) DEFAULT NULL,
                                        `is_superuser` tinyint(1) NOT NULL,
                                        `username` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
                                        `first_name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
                                        `last_name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
                                        `email` varchar(254) COLLATE utf8mb4_unicode_ci NOT NULL,
                                        `is_staff` tinyint(1) NOT NULL,
                                        `is_active` tinyint(1) NOT NULL,
                                        `date_joined` datetime(6) NOT NULL,
                                        PRIMARY KEY (`id`),
                                        UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
INSERT INTO `users` VALUES (1,'pbkdf2_sha256$180000$bjT96G8jlzfW$WtjIbBA1/2IBoGALiGzBWDM097FkqmE1OLPwmXfbXss=','2020-05-20 05:06:10.089832',1,'isak','Isak','Founder Ngapaen','isak@ngapaen.com',1,1,'2020-05-20 05:05:52.461293'),(2,'pbkdf2_sha256$180000$bjT96G8jlzfW$WtjIbBA1/2IBoGALiGzBWDM097FkqmE1OLPwmXfbXss=','2020-05-20 05:06:10.089832',0,'deili','Deili','Tandaju','deili@tandaju.com',0,1,'2020-05-20 05:05:52.461293'),(3,'pbkdf2_sha256$180000$bjT96G8jlzfW$WtjIbBA1/2IBoGALiGzBWDM097FkqmE1OLPwmXfbXss=','2020-05-20 05:06:10.089832',0,'eric','Eric','Sidarta','eric@sidarta.com',0,1,'2020-05-20 05:05:52.461293'),(4,'pbkdf2_sha256$180000$bjT96G8jlzfW$WtjIbBA1/2IBoGALiGzBWDM097FkqmE1OLPwmXfbXss=','2020-05-20 05:06:10.089832',0,'farida','Farida','Tjandra','farida@tjandra.com',0,1,'2020-05-20 05:05:52.461293');
