# Doors

A service for predicting the order of keys to use for opening doors in [Ladder Slasher](https://ladderslasher.d2jsp.org/).

# Run Locally

```bash
$ ./dev/run
Running on port :8080
```

Once running, pass the `class` query parameter to get a "suggestion" of key order to use:

```bash
$ curl "localhost:8080?class=monk"
map[keys:[1 2 4 3]]
```

Optionally, you may provide a `PORT` environment variable to change to listening on another port.

```bash
$ PORT=1337 ./dev/run
Running on port :1337
```