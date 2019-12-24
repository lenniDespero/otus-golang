package hw6

type List struct {
	last, first *Item
	length      int
}

type Item struct {
	data interface{}
	prev *Item
	next *Item
	list *List
}

// Returns List Length
func (l *List) Len() int {
	return l.length
}

// Returns first item from list
func (l *List) First() *Item {
	return l.first
}

// Returns last item from list
func (l *List) Last() *Item {
	return l.last
}

// Returns value
func (n *Item) Value() interface{} {
	return n.data
}

// Returns next item
func (n *Item) Next() *Item {
	return n.next
}

// Returns prev list
func (n *Item) Prev() *Item {
	return n.prev
}

// Add item as first
func (l *List) PushFront(v interface{}) *Item {
	item := Item{data: v, list: l}
	if l.last == nil {
		l.last = &item
	}
	if l.first != nil {
		item.next, l.first.prev = l.first, &item
	}
	l.first = &item
	l.length++
	return &item
}

// adds item as last
func (l *List) PushBack(v interface{}) *Item {
	item := Item{data: v, list: l}
	if l.last != nil {
		item.prev, l.last.next = l.last, &item
	}
	if l.first == nil {
		l.first = &item
	}

	l.last = &item
	l.length++
	return &item
}

func (l *List) Remove(i *Item) {
	if i.list == l {
		if i.prev != nil && i.next != nil {
			i.prev.next, i.next.prev = i.next, i.prev
		}
		if i.prev == nil && i.next != nil {
			i.next.prev, l.first = nil, i.next
		}
		if i.prev != nil && i.next == nil {
			i.prev.next, l.last = nil, i.prev
		}
		l.length--
	}
}
