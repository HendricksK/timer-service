package databaseconnector

// This might help
// https://medium.com/easyread/unit-test-sql-in-golang-5af19075e68e

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Construct DB entity, and return it, relatively simple. Can swap types of connectors more easily this way
// postgres://postgres:123456@127.0.0.1:5432/dummy
var postGresConnStr = os.Getenv("POSTGRES_DATABASE_URL")

var postGresDB, err = sql.Open("postgres", postGresConnStr)

// on localhost we need to disabled sslmode / I am lazy
// var postGresConnStr = "postgres://postgres:postgres@localhost:5432/timer?sslmode=disable"

// var mySQLConnStr = os.Getenv("MYSQL_DATABASE_URL")

func Init() {
	if postGresDB == nil {
		log.Fatal(err)
		panic("Database is down")
	}
}

func GetPostgresDatabaseHandler() *sql.DB {
	var postGresDB, err = sql.Open("postgres", postGresConnStr)
	if err != nil {
		log.Fatal(err)
		panic("Database is down")
	}

	return postGresDB
}

func CloseDBConnection(db *sql.DB) error {
	return db.Close()
}

// Will be using this to see how viable it would be to create a custom ORM
// https://trello.com/c/gQyqoozM

// func Update() struct {
// 	row := db.QueryRow("SELECT id FROM "+tableName+" WHERE ref = $1", update.Ref)

// 	if row == nil {
// 		database.CloseDBConnection(db)
// 		return timer
// 	}

// 	row.Scan(&update.Id)

// 	// iterate over timer struct to create prepared update, do not want to update fields that we do not have
// 	// Very helpful in my persuit https://www.golangprograms.com/how-to-get-struct-variable-information-using-reflect-package.html
// 	// Can build all of this into a function in the database package
// 	updateReflection := reflect.ValueOf(&update).Elem()

// 	var fieldsBeingUpdated string
// 	noUpdatesAllowed := []string{"Id", "Ref"}
// 	for i := 0; i < updateReflection.NumField(); i++ {

// 		if updateReflection.Type().Field(i).Name != "" && !slices.Contains(noUpdatesAllowed, updateReflection.Type().Field(i).Name) && updateReflection.Field(i).Interface() != "" {
// 			// https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
// 			// want to pass the value type, and then parse in the actual value
// 			// Will need to look into this to ensure we cannot SQL inject.
// 			// What we can do is rather build out an array of no type (hopefully and add values to that, only the values we want to update)
// 			if strings.Contains(updateReflection.Type().Field(i).Type.String(), "int") { // integer
// 				fieldsBeingUpdated += fmt.Sprintf("%v = %v,", strings.ToLower(updateReflection.Type().Field(i).Name), updateReflection.Field(i).Int())
// 			}

// 			if strings.Contains(updateReflection.Type().Field(i).Type.String(), "string") { // string
// 				fieldsBeingUpdated += fmt.Sprintf("%v = '%s',", strings.ToLower(updateReflection.Type().Field(i).Name), sanitize.XSS(updateReflection.Field(i).String()))
// 			}
// 		}
// 	}

// 	fieldsBeingUpdated = strings.TrimRight(fieldsBeingUpdated, ",")

// 	sqlStatement := `
// 		UPDATE ` + tableName + `
// 		SET ` + fieldsBeingUpdated + `
// 		WHERE ref = $1`

// 	// Need to get all non empty fields and then drop them in preparedQuery
// 	// https://stackoverflow.com/questions/49965387/use-slice-reflect-type-as-map-key
// 	updateResult, updateErr := db.Exec(sqlStatement, update.Ref)
// 	// updateResult, updateErr := db.Query(sqlStatement, update.Ref)
// }
