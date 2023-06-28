package delivery

import (
	"fmt"
	question_1 "question-1-2-3-4/question1"
	question_2 "question-1-2-3-4/question2"
	question "question-1-2-3-4/question3"
	question_4 "question-1-2-3-4/question4"
)

func De() {
	fmt.Println("Select an option:")
	fmt.Println("1. Question 1")
	fmt.Println("2. Question 2")
	fmt.Println("3. Question 3a")
	fmt.Println("4. Question 3b")
	fmt.Println("5. Question 4")
	fmt.Println("6. Exit")

	var choice int
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		question_1.Question1()
	case 2:
		question_2.Question2()
	case 3:
		question.Question3a()
	case 4:
		question.Question3b()
	case 5:
		question_4.Question4()
	default:
		fmt.Println("Invalid choice")
	}
}
