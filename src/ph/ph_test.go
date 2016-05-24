package main

import (
  "testing"
  "reflect"
)

// DecodeCommand
var parseArgsTests = []struct {
  In []string
  Action string
  Output RemoteBranchGroup
  Flags string
} {
  {
    []string{"ph", "origin", "master"},
    "push",
    RemoteBranchGroup{Remote: []string{"origin"}, Branch: []string{"master"}},
    "",
  },
  {
    []string{"ph", "origin", "m"},
    "push",
    RemoteBranchGroup{Remote: []string{"origin"}, Branch: []string{"master"}},
    "",
  },
  {
    []string{"ph", "o", "master"},
    "push",
    RemoteBranchGroup{Remote: []string{"origin"}, Branch: []string{"master"}},
    "",
  },
  {
    []string{"ph", "om"},
    "push",
    RemoteBranchGroup{Remote: []string{"origin"}, Branch: []string{"master"}},
    "",
  },
  {
    []string{"ph", "pull", "origin", "master"},
    "pull",
    RemoteBranchGroup{Remote: []string{"origin"}, Branch: []string{"master"}},
    "",
  },
  {
    []string{"ph", "pull", "origin", "master", "f"},
    "pull",
    RemoteBranchGroup{Remote: []string{"origin"}, Branch: []string{"master"}},
    "--force",
  },
  {
    []string{"ph", "pull", "origin", "master", "n"},
    "pull",
    RemoteBranchGroup{Remote: []string{"origin"}, Branch: []string{"master"}},
    "--dry-run",
  },
  {
    []string{"ph", "pull", "origin", "master", "t"},
    "pull",
    RemoteBranchGroup{Remote: []string{"origin"}, Branch: []string{"master"}},
    "--set-upstream",
  },
  {
    []string{"ph", "pull", "origin", "master", "q"},
    "pull",
    RemoteBranchGroup{Remote: []string{"origin"}, Branch: []string{"master"}},
    "--quiet",
  },
  {
    []string{"ph", "pull", "origin", "master", "v"},
    "pull",
    RemoteBranchGroup{Remote: []string{"origin"}, Branch: []string{"master"}},
    "--verbose",
  },
  {
    []string{"ph", "pull", "om", "f"},
    "pull",
    RemoteBranchGroup{Remote: []string{"origin"}, Branch: []string{"master"}},
    "--force",
  },
  {
    []string{"ph", "oml"},
    "pull",
    RemoteBranchGroup{Remote: []string{"origin"}, Branch: []string{"master"}},
    "",
  },
}

func TestParseArgs(t *testing.T) {
  for _, tt := range parseArgsTests {
    action, output, flags := ParseArgs(tt.In)
    if !(
      reflect.DeepEqual(action, tt.Action) &&
      reflect.DeepEqual(output, tt.Output) &&
      reflect.DeepEqual(flags, tt.Flags)) {
      t.Errorf("ParseArgs: %s didn't run to '%s', %s, '%s'", tt.In, tt.Action, tt.Output, tt.Flags)
    }
  }
}
