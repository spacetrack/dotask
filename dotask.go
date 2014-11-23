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

	var loc = time.Now().Local().Location()

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
	case "-v", "--version", "version":
		writeUpdate = false

		fmt.Println("dotask version 0.9.3 - 2014-10-19; (c) 2014 by Björn Winkler")
		fmt.Println("/* debug */")
		fmt.Println(time.Now().Local().Zone())
		fmt.Println(os.Getenv("TZ"))
		os.Exit(0)

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
			os.Exit(-1)
		}

		fmt.Println("DELETED:")
		fmt.Println(t)

	// ---
	// --- command "new"
	// --- ... create a new task and store it
	// ---
	case "n", "now":
		writeUpdate = true

		t = task.NewTask()
		t.Timestamp = time.Now().Local()
		t.Title = strings.Join(os.Args[2:], " ")

		tasks[t.Id] = t
		fmt.Println(t)

// ---
// --- command "shutdown"
// --- ... creates a new taks to indicate end of day.
// --  almost the same as "dotask now shutdown"
// ---
// --- example: dotask shutdown now
// --- example: dotask shutdown 20:30
// ---
case "shutdown":
	writeUpdate = true
	t = task.NewTask()
	t.Title = "shutdown"

	if(len(os.Args) > 2) {
		if(os.Args[2] == "now") {
			t.Timestamp = time.Now().Local()
		} else {
			// time only: "HH:mm"
			if len(os.Args[2]) == 5 {
				today := time.Now().Format("2006-01-02")
				t.Timestamp, err = time.ParseInLocation("2006-01-02 15:04", today+" "+os.Args[2], loc)
			}

			// date + time: "dd.mm.YYYY-HH:mm"
			if len(os.Args[2]) == 16 {
				t.Timestamp, err = time.ParseInLocation("02.01.2006-15:04", os.Args[2], loc)
			}

			// date + time: "dd.mm.YY-HH:mm"
			if len(os.Args[2]) == 14 {
				t.Timestamp, err = time.ParseInLocation("02.01.06-15:04", os.Args[2], loc)
			}
		}
	} else {
		t.Timestamp = time.Now().Local()
	}

	tasks[t.Id] = t
	fmt.Println(t)


	// ---
	// --- command "clone" / "continue"
	// --- ... clone existing task and store it with new timestamp
	// ---
	// --- example: dotask c 1413491670-9055 now
	// ---
	case "c", "clone", "continue":
		writeUpdate = true

		tSource, ok := tasks[os.Args[2]]

		if ok {
			t = task.NewTask()
			t.Title = tSource.Title

			if(len(os.Args) > 3) {
				if(os.Args[3] == "now") {
					t.Timestamp = time.Now().Local()
				} else {
					// time only: "HH:mm"
					if len(os.Args[3]) == 5 {
						today := time.Now().Format("2006-01-02")
						t.Timestamp, err = time.ParseInLocation("2006-01-02 15:04", today+" "+os.Args[3], loc)
					}

					// date + time: "dd.mm.YYYY-HH:mm"
					if len(os.Args[3]) == 16 {
						t.Timestamp, err = time.ParseInLocation("02.01.2006-15:04", os.Args[3], loc)
					}

					// date + time: "dd.mm.YY-HH:mm"
					if len(os.Args[3]) == 14 {
						t.Timestamp, err = time.ParseInLocation("02.01.06-15:04", os.Args[3], loc)
					}
				}
			} else {
				t.Timestamp = time.Now().Local()
			}
		} else {
			fmt.Println("CLONE failed")
			fmt.Println(ok)
			os.Exit(-1)
		}

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
			t.Timestamp = time.Now().Local()

		default:
			// time only: "HH:mm"
			if len(os.Args[2]) == 5 {
				today := time.Now().Format("2006-01-02")
				t.Timestamp, err = time.ParseInLocation("2006-01-02 15:04", today+" "+os.Args[2], loc)
			}

			// date + time: "dd.mm.YYYY-HH:mm"
			if len(os.Args[2]) == 16 {
				t.Timestamp, err = time.ParseInLocation("02.01.2006-15:04", os.Args[2], loc)
			}

			// date + time: "dd.mm.YY-HH:mm"
			if len(os.Args[2]) == 14 {
				t.Timestamp, err = time.ParseInLocation("02.01.06-15:04", os.Args[2], loc)
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
