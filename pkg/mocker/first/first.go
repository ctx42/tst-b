package first

import (
	"github.com/ctx42/tst-a/pkg/mocker/third"
)

type First interface {
	Method0() error
	Method1(a *third.Third) error
}
