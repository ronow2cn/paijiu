package c_gs

import (
	"client/msg"
)

func GS_TableCreate_R(message msg.Message, ctx interface{}) {
	req := message.(*msg.GS_TableCreate_R)
	log.Info("Res:", req.ErrorCode)
}
