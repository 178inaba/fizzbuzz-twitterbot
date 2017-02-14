DROP TABLE IF EXISTS fizzbuzz_tweets;
CREATE TABLE fizzbuzz_tweets (
  number BIGINT UNSIGNED NOT NULL,
  is_fizz BOOL NOT NULL,
  is_buzz BOOL NOT NULL,
  tweet VARCHAR(140) NOT NULL,
  twitter_tweet_id BIGINT UNSIGNED NOT NULL DEFAULT 0,
  updated_at DATETIME NOT NULL,
  created_at DATETIME NOT NULL,
  PRIMARY KEY (number)
) ENGINE InnoDB CHARSET utf8;

DROP TABLE IF EXISTS post_errors;
CREATE TABLE post_errors (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  fizzbuzz_tweet_number BIGINT UNSIGNED NOT NULL,
  error_message TEXT NOT NULL,
  updated_at DATETIME NOT NULL,
  created_at DATETIME NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (fizzbuzz_tweet_number) REFERENCES fizzbuzz_tweets (number)
) ENGINE InnoDB CHARSET utf8;
