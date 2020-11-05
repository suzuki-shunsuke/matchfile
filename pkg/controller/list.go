package controller

import (
	"context"
	"fmt"
	"strings"
)

func (ctrl Controller) List(ctx context.Context, params Params) error {
	checkedFiles, err := ctrl.readCheckedFile(params.CheckedFilePath)
	if err != nil {
		return fmt.Errorf("read a checked file "+params.CheckedFilePath+": %w", err)
	}

	conditionLines, err := ctrl.readConditionFile(params.ConditionFilePath)
	if err != nil {
		return fmt.Errorf("read a condition file "+params.CheckedFilePath+": %w", err)
	}

	matchedFiles, err := ctrl.list(checkedFiles, conditionLines)
	if err != nil {
		return fmt.Errorf("matching files: %w", err)
	}

	if len(matchedFiles) > 0 {
		fmt.Fprintln(ctrl.Stdout, strings.Join(matchedFiles, "\n"))
	}

	return nil
}

func (ctrl Controller) list(checkedFiles, conditionLines []string) ([]string, error) {
	conditions, err := getConditions(conditionLines)
	if err != nil {
		return nil, fmt.Errorf("parse checked files: %w", err)
	}

	matchedFiles := make([]string, 0, len(checkedFiles))
	for _, checkedFile := range checkedFiles {
		matched := false
		for _, condition := range conditions {
			if matched && !condition.Exclude {
				continue
			}
			if !matched && condition.Exclude {
				continue
			}
			b, err := condition.Match(checkedFile)
			if err != nil {
				return nil, fmt.Errorf("condition matching error: %w", err)
			}
			if b {
				matched = !condition.Exclude
				continue
			}
		}
		if matched {
			matchedFiles = append(matchedFiles, checkedFile)
			continue
		}
	}
	return matchedFiles, nil
}
