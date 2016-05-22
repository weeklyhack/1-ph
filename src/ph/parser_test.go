package main

import (
  "fmt"
  "testing"
)

// cases := []struct {
//   remote []string
//   branch []string
//   input string
//   result []string
// }{
//   {}
// }

func TestParse(t *testing.T) {
  // start with the remotes and origins that already exist
  config := RemoteBranchGroup{
    Remote: []string { "origin", "heroku" },
    Branch: []string { "master", "dev" },
  }

  output := Parse(config, "abc")
  if output.Remote == []string {"abc", "def"} {
    t.Fatalf("bad")
  }
}
