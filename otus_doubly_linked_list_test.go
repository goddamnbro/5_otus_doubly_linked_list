package otus_doubly_linked_list

import (
	"testing"
)

func areListsEqual(expected, actual []interface{}) bool {
	for index, value := range actual {
		if expected[index] != value {
			return false
		}
	}
	return true
}

func TestDoublyLinkedList(t *testing.T) {
	dl := DoublyLinkedList{}
	dl.InsertToBack("A")
	dl.InsertToBack("B")
	dl.InsertToBack("C")
	dl.InsertToFront(3)
	dl.InsertToFront(2)
	dl.InsertToFront(1)

	expected := []interface{}{1, 2, 3, "A", "B", "C"}
	actual := dl.Values()
	if !areListsEqual(expected, actual) {
		t.Errorf("An order of dobly linked list is incorrect!")
	}
	if dl.Length != 6 {
		t.Errorf("Length is incorrect!")
	}

	itemA := dl.Head.next.next.next
	dl.Remove(itemA)
	expected = []interface{}{1, 2, 3, "B", "C"}
	actual = dl.Values()
	if !areListsEqual(expected, actual) {
		t.Errorf("An order of dobly linked list is incorrect!")
	}
	if dl.Length != 5 {
		t.Errorf("Length is incorrect!")
	}

	item1 := dl.Head
	dl.Remove(item1)
	expected = []interface{}{2, 3, "B", "C"}
	actual = dl.Values()
	if !areListsEqual(expected, actual) {
		t.Errorf("An order of dobly linked list is incorrect!")
	}
	if dl.Length != 4 {
		t.Errorf("Length is incorrect!")
	}

	itemC := dl.Tail
	dl.Remove(itemC)
	expected = []interface{}{2, 3, "B"}
	actual = dl.Values()
	if !areListsEqual(expected, actual) {
		t.Errorf("An order of dobly linked list is incorrect!")
	}
	if dl.Length != 3 {
		t.Errorf("Length is incorrect!")
	}
}
