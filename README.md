# tic-init :sparkles:

This is a command line tool for the [tic](https://github.com/peterhellberg/tic)
Zig :zap: module, allowing you to create a directory with code that allows you
to build a [WebAssembly](https://webassembly.org/) cart for the
[TIC-80](https://tic80.com/) :tada:

## Installation

(Requires you to have [Go](https://go.dev/) installed)

```sh
go install github.com/peterhellberg/tic-init@latest
```

## Usage

```sh
tic-init mycart
cd mycart
zig build run
```

> [!Note]
> There is also a `zig build spy` command.
