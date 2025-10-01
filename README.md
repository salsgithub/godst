# godst (Go Data Structures)

Implementation of well known data structures and algorithms with the intent of having zero external dependencies (even for test libraries).

## Features

- **Type Safe & Generics** - Built with go1.18+generics
- **Zero dependencies** - Pure Go, no external libraries are required
- **Fully tested** - 100% code coverage across all data structures & benchmarks where necessary
- **Classic Algorithms** - Includes implementations of fundamental graph and tree traversal algorithms, pathfinding and more.

## Data Structures

- [**Array list**](./arraylist)
- [**Bi-directional map**](./bidimap)
- [**Binary search tree**](./binarytree)
- [**LRU cache**](./cache/lru.go)
- [**Doubly linked list**](./doublylinkedlist)
- [**Graph**](./graph)
- [**Heap**](./heap)
- [**Queue**](./queue)
- [**Set**](./set)
- [**Singly linked list**](./singlylinkedlist)
- [**Stack**](./stack)

## Running tests

```shell
# Run all unit tests
make test

# Run all benchmarks
make bench
```

## License

This project is licensed under the MIT license - see the [LICENSE](./LICENSE) file for details.