package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TicTacToe() {
	introMsg()

	r := bufio.NewReader(os.Stdin)
	p1, p2 := createPlayers(r)
	b := createBoard(p1, p2)
	b.Show()
	fmt.Printf("\n\n--> select a tile to play | %s: ", b.currentPlayer)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "e" {
			fmt.Println("ended")
			break
		}

		// @todo: complete implementation of the different cases
		switch s := scanner.Text(); s {
		case "d":
			// details
			fmt.Println("details")
		case "r":
			// rematch
			fmt.Println("rematch")
		case "h":
			// help
			fmt.Println("help")
		case "s":
			b.Show()
		case "":
			fmt.Println("\n please select a tile to play!")
		case "a1", "a2", "a3", "b1", "b2", "b3", "c1", "c2", "c3":
			b.PlayTile(s)
		default:
			fmt.Println("\n invalid selection!")
		}

		fmt.Printf("\n\n--> select a tile to play | %s: ", b.currentPlayer)
	}
}

func introMsg() {
	fmt.Println(`

		         TAC  
		     TIC     TOE 
		         TAC  

				 	created_by_Bada
	

	details (d), rematch (r), show board (s), end (e), help (h)  
	
	`)
}

type Player struct {
	name   string
	weapon string
}

// createPlayers returns players created by user.
func createPlayers(r *bufio.Reader) (Player, Player) {
	player := make([]Player, 2)
	weapon := []string{"X", "O"}

	for i := 0; i < 2; i++ {
		fmt.Print("--> enter player name: ")
		// @todo: accept only alphanumeric characters and underscores (no spaces, no dashes)
		// @todo: add limit of 12 characters
		name, _ := r.ReadString('\n')
		name = normalizeString(name)
		if name == "" {
			name = "player" + fmt.Sprint(i)
		}
		player[i] = Player{name, weapon[i]}
		fmt.Printf(" player `%s` created! %s \n \n", name, weapon[i])
	}

	return player[0], player[1]
}

func normalizeString(s string) string {
	s = strings.TrimSpace(s)
	return strings.ToLower(s)
}

type Board struct {
	currentPlayer Player
	tiles         map[string]string
	p1            Player
	p2            Player
}

func createBoard(p1 Player, p2 Player) Board {
	tiles := make(map[string]string, 9)
	emptyTile := "___"

	for _, t := range []string{"a1", "a2", "a3", "b1", "b2", "b3", "c1", "c2", "c3"} {
		tiles[t] = emptyTile
	}

	return Board{p1, tiles, p1, p2}
}

func (b *Board) Show() {
	fmt.Printf(` 
  	  | 1 | 2 | 3
  	--------------
  	a |%s|%s|%s
  	b |%s|%s|%s
  	c |%s|%s|%s
	`,
		b.tiles["a1"],
		b.tiles["a2"],
		b.tiles["a3"],
		b.tiles["b1"],
		b.tiles["b2"],
		b.tiles["b3"],
		b.tiles["c1"],
		b.tiles["c2"],
		b.tiles["c3"],
	)
}

func (b *Board) PlayTile(t string) {
	t = normalizeString(t)

	if b.tiles[t] == "___" {
		b.tiles[t] = fmt.Sprintf(" %s ", b.currentPlayer.weapon)
		fmt.Printf("\n  `%s` played \"%s\" in position %s \n",
			b.currentPlayer.name,
			b.currentPlayer.weapon,
			strings.ToUpper(t),
		)
		// @todo: check for win or draw
		b.checkWinOrDraw()
		b.Show() // @todo: show different stuff when there is a win
		b.ChangePlayer()
	} else {
		fmt.Printf("\n  position %s is not empty! \n", strings.ToUpper(t))
	}
}

func (b *Board) checkWinOrDraw() {}

func (b *Board) ChangePlayer() {
	if b.currentPlayer == b.p1 {
		b.currentPlayer = b.p2
	} else {
		b.currentPlayer = b.p1
	}
}
