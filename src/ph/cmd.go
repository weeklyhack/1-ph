package main

import (
  "fmt"
  "os/exec"
)

func RunGit(action string, output RemoteBranchGroup, flags string) {
  for _, remote := range output.Remote {
    for _, branch := range output.Branch {
      fmt.Println("-> $ git", action, remote, branch, flags)

      // run the command
      var cmd *exec.Cmd

      if len(flags) > 0 {
        cmd = exec.Command("git", action, remote, branch, flags)
      } else {
        cmd = exec.Command("git", action, remote, branch)
      }
      out, err := cmd.CombinedOutput()

      if err != nil {
        fmt.Println("An error occurred D:")
        fmt.Printf("%s\n", err)
      }
      fmt.Printf("%s", out)
    }
  }
}
