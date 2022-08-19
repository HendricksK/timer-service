package timer

import (
	"time"

	"github.com/gin-gonic/gin"
)

type timer struct {
	Id         int    `json:"id"`
	Ref        string `json:"ref"`
	Created    string `json:"created"`
	ModifiedAt string `json:"modified_at"`
	Deleted    bool   `json:"deleted"`
}

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
		Ref:        "AQFs1ggyP8sXqyfghi9g",
		Created:    time.Now().String(),
		ModifiedAt: "",
		Deleted:    false,
	},
}

func Init() string {
	return "yes"
}

func Read() []timer {
	return timers
}

func ReadById(ref string) []timer {
	return timers
}

func Create(c *gin.Context) []timer {
	return timers
}

func Update(ref string, c *gin.Context) []timer {
	return timers
}

func Delete(ref string) []timer {
	return timers
}
