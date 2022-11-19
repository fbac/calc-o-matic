# calc-o-matic

<img src="logo.png"
     style="float: left; margin-right: 10px;"
     width="375" height="395" />

**calc-o-matic** is a basic calculator based in [Shunting yard algorithm](https://en.wikipedia.org/wiki/Shunting_yard_algorithm)

## Usage

### Makefile targets

```shell
# build
make build

# test
make test

# clean
make clean

# generate coverage html report
make coverage

# generate docs
make docs
```

### calc-o-matic CLI

```shell
# write math expression
calc-o-matic /> 1+ 1 * (2/2) - 10^7

# cancel with Ctrl+C
calc-o-matic /> ^C
exiting calc-o-matic
```

## Shunting Yard algorithm

<img src="shunting-yard.png"
     alt="Markdown Monster icon"
     style="float: left; margin-right: 10px;"
     width="800" height="500" />
