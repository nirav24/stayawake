# Stay Awake

This is a small program which will move mouse up and down to keep screen on. This stimulates like user is doing something.

## Motivation

I had to run something over night but after some time, my VPN was getting disconnected. One of my friends had shown to me
that he opens notepad and puts something on spacebar so that it will keep writing spaces in blank notepad which will keep the screen on.

## Installation

### Prerequisite
- [Go](https://golang.org/doc/install)


```shell
go get github.com/nirav24/stayawake
```

## Usage

```shell
# See help
stayawake -help

# Stay awake for 30min and move every 10 seconds
stayawake -duration 30m -interval 10s
```