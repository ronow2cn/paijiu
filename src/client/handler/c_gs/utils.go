package c_gs

import (
	"client/msg"
	"comm/logger"
	"time"
)

var log = logger.DefaultLogger

func PrintTableData(m *msg.TableData) {
	if m == nil {
		log.Error("PrintTableData nil")
		return
	}
	log.Warningf("===========Table: %d Create time %v============", m.Id, time.Unix(m.CreateTs, 0).Format("2006-01-02 03:04:05 PM"))
	log.Warningf("cur play idx: %d, dicenum: %d", m.PlayIdx, m.DiceNum)

	log.Warningf("----------------------------------------")
	for k, v := range m.Plrs {
		log.Warningf("| Table Player: %s, name: %s, head: %s, score: %d |", k, v.Name, v.Head, v.Score)
	}
	log.Warningf("----------------------------------------")

	log.Warningf("----------------------------------------")
	for k, v := range m.Pos {
		log.Warningf("| Seat Player: pos %d, player: %s |", k, v)
	}
	log.Warningf("----------------------------------------")

	if m.CurPlay == nil {
		log.Error("PrintTableData CurPlay nil")
		log.Warningf("====================================================================")

		return
	}

	for k, v := range m.CurPlay.Chips {
		log.Warningf("+++++++++++ CurPlay chip: pos %d ++++++", k)
		for kk, vv := range v.Bets {
			log.Warningf("bets: player %s, betnum %d", kk, vv)
		}
		log.Warningf("++++++++++++++++++++++++++++++++++++++++")

	}

	log.Error("====================================================================")
}
