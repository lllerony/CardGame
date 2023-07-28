package main

import (
	"fmt"
	"math/rand"
)

const (
	Worms    string = "Черви"
	Diamonds        = "Буби"
	Crosses         = "Крести"
	Peaks           = "Пики"
)

type card struct {
	ownerID     int
	number      int
	picture     string
	sort        string
	isTrumpCard bool
}

type player struct {
	playerID int
	name     string
	cards    []card
}

type field struct {
	layID      int
	cardsField []card
}

var allCards = []card{
	{sort: Worms, number: 6, picture: ""},
	{sort: Worms, number: 7, picture: ""},
	{sort: Worms, number: 8, picture: ""},
	{sort: Worms, number: 9, picture: ""},
	{sort: Worms, number: 10, picture: ""},
	{sort: Worms, number: 11, picture: "Валет"},
	{sort: Worms, number: 12, picture: "Дама"},
	{sort: Worms, number: 13, picture: "Король"},
	{sort: Worms, number: 14, picture: "Туз"},

	{sort: Diamonds, number: 6, picture: ""},
	{sort: Diamonds, number: 7, picture: ""},
	{sort: Diamonds, number: 8, picture: ""},
	{sort: Diamonds, number: 9, picture: ""},
	{sort: Diamonds, number: 10, picture: ""},
	{sort: Diamonds, number: 11, picture: "Валет"},
	{sort: Diamonds, number: 12, picture: "Дама"},
	{sort: Diamonds, number: 13, picture: "Король"},
	{sort: Diamonds, number: 14, picture: "Туз"},

	{sort: Crosses, number: 6, picture: ""},
	{sort: Crosses, number: 7, picture: ""},
	{sort: Crosses, number: 8, picture: ""},
	{sort: Crosses, number: 9, picture: ""},
	{sort: Crosses, number: 10, picture: ""},
	{sort: Crosses, number: 11, picture: "Валет"},
	{sort: Crosses, number: 12, picture: "Дама"},
	{sort: Crosses, number: 13, picture: "Король"},
	{sort: Crosses, number: 14, picture: "Туз"},

	{sort: Peaks, number: 6, picture: ""},
	{sort: Peaks, number: 7, picture: ""},
	{sort: Peaks, number: 8, picture: ""},
	{sort: Peaks, number: 9, picture: ""},
	{sort: Peaks, number: 10, picture: ""},
	{sort: Peaks, number: 11, picture: "Валет"},
	{sort: Peaks, number: 12, picture: "Дама"},
	{sort: Peaks, number: 13, picture: "Король"},
	{sort: Peaks, number: 14, picture: "Туз"},
}

var players []player
var cardsOnField field

func createPlayer(numOfPlayers int) {
	var gamer player
	for i := 0; i < numOfPlayers; i++ {
		players = append(players, gamer)
	}
	for i := 0; i < numOfPlayers; i++ {
		players[i].playerID = i
	}
}
func dealCards(numOfPlayers int) {
	for i := 0; i < numOfPlayers; i++ {

		for j := 0; j < 6; j++ {
			k := rand.Intn(len(allCards))
			players[i].cards = append(players[i].cards, allCards[k])
			allCards = Remove(allCards, k)
		}

	}
	for i := 0; i < numOfPlayers; i++ {
		fmt.Printf("Карты %d-го игрока: ", players[i].playerID)
		for j := 0; j < 6; j++ {
			players[i].cards[j].ownerID = players[i].playerID
			if players[i].cards[j].number > 10 {
				fmt.Printf("%s %s, ", players[i].cards[j].picture, players[i].cards[j].sort)
			} else {
				fmt.Printf("%d %s, ", players[i].cards[j].number, players[i].cards[j].sort)
			}
		}
		fmt.Printf("\n")
	}
}

var step int
var counter int
var numOfPlayers int

func firstStep(numOfPlayers int) {
	k := rand.Intn(len(allCards))
	var minInd int
	var trumpCard = allCards[k]
	if trumpCard.number > 10 {
		fmt.Printf("Козырной картой является: %s %s", trumpCard.picture, trumpCard.sort)
	} else {
		fmt.Printf("Козырной картой является: %d %s", trumpCard.number, trumpCard.sort)
	}
	fmt.Printf("\n")
	var trumps []card
	for i := 0; i < numOfPlayers; i++ {
		for j := 0; j < 6; j++ {
			if players[i].cards[j].sort == trumpCard.sort {
				trumps = append(trumps, players[i].cards[j])
			}
		}
	}

	for i := 0; i < len(allCards); i++ {
		if allCards[i].sort == trumpCard.sort {
			allCards[i].isTrumpCard = true
		}
	}

	min := trumps[0].number
	for ind, v := range trumps {
		if v.number < min {
			min = v.number
			minInd = ind
		}
	}
	counter = minInd
	for i := 0; i < len(trumps); i++ {
		if trumps[i].number == min {
			players[counter].playerID = trumps[i].ownerID
		}
	}
	fmt.Println("Игру начинает игрок", players[counter].playerID)
}

func Remove(slice []card, s int) []card {
	copy(slice[s:], slice[s+1:])
	slice[len(slice)-1] = card{}
	slice = slice[:len(slice)-1]
	return slice
}

func Counter(counter *int, numOfPlayers int) {
	if *counter+1 == numOfPlayers {
		*counter = 0
	} else {
		*counter++
	}
}

func formatPrintCards() {
	for j := 0; j < len(players[counter].cards); j++ {
		if players[counter].cards[j].number > 10 {
			fmt.Printf("%s %s, ", players[counter].cards[j].picture, players[counter].cards[j].sort)
		} else {
			fmt.Printf("%d %s, ", players[counter].cards[j].number, players[counter].cards[j].sort)
		}
	}
}

func threwTheCard() {
	if players[counter].cards[step].number > 10 {
		fmt.Printf("%d игрок бросил карту: %s %s\n", players[counter].playerID, players[counter].cards[step].picture, players[counter].cards[step].sort)
	} else {
		fmt.Printf("%d игрок бросил карту: %d %s\n", players[counter].playerID, players[counter].cards[step].number, players[counter].cards[step].sort)
	}
}

func fightBack() {
	if cardsOnField.cardsField[0].sort > players[counter].cards[step].sort && len(cardsOnField.cardsField) != 0 && cardsOnField.cardsField[step].sort != players[counter].cards[step].sort {
		fmt.Println("Эта карта не соответствует карте противника. Выберите другую")
		fmt.Scan(&step)
	} else if (!players[counter].cards[step].isTrumpCard && !cardsOnField.cardsField[0].isTrumpCard) || (players[counter].cards[step].isTrumpCard && cardsOnField.cardsField[0].isTrumpCard) && (players[counter].cards[step].number > cardsOnField.cardsField[0].number) {
		fmt.Println("Бито:")
		players[counter].cards = Remove(players[counter].cards, step)
		fmt.Println("Берёте недостаущую карту:")
		k := rand.Intn(len(allCards))
		players[counter].cards = append(players[counter].cards, allCards[k])
		allCards = Remove(allCards, k)
	} else {
		fmt.Println("Вы не можете отбиться. Берёте карту противника")
		players[counter].cards = append(players[counter].cards, cardsOnField.cardsField[0])
	}
}

func playingTheFool() {
	cardsOnField.cardsField = []card{} //карта на поле сбрасывается
	fmt.Scan(&step)
	cardsOnField.cardsField = append(cardsOnField.cardsField, players[counter].cards[step])
	threwTheCard()
	fightBack()
	fmt.Printf("У %d-го игрока остались карты: ", players[counter].playerID)
	formatPrintCards()
	fmt.Printf("\n")
	Counter(&counter, numOfPlayers)
	if len(players[counter].cards) == 0 {
		fmt.Println("Игрок", players[counter].playerID, "победил!")
	}
	fmt.Println("Ходит игрок", players[counter].playerID, "его карты: ")
	formatPrintCards()
}

func main() {
	fmt.Println("Введите кол-во игроков: ")
	fmt.Scan(&numOfPlayers)
	createPlayer(numOfPlayers)
	dealCards(numOfPlayers)
	firstStep(numOfPlayers)
	for {
		playingTheFool()
	}
}
