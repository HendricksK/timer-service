package timer

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	database "github.com/HendricksK/timer-service/database-connector"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

// "project_ref": "qwerty12123",
// "name": "qwerty",
// "description": "qwerty",
// "notes": "qwerty",
// "timezone": "UTC"

type Timer struct {
	Id             uint64 `json:"id"`
	Ref            string `json:"ref"`
	Project_ref    string `json:"project_ref,omitempty"`
	User_ref       string `json:"user_ref,omitempty"`
	Previous_value string `json:"previous_value,omitempty"`
	Current_value  string `json:"current_value,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Notes          string `json:"notes,omitempty"`
	Created        string `json:"created,omitempty"`
	Modified_at    string `json:"modified_at,omitempty"`
	Deleted        int    `json:"deleted,omitempty"`
	Timezone       string `json:"timezone,omitempty"`
}

var tableName string = "timer"
var env string

var timers = []Timer{
	{
		Id:          1,
		Ref:         "AQFs1ggyP8sXqyfghi9g",
		Created:     time.Now().String(),
		Modified_at: "",
		Deleted:     0,
	},
	{
		Id:          2,
		Ref:         "wqdwqdwd878736gefduh",
		Created:     time.Now().String(),
		Modified_at: "",
		Deleted:     0,
	},
}

func Init() string {
	// At some point I would like to implement an interface
	// that will run tests for local development based on this ENV
	env = os.Getenv("ENV")
	fmt.Println(env)
	return "yes"
}

// We set mockdata here
func GetTestTimer() []Timer {
	return timers
}

func Read(limit string) []Timer {
	var db = database.GetPostgresDatabaseHandler()
	var data []Timer

	// https://go.dev/doc/database/sql-injection
	// rows, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	rows, err := db.Query("SELECT * FROM timer ORDER BY id DESC LIMIT $1", limit)

	if err != nil {
		log.Println(err)
		fmt.Println(err)
		database.CloseDBConnection(db)
		return []Timer{}
	}

	defer rows.Close()

	for rows.Next() {
		var timer Timer
		err = rows.Scan(
			&timer.Id,
			&timer.Ref,
		)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(timer)
		data = append(data, timer)
	}

	database.CloseDBConnection(db)

	return data
}

// https://github.com/golang/go/wiki/SliceTricks
func ReadByRef(ref string) Timer {

	var db = database.GetPostgresDatabaseHandler()
	var timer Timer

	// Id             uint64 `json:"id"`
	// Ref            string `json:"ref"`
	// Project_ref    string `json:"project_ref"`
	// Previous_value string `json:"previous_value"`
	// Current_value  string `json:"current_value"`
	// Name           string `json:"name"`
	// Description    string `json:"description"`
	// Notes          string `json:"notes"`
	// Created        string `json:"created"`
	// Modified_at    string `json:"modified_at"`
	// Deleted        int    `json:"deleted"`
	// Timezone       string `json:"timezone"`

	err := db.QueryRow("SELECT * FROM timer WHERE ref = $1", ref).Scan(
		&timer.Id,
		&timer.Ref,
		&timer.Project_ref,
		&timer.User_ref,
		&timer.Previous_value,
		&timer.Current_value,
		&timer.Name,
		&timer.Description,
		&timer.Notes,
		&timer.Created,
		&timer.Modified_at,
		&timer.Deleted,
		&timer.Timezone,
	)

	if err != nil {
		log.Println(err)
		fmt.Println(err)
		database.CloseDBConnection(db)
		return timer
	}

	database.CloseDBConnection(db)
	return timer
}

func Create(c *gin.Context) Timer {
	var db = database.GetPostgresDatabaseHandler()
	var timer Timer
	var data Timer
	var id uint64

	c.BindJSON(&data)

	sqlStatement := `
		INSERT INTO timer (data
			$2, 
			$3, 
			$4,
			$5
		)
		RETURNING "id"`

	preparedQuery, err := db.Prepare(sqlStatement)

	if err != nil {
		log.Println(err)
		database.CloseDBConnection(db)
	}

	defer preparedQuery.Close()

	errInsert := preparedQuery.QueryRow(
		data.Project_ref,
		data.Name,
		data.Description,
		data.Notes,
		data.Timezone,
	).Scan(&id)

	fmt.Println(id)

	if errInsert != nil {
		log.Println(err)
		database.CloseDBConnection(db)
		return timer
	}

	getErr := db.QueryRow("SELECT * FROM timer WHERE id = $1", id).Scan(
		&timer.Id,
		&timer.Ref,
		&timer.Project_ref,
		&timer.User_ref,
		&timer.Previous_value,
		&timer.Current_value,
		&timer.Name,
		&timer.Description,
		&timer.Notes,
		&timer.Created,
		&timer.Modified_at,
		&timer.Deleted,
		&timer.Timezone,
	)

	if getErr != nil {
		log.Println(err)
		database.CloseDBConnection(db)
		return timer
	}

	database.CloseDBConnection(db)

	return timer
}

func Update(c *gin.Context) Timer {

	var db = database.GetPostgresDatabaseHandler()
	var update, timer Timer

	// Id             uint64 `json:"id"`
	// Ref            string `json:"ref"`
	// Project_ref    string `json:"project_ref"`
	// Previous_value string `json:"previous_value"`
	// Current_value  string `json:"current_value"`
	// Name           string `json:"name"`
	// Description    string `json:"description"`
	// Notes          string `json:"notes"`
	// Created        string `json:"created"`
	// Modified_at    string `json:"modified_at"`
	// Deleted        int    `json:"deleted"`
	// Timezone       string `json:"timezone"`

	// var project_ref, current_value, name, description, notes, timezone, user_ref string
	// var deleted int

	// Need to figure out how to update only the fields that come via c.BindJSON

	c.BindJSON(&update)

	if update.Ref == "" {
		fmt.Println("for now I am just chilling" + update.Ref)
		return timer
	}

	row, err := db.Query("SELECT count(id) FROM timer WHERE ref = $1", update.Ref)

	if row == nil {
		database.CloseDBConnection(db)
		return timer
	}

	if err != nil {
		log.Println(err)
		database.CloseDBConnection(db)
		return timer
	}

	// iterate over timer struct to create prepared update, do not want to update fields that we do not have
	// Very helpful in my persuit https://www.golangprograms.com/how-to-get-struct-variable-information-using-reflect-package.html
	// Can build all of this into a function in the database package
	updateReflection := reflect.ValueOf(&update).Elem()

	var fieldsBeingUpdated string
	// var valuesBeingUpdated string
	// valuesBeingUpdatedCounter := 0
	noUpdatesAllowed := []string{"Id", "Ref"}

	for i := 0; i < updateReflection.NumField(); i++ {

		if updateReflection.Type().Field(i).Name != "" && !slices.Contains(noUpdatesAllowed, updateReflection.Type().Field(i).Name) && updateReflection.Field(i).Interface() != "" {
			// valuesBeingUpdatedCounter++
			// https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
			// want to pass the value type, and then parse in the actual value
			// Will need to look into this to ensure we cannot SQL inject.
			// What we can do is rather build out an array of no type (hopefully and add values to that, only the values we want to update)

			if strings.Contains(updateReflection.Type().Field(i).Type.String(), "int") { // integer
				fieldsBeingUpdated += fmt.Sprintf("%v = %v,", strings.ToLower(updateReflection.Type().Field(i).Name), updateReflection.Field(i).Interface())
			}

			if strings.Contains(updateReflection.Type().Field(i).Type.String(), "string") { // string
				fieldsBeingUpdated += fmt.Sprintf("%v = '%s',", strings.ToLower(updateReflection.Type().Field(i).Name), updateReflection.Field(i).Interface())
			}
		}
	}

	fieldsBeingUpdated = strings.TrimRight(fieldsBeingUpdated, ",")
	// valuesBeingUpdated = strings.TrimRight(valuesBeingUpdated, ",")

	sqlStatement := `
		UPDATE ` + tableName + `
		SET ` + fieldsBeingUpdated + `
		WHERE ref = $1`

	// fmt.Println(sqlStatement)
	// fmt.Println(valuesBeingUpdated)

	// Need to get all non empty fields and then drop them in preparedQuery
	// https://stackoverflow.com/questions/49965387/use-slice-reflect-type-as-map-key
	updateErr, updateResult := db.Query(sqlStatement, update.Ref)

	fmt.Println(updateErr)
	fmt.Println(updateResult)
	// defer preparedQuery.Close()

	// errUpdate := preparedQuery.QueryRow(
	// 	update,
	// )

	// fmt.Println(errUpdate)

	// Can build all of this into a function in the database package

	// preparedQuery, err := db.Prepare(sqlStatement)

	return timer

	// updateErr := db.QueryRow("SELECT * FROM timer WHERE id = $1", id).Scan(
	// 	&update.Id,
	// 	&update.Ref,
	// 	&update.Project_ref,
	// 	&update.Previous_value,
	// 	&update.Current_value,
	// 	&update.Name,
	// 	&update.Description,
	// 	&update.Notes,
	// 	&update.Created,
	// 	&update.Modified_at,
	// 	&update.Deleted,
	// 	&update.Timezone,
	// )

	if err != nil {
		log.Println(err)
		fmt.Println(err)
		database.CloseDBConnection(db)
		return timer
	}

	return timer
}

func Delete(ref string) bool {
	var db = database.GetPostgresDatabaseHandler()
	res, err := db.Exec("DELETE FROM timer WHERE ref = $1", ref)

	if err == nil {

		count, err := res.RowsAffected()
		if err != nil {
			database.CloseDBConnection(db)
			return false
		}
		if count == 0 {
			database.CloseDBConnection(db)
			return false
		}

	}

	database.CloseDBConnection(db)

	return true
}

// Tests
func TestRead() []Timer {
	return timers
}

// https://github.com/golang/go/wiki/SliceTricks
func TestReadById(ref string) Timer {
	var data Timer

	for _, timer := range timers {
		if timer.Ref == ref {
			data = timer
		}
	}

	return data
}

func TestCreate(newTimers []Timer) []Timer {

	timers = append(timers, newTimers...)
	return timers
}

func TestUpdate(ref string, c *gin.Context) []Timer {

	var dataUpdate Timer
	var data Timer

	dataUpdate.Modified_at = time.Now().String()
	dataUpdate.Notes = "Hello there"
	// https://forum.golangbridge.org/t/update-values-in-a-struct-through-a-method/19589
	// https://developer20.com/pointer-and-value-semantics-in-go/
	for _, timer := range timers {
		if timer.Ref == ref {
			data = timer

		}
	}

	fmt.Println(data)

	return timers

}

func TestDelete(ref string) []Timer {

	return timers
}
