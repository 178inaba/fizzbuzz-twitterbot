drop database if exists fizzbuzz_twitterbot;
create database fizzbuzz_twitterbot character set = utf8;
use fizzbuzz_twitterbot;

drop table if exists fizzbuzz;
create table fizzbuzz (
  id bigint not null auto_increment,
  tweet_id bigint not null,
  num bigint not null,
  updated_at datetime not null default current_timestamp,
  created_at datetime not null,
  primary key (id),
) engine InnoDB charset utf8;
