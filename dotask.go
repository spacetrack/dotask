package main

import (
	"encoding/json"
	"fmt"
	"github.com/spacetrack/dotask/task"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"sort"
)

func main() {
	var tasks = map[string]*task.Task{}
	var t *task.Task

	var writeUpdate = false
	var blob []byte
	var err error

	// only 1 argument? Fail!
	if len(os.Args) < 2 {
		fmt.Println("ERROR: invalid number of arugments; use \"dotask help\" for list of valid arguments")
		os.Exit(-1)
	}

	blob, err = ioutil.ReadFile("my_tasks.json")
	err = json.Unmarshal(blob, &tasks)

	switch os.Args[1] {

	// ---
	// --- command "help"
	// --- ... print the help text
	// ---
	case "--help", "help":
		writeUpdate = false

		fmt.Println("dotask (l)ist|(s)how")
		fmt.Println("dotask (n)ow <title>")
		fmt.Println("dotask (c)ontinue ID [_now_|<timestamp>]")
		fmt.Println("dotask delete ID")
		fmt.Println("dotask shutdown")
		fmt.Println()
		fmt.Println("dotask ID")
		fmt.Println("dotask ID asis|now|<timestamp> [<title> ...]")

		os.Exit(0)

	// ---
	// --- command "list"
	// --- ... print the help text
	// ---
	case "l", "list", "s", "show":
		writeUpdate = false
		thisDay := ""

		timeline := []*task.Task{}

		fmt.Println("Your tasks:")

		for _, aTask := range tasks {
			timeline = append(timeline, aTask)
		}

		sort.Sort(task.ByDate(timeline))

		for _, aTask := range timeline {

			if thisDay != aTask.Timestamp.Format("2006-01-02") {
				thisDay = aTask.Timestamp.Format("2006-01-02")
				fmt.Println("----------------------------------------")
			}

			fmt.Println(aTask)
		}

		os.Exit(0)

	case "delete":
		writeUpdate = true

		t, ok := tasks[os.Args[2]]

		if ok {
			delete(tasks, os.Args[2])
		} else {
			fmt.Println("DELETE failed")
			fmt.Println(ok)
		}

		fmt.Println("DELETED:")
		fmt.Println(t)

	// ---
	// --- command "new"
	// --- ... create a new task and store it
	// ---
	case "now":
		writeUpdate = true

		t = task.NewTask()
		t.Timestamp = time.Now()
		t.Title = strings.Join(os.Args[2:], " ")

		tasks[t.Id] = t
		fmt.Println(t)

	// ---
	// --- no command, but "<ID>"
	// --- ... create a new task if ID == 0 or
	// --  update existing task and store it
	// ---
	default:
		writeUpdate = true

		// first: the task ID
		t, ok := tasks[os.Args[1]]

		if ok {
			fmt.Println(t)
		} else {
			if os.Args[1] == "0" {
				t = task.NewTask()
			} else {
				fmt.Println("unkown task id \"" + os.Args[1] + "\"; use \"0\" to create a new task")
				os.Exit(-1)
			}
		}

		// second: asis, now, time or date + time
		switch os.Args[2] {
		case "asis":
			// nothing

		case "now":
			t.Timestamp = time.Now()

		default:
			// time only: "HH:mm"
			if len(os.Args[2]) == 5 {
				today := time.Now().Format("2006-01-02")
				t.Timestamp, err = time.Parse("2006-01-02 15:04", today+" "+os.Args[2])
			}

			// date + time: "dd.mm.YYYY-HH:mm"
			if len(os.Args[2]) == 16 {
				t.Timestamp, err = time.Parse("02.01.2006-15:04", os.Args[2])
			}

			// date + time: "dd.mm.YY-HH:mm"
			if len(os.Args[2]) == 14 {
				t.Timestamp, err = time.Parse("02.01.06-15:04", os.Args[2])
			}

		}

		// third: the title
		if(len(os.Args) > 3) {
			if(os.Args[3] != "asis") {
				t.Title = strings.Join(os.Args[3:], " ")
			}
		}

		tasks[t.Id] = t
		fmt.Println(t)

	}

	// ---
	// --- store data ...
	// ---

	if writeUpdate {
		blob, err = json.MarshalIndent(tasks, "", "	")

		if err != nil {
			fmt.Println("ERROR: can't print tasks as JSON [", err, "]")
			os.Exit(-1)
		}

		ioutil.WriteFile("my_tasks.json", blob, 0777)
	}
}
