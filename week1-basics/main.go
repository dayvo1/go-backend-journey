package main

func main() {
	//backendJourney()
	//guessingGame()
	//calculator()

	student := Student{
		Name: "Test Name",
		Id:   "12345",
		Gpa:  3.5,
	}

	student.Display()
	student.UpdateGpa(4.0)
	student.Display()

}
