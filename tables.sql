create table users
(
  ID varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  birth DATE NOT NULL,
  email varchar(255) NOT NULL,
  active boolean NOT NULL DEFAULT 1,
  password varchar(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  created_id varchar(255) NOT NULL,
  updated_id varchar(255) NOT NULL,
  PRIMARY KEY (ID)
);
