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
  { // search for current branch
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origincurrent",
    []string{"origin"}, []string{"current"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "oc",
    []string{"origin"}, []string{"current"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origincurrent",
    []string{"origin"}, []string{"current"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "ocurrent",
    []string{"origin"}, []string{"current"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "originc",
    []string{"origin"}, []string{"current"},
  },
  { // cannot mix two remotes/branches
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "masherokuter",
    []string{"heroku"}, []string{"master"},
  },
  { // colon-delimited branches
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origin dev:master",
    []string{"origin"}, []string{"dev:master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "od:m",
    []string{"origin"}, []string{"dev:master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "oc:m",
    []string{"origin"}, []string{"current:master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origin c:m",
    []string{"origin"}, []string{"current:master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "dev:master origin",
    []string{"origin"}, []string{"dev:master"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origin master:dev",
    []string{"origin"}, []string{"master:dev"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "heroku m:d",
    []string{"heroku"}, []string{"master:dev"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "hm:d",
    []string{"heroku"}, []string{"master:dev"},
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
