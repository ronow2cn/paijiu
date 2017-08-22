package dbmgr

import (
	"comm/db"
	"time"
)

// ============================================================================

type UserInfo struct {
	UserId       string    `bson:"_id"`         //游戏中玩家唯一id
	Channel      int32     `bson:"channel"`     //账号渠道类型
	ChannelUid   string    `bson:"channel_uid"` //对应渠道UID
	BanTs        time.Time `bson:"ban_ts"`      //玩家封号时间
	Name         string    `bson:"name"`        //玩家名字
	Head         string    `bson:"head"`        //头像
	Token        string    `bson:"token"`       //账号token
	ExpireT      time.Time `bson:"expire_t"`    //token过期时间
	RefreshToken string    `bson:"r_token"`     //刷新token(主要微信，支付宝用到)
}

// ============================================================================

func CenterGetUserInfo(channel int32, authid string) *UserInfo {
	var obj UserInfo

	err := DBCenter.GetObjectByCond(
		CTabNameUserinfo,
		db.M{
			"channel":     channel,
			"channel_uid": authid,
		},
		&obj,
	)
	if err == nil {

		return &obj

	} else if db.IsNotFound(err) {
		// allocate user db
		dbname := CenterAllocUserDB()
		if dbname == "" {
			return nil
		}

		obj.UserId = CenterGenUserId(dbname)
		obj.Channel = channel
		obj.ChannelUid = authid
		obj.BanTs = time.Unix(0, 0)
		obj.Name = ""
		obj.Head = ""

		// flush to db
		err = DBCenter.Insert(CTabNameUserinfo, &obj)
		if err == nil {
			// update user load
			CenterIncUserLoad(dbname)

			// return new userinfo
			return &obj
		} else {
			// failed
			return nil
		}
	} else {
		// failed
		return nil
	}
}

func CenterGetUserNameHead(channel int32, channeluid string) (name, head string) {
	obj := CenterGetUserInfo(channel, channeluid)
	if obj == nil {
		return
	}

	name, head = obj.Name, obj.Head
	return
}

// ============================================================================

func CenterBanUser(userid string, min int32) {
	// > 0: 封号
	// < 0: 解封

	if min == 0 {
		return
	}

	bants := time.Unix(0, 0)
	if min > 0 {
		bants = time.Now().Add(time.Duration(min) * time.Minute)
	}

	err := DBCenter.Update(
		CTabNameUserinfo,
		userid,
		db.M{"$set": db.M{"bants": bants}},
	)
	if err != nil {
		log.Error("dbmgr.CenterBanUser() failed:", err)
	}
}

func CenterGetBanInfo(userid string) time.Time {
	var obj UserInfo

	err := DBCenter.GetProjectionByCond(
		CTabNameUserinfo,
		db.M{"_id": userid},
		db.M{"ban_ts": 1},
		&obj,
	)
	if err != nil {
		log.Error("dbmgr.CenterGetBanInfo() failed:", err)
		return time.Unix(0, 0)
	}

	return obj.BanTs
}

func CenterUpdateUserToken(channel int32, uid string, token, refrtoken string, expire int64) {
	obj := CenterGetUserInfo(channel, uid)
	if obj == nil {
		log.Warning("CenterGetUserInfo error", obj)
		return
	}

	err := DBCenter.Upsert(
		CTabNameAccount,
		db.M{
			"channel":     channel,
			"channel_uid": uid,
		},
		db.M{
			"$set": db.M{
				"token":    token,
				"r_token":  refrtoken,
				"expire_t": time.Now().Add(time.Duration(expire) * time.Second),
			},
		},
	)

	if err != nil {
		log.Warning("CenterUpdateAccountInfo error", err)
		return
	}
}

func CenterUpdateUserNameHead(channel int32, uid string, name, head string) {
	obj := CenterGetUserInfo(channel, uid)
	if obj == nil {
		log.Warning("CenterGetUserInfo error", obj)
		return
	}

	err := DBCenter.Update(
		CTabNameAccount,
		db.M{
			"channel":     channel,
			"channel_uid": uid,
		},
		db.M{
			"$set": db.M{
				"name": name,
				"head": head,
			},
		},
	)

	if err != nil {
		log.Warning("CenterUpdateAccountInfo error", err)
		return
	}
}
