# gitignoreio
CLI interface for [https://www.gitignore.io/](https://www.gitignore.io/)

## Install from source
Get the source code:

```bash
git clone https://github.com/jasonblanchard/gitignoreio && cd gitignorio
```

Install:

```bash
make install
```

## Usage

```
$ gitignoreio

A CLI interface for gitignore.io

Usage:
  gitignoreio [command]

Available Commands:
  generate    Print .gitignore file to std out
  help        Help about any command
  list        List valid options for gitignore.io
  open        Open gitignore.io in a browser

Flags:
      --config string   config file (default is $HOME/.gitignoreio.yaml)
  -h, --help            help for gitignoreio
  -t, --toggle          Help message for toggle

Use "gitignoreio [command] --help" for more information about a command.
```
