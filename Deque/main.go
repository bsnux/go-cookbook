// Package deque represents a double-ended queue, which is an indexed sequence container
// that allows fast insertion and deletion at both its beginning and its end.
package deque

// Deque represents a deque as a slices of integers
type Deque struct {
	Items []int
}

// NewDeque creates a deque
func NewDeque() *Deque {
	return &Deque{}
}

// PushRight insert san element at the end
func (d *Deque) PushRight(item int) {
	d.Items = append(d.Items, item)
}

// PushLeft inserts an element at the beginning
func (d *Deque) PushLeft(item int) {
	temp := []int{item}
	d.Items = append(temp, d.Items...)
}

// PopLeft deletes first element
func (d *Deque) PopLeft() int {
	left := d.Items[0]
	d.Items = d.Items[1:]
	return left
}

// PopRight deletes last element
func (d *Deque) PopRight() int {
	last := len(d.Items) - 1
	temp := d.Items[last]
	// Using variadic arguments to append two slices to another one
	d.Items = append(d.Items[:last], d.Items[last+1:]...)
	return temp
}

// IsEmpty returns true if deque contains no data
func (d *Deque) IsEmpty() bool {
	if len(d.Items) == 0 {
		return true
	}
	return false
}
