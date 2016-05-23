package main

import (
  "testing"
  "reflect"
)

// ParseGitBranches
var branchTests = []struct {
  In string
  Out []string
} {
  {"* master\n  dev\n  other", []string{"master", "dev", "other"}},
  {"  master\n* other\n  dev", []string{"master", "other", "dev"}},
  {"  f04%*()\n* other\n  dev", []string{"f04%*()", "other", "dev"}},
}

func TestBranch(t *testing.T) {
  for _, tt := range branchTests {
    out := ParseGitBranches(tt.In)
    if !reflect.DeepEqual(out, tt.Out) {
      t.Errorf("GitBranch: should have been %s, was really %s", tt.Out, out)
    }
  }
}
