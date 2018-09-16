/*
Copyright (c) Shajal Ahamed, All rights reserved.


This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 3.0 of the License, or (at your option) any later version.
This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.
You should have received a copy of the GNU Lesser General Public
License along with this library. See the file LICENSE included
with this distribution for more information.


*/

/*
Implementation of stack using a singly linked list
*/

package linkedListStack

type Stack struct {
	size int
	top  *Element
}

type Element struct {
	value interface{}
	next  *Element
}

// A new empty stack
func NewStack() *Stack {
	return &Stack{}
}

//Pushes a value onto the top of the stack
func (stack *Stack) Push(value interface{}) {
	e := &Element{}
	e.value = value
	e.next = stack.top
	stack.top = e

	stack.size++
}

// Pops (removes) top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *Stack) Pop() (value interface{}, ok bool) {
	if stack.size > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--

		return value, true
	}
	return nil, false
}

// Returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (stack *Stack) Peek() (value interface{}, ok bool) {
	if stack.size > 0 {
		return stack.top.value, true
	}
	return nil, false
}

func (stack *Stack) Empty() bool {
	return stack.size == 0
}

func (stack *Stack) Size() int {
	return stack.size
}
