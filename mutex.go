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
		instance *mutex.Instance
		setting  defaultSetting
		locks    sync.Map
	}
	defaultSetting struct {
	}
	defaultValue struct {
		Expire time.Time
	}
)

func (driver *defaultDriver) Connect(inst *mutex.Instance) (mutex.Connect, error) {
	setting := defaultSetting{}
	return &defaultConnect{
		instance: inst, setting: setting,
	}, nil
}

// 打开连接
// 待处理，需要一个定时器，定期清理过期的数据
func (this *defaultConnect) Open() error {
	return nil
}

// 关闭连接
func (this *defaultConnect) Close() error {
	return nil
}

// 待优化，加上超时设置
func (this *defaultConnect) Lock(key string, expire time.Duration) error {
	now := time.Now()

	if vv, ok := this.locks.Load(key); ok {
		if tm, ok := vv.(defaultValue); ok {
			if tm.Expire.UnixNano() > now.UnixNano() {
				return errors.New("existed")
			}
		}
	}

	if expire <= 0 {
		expire = this.instance.Config.Expire
	}

	value := defaultValue{
		Expire: now.Add(this.instance.Config.Expire),
	}

	this.locks.Store(key, value)

	return nil
}
func (this *defaultConnect) Unlock(key string) error {
	this.locks.Delete(key)
	return nil
}
