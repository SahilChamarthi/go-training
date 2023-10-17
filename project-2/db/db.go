package db

type Conn struct {
	db string
}

func NewConn(dbname string) (Conn, bool) {

	if dbname == "" {
		return Conn{}, false
	}
	return Conn{db: dbname}, true
}
