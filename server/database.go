package server

type account map[string]string

type database struct {
	Accounts account
}

func NewDatabase() database {
	return database{make(map[string]string)}
}
