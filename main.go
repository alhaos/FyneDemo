package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"strconv"
	"time"
)

var (
	a fyne.App
	w fyne.Window
	l *widget.Label
	e *widget.Entry
	b *widget.Button
	r *rand.Rand
	n int
)

func main() {

	src := rand.NewSource(time.Now().UnixNano())
	r = rand.New(src)
	n = r.Intn(99) + 1

	a = app.New()
	w = a.NewWindow("Guess the number")

	w.Resize(fyne.NewSize(400, 300))
	w.SetFixedSize(true)

	l = widget.NewLabel(
		fmt.Sprintf("I made a number from 1 to 100 guess it %d", n),
	)

	l.Alignment = fyne.TextAlignCenter

	e = widget.NewEntry()

	b = widget.NewButton("Submit", ButtonPress)

	e.OnSubmitted = OnSubmitted

	w.SetContent(
		container.NewVBox(
			l,
			e,
			b,
		),
	)

	w.ShowAndRun()
}

func OnSubmitted(text string) {
	ButtonPress()
}

func ButtonPress() {

	b.SetText("Submit")
	en := e.Text
	e.SetText("")
	e.PlaceHolder = en

	enInt, err := strconv.Atoi(en)
	if err != nil {
		l.SetText("Invalid input")
		return
	}
	if enInt > n {
		l.SetText("this number is higher than expected")
		return
	}

	if enInt < n {
		l.SetText("this number is less than expected")
		return
	}

	if enInt == n {
		l.SetText("number guessed")
		b.SetText("Again")
		n = r.Intn(99) + 1
		return
	}
}
