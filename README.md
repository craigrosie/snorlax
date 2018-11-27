# Snorlax

For when your data is just _too fast_.

Use snorlax to make your streaming data take a nap, so you can actually read the output in your terminal.

## Installation

```bash
$ go get github.com/craigrosie/snorlax
```

## Usage

```bash
$ snorlax -help
Usage of snorlax:
  -s float
        seconds to sleep for (default 1)

$ cat really_big_file.txt | snorlax

$ cat really_big_file.txt | snorlax -s 0.5
```
