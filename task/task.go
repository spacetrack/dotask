package task

import (
	"math/rand"
	"strconv"
	"time"
	"strings"
	"sort"
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

func GetTask(id string, tasks map[string]*Task) (*Task, bool) {
	t, ok := tasks[id]

	if !ok {
		// add "-" if not part of the string
		if !strings.Contains(id, "-") {
			id = "-" + id
		}

		timeline := []*Task{}

		// find the elements with the id as substring of the key
		for key, aTask := range tasks {
			if strings.Contains(key, id) {
				timeline = append(timeline, aTask)
			}
		}

		// sort timeline
		sort.Sort(ByDate(timeline))

		// return last element of timeline
		if len(timeline) > 0 {
			return timeline[len(timeline)-1], true
		} else {
			return nil, false
		}
	}

	return t, ok
}

func (t *Task) String() string {
	var result = ""

	result = result + t.Id + " | "
	result = result + t.Timestamp.Format("2006-01-02 15:04") + " | "
	result = result + t.Title

	return result
}


// JSON support
// https://eager.io/blog/go-and-json/

//func (t *Task) MarshalJSON() ([]byte, error) {
//	return []byte(fmt.Sprintf("%d/%d", m.MonthNumber, m.YearNumber)), nil
//}

//func (t *Task) UnmarshalJSON([]byte value) error {
//    parts := strings.Split(string(value), "/")
//    m.MonthNumber = fmt.ParseInt(parts[0], 10, 32)
//    m.YearNumber = fmt.ParseInt(parts[1], 10, 32)
//
//    return nil
//}

// list of tasks sorted by date
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

