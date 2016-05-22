package main

import (
  "fmt"
  "regexp"
)

type RemoteBranchGroup struct {
  Remote []string
  Branch []string
}

// given a slug used array, start and end, this function returns a bool
// singifying if the region specified is clear (ie, all false)
func indexAreaFree(slugCharsUsed []bool, start int, end int) bool {
  for ct, i := range slugCharsUsed {
    if ct > start && ct < end && i == true {
      return false
    }
  }
  return true
}

// given a haystack and a collection of needles, return all items that are
// contained within. Also, mark the used charaters within haystackCharsUsed
func itemsWithin(haystack string, collection []string, haystackCharsUsed []bool) []string {
  var totals []string

  for _, item := range collection {
    match := regexp.MustCompile(item).FindStringIndex(haystack)
    fmt.Println(item, match)
    if (len(match) != 0) {
      // if the specified area hasn't been used, then we're good
      if indexAreaFree(haystackCharsUsed, match[0], match[1]) {
        totals = append(totals, item)

        // mark chars as used
        for k := match[0]; k < match[1]; k++ {
          haystackCharsUsed[k] = true;
        }
      }
    }
  }

  return totals
}

func Parse(config RemoteBranchGroup, slug string) RemoteBranchGroup {
  var branches []string

  // create an array, specifying whether a charater has been used
  var slugCharsUsed []bool
  for k := 0; k < len(slug); k++ {
    slugCharsUsed = append(slugCharsUsed, false)
  }

  // get all matching full-length remotes
  remotes := itemsWithin(slug, config.Remote, slugCharsUsed)
  fmt.Println(slugCharsUsed)

  // For branches

  return RemoteBranchGroup{Remote: remotes, Branch: branches}
}
