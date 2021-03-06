package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	a1, a2, a3, a4 string //user's answers for each of 4 (N) questions
	f, rep         *os.File
	// f is a unique temporary file that stores the current state machine's condition
	// rep is a unique report file that gathers user's answers
	inner, inrep int
)

//infile is the function to be used as state machine
//unique temporary file will be created for each user to store the current state machine condition
func infile(fname, state string) bool {

	cont, _ := ioutil.ReadFile(fname)

	if strings.Contains(string(cont), state) {

		return true

	} else {

		return false
	}

}

func main() {

	b, err := tb.NewBot(tb.Settings{

		URL: "https://api.telegram.org",

		Token:  "put token here",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {

		if m.Text == "/start" {

			b.Send(m.Sender, "Question 1")

			//the name of the file is same as user's telegram id
			f, _ = os.Create(strconv.Itoa(m.Sender.ID))
			inner, _ = f.WriteString("q2")

			a1 = m.Text

			f.Close()

			//the name of the report file is same as user's telegram id with "rep" suffix
			rep, _ = os.Create(strconv.Itoa(m.Sender.ID) + "rep")

			rep.Close()
		} else if infile(strconv.Itoa(m.Sender.ID), "q2") && m.Text != a1 {
			b.Send(m.Sender, "Question 2")

			f, _ = os.Create(strconv.Itoa(m.Sender.ID))
			inner, _ = f.WriteString("q3")

			a2 = m.Text

			f.Close()

			rep, _ = os.OpenFile(strconv.Itoa(m.Sender.ID)+"rep", os.O_APPEND, 0644)
			inrep, _ = rep.WriteString(m.Text + "\n")

			rep.Close()

		} else if infile(strconv.Itoa(m.Sender.ID), "q3") && m.Text != a2 {
			b.Send(m.Sender, "Question 3")

			f, _ = os.Create(strconv.Itoa(m.Sender.ID))
			inner, _ = f.WriteString("q4")

			a3 = m.Text

			f.Close()

			rep, _ = os.OpenFile(strconv.Itoa(m.Sender.ID)+"rep", os.O_APPEND, 0644)
			inrep, _ = rep.WriteString(m.Text + "\n")

			rep.Close()

		} else if infile(strconv.Itoa(m.Sender.ID), "q4") && m.Text != a3 {
			b.Send(m.Sender, "Question 4")

			f, _ = os.Create(strconv.Itoa(m.Sender.ID))
			inner, _ = f.WriteString("done")

			a4 = m.Text

			f.Close()

			rep, _ = os.OpenFile(strconv.Itoa(m.Sender.ID)+"rep", os.O_APPEND, 0644)
			inrep, _ = rep.WriteString(m.Text + "\n")

			rep.Close()

		} else if infile(strconv.Itoa(m.Sender.ID), "done") && m.Text != a4 {
			b.Send(m.Sender, "Done")

			e := os.Remove(strconv.Itoa(m.Sender.ID))
			fmt.Println(e)

			f.Close()

			rep, _ = os.OpenFile(strconv.Itoa(m.Sender.ID)+"rep", os.O_APPEND, 0644)
			inrep, _ = rep.WriteString(m.Text + "\n")

			rep.Close()

		}

	})

	b.Start()
}
