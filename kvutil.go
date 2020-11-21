package kvutil

type KV struct {
	Key   string
	Value string
}

func Get(src *[]KV, key string) *KV {
	for _, item := range *src {
		if item.Key == key {
			return &item
		}
	}
	return nil
}

func Set(dest *[]KV, kv KV) {
	for idx, item := range *dest {
		if item.Key == kv.Key {
			(*dest)[idx] = kv
			break
		}
	}
}
