Face detection with go(lang) and opencv
==================

## Install

### Linux & Mac OS X

Install [Go](https://golang.org/doc/install#install), [Glide](https://github.com/Masterminds/glide#install) and [OpenCV](http://opencv.org/). You might want to install all of them via `apt-get` or `homebrew`.

```
cd $GOPATH/src/github.com/[username]
git clone https://github.com/mikitu/go-opencv-fd.git
cd go-opencv-fd
glide install
go run src/main.go
```
