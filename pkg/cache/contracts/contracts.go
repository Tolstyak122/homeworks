package contracts

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string)
}
