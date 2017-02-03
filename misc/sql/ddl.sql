DROP DATABASE IF EXISTS fizzbuzz_twitterbot;
CREATE DATABASE fizzbuzz_twitterbot CHARACTER SET utf8;
USE fizzbuzz_twitterbot;

DROP TABLE IF EXISTS fizzbuzz;
CREATE TABLE fizzbuzz (
  id BIGINT NOT NULL AUTO_INCREMENT,
  num BIGINT NOT NULL,
  tweet VARCHAR(140) NOT NULL,
  tweet_id BIGINT,
  updated_at DATETIME NOT NULL,
  created_at DATETIME NOT NULL,
  PRIMARY KEY (id)
) ENGINE InnoDB CHARSET utf8;

DROP TABLE IF EXISTS error_log;
CREATE TABLE error_log (
  id BIGINT NOT NULL AUTO_INCREMENT,
  fizzbuzz_id BIGINT NOT NULL,
  error_message VARCHAR(255) NOT NULL,
  updated_at DATETIME NOT NULL,
  created_at DATETIME NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (fizzbuzz_id) REFERENCES fizzbuzz (id)
) ENGINE InnoDB CHARSET utf8;
