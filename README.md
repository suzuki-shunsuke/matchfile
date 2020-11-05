# matchfile

[![Build Status](https://github.com/suzuki-shunsuke/matchfile/workflows/CI/badge.svg)](https://github.com/suzuki-shunsuke/matchfile/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/matchfile)](https://goreportcard.com/report/github.com/suzuki-shunsuke/matchfile)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/matchfile.svg)](https://github.com/suzuki-shunsuke/matchfile)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/matchfile/master/LICENSE)

CLI tool to check file paths are matched to the condition

## Motivation

In CI of Pull Request (PR), sometimes we want to run the test, lint, and build for only updated code to save time and prevent unrelated failures.
`matchfile` judges whether we have to run the job with the two arguments.

* which files are updated
* which files depend on the job

Of course we can use `matchfile` for the other objective, because `matchfile` doesn't depend on CI and PR.

## Install

Download from [GitHub Releases](https://github.com/suzuki-shunsuke/matchfile/releases)

```
$ matchfile --version
matchfile version 0.1.0
```

## How to use

`matchfile run` and `matchfile list` takes two positional arguments.

```
$ matchfile run <checked file> <condition file>
```

`<checked file>` is the path to the file whose content is the list of checked file paths.
`<condition file>` is the path to the file whose content is the condition.

If there is a file path which matches to the condition of `<condition file>` in the `<checked file>`, `matchfile run` outputs `true` as the standard output, otherwise outputs `false`.

`matchfile list` outputs the file paths which matches to the condition.

## Example

Prepare the two file `checked_files.txt` and `condition.txt` and run `matchfile run checked_files.txt condition.txt`.

```
$ cat checked_files.txt
service/foo/main.go
README.md

$ cat condition.txt
# comment. This line is ignored.
regexp scripts/.*
service/foo

$ matchfile run checked_files.txt condition.txt
true
```

`true` is outputted because `service/foo/main.go` matches to the condition `service/foo`.

```
$ matchfile list checked_files.txt condition.txt
service/foo/main.go
```

## Condition File Format

The format is `matchfile` specific.
The format is inspired by [gitignore](https://git-scm.com/docs/gitignore).

```
[#][!][<kind>,...] <path>
...
```

When the multiple kinds are specified, the condition matches when either of them matches.

The line starts with "#" is ignored as code comment.
Note that the comment in the middle of the line isn't supported.

`[<kind>,...]` is optional, and the default value is `equal,dir,glob`.

### kind

* equal: check the equality
* dir: [strings.HasPrefix](https://golang.org/pkg/strings/#HasPrefix)
* regexp: [regexp.MatchString](https://golang.org/pkg/regexp/#Regexp.MatchString)
* glob: [filepath.Match](https://golang.org/pkg/path/filepath/#Match)

## Usage

```
$ matchfile help
NAME:
   matchfile - Check file paths are matched to the condition. https://github.com/suzuki-shunsuke/matchfile

USAGE:
   matchfile [global options] command [command options] [arguments...]

VERSION:
   0.1.3

COMMANDS:
   run      Check file paths are matched to the condition
   list     List file paths which matches to the condition
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

```
$ matchfile run --help
NAME:
   matchfile run - Check file paths are matched to the condition

USAGE:
   matchfile run [command options] <checked file> <condition file>

OPTIONS:
   --log-level value  log level
   --help, -h         show help (default: false)
```

```
$ matchfile list --help
NAME:
   matchfile list - List file paths which matches to the condition

USAGE:
   matchfile list [command options] <checked file> <condition file>

OPTIONS:
   --log-level value  log level
   --help, -h         show help (default: false)
```

## LICENSE

[MIT](LICENSE)
