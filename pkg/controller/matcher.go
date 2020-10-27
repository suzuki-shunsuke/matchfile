package controller

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

type MatcherParam struct {
	Kinds   []string
	Path    string
	Exclude bool
	Comment bool
}

func parseLine(line string) MatcherParam {
	if strings.HasPrefix(line, "#") {
		return MatcherParam{Comment: true}
	}
	param := MatcherParam{}
	if strings.HasPrefix(line, "!") {
		param.Exclude = true
		line = line[1:]
	}
	idx := strings.Index(line, " ")
	if idx == -1 {
		param.Kinds = []string{"equal", "dir", "glob"}
		param.Path = line
		return param
	}
	param.Kinds = strings.Split(line[:idx], ",")
	param.Path = line[idx+1:]
	return param
}

type Matcher interface {
	Match(string) (bool, error)
}

type combinedMatcher struct {
	matchers []Matcher
}

func (m combinedMatcher) Match(p string) (bool, error) {
	for _, matcher := range m.matchers {
		b, err := matcher.Match(p)
		if err != nil {
			return b, fmt.Errorf("matching error: %w", err)
		}
		if b {
			return true, nil
		}
	}
	return false, nil
}

type dirMatcher struct {
	dir string
}

func (matcher dirMatcher) Match(p string) (bool, error) {
	return strings.HasPrefix(p, matcher.dir), nil
}

type globMatcher struct {
	pattern string
}

func (matcher globMatcher) Match(p string) (bool, error) {
	return filepath.Match(matcher.pattern, p)
}

type regexpMatcher struct {
	pattern *regexp.Regexp
}

func (matcher regexpMatcher) Match(p string) (bool, error) {
	return matcher.pattern.MatchString(p), nil
}

type equalMatcher struct {
	pattern string
}

func (matcher equalMatcher) Match(p string) (bool, error) {
	return p == matcher.pattern, nil
}

func NewMatcher(pattern, kind string) (Matcher, error) {
	switch kind {
	case "dir":
		return dirMatcher{dir: filepath.Clean(pattern) + "/"}, nil
	case "regexp":
		exp, err := regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("compile a regular expression: %w", err)
		}
		return regexpMatcher{pattern: exp}, nil
	case "glob":
		return globMatcher{pattern: pattern}, nil
	case "equal":
		return equalMatcher{pattern: pattern}, nil
	default:
		return nil, errors.New("invalid kind: " + kind)
	}
}
