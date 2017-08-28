package c_gs

import (
    "client/msg"
)

func GS_TableInfoNotify(message msg.Message, ctx interface{}) {
    req := message.(*msg.GS_TableInfoNotify)
    req = req
}
