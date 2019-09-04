# Go WebSocket Server

A simple WebSocket server that echoes text messages,
however it hates questions marks (`?`) and makes them into exclamation marks`!` 

## Acceptance Criteria

- [x] HTTP Server with custom port
- [x] Port is passed via flag
- [x] Endpoint `/ws` for WebSocket handshake
- [x] WebSocket upgrade
- [x] Listen to Text messages
- [x] Echo to Text messages
- [x] Replace `?` to `!` in echoed message
- [x] Close WebSocket on either client close message or upon receiving non-text message
- [x] Should run on ~~OS X~~, ~~Windows~~, Linux
- [x] Readme.MD

### Running

In order to run this application, first build it

```
go build
```

And then run it

```
./go-websocker-server
```

**NOTE**: By default port is `1234`.

If you want to run it with a custom port, use


```
./go-websocker-server -port 3322
```

## Running the tests

This project uses default test framework


```
go test -cover
```

**Test coverage**:
`78.9% of statements`

### Linting

This project uses [golint](https://github.com/golang/lint)

```
golint
```