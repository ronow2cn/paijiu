package c_gs

import (
	"client/msg"
)

func GS_TableNextPlay_R(message msg.Message, ctx interface{}) {
	req := message.(*msg.GS_TableNextPlay_R)
	log.Info("Res:", req.ErrorCode)
}
