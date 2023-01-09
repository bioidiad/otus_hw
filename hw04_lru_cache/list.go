package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int // Place your code here.
	front *ListItem
	back  *ListItem
}

func NewList() List {
	return new(list)
}

func (lst *list) Len() int {
	return lst.len
}

func (lst *list) Front() *ListItem {
	return lst.front
}

func (lst *list) Back() *ListItem {
	return lst.back
}

func (lst *list) firstElem(v interface{}) *ListItem {
	firstElem := ListItem{v, nil, nil}
	lst.front = &firstElem
	lst.back = &firstElem
	lst.len = 1
	fmt.Println("First element created", firstElem.Value, firstElem.Next, firstElem.Prev, lst.Len(), lst.Front(), lst.Back()) //nolint:lll
	return &firstElem
}

func (lst *list) PushFront(v interface{}) *ListItem {
	if lst.len == 0 {
		return lst.firstElem(v)
	}
	var newElem ListItem

	newElem.Value = v
	newElem.Next = lst.Front()
	lst.Front().Prev = &newElem
	lst.front = &newElem
	lst.len++
	return &newElem
}

func (lst *list) PushBack(v interface{}) *ListItem {
	if lst.len == 0 {
		return lst.firstElem(v)
	}
	var newElem ListItem

	newElem.Value = v
	newElem.Prev = lst.Back()
	lst.Back().Next = &newElem
	lst.back = &newElem
	lst.len++
	return &newElem
}

func (lst *list) Remove(i *ListItem) {
	if i != lst.front {
		i.Prev.Next = i.Next
	}
	if i != lst.back {
		i.Next.Prev = i.Prev
	}
	if lst.len > 0 {
		lst.len--
	}
}

func (lst *list) MoveToFront(i *ListItem) {
	lst.PushFront(i.Value)
	lst.Remove(i)
}
