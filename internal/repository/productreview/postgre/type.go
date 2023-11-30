package postgre

import "database/sql"

type (
	ProductReviewRepo struct {
		db        DBList
		statement map[string]*sql.Stmt
	}

	DBList struct {
		TelkomRead  *sql.DB
		TelkomWrite *sql.DB
	}

	InitAttribute struct {
		DB DBList
	}
)
