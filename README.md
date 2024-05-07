# nanoid-cli

A tiny cli wrapper around [go-nanoid](https://github.com/jaevor/go-nanoid) to generate unique ids from the command line.

### Install

With go 1.22 or higher:

```sh
go install github.com/37sprinkles/nanoid-cli@latest
```

### Usage

```sh
Usage: nanoid <count> -- generates canonic ids
       nanoid <charset> <length> [count] -- generates ids with a specific charset and length
  charset (string): required, "hex", "alpha", "numeric", "base64", "base90" or a string of characters
  length (byte): required, must be between 2 and 255
  count (int32): optional, must be at least 1
```

### Examples

```sh
$ nanoid 
HqNAU-6Echg2J-AYSk4PG
```

```sh
$ nanoid 3
IeBZmiYMF-PxXk2QQAPtr
teDcoIqIOGOcc8xGBOhZb
VVQ64IbKU1R1PzqOOjTMM
```

```sh
$ nanoid hex 32
7C557D76C60CD45561C3352D5ECB2089
```

```sh
$ nanoid abcdefghijkl 16 5
fchfdkiehkdkibij
gafkidechhelccdi
hjacfkbilagdlhfh
hbfllkgcdkeacfkc
jehidagfhhkkbaji
```
