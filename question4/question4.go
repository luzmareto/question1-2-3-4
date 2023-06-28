package question_4

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(num int) {
	s.data = append(s.data, num)
}

func (s *Stack) Pop() int {
	if len(s.data) > 0 {
		index := len(s.data) - 1
		num := s.data[index]
		s.data = s.data[:index]
		return num
	}
	return 0
}

func (s *Stack) Peek() int {
	if len(s.data) > 0 {
		return s.data[len(s.data)-1]
	}
	return 0
}

func Question4() {
	stack := Stack{}

	bilangan := []int{2, 5, 10, 3, 8}

	for _, num := range bilangan {
		stack.Push(num)
	}

	fmt.Println("Stack:", stack.data)

	// Pop the topmost element from the stack
	popped := stack.Pop()
	fmt.Println("Popped:", popped)

	fmt.Println("Stack after pop:", stack.data)

	// Peek the topmost element from the stack
	peeked := stack.Peek()
	fmt.Println("Peeked:", peeked)

	fmt.Println("Stack after peek:", stack.data)
}
