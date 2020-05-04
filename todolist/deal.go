package todolist

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Deal - deal to do
type Deal struct {
	ID                 string
	Name               string `json:"name"`
	Description        string `json:"description"`
	Date               string `json:"date"`
	StartDate, EndDate time.Time
	IsComplete         bool
}

// Init - initialize deal
func (d *Deal) Init() {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	d.ID = id.String()
	d.StartDate = time.Now()
	date, err := time.Parse("02.01.2006 15:4", d.Date)
	if err != nil {
		panic(err)
	}
	d.EndDate = date
}

func (d Deal) String() string {
	return fmt.Sprintf("ID: %s Name: %s", d.ID, d.Name)
}

// Complete - complete deal
func (d *Deal) Complete() {
	d.IsComplete = true
}
