package main

import (
  "fmt"
  "strings"
  "os/exec"
)

func RunGit(action string, output RemoteBranchGroup, flags string) {
  for _, remote := range output.Remote {
    for _, branch := range output.Branch {
      // replace any instances of current with the active branch
      activeBranch := GetCwdActiveGitBranch()
      branch = strings.Replace(branch, "current", activeBranch, -1)

      // log what we are about to do
      fmt.Println("-> $ git", action, remote, branch, flags)

      // run the command
      if len(flags) > 0 {
        RunCmd("git", []string{action, remote, branch, flags})
      } else {
        RunCmd("git", []string{action, remote, branch})
      }
    }
  }
}

func RunCmd(bin string, args []string) error {
  cmd := exec.Command(bin, args...)
  out, err := cmd.CombinedOutput()

  fmt.Printf("%s", out)
  if err != nil {
    fmt.Printf("Error: %s\n", err)
    return err
  }
  return nil
}
