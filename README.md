# FizzBuzz Twitter Bot

[![Build Status](https://travis-ci.org/178inaba/fizzbuzz-twitterbot.svg?branch=master)](https://travis-ci.org/178inaba/fizzbuzz-twitterbot)
[![Coverage Status](https://coveralls.io/repos/github/178inaba/fizzbuzz-twitterbot/badge.svg?branch=master)](https://coveralls.io/github/178inaba/fizzbuzz-twitterbot?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/178inaba/fizzbuzz-twitterbot)](https://goreportcard.com/report/github.com/178inaba/fizzbuzz-twitterbot)

## Docker

### Build

```console
$ docker build --force-rm --no-cache -t 178inaba/fizzbuzz-twitterbot .
```

### Run

```console
$ docker run -d --restart unless-stopped -e CONSUMER_KEY=... -e CONSUMER_SECRET=... -e ACCESS_TOKEN=... -e ACCESS_TOKEN_SECRET=... --name fizzbuzz-twitterbot 178inaba/fizzbuzz-twitterbot
```

### Compose

```console
$ docker-compose up -d --build
```

## Test

Require MySQL or MariaDB.

```console
$ mysql -u root < misc/sql/create_test_db.sql
$ mysql -u root fizzbuzz_twitterbot_test < misc/sql/ddl.sql
$ go test ./...
```

## License

[MIT](LICENSE)
