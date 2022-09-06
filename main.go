package main

import (
	"fmt"
	"log"
	"os"
	"proxx/proxx"
	"strconv"
	"strings"
	"time"
)

func main() {

	colored := true
	if strings.ToLower(os.Getenv("COLORED")) == "false" {
		colored = false
	}

	rowsNum, columnsNum, holesNum := scanGameParams()
	game, err := proxx.NewGame(rowsNum, columnsNum, holesNum)
	if err != nil {
		log.Fatal(err)
	}

	printLegend()
	fmt.Print(game.ASCIIBoard(colored))
	game.Start()

	var row, column int
	for {
		fmt.Print("Enter row and column numbers separated with space. E.g '4 5': ")
		n, err := fmt.Scanln(&row, &column)
		fmt.Println()
		if err != nil || n != 2 {
			log.Print("Invalid input. Should be in format 'rowNum columnNum'")
			continue
		}

		if row < 0 || column < 0 { //user decided to stop playing
			break
		}

		if row > rowsNum-1 || column > columnsNum-1 {
			fmt.Println("coordinates out of bound")
			continue
		}

		if game.IsHole(row, column) {
			game.OpenAllCells()
			fmt.Print(game.ASCIIBoard(colored))
			printStats(game.StepsCount(), game.StartTime())
			fmt.Println("GAME OVER")
			os.Exit(0)
		}

		if game.Won() {
			game.OpenAllCells()
			fmt.Print(game.ASCIIBoard(colored))
			printStats(game.StepsCount(), game.StartTime())
			fmt.Println("CONGRATS! YOU WON!")
			os.Exit(0)
		}

		if game.Opened(row, column) {
			fmt.Println("Already opened")
			continue
		}

		game.OpenCell(row, column)
		fmt.Print(game.ASCIIBoard(colored))
	}

}

func scanGameParams() (rowNum, columnNum, holesNum int) {
	var rowsNumStr, columnsNumStr, holesNumStr string
	for {
		fmt.Print("Enter number of rows: ")
		_, err := fmt.Scanln(&rowsNumStr)
		if err != nil {
			fmt.Printf("Can't scan rows number: %s\n", err.Error())
			continue
		}
		rowNum, err = strconv.Atoi(rowsNumStr)
		if err != nil || rowNum <= 0 {
			fmt.Println("Rows number should be a positive integer")
			continue
		}

		break
	}

	for {
		fmt.Print("Enter number of columns: ")
		_, err := fmt.Scanln(&columnsNumStr)
		if err != nil {
			fmt.Printf("Can't scan columns number: %s\n", err.Error())
			continue
		}
		columnNum, err = strconv.Atoi(columnsNumStr)
		if err != nil || columnNum < 0 {
			fmt.Println("Columns number should be a non negative integer")
			continue
		}

		if columnNum*rowNum <= 1 {
			fmt.Println("Number of cells should be bigger than 1")
			continue
		}

		break
	}

	for {
		fmt.Print("Enter number of holes: ")
		_, err := fmt.Scanln(&holesNumStr)
		if err != nil {
			fmt.Printf("Can't scan holes number: %s\n", err.Error())
			continue
		}
		holesNum, err = strconv.Atoi(holesNumStr)
		if err != nil || holesNum < 1 {
			fmt.Println("Holes number should be a non negative integer")
			continue
		}

		if holesNum >= rowNum*columnNum {
			fmt.Println("Holes number should be less than the total number of cells")
			continue
		}

		break
	}

	return
}

func printStats(stepsCount int, startTime time.Time) {
	fmt.Printf("Time taken: %s, steps: %d\n", time.Since(startTime), stepsCount)
}

func printLegend() {
	fmt.Println("Proxx")
	fmt.Println()
	fmt.Println("Legend: ")
	fmt.Println("X - closed")
	fmt.Println("<empty> - opened")
	fmt.Println("<number> - opened, number of surrounding holes")
	fmt.Println("@ - hole")
}
