package handler

import (
    "client/handler/c_gs"
    "client/handler/c_gw"
    "client/msg"
)

func Init() {
    msg.Handler(1001, c_gw.GW_Login_R)
    msg.Handler(3000, c_gs.GS_LoginError)
    msg.Handler(3001, c_gs.GS_UserInfo)
    msg.Handler(4001, c_gs.GS_TableCreate_R)
    msg.Handler(4003, c_gs.GS_TableEnter)
    msg.Handler(4004, c_gs.GS_TableInfoNotify)
    msg.Handler(4007, c_gs.GS_TableLeave_R)
    msg.Handler(4009, c_gs.GS_TableSeatDown_R)
    msg.Handler(4011, c_gs.GS_TableStandUp)
    msg.Handler(4013, c_gs.GS_TableDice)
    msg.Handler(101, c_gs.GS_Test_R)
}
