package c_gs

import (
	"client/msg"
)

func GS_TableGetRecord_R(message msg.Message, ctx interface{}) {
	req := message.(*msg.GS_TableGetRecord_R)
	log.Info("Res:", req.ErrorCode)

	log.Warning("===============record===============")
	for _, v := range req.Records {
		log.Warningf("---------------play idx %d----------------", v.Id)
		for k, vv := range v.PosCard {
			log.Warningf("pos :%d card1: {%d, %d}, card2: {%d, %d}", k, vv[0].T, vv[0].N, vv[1].T, vv[1].N)
		}
	}

	log.Warning("====================================")
}
