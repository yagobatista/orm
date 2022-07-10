package db

func AutoMiograte() {
	// if _, err := globalConnection.db.Exec("Create database orm;"); err != nil {
	// 	panic(err)
	// }

	if _, err := globalConnection.db.Exec("Create table users (name VARCHAR(128));"); err != nil {
		panic(err)
	}

}
