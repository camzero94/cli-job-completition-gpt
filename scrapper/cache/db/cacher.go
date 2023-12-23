package db

type Cacher interface {
	Get(key string) (string, bool)
	Set(key string, val string) error
	Remove(key string) error
}

