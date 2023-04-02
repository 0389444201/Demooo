CREATE TABLE `user_table` (
  `id` int AUTO_INCREMENT,
  `email` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `gender` varchar(20) NOT NULL,
  `password` varchar(200) NOT NULL,
  PRIMARY KEY (id)
);