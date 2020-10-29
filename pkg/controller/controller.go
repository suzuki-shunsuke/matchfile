package controller

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type Params struct {
	LogLevel          string
	CheckedFilePath   string
	ConditionFilePath string
}

type Condition struct {
	Exclude bool
	matcher Matcher
}

func (cond Condition) Match(p string) (bool, error) {
	return cond.matcher.Match(p)
}

func New(ctx context.Context, params Params) (Controller, Params, error) {
	if params.LogLevel != "" {
		lvl, err := logrus.ParseLevel(params.LogLevel)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"log_level": params.LogLevel,
			}).WithError(err).Error("the log level is invalid")
		}
		logrus.SetLevel(lvl)
	}

	return Controller{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}, params, nil
}

func (ctrl Controller) readCheckedFile(p string) ([]string, error) {
	checkedFile, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("open a checked file "+p+": %w", err)
	}
	defer checkedFile.Close()
	scanner := bufio.NewScanner(checkedFile)
	checkedFiles := []string{}
	for scanner.Scan() {
		checkedFiles = append(checkedFiles, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error: %w", err)
	}
	return checkedFiles, nil
}

func (ctrl Controller) readConditionFile(p string) ([]string, error) {
	file, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("open a condition file "+p+": %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	conditions := []string{}
	for scanner.Scan() {
		conditions = append(conditions, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error: %w", err)
	}
	return conditions, nil
}

func (ctrl Controller) match(checkedFiles, conditionLines []string) (bool, error) {
	conditions := make([]Condition, 0, len(conditionLines))
	for _, conditionLine := range conditionLines {
		matchParam := parseLine(conditionLine)
		if matchParam.Comment {
			continue
		}
		matchers := make([]Matcher, len(matchParam.Kinds))
		for j, kind := range matchParam.Kinds {
			matcher, err := NewMatcher(matchParam.Path, kind)
			if err != nil {
				return false, fmt.Errorf("initialize a matcher: %w", err)
			}
			matchers[j] = matcher
		}
		conditions = append(conditions, Condition{
			Exclude: matchParam.Exclude,
			matcher: combinedMatcher{matchers: matchers},
		})
	}

	matched := false
	for _, checkedFile := range checkedFiles {
		for _, condition := range conditions {
			if matched && !condition.Exclude {
				continue
			}
			if !matched && condition.Exclude {
				continue
			}
			b, err := condition.Match(checkedFile)
			if err != nil {
				return false, fmt.Errorf("condition matching error: %w", err)
			}
			if b {
				matched = !condition.Exclude
				continue
			}
		}
		if matched {
			break
		}
	}
	return matched, nil
}

func (ctrl Controller) Run(ctx context.Context, params Params) error {
	checkedFiles, err := ctrl.readCheckedFile(params.CheckedFilePath)
	if err != nil {
		return fmt.Errorf("read a checked file "+params.CheckedFilePath+": %w", err)
	}

	conditionLines, err := ctrl.readConditionFile(params.ConditionFilePath)
	if err != nil {
		return fmt.Errorf("read a condition file "+params.CheckedFilePath+": %w", err)
	}

	matched, err := ctrl.match(checkedFiles, conditionLines)
	if err != nil {
		return fmt.Errorf("matching files: %w", err)
	}

	msg := "false"
	if matched {
		msg = "true"
	}
	fmt.Fprintln(ctrl.Stdout, msg)
	return nil
}
