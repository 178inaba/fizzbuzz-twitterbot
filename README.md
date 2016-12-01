# FizzBuzz Twitter Bot

## Docker

### Build

```console
$ docker build --force-rm --no-cache -t 178inaba/fizzbuzz-twitterbot .
```

### Run

```console
$ docker run -d --restart unless-stopped -e CONSUMER_KEY=... -e CONSUMER_SECRET=... -e ACCESS_TOKEN=... -e ACCESS_TOKEN_SECRET=... --name fizzbuzz-twitterbot 178inaba/fizzbuzz-twitterbot
```
