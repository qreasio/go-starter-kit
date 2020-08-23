CREATE TABLE IF NOT EXISTS `users`  (
        `id` varchar(36) NOT NULL,
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
        UNIQUE KEY `username` (`username`),
        UNIQUE KEY `email_UNIQUE` (`email`)
        ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


