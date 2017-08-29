package macrocode

// -----------------------------------------------------
// Define macro that used for client and server syn.
// -----------------------------------------------------
const (
	// --------------------------------
	// 渠道类型 [100, 200)
	// --------------------------------
	ChannelType_Test   = 100
	ChannelType_WeiXin = 101

	// --------------------------------
	// 登陆认证类型 [200, 250)
	// --------------------------------
	LoginType_Default     = 200
	LoginType_WeiXinCode  = 201
	LoginType_WeiXinToken = 202

	// --------------------------------
	// 牌型 [0, 32)
	// --------------------------------
	CardType_Tian   = 32
	CardType_Di     = 31
	CardType_Ren    = 30
	CardType_He     = 29
	CardType_Mei    = 28
	CardType_Chang  = 27
	CardType_Ban    = 26
	CardType_Fu     = 25
	CardType_46     = 24
	CardType_16     = 23
	CardType_15     = 22
	CardType_Dian   = 21
	CardType_ZhiZun = 20
)
