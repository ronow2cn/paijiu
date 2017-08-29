package table

import (
	"game/app/gconst"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"proto/macrocode"
	"time"
)

var randCard = rand.New(rand.NewSource(time.Now().Unix()))

// ============================================================================

type Chip struct {
	Bets map[string]int32 `bson:"num"` //[玩家]玩家下注数量
}

type chips map[int32]*Chip     //[位置]位置上筹码
type poscard map[int32][]*Card //[位置]位置上筹码

type Card struct {
	T int32 //类型
	N int32 //数值
}

//一场牌局
type Play struct {
	Chips   chips   `bson:"chips"`   //下注情况
	Cards   []*Card `bson:"cards"`   //牌
	PosCard poscard `bson:"poscard"` //位置上的牌
	Idx     int32   `bson:"idx"`     //牌的位置
}

// ============================================================================
// marshalling

func (self chips) GetBSON() (interface{}, error) {
	type chip_t struct {
		Id  int32
		Pos *Chip
	}
	var arr []*chip_t

	for id, val := range self {
		arr = append(arr, &chip_t{id, val})
	}

	return arr, nil
}

func (self *chips) SetBSON(raw bson.Raw) error {
	type chip_t struct {
		Id  int32
		Pos *Chip
	}
	var arr []*chip_t

	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}

	*self = make(chips)
	for _, v := range arr {
		(*self)[v.Id] = v.Pos
	}

	return nil
}

func (self poscard) GetBSON() (interface{}, error) {
	type poscard_t struct {
		Id      int32
		PosCard []*Card
	}
	var arr []*poscard_t

	for id, val := range self {
		arr = append(arr, &poscard_t{id, val})
	}

	return arr, nil
}

func (self *poscard) SetBSON(raw bson.Raw) error {
	type poscard_t struct {
		Id      int32
		PosCard []*Card
	}
	var arr []*poscard_t

	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}

	*self = make(poscard)
	for _, v := range arr {
		(*self)[v.Id] = v.PosCard
	}

	return nil
}

// ============================================================================

func NewPlay() *Play {
	return &Play{
		Chips:   make(chips),
		Cards:   make([]*Card, 32),
		PosCard: make(poscard),
		Idx:     0,
	}
}

func (self *Play) Init() {

}

func (self *Play) Reset() {
	self.Chips = make(chips)
	self.Cards = make([]*Card, 32)
	self.PosCard = make(poscard)
	self.Idx = 0
}

//洗牌
func (self *Play) Shuffle() {
	self.Cards = NewCards()

	n1 := randCard.Int31n(32)
	n2 := randCard.Int31n(32)

	for i := 0; i < 100; i++ {
		self.Cards[n1], self.Cards[n2] = self.Cards[n2], self.Cards[n1]
		n1 = randCard.Int31n(32)
		n2 = randCard.Int31n(32)
	}
}

func (self *Play) Deal() {
	self.PosCard[gconst.TablePosBanker] = self.Cards[self.Idx : self.Idx+2]
	self.PosCard[gconst.TablePosPlayer1] = self.Cards[self.Idx+2 : self.Idx+4]
	self.PosCard[gconst.TablePosPlayer2] = self.Cards[self.Idx+4 : self.Idx+6]
	self.PosCard[gconst.TablePosPlayer3] = self.Cards[self.Idx+6 : self.Idx+8]

	self.Idx += 8
}

func (self *Play) isDuiZi(d1, d2 *Card) bool {
	return (d1.T == d2.T) && (d1.N == d2.N)
}

func (self *Play) Compere(c1, c2 []*Card) int32 {
	if c1[0].T == macrocode.CardType_ZhiZun && c1[1].T == macrocode.CardType_ZhiZun {
		return 1
	}

	if c2[0].T == macrocode.CardType_ZhiZun && c2[1].T == macrocode.CardType_ZhiZun {
		return 2
	}

	if self.isDuiZi(c1[0], c1[1]) {
		if self.isDuiZi(c2[0], c2[1]) {
			if c1[0].T > c2[0].T {
				return 1
			} else {
				return 2
			}
		} else {
			return 1
		}
	} else {
		if self.isDuiZi(c2[0], c2[1]) {
			return 2
		} else {

			if (c1[0].N + c1[1].N) == (c2[0].N + c2[1].N) {

			} else if (c1[0].N + c1[1].N) > (c2[0].N + c2[1].N) {
				return 1
			} else {
				return 2
			}
		}
	}

	return 0
}
