// Package mypkg2 tests that it can import a vendored dep.
package mypkg2

import "github.com/alice/mypkg"

func init() {
  mypkg.Foo()
}
