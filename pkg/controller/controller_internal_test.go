package controller

import (
	"os"
	"testing"
)

//nolint:funlen
func TestController_match(t *testing.T) {
	data := []struct {
		title        string
		checkedFiles []string
		conditions   []string
		exp          bool
		isErr        bool
	}{
		{
			title:        "equal",
			checkedFiles: []string{"foo.txt"},
			conditions:   []string{"foo.txt"},
			exp:          true,
		},
		{
			title:      "no checked files",
			conditions: []string{"foo.txt"},
			exp:        false,
		},
		{
			title:        "no conditions",
			checkedFiles: []string{"foo.txt"},
			// TODO consider the specification
			exp: false,
		},
		{
			title:        "default dir",
			checkedFiles: []string{"services/foo.txt"},
			conditions:   []string{"services"},
			exp:          true,
		},
		{
			title:        "dir",
			checkedFiles: []string{"services/foo.txt"},
			conditions:   []string{"dir services"},
			exp:          true,
		},
		{
			title:        "glob",
			checkedFiles: []string{"services/foo.txt"},
			conditions:   []string{"glob services/*"},
			exp:          true,
		},
		{
			title:        "regexp",
			checkedFiles: []string{"services/foo.txt"},
			conditions:   []string{"regexp services/.*"},
			exp:          true,
		},
		{
			title:        "exclude 1",
			checkedFiles: []string{"services/foo.txt"},
			conditions:   []string{"!dir services"},
			exp:          false,
		},
		{
			title:        "exclude 2",
			checkedFiles: []string{"services/foo.txt"},
			conditions: []string{
				"dir services",
				"!dir services",
			},
			exp: false,
		},
		{
			title:        "exclude 3",
			checkedFiles: []string{"services/foo.txt"},
			conditions: []string{
				"dir services",
				"!dir services",
				"dir services",
			},
			exp: true,
		},
		{
			title: "exclude 4",
			checkedFiles: []string{
				".akoi.yml",
				".github-comment.yml",
				"ci/build.sh",
				"ci/install-aws-cli.sh",
				"ci/install.sh",
				"ci/merge.sh",
				"ci/validate-renovate.sh",
				"go-dependencies.txt",
				"pkg/k8s/k8s.go",
				"renovate.json",
			},
			conditions: []string{
				"regexp .*",
				"!equal README.md",
				"!equal renovate.json",
			},
			exp: true,
		},
	}
	ctrl := Controller{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			b, err := ctrl.match(d.checkedFiles, d.conditions)
			if d.isErr {
				if err == nil {
					t.Fatal("error should occur")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if d.exp != b {
				t.Fatalf("ctrl.match() = %t, wanted %t", b, d.exp)
			}
		})
	}
}
