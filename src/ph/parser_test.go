package main

import (
  "fmt"
  "testing"
  "reflect"
)

var parseTests = []struct {
  remote []string
  branch []string
  slug string
  expectedRemote []string
  expectedBranch []string
} {
  { // different permutations of remotes and branches
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origin master",
    []string{"origin"}, []string{"master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origin m",
    []string{"origin"}, []string{"master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "o master",
    []string{"origin"}, []string{"master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "o m",
    []string{"origin"}, []string{"master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "om",
    []string{"origin"}, []string{"master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origin heroku master",
    []string{"origin", "heroku"}, []string{"master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "ohm",
    []string{"origin", "heroku"}, []string{"master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "originhm",
    []string{"origin", "heroku"}, []string{"master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "originhmaster",
    []string{"origin", "heroku"}, []string{"master"},
  },
  { // cannot mix two remotes/branches
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "masherokuter",
    []string{"heroku"}, []string{"master"},
  },
}

func TestParse(t *testing.T) {
  for _, tt := range parseTests {
    // start with the remotes and origins that already exist
    config := RemoteBranchGroup{
      Remote: tt.remote,
      Branch: tt.branch,
    }

    output, left := Parse(config, tt.slug)
    fmt.Println(output, left)

    if !reflect.DeepEqual(output.Remote, tt.expectedRemote) {
      t.Errorf("Parse: Remote should be %s, was actually %s", tt.expectedRemote, output.Remote)
    }
    if !reflect.DeepEqual(output.Branch, tt.expectedBranch) {
      t.Errorf("Parse: Branch should be %s, was actually %s", tt.expectedBranch, output.Branch)
    }
  }
}
