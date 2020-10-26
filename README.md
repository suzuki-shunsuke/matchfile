# matchfile

[![Build Status](https://github.com/suzuki-shunsuke/matchfile/workflows/CI/badge.svg)](https://github.com/suzuki-shunsuke/matchfile/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/matchfile)](https://goreportcard.com/report/github.com/suzuki-shunsuke/matchfile)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/matchfile.svg)](https://github.com/suzuki-shunsuke/matchfile)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/matchfile/master/LICENSE)

CLI tool to check file paths are matched to the condition

## Install

Download from [GitHub Releases](https://github.com/suzuki-shunsuke/matchfile/releases)

```
$ matchfile --version
matchfile version 0.1.0
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

### kind

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
   0.1.0

COMMANDS:
   run      Check file paths are matched to the condition
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

## LICENSE

[MIT](LICENSE)
