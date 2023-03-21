package mutex

import (
	"github.com/infrago/infra"
	"github.com/infrago/mutex"
)

func Driver() mutex.Driver {
	return &defaultDriver{}
}

func init() {
	infra.Register("default", Driver())
}
