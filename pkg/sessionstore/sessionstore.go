package sessionstore

import (
	"github.com/gin-contrib/sessions"
	"github.com/yyyyymmmmm/Test/pkg/cache"
)

type Store interface {
	sessions.Store
}

func NewStore(driver cache.Driver, keyPairs ...[]byte) Store {
	return &store{newKvStore("cd_session_", driver, keyPairs...)}
}

type store struct {
	*kvStore
}

func (c *store) Options(options sessions.Options) {
	c.kvStore.Options = options.ToGorillaOptions()
}
