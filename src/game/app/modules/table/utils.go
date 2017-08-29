package table

import (
	"comm/logger"
	"proto/macrocode"
)

var log = logger.DefaultLogger

// ============================================================================

func NewCards() []*Card {
	ret := []*Card{
		&Card{macrocode.CardType_Tian, 12}, &Card{macrocode.CardType_Tian, 12},
		&Card{macrocode.CardType_Di, 2}, &Card{macrocode.CardType_Di, 2},
		&Card{macrocode.CardType_Ren, 8}, &Card{macrocode.CardType_Ren, 8},
		&Card{macrocode.CardType_He, 4}, &Card{macrocode.CardType_He, 4},
		&Card{macrocode.CardType_Mei, 10}, &Card{macrocode.CardType_Mei, 10},
		&Card{macrocode.CardType_Chang, 6}, &Card{macrocode.CardType_Chang, 6},
		&Card{macrocode.CardType_Ban, 4}, &Card{macrocode.CardType_Ban, 4},
		&Card{macrocode.CardType_Fu, 11}, &Card{macrocode.CardType_Fu, 11},
		&Card{macrocode.CardType_46, 10}, &Card{macrocode.CardType_46, 10},
		&Card{macrocode.CardType_16, 7}, &Card{macrocode.CardType_16, 7},
		&Card{macrocode.CardType_15, 6}, &Card{macrocode.CardType_15, 6},
		&Card{macrocode.CardType_Dian, 9}, &Card{macrocode.CardType_Dian, 9},
		&Card{macrocode.CardType_Dian, 8}, &Card{macrocode.CardType_Dian, 8},
		&Card{macrocode.CardType_Dian, 7}, &Card{macrocode.CardType_Dian, 7},
		&Card{macrocode.CardType_Dian, 5}, &Card{macrocode.CardType_Dian, 5},
		&Card{macrocode.CardType_ZhiZun, 6}, &Card{macrocode.CardType_ZhiZun, 3},
	}

	return ret
}
