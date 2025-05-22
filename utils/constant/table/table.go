package table

import "github.com/doug-martin/goqu/v9"

var (
	Users = goqu.T("tm_accounts")
	Roles = goqu.T("tm_roles")
)
