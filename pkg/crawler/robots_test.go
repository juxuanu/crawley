package crawler

import (
	"errors"
	"testing"
)

func TestParseAction(t *testing.T) {
	t.Parallel()

	type testCase struct {
		Have string
		Want RobotsAction
	}

	cases := []testCase{
		{Have: "ignore", Want: RobotsIgnore},
		{Have: "crawl", Want: RobotsCrawl},
		{Have: "respect", Want: RobotsRespect},
	}

	for i, tc := range cases {
		got, err := ParseAction(tc.Have)
		if err != nil {
			t.Errorf("case[%d]: got error: %v", i+1, err)
		}

		if got != tc.Want {
			t.Errorf("case[%d]: unexpected result want: %d got: %d", i+1, tc.Want, got)
		}
	}
}

func TestParseActionErr(t *testing.T) {
	t.Parallel()

	_, err := ParseAction("dsf")
	if err == nil {
		t.Error("no error")
	}

	if !errors.Is(err, ErrActionUnknown) {
		t.Error("unexpected error")
	}
}
