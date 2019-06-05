# Setup & Run

## Setup

Setup go & set Enviroment GOROOT, GOPATH

```bash
chmod +x ./install.sh
./install.sh
```

Get package for DEV

```bash
chmod +x ./sh/go_get_dev.sh
./sh/go_get_dev.sh
```

Get vender with [dep](github.com/golang/dep/cmd/dep)

```bash
chmod +x ./sh/go_dep.sh
./sh/go_dep.sh
```

## Run

Run file main.go

```bash
go run main.go
```

Run with [Fresh](https://github.com/gravityblast/fresh)

```bash
fresh
```

Testing with [Goconvey](github.com/smartystreets/goconvey)

```bash
cd tut
goconvey
```
