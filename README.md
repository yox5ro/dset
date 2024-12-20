# dset

Perform set operations (intersection, union, subtraction) on lexicographically sorted text files without Out Of Memory error.

dset works with text file and gzipped text file.

## Installation

```bash
$ go install github.com/yox5ro/dset
```

## Usage

```bash
$ dset --help
Perform set operations (intersection, union, subtraction) on lexicographically sorted text files.

Usage:
  dset [command]

Available Commands:
  help        Help about any command
  intersect   Perform set intersection on lexicographically sorted text files.
  subtract    Perform set subtraction on lexicographically sorted text files.
  union       Perform set union on lexicographically sorted text files.

Flags:
  -h, --help   help for dset

Use "dset [command] --help" for more information about a command.
```

## Example

```bash
$ cat a.txt
apple
banana
cherry
dog
elephant
frog
gorilla

$ cat b.txt
apple
bravo
cherry
dog
elephant
flower
gorilla

$ dset intersect a.txt b.txt
apple
cherry
dog
elephant
gorilla

$ dset union a.txt b.txt
apple
banana
bravo
cherry
dog
elephant
flower
frog
gorilla

$ dset subtract a.txt b.txt
banana
frog
```
