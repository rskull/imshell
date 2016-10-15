# imshell

Display an image on the terminal.

## Install

```
$ go get -u github.com/rskull/imshell/...
```

## Build from source code

```
$ go get -u github.com/Masterminds/glide
$ make bundle
$ make
```

## Usage

```
Usage of ./bin/imshell:
  -i string
        image file path
  -t string
        given text replace
  -v    imshell version
  -w float
        output width (default 50)
```

## Example

```
$ imshell -i icon.png -w 30
```

![icon](https://github.com/rskull/imshell/blob/master/testdate/icon.png)

![sample](https://cloud.githubusercontent.com/assets/886268/19393757/142e6fb2-9271-11e6-8646-aa2f0b32bb82.png)
