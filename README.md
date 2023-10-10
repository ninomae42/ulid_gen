# ULID gen
ULID generator tool

## Requirements
- Go

## Install
```
go install github.com/ninomae42/ulid_gen@latest
```

## Usage
```shell
$ ulid_gen
01HCC3PCB3J06WGVAM192EBDB1

$ ulid_gen -c 2 (with num of ulids to generate)
01HCC3R6ACMJCTCRP2T4BCZDFW
01HCC3R6ACMJCTCRP2T5STAKKT

$ ulid_gen -n hoge (with prefix)
hoge+01HCC3PYZ8ZZMB4K3C3ZQ1YDWS

$ ulid_gen -c 2 -n hoge (with prefix)
hoge+01HCC3RZD0MMM56D8WPJM8NKRS
hoge+01HCC3RZD0MMM56D8WPN6JZHMW
```
