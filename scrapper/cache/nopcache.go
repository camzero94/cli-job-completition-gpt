package main 

type Cache interface{
	Get (string) (string, bool)
	Set (string, string) error
	Remove (string) error
}

type NopCache struct{}

func (c NopCache) Get(string) (string, bool) {
	return "", false 
}

func (c NopCache) Set(string, string) error {
	return nil
}

func (c NopCache) Remove(string) error {
	return nil
}

