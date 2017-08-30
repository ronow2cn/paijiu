package Err

const (
	OK     = 0
	Failed = 1

	// --------------------------------
	// 登录 [2, 100)
	// --------------------------------

	Login_InvalidVersion = 2
	Login_Failed         = 3
	Login_UserBanned     = 4
	Login_UserInfo       = 5

	// --------------------------------
	// 牌桌 [100, 200)
	// --------------------------------

	Table_Full                 = 100
	Table_IsBanker             = 101
	Table_ErrorPos             = 102
	Table_PosOccupy            = 103
	Table_NotExist             = 104
	Table_IsNotBanker          = 105
	Table_NotInTable           = 106
	Table_IdIsOver             = 107
	Table_PosPlrError          = 108
	Table_ScoreError           = 109
	Table_IsInOtherTable       = 110
	Table_PlayerHaveTable      = 111
	Table_BankerScoreNotEnough = 112
	Table_NotCurPlay           = 113
	Table_CurPlayIng           = 114
	Table_PlayerNotFound       = 115
)
