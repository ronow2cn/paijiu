package c_gs

import (
	"client/msg"
)

func GS_TableLeave_R(message msg.Message, ctx interface{}) {
	req := message.(*msg.GS_TableLeave_R)
	log.Info("Res:", req.ErrorCode)
}
