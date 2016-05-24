package main

import (
  "testing"
  "reflect"
)

// DecodeCommand
var decoderTests = []struct {
  In string
  Out string
} {
  {"git push origin master", "om"},
  {"git push origin dev", "od"},
  {"git push origin dev --flag", "od"},
  {"git push origin BrAnCh", "oB"},

  {"git pull origin master", "lom"},
  {"git pull origin dev", "lod"},
  {"git pull origin BrAnCh", "loB"},
  {"git pull origin dev --flag", "lod"},

  {"git push origin", ""},
  {"git pull origin", ""},
  {"git push", ""},
  {"git pull", ""},
  {"git log", ""},
  {"not a git command", ""},
}

func TestDecodeCommand(t *testing.T) {
  for _, tt := range decoderTests {
    out := DecodeCommand(tt.In)
    if !reflect.DeepEqual(out, tt.Out) {
      t.Errorf("DecodeCommand: should have been %s, was really %s", tt.Out, out)
    }
  }
}
