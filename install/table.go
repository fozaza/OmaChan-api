package install

import (
	"github.com/OmaChan/database"
	"github.com/OmaChan/database/table"
)

func Install_table() {
	db := database.Get_db()
	table_all := []any{&table.User{}}
	db.AutoMigrate(table_all...)
}
