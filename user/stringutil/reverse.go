package stringutil
// package stringutil contains utility functions
// for working with package stringutil

// reverse returns argument rune-wise left
func Reverse(s string) string {
  r := []rune(s)

  for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
    r[i], r[j] = r[j], r[i]
  }

  // string is readonly slice of bytes
  // does not need to be unicode format
  // indexing a string accesses individual bytes, not characters
  // only string literals are utf-8
  return string(r)
}

// `go build` --> test that package compiles
// `go install` --> produce output file
