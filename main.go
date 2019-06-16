package otus_doubly_linked_list

//Item   // элемент списка
//Value() interface{}  // возвращает значение
//Nex() *Item          // следующий Item
//Prev() *Item         // предыдущий
//Remove()

//List      // тип контейнер
//Len()   // длинна списка
//First() // первый Item
//Last()  // последний Item
//PushFront(v interface{}) // добавить значение в начало
//PushBack(v interface{})  // добавить значение в конец
//

// DoublyLinkedList and his methods
type DoublyLinkedList struct {
	Length uint
	Head   *Item
	Tail   *Item
}

func (dl *DoublyLinkedList) InsertToFront(value interface{}) {
	item := &Item{value: value}

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
	item := &Item{value: value}

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

	// looking for item
	listItem := dl.Head
	for listItem.next != nil {
		if listItem.next == item {
			listItem.next = listItem.next.next
			listItem.next.prev = listItem
			dl.Length--
			return
		}
		listItem = listItem.next
	}
}

func (dl *DoublyLinkedList) Values() []interface{} {
	var values []interface{}

	if dl.Length > 0 {
		item := dl.Head
		values = append(values, item.value)
		for item.Next() != nil {
			item = item.Next()
			values = append(values, item.value)
		}
	}

	return values
}

// Item and his methods
type Item struct {
	value interface{}
	next  *Item
	prev  *Item
}

func (item *Item) Value() interface{} {
	return item.Value
}

func (item *Item) Next() *Item {
	return item.next
}

func (item *Item) Prev() *Item {
	return item.prev
}

//func main() {
//
//	dl := DoublyLinkedList{}
//
//	dl.InsertToBack("A")
//	dl.InsertToBack("B")
//	dl.InsertToFront(2)
//	dl.InsertToFront(1)
//
//	fmt.Println("dl -->", dl)
//	fmt.Println("values -->", dl.Values())
//	fmt.Println("len -->", dl.Length)
//
//	item := dl.Head
//	fmt.Println("item -->", item.value)
//	dl.Remove(item)
//	fmt.Println("values -->", dl.Values())
//	fmt.Println("len -->", dl.Length)
//
//}
