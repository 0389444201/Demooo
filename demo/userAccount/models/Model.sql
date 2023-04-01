CREATE TABLE `user_table` (
  `id` int AUTO_INCREMENT,
  `email` varchar(50) NOT NULL,
  `name` varchar(24) NOT NULL,
  `gender` varchar(10) NOT NULL,
  `password` varchar(30) NOT NULL,
  PRIMARY KEY (id)
);