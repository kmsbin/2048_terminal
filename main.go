package main

import (
	"os"

	term "github.com/nsf/termbox-go"
)

func main() {

	_ = term.Init()

	gameLogic := GameBusinessLogic{[][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
	}, 0}
	positionsMap := gameLogic.randowNewCellPosition()
	gameLogic.SetNewValue(positionsMap["column"], positionsMap["row"])
	draw(template(gameLogic.table))

	defer term.Close()
	eventQueue := make(chan term.Event)

	go func() {
		for {
			eventQueue <- term.PollEvent()
		}
	}()
	// Game entry point
	for {
		ev := <-eventQueue
		if ev.Type == term.EventKey {
			switch ev.Key {
			case term.KeyEsc:
				term.Close()
				os.Exit(0)
			case term.KeyArrowDown:
				term.Clear(term.ColorDefault, term.ColorDefault)
				nextTable := gameLogic.DownMove()
				draw(template(nextTable))

			case term.KeyArrowUp:
				term.Clear(term.ColorDefault, term.ColorDefault)
				nextTable := gameLogic.UpMove()
				draw(template(nextTable))
				// fmt.Println(term.KeyArrowUp)
			case term.KeyArrowLeft:

				term.Clear(term.ColorDefault, term.ColorDefault)
				nextTable := gameLogic.LeftMove()
				draw(template(nextTable))
				// fmt.Println(term.KeyArrowLeft)
			case term.KeyArrowRight:
				nextTable := gameLogic.RightMove()
				draw(template(nextTable))
			case term.KeyCtrlC:
				term.Close()
				os.Exit(0)
			}
			positionsMap := gameLogic.randowNewCellPosition()
			gameLogic.SetNewValue(positionsMap["column"], positionsMap["row"])
		}

	}
}
