package collection

type OrderedDict struct {
	lookup map[string]*LinkedListNode
	list   *LinkedList
}

func NewOrderedDict() *OrderedDict {
	return &OrderedDict{
		lookup: make(map[string]*LinkedListNode),
		list:   NewLinkedList(),
	}
}

func (d *OrderedDict) Set(key string, value interface{}) {
	if n, ok := d.lookup[key]; ok {
		d.list.Remove(n)
	}

	d.lookup[key] = d.list.Append(value)
}

func (d *OrderedDict) Get(key string) interface{} {
	if n, ok := d.lookup[key]; ok {
		return n.Value()
	}
	return nil
}

func (d *OrderedDict) Remove(key string) bool {
	if n, ok := d.lookup[key]; ok {
		if ok := d.list.Remove(n); !ok {
			return false
		}
		delete(d.lookup, key)
		return true
	}
	return false
}

func (d *OrderedDict) Iterate() chan interface{} {
	ch := make(chan interface{})

	go func() {
		for v := range d.list.Iterate() {
			ch <- v.Value()
		}

		close(ch)
	}()

	return ch
}
