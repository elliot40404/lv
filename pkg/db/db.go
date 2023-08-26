package db

type DB struct {
	Logs []string
}

var db *DB

func InitDB() int {
	db = &DB{
		Logs: []string{},
	}
	return len(db.Logs)
}

func Open() *DB {
	return db
}

func (db *DB) AddLog(log string) {
	db.Logs = append(db.Logs, log)
}

func (db *DB) GetLastLog() string {
	return db.Logs[len(db.Logs)-1]
}

func (db *DB) GetLogs() []string {
	return db.Logs
}
