package main

import "database/sql"

type NS struct{ sql.NullString }

func (s NS) String() string {
	return s.NullString.String
}
