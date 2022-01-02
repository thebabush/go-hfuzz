# go-hfuzz

Go bindings for [honggfuzz](https://github.com/google/honggfuzz).

**NOTE:** this module does NOT instrument your code automatically.
          If you want that, you probably want to use go's built-in fuzzing support.

## Install

```sh
make clean
make
```

## Example

[cmd/simple-test/main.go](cmd/simple-test/main.go).

## Test

```sh
mkdir -p /tmp/in
./honggfuzz/honggfuzz -P -i /tmp/in -o /tmp/out --crashdir /tmp/crash -- ./cmd/simple-test/simple-test
```

...wait until the crash and then:

```sh
$ xxd /tmp/crash/*
00000000: c0ca f16a dead beef [...]
[...]
```
