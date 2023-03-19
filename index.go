package mutex

import (
	"github.com/infrago/mutex"
)

func Driver() mutex.Driver {
	return &defaultDriver{}
}

func init() {
	mutex.Register("default", Driver())
}
