package pkg

import (
	assist "CheckExplicitEmbeddedSelectorassist"
	"io"
)

func fnQualified() {
	_ = io.EOF.Error  // minimal form
	_ = assist.V.T2.F //@ diag(`could remove embedded field "T2" from selector`)
	_ = assist.V.F    // minimal form
}
