package migrations

import (
	"github.com/Dalot/goddd/internal/app/adapter/mysql"
	"github.com/Dalot/goddd/internal/app/adapter/mysql/model"
)

func Migrate() {

	mysql.Connection().AutoMigrate(model.User{}, model.Project{})
	//we take the structs we created earlier to represent tables and create tables from them.
	//for example models.Users{} will create a table called users  with the fields we defined in that struct as the table fields.
	///if by any chance you ever add another struct you need to create a table for you can add it here.

}
