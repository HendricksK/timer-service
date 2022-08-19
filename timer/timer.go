package timer

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type timer struct {
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

var timers = []timer{
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
	env = os.Getenv("ENV")
	fmt.Println(env)
	return "yes"
}

// We set mockdata here
func GetTestTimer() []timer {
	return timers
}

func Read() []timer {
	return timers
}

// https://github.com/golang/go/wiki/SliceTricks
func ReadById(ref string) timer {
	var data timer

	for _, timer := range timers {
		if timer.Ref == ref {
			data = timer
		}
	}

	return data
}

func Create(c *gin.Context) []timer {
	var data timer

	// data.Id = timers[len(timers)-1].Id + 1
	// Id will be set on insert
	data.Ref = c.PostForm("ref")
	data.Created = time.Now().String()
	data.ModifiedAt = ""
	data.Deleted = false

	return timers
}

func Update(ref string, c *gin.Context) []timer {

	return timers

}

func Delete(ref string) []timer {

	return timers
}

// Tests
func TestRead() []timer {
	return timers
}

// https://github.com/golang/go/wiki/SliceTricks
func TestReadById(ref string) timer {
	var data timer

	for _, timer := range timers {
		if timer.Ref == ref {
			data = timer
		}
	}

	return data
}

func TestCreate(newTimers []timer) []timer {

	timers = append(timers, newTimers...)
	return timers
}

func TestUpdate(ref string, c *gin.Context) []timer {

	var dataUpdate timer
	var data timer

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

func TestDelete(ref string) []timer {

	return timers
}
