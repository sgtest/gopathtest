// Package gopathtest tests that it can access vendored deps in vendor/.
package gopathtest

import (
  "github.com/alice/mypkg"
  _ "github.com/alice/mypkg2"
)

func init() {
  mypkg.Foo()
}
