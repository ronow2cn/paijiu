package c_gs

import (
	"client/msg"
)

func GS_TableEnter_R(message msg.Message, ctx interface{}) {
	req := message.(*msg.GS_TableEnter_R)
	log.Info("Res:", req.ErrorCode)
}
