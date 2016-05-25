package main

import (
  // "fmt"
  "testing"
  "reflect"
)

var parseTests = []struct {
  remote []string
  branch []string
  slug string
  expectedRemote []string
  expectedBranch []string
  expectedFlags []string
} {
  { // different permutations of remotes and branches
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origin master",
    []string{"origin"}, []string{"master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origin m",
    []string{"origin"}, []string{"master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "o master",
    []string{"origin"}, []string{"master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "o m",
    []string{"origin"}, []string{"master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "om",
    []string{"origin"}, []string{"master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origin heroku master",
    []string{"origin", "heroku"}, []string{"master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "ohm",
    []string{"origin", "heroku"}, []string{"master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "originhm",
    []string{"origin", "heroku"}, []string{"master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "originhmaster",
    []string{"origin", "heroku"}, []string{"master"}, []string{},
  },
  { // search for current branch
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origincurrent",
    []string{"origin"}, []string{"current"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "oc",
    []string{"origin"}, []string{"current"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origincurrent",
    []string{"origin"}, []string{"current"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "ocurrent",
    []string{"origin"}, []string{"current"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "originc",
    []string{"origin"}, []string{"current"}, []string{},
  },
  { // cannot mix two remotes/branches
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "masherokuter",
    []string{"heroku"}, []string{"master"}, []string{},
  },
  { // colon-delimited branches
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origin dev:master",
    []string{"origin"}, []string{"dev:master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "od:m",
    []string{"origin"}, []string{"dev:master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "oc:m",
    []string{"origin"}, []string{"current:master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origin c:m",
    []string{"origin"}, []string{"current:master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "dev:master origin",
    []string{"origin"}, []string{"dev:master"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "origin master:dev",
    []string{"origin"}, []string{"master:dev"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "heroku m:d",
    []string{"heroku"}, []string{"master:dev"}, []string{},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "hm:d",
    []string{"heroku"}, []string{"master:dev"}, []string{},
  },
  { // support flags in addition to normal arguments
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "hm:d --flag",
    []string{"heroku"}, []string{"master:dev"}, []string{"--flag"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "hm:d --flag value",
    []string{"heroku"}, []string{"master:dev"}, []string{"--flag value"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "hm:d -f",
    []string{"heroku"}, []string{"master:dev"}, []string{"-f"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "hm:d -f value",
    []string{"heroku"}, []string{"master:dev"}, []string{"-f value"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "hm:d --with-dash",
    []string{"heroku"}, []string{"master:dev"}, []string{"--with-dash"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "hm:d --with-dash value",
    []string{"heroku"}, []string{"master:dev"}, []string{"--with-dash value"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "hm:d -- -a",
    []string{"heroku"}, []string{"master:dev"}, []string{"--", "-a"},
  },
  {
    []string{"origin", "heroku"}, []string{"master", "dev"},
    "hm:d -- --flag",
    []string{"heroku"}, []string{"master:dev"}, []string{"--", "--flag"},
  },
}

func TestParse(t *testing.T) {
  for _, tt := range parseTests {
    // start with the remotes and origins that already exist
    config := RemoteBranchGroup{
      Remote: tt.remote,
      Branch: tt.branch,
    }

    output, _, flags := Parse(config, tt.slug)

    if !reflect.DeepEqual(output.Remote, tt.expectedRemote) {
      t.Errorf("Parse: Remote should be %s, was actually %s", tt.expectedRemote, output.Remote)
    }
    if !reflect.DeepEqual(output.Branch, tt.expectedBranch) {
      t.Errorf("Parse: Branch should be %s, was actually %s", tt.expectedBranch, output.Branch)
    }
    if len(flags) > 0 && !reflect.DeepEqual(flags, tt.expectedFlags) {
      t.Errorf("Parse: Flags should be %s, was actually %s", tt.expectedFlags, flags)
    }
  }
}
