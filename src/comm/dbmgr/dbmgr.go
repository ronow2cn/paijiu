package dbmgr

import (
	"comm/config"
	"comm/db"
)

// ============================================================================

const (
	// center
	CTabNameUserinfo = "userinfo"
	CTabNameUserload = "userload"
	CTabNameSeqid    = "seqid"
	CTabNameNames    = "names"
	CTabNameAccount  = "account"

	// game
	CTabNameWorlddata = "worlddata"

	// log
	CTabNameLog = "log"

	// user
	CTabNameUser = "user"
)

// ============================================================================

var (
	DBCenter *db.Database
	DBBill   *db.Database
	DBGame   *db.Database
	DBLog    *db.Database
)

var dbUser = map[string]*db.Database{}

// ============================================================================

func GameOpen() {
	// 初始化 中心 数据库
	if DBCenter == nil {
		DBCenter = db.NewDatabase()
		DBCenter.Open(config.Common.DBCenter, false)
	}

	CenterCreateSeqId()
	CenterCreateUserLoad()

	DBCenter.CreateIndex(CTabNameUserinfo, "idx_svr", []string{"svr"}, false)
	DBCenter.CreateIndex(CTabNameUserinfo, "idx_name", []string{"name"}, false)

	DBCenter.CreateIndex(CTabNameNames, "uk_name", []string{"name"}, true)

	// 初始化 游戏 数据库
	if DBGame == nil {
		DBGame = db.NewDatabase()
		DBGame.Open(config.DefaultGame.DBGame, false)
	}

	// 初始化 日志 数据库
	if DBLog == nil {
		DBLog = db.NewDatabase()
		DBLog.Open(config.DefaultGame.DBLog, false)
	}

	DBLog.CreateIndex(CTabNameLog, "idx_op", []string{"op"}, false)
	DBLog.CreateIndex(CTabNameLog, "idx_ts", []string{"ts"}, false)
	DBLog.CreateIndex(CTabNameLog, "idx_uid", []string{"uid"}, false)

	// 初始化 用户 数据库
	for k, v := range config.Common.DBUser {
		if dbUser[k] == nil {
			db := db.NewDatabase()
			db.Open(v, false)

			dbUser[k] = db
		}
	}
}

func GameClose() {
	DBCenter.Close()
	DBBill.Close()
	DBGame.Close()
	DBLog.Close()

	for _, db := range dbUser {
		db.Close()
	}
}

func UserDB(dbname string) *db.Database {
	return dbUser[dbname]
}

// ============================================================================

func GateOpen() {
	// init center db
	if DBCenter == nil {
		DBCenter = db.NewDatabase()
		DBCenter.Open(config.Common.DBCenter, true)
	}
}

func GateClose() {
	DBCenter.Close()
}

// ============================================================================

func AuthOpen() {
	// init center db
	if DBCenter == nil {
		DBCenter = db.NewDatabase()
		DBCenter.Open(config.Common.DBCenter, true)
	}
}

func AuthClose() {
	DBCenter.Close()
}
