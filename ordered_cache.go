package gcache

type OrderedCache interface {
	//for ordered (queues) cache,
	EnQueue(interface{}, interface{}) error
	EnQueueBatch([]interface{}, []interface{}) error //EnQueueBatch if searchFunc is not nil , keys are  required sorted
	DeQueue() (interface{}, interface{}, error)
	DeQueueBatch(count int) ([]interface{}, []interface{}, error)
	OrderedKeys() []interface{}
	MoveFront(interface{}) error               //move an element to front
	GetTop() (interface{}, interface{}, error) //get top element
	Prepend(interface{}, interface{}) error
	PrependBatch([]interface{}, []interface{}) error //PrependBatch if searchFunc is not nil , keys are  required sorted
	RemoveExpired(allowFailCount int) error
	Get(interface{}) (interface{}, error)
	GetIFPresent(interface{}) (interface{}, error)
	GetALL() map[interface{}]interface{}
	GetKeysAndValues() ([]interface{}, []interface{})
	get(interface{}, bool) (interface{}, error)
	Remove(interface{}) bool
	PrintValues(int)
	Purge()
	Refresh()
	Keys() []interface{}
	Len() int
	Sort()
	statsAccessor
}

