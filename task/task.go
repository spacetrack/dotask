package task

import (
	"math/rand"
	"strconv"
	"time"
)

type Task struct {
	Id          string
	Title       string
	Timestamp   time.Time
	External_ID string
	Parent_ID   string
}

func NewTask() *Task {
	rand.Seed(time.Now().UnixNano())
	r := 1000 + rand.Intn(9999-1000)

	return &Task{
		Id: strconv.FormatInt(time.Now().Unix(), 10) + "-" + strconv.Itoa(r),
	}
}

func (t *Task) String() string {
	var result = ""

	result = result + t.Id + " | "
	result = result + t.Timestamp.Format("2006-01-02 15:04") + " | "
	result = result + t.Title

	return result
}

type ByDate []*Task

func (t ByDate) Len() int {
	return len(t)
}

func (t ByDate) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t ByDate) Less(i, j int) bool {
	return t[i].Timestamp.Before(t[j].Timestamp)
}
