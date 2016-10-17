# imshell

Displaying an image on the terminal.

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

```
$ imshell -i icon.png -w 30 -t Gopher
```

![gopher](https://cloud.githubusercontent.com/assets/886268/19423369/ab784aec-945a-11e6-8a38-8a9a533da26e.png)

![sample2](https://cloud.githubusercontent.com/assets/886268/19423372/b134edd2-945a-11e6-9b46-a22b7483a4e8.png)
