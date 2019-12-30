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
	item := Item{next: l.first, data: v, list: l}
	if l.last == nil {
		l.last = &item
	}
	if l.first != nil {
		l.first.prev = &item
	}
	l.first = &item
	l.length++
	return &item
}

// adds item as last
func (l *List) PushBack(v interface{}) *Item {
	item := Item{prev: l.last, data: v, list: l}
	if l.last != nil {
		l.last.next = &item
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
			l.removeFromMiddle(i)
		} else if i.prev == nil && i.next != nil {
			l.removeFromFront(i)
		} else if i.prev != nil && i.next == nil {
			l.removeFromBack(i)
		}
	}
}

func (l *List) removeFromMiddle(i *Item) {
	if i.prev.next == i && i.next.prev == i {
		i.prev.next, i.next.prev = i.next, i.prev
		i.list = nil
		l.length--
	}
}

func (l *List) removeFromFront(i *Item) {
	if i.next.prev == i {
		i.next.prev, l.first = nil, i.next
		i.list = nil
		l.length--
	}
}

func (l *List) removeFromBack(i *Item) {
	if i.prev.next == i {
		i.prev.next, l.last = nil, i.prev
		i.list = nil
		l.length--
	}
}
