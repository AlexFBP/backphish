# PhishPayback: Phising Spammer

The idea is to combat phishing with spam, this is a hacking repo (needless to add "ethical").

Inspired in

- GitHub - [v3lip/PhishingSpammer](https://github.com/v3lip/PhishingSpammer/)
- YouTube - [@ScammerPayback](https://www.youtube.com/@ScammerPayback)

## Types

1. Send random data to phishing backends
2. TODO: Create DDOs attacks for frontends
3. TODO: Create DDOs attacks for backends

## Targets

For safety reasons, no URL will be in this readme, please check source code for details.

And also, for non-safety reasons, Why should I care about any exposed token in this repository? ¯\_(ツ)_/¯

- (Type 1) Daviplata phishing page 1

## Usage

Requirements:

- This repo cloned. All commands here assumes that you are within your cloned copy of this repo
- [Golang installed](https://go.dev/doc/install) (For now, perhaps in future i'd upload binaries...)
- (Optional for development) [docker installed](https://docs.docker.com/engine/install/) (To run a local mockserver)

### Init/Install

```sh
go get
```

### Build

```sh
go build
```

### Run

```sh
./backphish
```
If no options (as above), will list the available targets.

Available CLI options (bypass the menu):
```
-t TARGET       Use a listed TARGET
```


### (Miscellaneous)

#### Local mock server (for development only)

Runs a mock/dummy HTTP server in your local 1080 port - Based in [mockserver](https://www.mock-server.com/). With docker:

```sh
docker run -it --rm -p 1080:1080 mockserver/mockserver
```

And then [go to the dashboard](http://localhost:1080/mockserver/dashboard)
