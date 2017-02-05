DROP DATABASE IF EXISTS fizzbuzz_twitterbot;
CREATE DATABASE fizzbuzz_twitterbot CHARACTER SET utf8;
USE fizzbuzz_twitterbot;

DROP TABLE IF EXISTS fizzbuzz_tweets;
CREATE TABLE fizzbuzz_tweets (
  id BIGINT NOT NULL AUTO_INCREMENT,
  number BIGINT UNSIGNED NOT NULL,
  tweet VARCHAR(140) NOT NULL,
  twitter_tweet_id BIGINT,
  updated_at DATETIME NOT NULL,
  created_at DATETIME NOT NULL,
  PRIMARY KEY (id)
) ENGINE InnoDB CHARSET utf8;

DROP TABLE IF EXISTS post_errors;
CREATE TABLE post_errors (
  id BIGINT NOT NULL AUTO_INCREMENT,
  fizzbuzz_tweet_id BIGINT NOT NULL,
  error_message TEXT NOT NULL,
  updated_at DATETIME NOT NULL,
  created_at DATETIME NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (fizzbuzz_tweet_id) REFERENCES fizzbuzz_tweets (id)
) ENGINE InnoDB CHARSET utf8;
