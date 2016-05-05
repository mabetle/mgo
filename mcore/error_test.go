package mcore

import (
	"testing"
	. "github.com/mabetle/mcore/mtest"
	"errors"
)


func TestIsReturnHasError(t *testing.T) {
	AssertFalse(IsReturnHasError())
	AssertFalse(IsReturnHasError("Hello"))

	AssertTrue(IsReturnHasError(errors.New("err")))
	AssertTrue(IsReturnHasError("Hello", errors.New("err")))

	AssertFalse(IsReturnHasError("Hello", nil))

}

