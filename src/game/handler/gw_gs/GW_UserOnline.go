package gw_gs

import (
	"comm/dbmgr"
	"game/app"
	"game/msg"
	"proto/errorcode"
)

func GW_UserOnline(message msg.Message, ctx interface{}) {
	req := message.(*msg.GW_UserOnline)

	// * try load player
	// * if not found, create new
	plr := app.PlayerMgr.LoadPlayer(req.UserId, true)
	if plr == nil {
		// create new user
		plr = app.PlayerMgr.CreatePlayer(req.UserId, func(user *app.User) {
			user.Channel = req.Channel
			user.ChannelUid = req.ChannelUid
			user.Name, user.Head = dbmgr.CenterGetUserNameHead(req.Channel, req.ChannelUid)

		})
		if plr == nil {
			app.NetMgr.Send2Player(req.Sid, &msg.GS_LoginError{ErrorCode: Err.Login_UserInfo})
			return
		}
	} else {
		// user exists. check double login
		if plr.IsOnline() {
			plr.Logout()
		}
	}

	// send userinfo
	user := plr.User()

	app.NetMgr.Send2Player(req.Sid, &msg.GS_UserInfo{
		UserId: user.Id,
		Name:   user.Name,
		Head:   user.Head,
		Lv:     user.Lv,
		Exp:    user.Exp,
	})

	// set online
	app.PlayerMgr.SetOnline(plr, req.Sid, req.LoginIP)
}
