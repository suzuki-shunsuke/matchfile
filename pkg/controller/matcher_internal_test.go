package controller

import (
	"testing"
)

func TestMatcher_Match(t *testing.T) {
	t.Parallel()
	regexpMatcher1, err := NewMatcher(".*", "regexp")
	if err != nil {
		t.Fatal(err)
	}
	data := []struct {
		title       string
		checkedFile string
		matcher     Matcher
		exp         bool
		isErr       bool
	}{
		{
			title:       "dir",
			checkedFile: "foo.txt",
			matcher:     dirMatcher{dir: "template"},
			exp:         false,
		},
		{
			title:       "regexp",
			checkedFile: "ci/build.sh",
			matcher:     regexpMatcher1,
			exp:         true,
		},
		{
			title:       "regexp",
			checkedFile: "ci/build.sh",
			matcher:     regexpMatcher1,
			exp:         true,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			b, err := d.matcher.Match(d.checkedFile)
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
