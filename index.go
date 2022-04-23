package mutex_redis

import (
	"github.com/chefsgo/mutex"
)

func Driver() mutex.Driver {
	return &redisDriver{}
}

func init() {
	mutex.Register("redis", Driver())
}
