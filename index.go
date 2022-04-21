package mutex_redis

import (
	"github.com/chefsgo/chef"
)

func Driver() chef.MutexDriver {
	return &redisMutexDriver{}
}

func init() {
	chef.Register("redis", Driver())
}
