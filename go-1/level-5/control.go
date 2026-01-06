package main

import (
	"fmt"
	"sort"
)

type Player struct {
    Name      string
    Goals     int
    Misses    int
    Assists   int
    Rating    float64
}

func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
	}
	p.calculateRating()
	return p
}

func (p *Player) calculateRating() {
	if p.Misses != 0 {
		p.Rating = (float64(p.Goals) + float64(p.Assists) / 2) / float64(p.Misses)
	} else {
		p.Rating = float64(p.Goals) + float64(p.Assists) / 2
	}
}

// Сортировка по: Убыванию количества голов
func goalsSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Goals == players[j].Goals {
			return players[i].Name < players[j].Name
		}
		return players[i].Goals > players[j].Goals
	})
	return players
}

// Сортировка по: Убыванию рейтинга
func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Rating == players[j].Rating {
			return players[i].Name < players[j].Name
		}
		return players[i].Rating > players[j].Rating
	})
	return players
}

// Сортировка по: Убыванию отношения голов к промахам 
func gmSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		var gi, gj float64
		if players[i].Misses == 0 {
			gi = float64(players[i].Goals)
		} else {
			gi = float64(players[i].Goals) / float64(players[i].Misses)
		}
		if players[j].Misses == 0 {
			gj = float64(players[j].Goals)
		} else {
			gj = float64(players[j].Goals) / float64(players[j].Misses)
		}
		if gi != gj {
			return gi > gj
		}
		return players[i].Name < players[j].Name
	})
	return players
}

func main() {
	myTeam := []Player{
		NewPlayer("A", 900, 50, 100000),
		NewPlayer("B", 700, 90, 10000),
		NewPlayer("C", 800, 91, 1000),
		NewPlayer("D", 900, 49, 100),
	}
	fmt.Println(goalsSort(myTeam))
	fmt.Println(ratingSort(myTeam))
	fmt.Println(gmSort(myTeam))
}