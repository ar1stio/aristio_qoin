package main

import (
	"fmt"
	"math/rand"
	"time"
)

// struct untuk player
type Player struct {
	id     int
	dice   []int
	points int
	inGame bool
}

// random nomor untuk dadu
func (p *Player) rollDice() {
	for i := range p.dice {
		p.dice[i] = rand.Intn(6) + 1
	}
}

// penentu untuk hasil dari dadu yang didapat
func (p *Player) evaluateDice() (giveDice []int) {
	newDice := []int{}
	for _, d := range p.dice {
		switch d {
		case 6:
			p.points++
		case 1:
			giveDice = append(giveDice, d)
		default:
			newDice = append(newDice, d)
		}
	}
	p.dice = newDice
	if len(p.dice) == 0 {
		p.inGame = false
	}
	return giveDice
}

// opsi pengiriman array int jika dikirim ke player berikutnya
func (p *Player) addDice(dice []int) {
	p.dice = append(p.dice, dice...)
	if len(p.dice) > 0 {
		p.inGame = true
	}
}

func main() {
	var N, M int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&N)
	fmt.Print("Masukkan jumlah dadu per pemain: ")
	fmt.Scan(&M)

	rand.Seed(time.Now().UnixNano())
	players := make([]*Player, N)

	// membuat player sesuai dengan yang di inputkan
	for i := 0; i < N; i++ {
		players[i] = &Player{
			id:     i + 1,
			dice:   make([]int, M),
			points: 0,
			inGame: true,
		}
	}

	round := 1
	for {
		fmt.Printf("==================\nGiliran %d lempar dadu:\n", round)
		activePlayers := 0

		// process melemparkan dadu
		for _, player := range players {
			if player.inGame {
				player.rollDice()
				fmt.Printf("Pemain #%d (%d): %v\n", player.id, player.points, player.dice)
				activePlayers++
			}
		}

		mapDice := make(map[int][]int)

		// menampung array map yang akan dilemparkan ke player berikutnya
		fmt.Println("Setelah evaluasi:")
		for i, player := range players {
			if player.inGame {
				giveDice := player.evaluateDice()
				mapDice[(i+1)%N] = giveDice
			}
		}

		// proses melemparkan dadu ke player berikutnya
		for i, giveDice := range mapDice {
			nextPlayer := players[i]
			nextPlayer.addDice(giveDice)
		}
		for _, player := range players {
			if player.inGame {
				fmt.Printf("Pemain #%d (%d): %v\n", player.id, player.points, player.dice)
			} else {
				fmt.Printf("Pemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", player.id, player.points)
			}
		}

		if activePlayers <= 1 {
			break
		}

		round++
	}

	// penentuan pemenang
	winner := players[0]
	playerID := 0
	for _, player := range players {
		if player.points > winner.points {
			winner = player
		}
		if len(player.dice) > 0 {
			playerID = player.id
		}
	}

	fmt.Printf("==================\nGame berakhir karena hanya pemain #%d yang memiliki dadu.\n", playerID)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak (%d poin) dari pemain lainnya.\n", winner.id, winner.points)
}
