package table

import (
	"comm/logger"
	"proto/macrocode"
)

var log = logger.DefaultLogger

// ============================================================================

func NewCards() []*Card {
	ret := []*Card{
		&Card{macrocode.CardType_Tian, 12, 0}, &Card{macrocode.CardType_Tian, 12, 1},
		&Card{macrocode.CardType_Di, 2, 0}, &Card{macrocode.CardType_Di, 2, 1},
		&Card{macrocode.CardType_Ren, 8, 0}, &Card{macrocode.CardType_Ren, 8, 1},
		&Card{macrocode.CardType_He, 4, 0}, &Card{macrocode.CardType_He, 4, 1},
		&Card{macrocode.CardType_Mei, 10, 0}, &Card{macrocode.CardType_Mei, 10, 1},
		&Card{macrocode.CardType_Chang, 6, 0}, &Card{macrocode.CardType_Chang, 6, 1},
		&Card{macrocode.CardType_Ban, 4, 0}, &Card{macrocode.CardType_Ban, 4, 1},
		&Card{macrocode.CardType_Fu, 11, 0}, &Card{macrocode.CardType_Fu, 11, 1},
		&Card{macrocode.CardType_46, 10, 0}, &Card{macrocode.CardType_46, 10, 1},
		&Card{macrocode.CardType_16, 7, 0}, &Card{macrocode.CardType_16, 7, 1},
		&Card{macrocode.CardType_15, 6, 0}, &Card{macrocode.CardType_15, 6, 1},
		&Card{macrocode.CardType_Dian, 9, 0}, &Card{macrocode.CardType_Dian, 9, 1},
		&Card{macrocode.CardType_Dian, 8, 0}, &Card{macrocode.CardType_Dian, 8, 1},
		&Card{macrocode.CardType_Dian, 7, 0}, &Card{macrocode.CardType_Dian, 7, 1},
		&Card{macrocode.CardType_Dian, 5, 0}, &Card{macrocode.CardType_Dian, 5, 1},
		&Card{macrocode.CardType_ZhiZun, 6, 0}, &Card{macrocode.CardType_ZhiZun, 3, 0},
	}

	return ret
}
