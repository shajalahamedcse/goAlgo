package doublyLinkedList

type DoublyLinkedList struct {
	first *Element
	last  *Element
	size  int
}

type Element struct {
	value interface{}
	prev  *Element
	next  *Element
}

// Instantiates a new empty list
func NewList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Appends one or more at the end of the list
func (list *DoublyLinkedList) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &Element{}
		newElement.value = value
		newElement.prev = list.last

		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

// Appends a value (one or more) at the end of the list (same as Add())
func (list *DoublyLinkedList) Append(values ...interface{}) {
	list.Add(values...)
}

// Prepends a values (or more)

func (list *DoublyLinkedList) Prepend(values ...interface{}) {
	// in reverse to keep passed order i.e. ["3","4"] -> Prepend(["1","2"]) -> ["1","2","3","4"]

	for v := len(values) - 1; v >= 0; v-- {
		newElement := &Element{}
		newElement.value = values[v]
		newElement.next = list.first

		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.first.prev = newElement
			list.first = newElement
		}
		list.size++
	}
}
