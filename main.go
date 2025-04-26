package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	FrontHalfTicketPrice = 10
	BackHalfTicketPrice  = 8

	SeatAvailable = "S"
	SeatBooked    = "B"

	SmallRoomThreshold = 60

	MenuExit       = 0
	MenuShowSeats  = 1
	MenuBuyTicket  = 2
	MenuStatistics = 3
)

func main() {
	totalNumRows := getPromptedNumber("Enter the number of rows:")
	totalNumSeats := getPromptedNumber("Enter the number of seats in each row:")
	seats := createSeats(totalNumRows, totalNumSeats)

	for {
		printMenu()
		menuItem := getMenuSelection()

		switch menuItem {
		case MenuExit:
			return
		case MenuShowSeats:
			printSeats(seats)
		case MenuBuyTicket:
			bookSeat(seats, totalNumRows, totalNumSeats)
		case MenuStatistics:
			getStatistics(seats, totalNumRows, totalNumSeats)
		default:
			log.Println("Unknown option")
		}
	}
}

func getPromptedSeat(seats [][]string, totalNumRows, totalNumSeats int, numRow, numSeat *int, err *string) bool {
	*numRow = getPromptedNumber("Enter a row number:")
	*numSeat = getPromptedNumber("Enter a seat number in that row:")

	if *numRow > 9 || *numRow < 1 {
		*err = "Wrong input!"
		return false
	}

	if *numSeat > 9 || *numSeat < 1 {
		*err = "Wrong input!"
		return false
	}

	if seats[*numRow-1][*numSeat-1] == SeatBooked {
		*err = "That ticket has already been purchased!"
		return false
	}

	ticketPrice := getTicketPrice(totalNumRows, totalNumSeats, *numRow)
	fmt.Printf("Ticket price: $%d\n", ticketPrice)

	if *numRow < 1 || *numRow > totalNumRows || *numSeat < 1 || *numSeat > totalNumSeats {
		*err = "Wrong input!"
		return false
	}

	return true
}

func bookSeat(seats [][]string, totalNumRows, totalNumSeats int) {
	var numRow, numSeat int
	var err string

	for !getPromptedSeat(seats, totalNumRows, totalNumSeats, &numRow, &numSeat, &err) {
		fmt.Println(err)
	}

	seats[numRow-1][numSeat-1] = SeatBooked
}

func countPurchasedTickets(seats [][]string) int {
	count := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == SeatBooked {
				count++
			}
		}
	}
	return count
}

func createSeats(numRows, numSeats int) [][]string {
	seats := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		seats[i] = make([]string, numSeats)
		for j := 0; j < numSeats; j++ {
			seats[i][j] = SeatAvailable
		}
	}
	return seats
}

func getCurrentIncome(seats [][]string, numRows, numSeats int) int {
	currentIncome := 0
	for rowIndex, row := range seats {
		for _, seat := range row {
			if seat == SeatBooked {
				ticketPrice := getTicketPrice(numRows, numSeats, rowIndex+1)
				currentIncome += ticketPrice
			}
		}
	}

	return currentIncome
}

func getMenuSelection() int {
	return getPromptedNumber("Select an option:")
}

func getNumFromReader() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal("Invalid input:", err)
	}

	return num
}

func getPercentage(seats [][]string, num int) float64 {
	total := len(seats) * len(seats[0])
	return (float64(num) / float64(total)) * 100
}

func getPromptedNumber(prompt string) int {
	fmt.Println(prompt)
	return getNumFromReader()
}

func getStatistics(seats [][]string, totalNumRows, totalNumSeats int) {
	num := countPurchasedTickets(seats)
	fmt.Printf("Number of purchased tickets: %d\n", num)
	fmt.Printf("Percentage: %.2f%%\n", getPercentage(seats, num))
	fmt.Printf("Current income: $%d\n", getCurrentIncome(seats, totalNumRows, totalNumSeats))
	fmt.Printf("Total income: $%d\n", getTotalIncome(totalNumRows, totalNumSeats))
}

func getTicketPrice(numRows, numSeats, numRow int) int {
	totalSeats := numRows * numSeats

	if totalSeats < SmallRoomThreshold {
		return FrontHalfTicketPrice
	}
	if numRow <= numRows/2 {
		return FrontHalfTicketPrice
	}
	return BackHalfTicketPrice
}

func getTotalIncome(numRows, numSeats int) int {
	totalSeats := numRows * numSeats
	if totalSeats < SmallRoomThreshold {
		return totalSeats * FrontHalfTicketPrice
	}
	half := numRows / 2
	if numRows%2 == 0 {
		return half*numSeats*FrontHalfTicketPrice + half*numSeats*BackHalfTicketPrice
	}
	return half*numSeats*FrontHalfTicketPrice + (numRows-half)*numSeats*BackHalfTicketPrice
}

func printMenu() {
	fmt.Println(`
1. Show the seats
2. Buy a ticket
3. Statistics
0. Exit`)
}

func printSeats(seats [][]string) {
	fmt.Println("Cinema:")
	fmt.Print("  ")
	for i := 1; i <= len(seats[0]); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	for i, row := range seats {
		fmt.Printf("%d ", i+1)
		for _, seat := range row {
			fmt.Printf("%s ", seat)
		}
		fmt.Println()
	}
}
