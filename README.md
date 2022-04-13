# kvstore

A simple in-memory key-value store in Go.

## Launching Server

```
> go build kvstore.go

> ./kvstore -h 127.0.0.1 -p 7777
Listening on 127.0.0.1:7777
```

## Sending Commands

```
> telnet 127.0.0.1 7777
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
set mykey myvalue 0
OK
get mykey
myvalue
OK
del mykey
OK
get mykey
OK
```

## Commands

### Set

Stores a value with a given TTL. Use a TTL of `0` to never expire the entry.

```
set key value ttl
```

### Get

Retrieves a stored key-value pair.

```
get key
```

### Delete

Deletes a stored key-value pair.

```
del key
```
