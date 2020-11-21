package kvutil

type KeyValue struct {
	Key   string
	Value string
}

func Get(src *[]KeyValue, key string) *KeyValue {
	for _, item := range *src {
		if item.Key == key {
			return &item
		}
	}
	return nil
}

func Set(dest *[]KeyValue, kv KeyValue) {
	for idx, item := range *dest {
		if item.Key == kv.Key {
			(*dest)[idx] = kv
			break
		}
	}
}
