package main

import "fmt"

type Stack []int

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s *Stack) IsEmpty() bool {
	if len(*s) == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Clear() {
	*s = nil
}
func main() {
	var n int
	var s Stack
	var el int
	fmt.Println("Введите количество элементов стека n>2")
	fmt.Scan(&n)
	fmt.Printf("ДЕМОНСТРАЦИЯ PUSH\n")
	fmt.Printf("Введите %#v элементов (int)\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&el)
		s.Push(el)
	}
	fmt.Println("\nСтек s:")
	fmt.Printf("%v", s)

	fmt.Printf("\n\nДЕМОНСТРАЦИЯ POP")
	s.Pop()
	fmt.Println("\nСтек s:")
	fmt.Printf("%v", s)

	fmt.Printf("\n\nДЕМОНСТРАЦИЯ ISEMPTY")
	fmt.Println("\nСтек s:")
	fmt.Printf("%v", s)
	fmt.Printf("\ns.IsEmpty=%v", s.IsEmpty())

	fmt.Printf("\n\nДЕМОНСТРАЦИЯ SIZE")
	fmt.Println("\nСтек s:")
	fmt.Printf("%v", s)
	fmt.Printf("\ns.Size=%v", s.Size())

	fmt.Printf("\n\nДЕМОНСТРАЦИЯ CLEAR")
	s.Clear()
	fmt.Println("\nСтек s:")
	fmt.Printf("%v ", s)
	fmt.Printf(" s.Size=%v ", s.Size())
	fmt.Printf(" s.IsEmpty=%v", s.IsEmpty())
}
