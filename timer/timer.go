package timer

import (
	"fmt"
	"log"
	"os"
	"time"

	database "github.com/HendricksK/timer-service/database-connector"
	"github.com/gin-gonic/gin"
)

type Timer struct {
	Id            uint64 `json:"id"`
	Ref           string `json:"ref"`
	ProjectRef    string `json:"project_ref"`
	PreviousValue string `json:"previous_value"`
	CurrentValue  string `json:"current_value"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Notes         string `json:"notes"`
	Created       string `json:"created"`
	ModifiedAt    string `json:"modified_at"`
	Deleted       bool   `json:"deleted"`
}

var env string

var timers = []Timer{
	{
		Id:         1,
		Ref:        "AQFs1ggyP8sXqyfghi9g",
		Created:    time.Now().String(),
		ModifiedAt: "",
		Deleted:    false,
	},
	{
		Id:         2,
		Ref:        "wqdwqdwd878736gefduh",
		Created:    time.Now().String(),
		ModifiedAt: "",
		Deleted:    false,
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

	err := db.QueryRow("SELECT * FROM timer WHERE ref = $1", ref).Scan(
		&timer.Id,
		&timer.Ref,
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

func Create(c *gin.Context) []Timer {
	var data Timer

	// data.Id = timers[len(timers)-1].Id + 1
	// Id will be set on insert
	data.Ref = c.PostForm("ref")
	data.Created = time.Now().String()
	data.ModifiedAt = ""
	data.Deleted = false

	return timers
}

func Update(ref string, c *gin.Context) []Timer {

	return timers

}

func Delete(ref string) []Timer {

	return timers
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

	dataUpdate.ModifiedAt = time.Now().String()
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
