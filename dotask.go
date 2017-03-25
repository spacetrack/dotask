package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/spacetrack/dotask/task"
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
	case "v", "-v", "version", "--version":
		writeUpdate = false

		fmt.Println("dotask version 0.9.5 - 2017-03-25 - (c) 2015-2017 by Björn Winkler")
		os.Exit(0)

	case "?", "-?", "h", "-h", "help", "--help":
		writeUpdate = false

		fmt.Println("list all tasks:")
		fmt.Println("	dotask l|list")
		fmt.Println("")
		fmt.Println("create new task:")
		fmt.Println("	dotask now-|now|now+ <title>")
		fmt.Println("	dotask n|new|0 now-|now|now+|<timestamp> <title>")
		fmt.Println("")
		fmt.Println("update existing task:")
		fmt.Println("	dotask u|update ID asis|now-|now|now+|<timestamp> [<title>]")
		fmt.Println("")
		fmt.Println("clone existing task as new one:")
		fmt.Println("	dotask c|clone|continue ID [now-|now|now+|<timestamp>]")
		fmt.Println("")
		fmt.Println("delete a task:")
		fmt.Println("	dotask delete ID")
		fmt.Println("")
		fmt.Println("more commands:")
		fmt.Println("	dotask ?|-?|help|-h|--help")
		fmt.Println("	dotask sh|shutdown [now-|now|now+|<timestamp>]")
		fmt.Println("	dotask v|version")

		os.Exit(0)

	// ---
	// --- command "list"
	// --- ... ll: show all tasks
	// --- pending: filter tasks by number of days; default: -5
	// ---
	case "l", "ll", "list", "s", "show":
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
	// --- command "now"
	// --- ... create a new task and store it
	// ---
	case "now", "now+", "now-":
		writeUpdate = true

		t = task.NewTask()
		t.Timestamp, err = parseTime(os.Args[1])
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
	case "sdown", "shutdown":
		writeUpdate = true
		t = task.NewTask()
		t.Title = "shutdown"

		if len(os.Args) > 2 {
			t.Timestamp, err = parseTime(os.Args[2])
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
		//case "c", "clone", "continue", "create":
		writeUpdate = true

		// pending: if os.Args[2] == 0 then create a new task!

		tSource, ok := tasks[os.Args[2]]

		if ok {
			t = task.NewTask()
			t.Title = tSource.Title

			if len(os.Args) > 3 {
				t.Timestamp, err = parseTime(os.Args[3])
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
	// --- command "new"
	// --- ... create a new task and store it
	// ---
	case "n", "new", "0":
		writeUpdate = true
		t = task.NewTask()

		t.Timestamp, err = parseTime(os.Args[2])

		if len(os.Args) > 3 {
			t.Title = strings.Join(os.Args[3:], " ")
		}

		tasks[t.Id] = t
		fmt.Println(t)

	// ---
	// --- command "u" / "update"
	// --- ... update existing task
	// ---
	// --- example: dotask u 1413491670-9055 asis new text
	// ---
	case "u", "update":
		//case "c", "clone", "continue", "create":
		writeUpdate = true

		// first: the task ID
		t, ok := tasks[os.Args[2]]

		if ok {
			fmt.Println(t)
		} else {
			if os.Args[2] == "0" {
				t = task.NewTask()
			} else {
				fmt.Println("unkown task id \"" + os.Args[2] + "\"; use \"0\" to create a new task")
				os.Exit(-1)
			}
		}

		// second: asis, now, time or date + time
		switch os.Args[3] {
		case "asis":
			// nothing

		default:
			t.Timestamp, err = parseTime(os.Args[3])
		}

		// third: the title
		// ... or "-e", "--edit" for interactive editing (to avoid hassle
		// with quotes and ampersands)
		if len(os.Args) > 4 {
			t.Title = strings.Join(os.Args[4:], " ")
		}

		tasks[t.Id] = t
		fmt.Println(t)


	case "debug":
		fmt.Println(time.Now().Local().Zone())
		fmt.Println(os.Getenv("TZ"))

	// ---
	// --- no command, but "<ID>"
	// --- ... create a new task if ID == 0 or
	// --  update existing task and store it
	// ---
	default:
		writeUpdate = false
		fmt.Println("unkown comamand \"" + os.Args[2] + "\"")
		os.Exit(-1)
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

func parseTime(theTime string) (time.Time, error) {
	var t time.Time
	var err error
	var loc = time.Now().Local().Location()

	if theTime == "now" {
		// the time right now
		t = time.Now().Local()
	} else if theTime == "now+" {
		// the time of the next 5 min
		t = time.Now().Local()
		minute := t.Minute()

		if (minute % 5) > 0 {
			minute = minute + 5 - (minute % 5)
		}

		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minute, 0, 0, loc)
	} else if theTime == "now-" {
		// the time of the next 5 min
		t = time.Now().Local()

		minute := t.Minute()
		minute = minute - (minute % 5)

		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minute, 0, 0, loc)
	} else {
		var aTime string
		today := time.Now().Format("2006-01-02")

		// H:mm
		t, err = time.ParseInLocation("2006-01-02 3:04", today+" "+theTime, loc)

		// h:mmpm
		if err != nil {
			t, err = time.ParseInLocation("2006-01-02 3:04pm", today+" "+theTime, loc)
		}

		// h:mmPM
		if err != nil {
			t, err = time.ParseInLocation("2006-01-02 3:04PM", today+" "+theTime, loc)
		}

		// HH:mm
		if err != nil {
			t, err = time.ParseInLocation("2006-01-02 15:04", today+" "+theTime, loc)
		}

		// YY-MM-DD H:mm
		if err != nil {
			aTime = theTime[0:8] + " " + theTime[9:]
			t, err = time.ParseInLocation("06-01-02 3:04", aTime, loc)
		}

		// YY-MM-DD h:mmpm
		if err != nil {
			aTime = theTime[0:8] + " " + theTime[9:]
			t, err = time.ParseInLocation("06-01-02 3:04pm", aTime, loc)
		}

		// YY-MM-DD h:mmPM
		if err != nil {
			aTime = theTime[0:8] + " " + theTime[9:]
			t, err = time.ParseInLocation("06-01-02 3:04PM", aTime, loc)
		}

		// YY-MM-DD HH:mm
		if err != nil {
			aTime = theTime[0:8] + " " + theTime[9:]
			t, err = time.ParseInLocation("06-01-02 15:04", aTime, loc)
		}

		// YYYY-MM-DD H:mm
		if err != nil {
			aTime = theTime[0:10] + " " + theTime[11:]
			t, err = time.ParseInLocation("2006-01-02 3:04", aTime, loc)
		}

		// YYYY-MM-DD h:mmpm
		if err != nil {
			aTime = theTime[0:10] + " " + theTime[11:]
			t, err = time.ParseInLocation("2006-01-02 3:04pm", aTime, loc)
		}

		// YYYY-MM-DD h:mmPM
		if err != nil {
			aTime = theTime[0:10] + " " + theTime[11:]
			t, err = time.ParseInLocation("2006-01-02 3:04PM", aTime, loc)
		}

		// YYYY-MM-DD HH:mm
		if err != nil {
			aTime = theTime[0:10] + " " + theTime[11:]
			t, err = time.ParseInLocation("2006-01-02 15:04", aTime, loc)
		}
	}

	return t, err
}
