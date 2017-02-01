DROP DATABASE IF EXISTS fizzbuzz_twitterbot;
CREATE DATABASE fizzbuzz_twitterbot CHARACTER SET = utf8;
USE fizzbuzz_twitterbot;

DROP TABLE IF EXISTS fizzbuzz;
CREATE TABLE fizzbuzz (
  id BIGINT NOT NULL AUTO_INCREMENT,
  tweet_id BIGINT NOT NULL,
  num BIGINT NOT NULL UNIQUE,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at DATETIME NOT NULL,
  PRIMARY KEY (id)
) ENGINE InnoDB CHARSET utf8;
