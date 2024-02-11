package repository

func AddCollection(str string) {
	// Select database
	database := GlobalConnection.Database("admin")

	// Create collection (optional, MongoDB creates it implicitly if it doesn't exist)
	database.Collection(str)
}
