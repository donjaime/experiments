package approot

import (
	"github.com/donjaime/experiments/a"
	"github.com/donjaime/experiments/b"
)

func init() {
	a.UsedExternally()
	b.UsedExternally()
}
