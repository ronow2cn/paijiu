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

	Table_Full      = 100
	Table_IsBanker  = 101
	Table_ErrorPos  = 102
	Table_PosOccupy = 103
)
