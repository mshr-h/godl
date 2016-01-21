# godl [![Build Status](https://travis-ci.org/mshr-h/godl.svg?branch=master)](https://travis-ci.org/mshr-h/godl)

Simple file downloader written in Golang

## Features

- Simultaneous download using go channels
- Allow passing urls by argument or file

## Usage

```shell
$ godl http://someurl.com/somefile1.zip http://someurl.com/somefile2.zip

$ cat path_to_url_list.txt
http://someurl.com/somefile1.zip
http://someurl.com/somefile2.zip
http://someurl.com/somefile3.zip
http://someurl.com/somefile4.zip

$ godl -i path_to_url_list.txt
```
