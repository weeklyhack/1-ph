package main

import (
  // "fmt"
  "os/exec"
  "strings"
)

// is git currently installed and in the path?
func GitExists() bool {
  out, err := exec.Command("which", "git").Output()
  if err != nil {
    panic(err)
  }

  if len(out) > 0 {
    return true
  } else {
    return false
  }
}

// get the branches, as a slice of strings, in the repository within the current
// directory
func GetCwdBranches() []string {
  cmd := exec.Command("git", "branch")
  out, err := cmd.CombinedOutput()

  if err != nil {
    panic(err)
  }

  return ParseGitBranches(string(out))
}

// parse git branches from a string to a slice
func ParseGitBranches(branchesString string) []string {
  total := strings.Split(branchesString, "\n")
  var branches []string

  for _, branch := range total {
    if len(branch) > 2 {
      branches = append(branches, branch[2:])
    }
  }

  return branches
}

// get all the git remotes for the current working directory
func GetCwdRemotes() []string {
  cmd := exec.Command("git", "remote")
  out, err := cmd.CombinedOutput()

  if err != nil {
    panic(err)
  }

  return strings.Split(strings.Trim(string(out), "\n"), "\n")
}
