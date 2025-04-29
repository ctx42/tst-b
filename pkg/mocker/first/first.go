package first

import (
	"github.com/ctx42/tst-a/pkg/mocker/second"
)

type First interface {
	Method0() error
	Method1(a *second.T0) error
}
