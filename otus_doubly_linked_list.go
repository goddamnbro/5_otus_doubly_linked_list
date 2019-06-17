package otus_doubly_linked_list

// DoublyLinkedList and his methods
type DoublyLinkedList struct {
	Length uint
	Head   *Item
	Tail   *Item
}

func (dl *DoublyLinkedList) InsertToFront(value interface{}) {
	item := &Item{Value: value}

	if dl.Length == 0 {
		dl.Head = item
		dl.Tail = item
	} else {
		prevHead := dl.Head
		prevHead.prev = item
		dl.Head = item
		item.next = prevHead
	}

	dl.Length++
}

func (dl *DoublyLinkedList) InsertToBack(value interface{}) {
	item := &Item{Value: value}

	if dl.Length == 0 {
		dl.Head = item
		dl.Tail = item
	} else {
		prevTail := dl.Tail
		prevTail.next = item
		dl.Tail = item
		item.prev = prevTail
	}

	dl.Length++
}

func (dl *DoublyLinkedList) Remove(item *Item) {
	// if list is empty
	if dl.Length == 0 {
		return
	}

	// if list contains 1 item
	if dl.Length == 1 {
		if dl.Head == item {
			dl.Head = nil
			dl.Tail = nil
			dl.Length--
		}
		return
	}

	// if we need to remove head item
	if dl.Head == item {
		dl.Head = dl.Head.next
		dl.Head.prev = nil
		dl.Length--
		return
	}
	// if we need to remove tail item
	if dl.Tail == item {
		dl.Tail = dl.Tail.prev
		dl.Tail.next = nil
		dl.Length--
		return
	}

	// if we need to remove item somewhere in the middle of list
	item.prev.next = item.next
	item.next.prev = item.prev
	dl.Length--
}

func (dl *DoublyLinkedList) Values() []interface{} {
	var values []interface{}

	if dl.Length > 0 {
		item := dl.Head
		values = append(values, item.Value)
		for item.Next() != nil {
			item = item.Next()
			values = append(values, item.Value)
		}
	}

	return values
}

// Item and his methods
type Item struct {
	Value interface{}
	next  *Item
	prev  *Item
}

func (item *Item) Next() *Item {
	return item.next
}

func (item *Item) Prev() *Item {
	return item.prev
}
