package main

import "fmt"

type wwtbm struct {
	Question  string
	options   [4]string
	answer    string
	Correct   int
	Incorrect int
}

type WTBM [15]wwtbm

type Player struct {
	Username string
	Pass     string
	point    int
}

type Account [100]Player

var Count, CountP int
var Name string

func main() {
	var role int
	var pass string
	var game WTBM
	var player Account
	var validRole bool
	for {
		menu()
		fmt.Scan(&role)
		validRole = (role >= 1 && role <= 2)
		if !validRole {
			fmt.Println("Invalid role number. Please try again.")
		} else {
			if role == 1 {
				fmt.Print("Please enter the admin password: ")
				fmt.Scan(&pass)
				if pass != "admin" {
					fmt.Println("Incorrect password. Returning to menu...")
				} else {
					// Lakukan sesuatu jika password benar
					fmt.Println("Welcome, Admin!")
					adminMenu(&game, &Count, &player)
				}
			} else if role == 2 {
				// Lakukan sesuatu untuk pemain
				fmt.Println("Welcome, Player!")
				menuPlayer(&player, &CountP, &game, &Name)
			} else {
				fmt.Println("Invalid role number. Please try again.")
			}
		}
	}
}

func menu() {
	fmt.Println("-----------------------------------")
	fmt.Println("  Who Wants To Be a Millionaire")
	fmt.Println("-----------------------------------")
	fmt.Println("1. Admin")
	fmt.Println("2. Player")
	fmt.Println("-----------------------------------")
	fmt.Print("Please insert your role number: ")
}

func adminMenu(game *WTBM, Count *int, player *Account) {
	var choice int
	for choice != 7 {
		fmt.Println("\nAdmin Menu:")
		fmt.Println("1. Add Question")
		fmt.Println("2. Edit Question")
		fmt.Println("3. Delete Question")
		fmt.Println("4. View All Questions")
		fmt.Println("5. View Players")
		fmt.Println("6. View Top Questions")
		fmt.Println("7. Exit")
		fmt.Print("Please select an option: ")

		fmt.Scan(&choice)
		fmt.Println()
		switch choice {
		case 1:
			addQuestion(game, Count)
		case 2:
			editQuestion(game, Count)
		case 3:
			deleteQuestion(game, Count)
		case 4:
			viewQuestion(game, Count)
		case 5:
			viewPlayer(player)
		case 6:
			viewTopQuestions(game, Count)
		case 7:
			fmt.Println("Exiting Admin Menu....")
			return
		default:
			fmt.Println("Invalid choice, Please try again.")
		}
	}
}

func addQuestion(game *WTBM, Count *int) {
	var j, temp, JumlahSoal int

	if *Count == 0 {
		fmt.Println("How Many Question You Want to Add ")
		fmt.Scan(&JumlahSoal)
		j = 1
		for j <= JumlahSoal {
			fmt.Println("Enter the question: ")
			fmt.Scan(&game[j].Question)

			for i := 0; i < 4; i++ {
				fmt.Printf("Enter option %d: ", i+1)
				fmt.Scan(&game[j].options[i])
			}

			fmt.Println("Enter the Answer: ")
			fmt.Scan(&game[j].answer)
			j++
			fmt.Println("Question added successfully.")
			*Count++
		}
	} else {
		fmt.Println("How Many Question You Want to Add ")
		fmt.Scan(&JumlahSoal)
		temp = *Count
		j = *Count + 1
		for j <= JumlahSoal+temp {
			fmt.Println("Enter the question: ")
			fmt.Scan(&game[j].Question)

			for i := 0; i < 4; i++ {
				fmt.Printf("Enter option %d: ", i+1)
				fmt.Scan(&game[j].options[i])
			}

			fmt.Println("Enter the Answer: ")
			fmt.Scan(&game[j].answer)
			j++
			fmt.Println("Question added successfully.")
			*Count++
		}
	}
}

func editQuestion(game *WTBM, Count *int) {
	var index int
	var newQuestion, newOption, newAnswer string
	if *Count == 0 {
		fmt.Println("No questions available to edit.")
		return
	}
	viewQuestion(game, Count)

	fmt.Println("Enter the question number to edit (1 to", *Count, "):")
	fmt.Scan(&index)
	if index < 0 || index > *Count {
		fmt.Println("Invalid question number.")
		return
	}
	fmt.Println("Editing question:", game[index].Question)
	fmt.Print("Enter new question (input \"-\" to keep current):")
	fmt.Scan(&newQuestion)
	if newQuestion != "-" {
		game[index].Question = newQuestion
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("Enter new option %d (input \"-\" to keep current): ", i+1)
		fmt.Scan(&newOption)
		if newOption != "-" {
			game[index].options[i] = newOption
		}
	}

	fmt.Println("Enter new answer (input \"-\" to keep current):")
	fmt.Scan(&newAnswer)
	if newAnswer != "-" {
		game[index].answer = newAnswer
	}
	fmt.Println("Question edited successfully.")
}

func deleteQuestion(game *WTBM, Count *int) {
	var index int
	if *Count == 0 {
		fmt.Println("No questions available to delete.")
		return
	}

	viewQuestion(game, Count)

	fmt.Println("Enter the question number to delete (1 to", *Count, "):")
	fmt.Scan(&index)

	if index < 1 && index > *Count {
		fmt.Println("Invalid question number.")
		return
	}

	for i := index; i < *Count; i++ {
		game[i].Question = game[i+1].Question
		for j := 0; j < 4; j++ {
			game[i].options[j] = game[i+1].options[j]
		}
		game[i].answer = game[i+1].answer
	}
	*Count--
	fmt.Println("Question deleted successfully.")
}

func viewQuestion(game *WTBM, Count *int) {
	if *Count == 0 {
		fmt.Println("No questions available.")
		return
	}
	for i := 1; i <= *Count; i++ {
		fmt.Println("Question No.", i)
		fmt.Println(game[i].Question)
		for j := 0; j < 4; j++ {
			fmt.Printf("Option %d: ", j+1)
			fmt.Println(game[i].options[j])
		}
		fmt.Println("Correct Answer :", game[i].answer)
		fmt.Println()
	}
}

func viewPlayer(player *Account) {
	if CountP == 0 {
		fmt.Println("No registered players.")
		return
	}
	fmt.Println("Here are the registered players.")
	sortPlayersByScore(player)
	fmt.Println()
	for i := 0; i < CountP; i++ {
		fmt.Printf("Username: %s, Points: %d\n", player[i].Username, player[i].point)
	}
}

func sortPlayersByScore(player *Account) {
	for i := 1; i < CountP; i++ {
		key := player[i]
		j := i - 1
		for j >= 0 && player[j].point < key.point {
			player[j+1] = player[j]
			j = j - 1
		}
		player[j+1] = key
	}
}

func menuPlayer(player *Account, CountP *int, game *WTBM, Name *string) {
	var choice int
	for {
		fmt.Println("-----------------------------------")
		fmt.Println("  Who Wants To Be a Millionaire")
		fmt.Println("-----------------------------------")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Println("-----------------------------------")
		fmt.Print("Please insert your choice number: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			RegisPlayer(player, CountP, game, Name)
		case 2:
			LoginPlayer(player, CountP, game, Name)
		case 3:
			fmt.Println("Exiting Player Menu....")
			return
		default:
			fmt.Println("Invalid choice, Please try again.")
		}
	}
}

func RegisPlayer(player *Account, CountP *int, game *WTBM, Name *string) {
	if *CountP == 0 {
		fmt.Print("Enter Username: ")
		fmt.Scan(&player[0].Username)
		fmt.Print("Enter Password: ")
		fmt.Scan(&player[0].Pass)
		fmt.Println("Registration successful!")
		*Name = player[0].Username
		*CountP++
		gameMenu(game, player)
	} else {
		if *CountP >= 99 {
			fmt.Println("Registration successful!")
			fmt.Println("Registration failed. No more slots available.")
			return
		}
		fmt.Print("Enter Username: ")
		fmt.Scan(&player[*CountP].Username)

		for i := 0; i < *CountP; i++ {
			if player[*CountP].Username == player[i].Username {
				fmt.Println("Username already exists. Please try again.")
				return
			}
		}
		fmt.Print("Enter Password: ")
		fmt.Scan(&player[*CountP].Pass)
		*Name = player[*CountP].Username
		fmt.Println("Registration successful!")
		*CountP++
		gameMenu(game, player)
	}
}

func LoginPlayer(player *Account, CountP *int, game *WTBM, Name *string) {
	var User, Pas string
	if *CountP == 0 {
		fmt.Println("No registered players. Please register first.")
		return
	}
	fmt.Print("Enter Username: ")
	fmt.Scan(&User)
	fmt.Print("Enter Password: ")
	fmt.Scan(&Pas)
	for i := 0; i < *CountP; i++ {
		if User == player[i].Username && Pas == player[i].Pass {
			fmt.Println("Login successful! Welcome,", User)
			*Name = User
			gameMenu(game, player)
			return
		}
	}
	fmt.Println("Invalid username or password. Please try again.")
}

func gameMenu(game *WTBM, player *Account) {
	var choice int
	for {
		fmt.Println("-----------------------------------")
		fmt.Println("  Who Wants To Be a Millionaire")
		fmt.Println("-----------------------------------")
		fmt.Println("1. Play")
		fmt.Println("2. High Score")
		fmt.Println("3. Exit")
		fmt.Println("-----------------------------------")
		fmt.Print("Please insert your choice number: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			gameTime(game, player)
		case 2:
			highScore(player)
		case 3:
			fmt.Println("Exiting Game Menu....")
			return
		default:
			fmt.Println("Invalid choice, Please try again.")
		}
	}
}
func gameTime(game *WTBM, player *Account) {
	var num, k, idx int
	var Ans string
	for i := 0; i < CountP; i++ {
		if Name == player[i].Username {
			idx = i
		}
	}
	player[idx].point = 0
	k = 1
	for k <= Count {
		for i := 1; i <= Count; i++ {
			fmt.Println("Question No.", i)
			fmt.Println(game[i].Question)
		}
		fmt.Print("Choose a question number ")
		fmt.Scan(&num)
		if num >= 1 && num <= Count {
			fmt.Println("Question No.", num)
			fmt.Println(game[num].Question)
			for j := 0; j < 4; j++ {
				fmt.Printf("Option %d: ", j+1)
				fmt.Println(game[num].options[j])
			}
			fmt.Print("Enter your answer exactly as the options show ")
			fmt.Scan(&Ans)

			if Ans == game[num].answer {
				player[idx].point++
				game[num].Correct++
			} else {
				fmt.Println("Incorrect")
				fmt.Println("The correct answer is", game[num].answer)
				game[num].Incorrect++
			}
			fmt.Println("Choose a different question number")
			k++
		} else {
			fmt.Println("Invalid choice, Please try again.")
			return
		}
	}
	fmt.Println("Your final score :", player[idx].point)
}

func sortCorrect(game *WTBM, Count int) {
	for i := 1; i < Count; i++ {
		maxIdx := i
		for j := i + 1; j <= Count; j++ {
			if game[j].Correct > game[maxIdx].Correct {
				maxIdx = j
			}
		}
		if maxIdx != i {
			temp := game[i]
			game[i] = game[maxIdx]
			game[maxIdx] = temp
		}
	}
}

func sortIncorrect(game *WTBM, Count int) {
	for i := 2; i <= Count; i++ {
		temp := game[i]
		j := i - 1
		for j >= 1 && game[j].Incorrect < temp.Incorrect {
			game[j+1] = game[j]
			j--
		}
		game[j+1] = temp
	}
}

func viewTopQuestions(game *WTBM, Count *int) {
	if *Count == 0 {
		fmt.Println("No questions available.")
		return
	}
	if *Count >= 5 {
		fmt.Println("Top 5 questions by correct answers:")
		sortCorrect(game, *Count)

		for i := 1; i <= *Count && i <= 5; i++ {
			fmt.Printf("Question No. %d: %s (Correct: %d, Incorrect: %d)\n", i, game[i].Question, game[i].Correct, game[i].Incorrect)
		}

		fmt.Println("Top 5 questions by incorrect answers:")
		sortIncorrect(game, *Count)

		for i := 1; i <= *Count && i <= 5; i++ {
			fmt.Printf("Question No. %d: %s (Correct: %d, Incorrect: %d)\n", i, game[i].Question, game[i].Correct, game[i].Incorrect)
		}
	} else if *Count < 5 {
		fmt.Println("Top 5 questions by correct answers:")
		sortCorrect(game, *Count)

		for i := 1; i <= *Count; i++ {
			fmt.Printf("Question No. %d: %s (Correct: %d, Incorrect: %d)\n", i, game[i].Question, game[i].Correct, game[i].Incorrect)
		}

		fmt.Println("Top 5 questions by incorrect answers:")
		sortIncorrect(game, *Count)

		for i := 1; i <= *Count; i++ {
			fmt.Printf("Question No. %d: %s (Correct: %d, Incorrect: %d)\n", i, game[i].Question, game[i].Correct, game[i].Incorrect)
		}
	}
}

func highScore(player *Account) {
	sortPlayersByScore(player)
	var nama string
	fmt.Print("Enter the player's username to find the rank: ")
	fmt.Scan(&nama)
	rank := binarySearch(player, nama, CountP)
	if rank != -1 {
		fmt.Printf("The rank of player %s is: %d\n", nama, rank+1)
	} else {
		fmt.Println("Player not found.")
	}
}

func binarySearch(player *Account, username string, CountP int) int {
	low, high := 0, CountP-1
	for low <= high {
		mid := (low + high) / 2
		if player[mid].Username == username {
			return mid
		} else if player[mid].Username < username {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
