# golangredblack

Generic red-black tree implementation in Go.

## Layout

- `redblack/`: reusable tree package
- `cmd/demo/`: benchmark-style demo program

## Usage

```go
tree := redblack.New[int, string](cmp, 0)
tree.Insert(10, "ten")
value, ok := tree.Get(10)
```

Build the demo:

```sh
make build
```

Run tests:

```sh
make test
```

Run the benchmark:

```sh
make benchmark
```

Run the demo from the top-level `bin/` directory:

```sh
./bin/demo
```

Clean generated binaries and object files:

```sh
make clean
```
