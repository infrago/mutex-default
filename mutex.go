package mutex

import (
	"errors"
	"sync"
	"time"

	"github.com/infrago/mutex"
)

//默认mutex驱动

type (
	defaultDriver  struct{}
	defaultConnect struct {
		name    string
		config  mutex.Config
		setting defaultSetting
		locks   sync.Map
	}
	defaultSetting struct {
	}
	defaultValue struct {
		Expiry time.Time
	}
)

func (driver *defaultDriver) Connect(name string, config mutex.Config) (mutex.Connect, error) {
	setting := defaultSetting{}
	return &defaultConnect{
		name: name, config: config, setting: setting,
	}, nil
}

// 打开连接
// 待处理，需要一个定时器，定期清理过期的数据
func (connect *defaultConnect) Open() error {
	return nil
}

// 关闭连接
func (connect *defaultConnect) Close() error {
	return nil
}

// 待优化，加上超时设置
func (connect *defaultConnect) Lock(key string, expiry time.Duration) error {
	now := time.Now()

	if vv, ok := connect.locks.Load(key); ok {
		if tm, ok := vv.(defaultValue); ok {
			if tm.Expiry.UnixNano() > now.UnixNano() {
				return errors.New("existed")
			}
		}
	}

	if expiry <= 0 {
		expiry = connect.config.Expiry
	}

	value := defaultValue{
		Expiry: now.Add(connect.config.Expiry),
	}

	connect.locks.Store(key, value)

	return nil
}
func (connect *defaultConnect) Unlock(key string) error {
	connect.locks.Delete(key)
	return nil
}
