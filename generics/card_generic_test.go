package generics

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type PlayingCard struct {
	Suit string
	Rank string
}

func NewPlayingCard(suit string, card string) *PlayingCard {
	return &PlayingCard{Suit: suit, Rank: card}
}

func (pc *PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", pc.Rank, pc.Suit)
}

type Deck struct {
	cards []interface{}
}

func NewPlayingCardDeck() *Deck {
	suits := []string{"Diamond", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &Deck{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}
	return deck
}

func (d *Deck) AddCard(card interface{}) {
	d.cards = append(d.cards, card)
}

func (d *Deck) RandomCard() interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx]
}

type TradingCard struct {
	CollectableName string
}

func NewTradingCard(collectableName string) *TradingCard {
	return &TradingCard{CollectableName: collectableName}
}

func (tc *TradingCard) String() string {
	return tc.CollectableName
}

type DeckGeneric[C any] struct {
	cards []C
}

func (d *DeckGeneric[C]) AddCard(card C) {
	d.cards = append(d.cards, card)
}

func (d *DeckGeneric[C]) RandomCard() C {
	r := rand.New(rand.NewSource(int64(time.Now().UnixNano())))

	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx]
}

func NewPlayingCardDeckGeneric() *DeckGeneric[*PlayingCard] {
	suits := []string{"Diamond", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &DeckGeneric[*PlayingCard]{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}
	return deck
}

func NewTradingCardDeckGeneric() *DeckGeneric[*TradingCard] {
	collectables := []string{"Sammy", "Droplets", "Spaces", "App Platform"}

	deck := &DeckGeneric[*TradingCard]{}
	for _, collectable := range collectables {
		deck.AddCard(NewTradingCard(collectable))
	}
	return deck
}

func Test_Non_Generics(t *testing.T) {
	deck := NewPlayingCardDeck()

	t.Log("--- drawing playing card ---")
	card := deck.RandomCard()
	t.Logf("drew card: %s\n", card)

	playingCard, ok := card.(*PlayingCard)
	if !ok {
		t.Fatal("card received wasn't a playing card!")
	}

	t.Logf("card suit: %s\n", playingCard.Suit)
	t.Logf("card rank: %s\n", playingCard.Rank)
}

func Test_Generics(t *testing.T) {
	playingDeck := NewPlayingCardDeckGeneric()
	tradingDeck := NewTradingCardDeckGeneric()

	t.Log("--- drawing playing card ---")
	playingCard := playingDeck.RandomCard()

	t.Logf("card suit: %s\n", playingCard.Suit)
	t.Logf("card rank: %s\n", playingCard.Rank)

	t.Log("--- drawing trading card ---")
	tradingCard := tradingDeck.RandomCard()

	t.Logf("drew card: %s\n", tradingCard)
	t.Logf("card collectable name: %s\n", tradingCard.CollectableName)

}
