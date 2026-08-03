// package main

// import "fmt"

// func Find[T comparable](s []T, v T) (T, bool) {
// 	for _, elem := range s {
// 		if elem == v {
// 			return elem, true
// 		}
// 	}
// 	var empty T
// 	return empty, false
// }

//	func main() {
//		s := []int{1, 2, 3, 4, 5}
//		elem, ok := Find(s, 0)
//		fmt.Println(elem, ok)
//	}
package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) String() string {
	var result string
	var counter int
	current := l
	for current.next != nil {
		result += fmt.Sprintf("Node(%v, ", current.val)
		current = current.next
		counter++
	}
	result += fmt.Sprintf("Node(%v, nil)", current.val)

	return result
}

func (l *List[T]) Length() int {
	if l == nil {
		return 0
	}
	current := l
	length := 1
	for current.next != nil {
		current = current.next
		length++
	}
	return length
}

func (l *List[T]) Head() (T, bool) {
	if l == nil {
		var result T
		return result, false
	}
	return l.val, true
}

func (l *List[T]) Last() (T, bool) {
	if l == nil {
		var result T
		return result, false
	}
	current := l
	for current.next != nil {
		current = current.next
	}
	return current.val, true
}

func main() {
	l := &List[int]{val: 4, next: &List[int]{val: 5, next: nil}}
	l1 := &List[int]{val: 4, next: nil}
	l2 := &List[int]{val: 4, next: &List[int]{val: 5, next: &List[int]{val: 6, next: nil}}}
	l3 := &List[int]{}
	fmt.Println(l.Length())
	fmt.Println(l1.Length())
	fmt.Println(l2.Length())
	fmt.Println(l3.Length())
}
