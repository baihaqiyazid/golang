package simple

type Database struct{
	Name string
}

//Alias
type DatabasePostgreSQL Database
type DatabaseMongoDB Database

//Constructor or Provider
func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "PostgreSQL"})
}

type DatabaseRepository struct{
	*DatabasePostgreSQL
	*DatabaseMongoDB
}

func NewDatabaseRepository(PostgreSQL *DatabasePostgreSQL, MongoDB *DatabaseMongoDB) *DatabaseRepository  {
	return &DatabaseRepository{
		DatabasePostgreSQL: PostgreSQL,
		DatabaseMongoDB: MongoDB,
	}
}

