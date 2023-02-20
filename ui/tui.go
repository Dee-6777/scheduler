package ui

import (
	"github.com/Dee-6777/scheduler/cmd"
	//"github.com/stjohnjohnson/gifview"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var table = tview.NewTable().SetFixed(1, 1).SetSelectable(true, false)
var pages = tview.NewPages()
var jokeText = tview.NewTextView()
var app = tview.NewApplication()
var form = tview.NewForm()

//var todoList = tview.NewList().ShowSecondaryText(false)

var flex = tview.NewFlex()
var flex1 = tview.NewFlex()

//var flex2 = tview.NewFlex()

var text = tview.NewTextView().
	SetTextColor(tcell.ColorPaleVioletRed).
	SetText("(a) YOUR SCHEDULE \n(b) EDIT SCHEDULE \n(q) QUIT")

var text1 = tview.NewTextView().
	SetTextColor(tcell.ColorPaleVioletRed).
	SetText("(a) DELETE SCHEDULE \n(b) ADD SCHEDULE \n(c) UPDATE SCHEDULE \n(q) RETURN")

func Greet() {
	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(jokeText, 0, 4, false), 0, 6, false).
		AddItem(text, 0, 1, false)

	flex.SetBorder(true).SetTitle("S[red]C[yellow]H[green]E[darkcyan]D[blue]U[darkmagenta]L[red]E[white]R[red]").SetTitleAlign(tview.AlignCenter)
	// Adding conditions for different options
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 97 {
			ScheduleTable()
			pages.SwitchToPage("Your Schedule")
		} else if event.Rune() == 98 {
			form.Clear(true)
			Greet_()
			pages.SwitchToPage("Edit Schedule")
		} else if event.Rune() == 113 {
			app.Stop()
		}
		return event
	})
	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Your Schedule", flex, true, false)
	pages.AddPage("Edit Schedule", flex1, true, false)
	pages.AddPage("Delete Schedule", form, true, false)
	pages.AddPage("Add Schedule", form, true, false)
	pages.AddPage("Update Schedule", form, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func ScheduleTable() {
	pages.SwitchToPage("Your Schedule")
	table.Clear()

}

func Greet_() {
	pages.SwitchToPage("Edit Schedule")
	flex1.Clear()
	flex1.SetDirection(tview.FlexRow). // Adjusting jokelist & jokeText
						AddItem(text1, 0, 1, false)
	flex1.SetBorder(true).SetTitle("E[red]D[yellow]I[green]T[darkcyan]O[blue]R[darkmagenta]").SetTitleAlign(tview.AlignCenter)
	flex1.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 97 {
			form.Clear(true)
			DeleteForm()
			pages.SwitchToPage("Delete Schedule")
		} else if event.Rune() == 98 {
			form.Clear(true)
			AddForm()
			pages.SwitchToPage("Add Schedule")
		} else if event.Rune() == 99 {
			form.Clear(true)
			UpdateForm()
			pages.SwitchToPage("Update Schedule")
		} else if event.Rune() == 113 {
			pages.SwitchToPage("Menu")
		}
		return event
	})
}

func AddForm() *tview.Form {

	user_ := cmd.Todo{}

	form.SetBackgroundColor(tcell.ColorPeachPuff)

	form.AddInputField("ID", "", 20, nil, func(ID string) {
		user_.ID = ID
	})

	form.AddInputField("To Do", "", 20, nil, func(Title string) {
		user_.Item = Title
	})

	form.AddInputField("Time", "", 20, nil, func(Time string) {
		user_.Time = Time
	})

	form.AddCheckbox("Completed", false, func(Completed bool) {
		var ans string
		if Completed {
			ans = "false"
		} else {
			ans = "true"
		}
		user_.Completed = ans
	})

	form.AddButton("save", func() {
		arg := []string{}
		arg = append(arg, user_.ID, user_.Item, user_.Time, user_.Completed)
		cmd.AddTodo(arg)
		pages.SwitchToPage("Edit Schedule")
	})

	form.AddButton("cancel", func() {
		pages.SwitchToPage("Edit Schedule")
	})
	form.SetBorder(true).SetBorderColor(tcell.ColorDarkRed).SetTitle("E[red]D[yellow]I[green]T[darkcyan]O[blue]R[darkmagenta]").SetTitleAlign(tview.AlignCenter)
	return form
}

func DeleteForm() *tview.Form {
	user_ := cmd.Todo{}

	form.SetBackgroundColor(tcell.ColorAntiqueWhite)

	form.AddInputField("ID", "", 20, nil, func(ID string) {
		user_.ID = ID
	})

	form.AddButton("DELETE", func() {
		arg := []string{} // Declared a slice and initialised it with 3 items
		arg = append(arg, user_.ID)
		cmd.DelTodo(arg)
		pages.SwitchToPage("Edit Schedule")
	})
	form.AddButton("CANCEL", func() {
		pages.SwitchToPage("Edit Schedule")
	})
	form.SetBorder(true).SetBorderColor(tcell.ColorDarkRed).SetTitle("E[red]D[yellow]I[green]T[white]O[purple]R[blue]").SetTitleAlign(tview.AlignCenter)
	return form
}

func UpdateForm() *tview.Form {

	user_ := cmd.Todo{}

	//form.SetBackgroundColor(tcell.ColorOrangeRed)

	form.AddInputField("ID", "", 20, nil, func(ID string) {
		user_.ID = ID
	})

	form.AddInputField("To Do", "", 20, nil, func(Title string) {
		user_.Item = Title
	})

	form.AddInputField("Time", "", 20, nil, func(Time string) {
		user_.Time = Time
	})

	form.AddCheckbox("Completed", false, func(Completed bool) {
		var ans string
		if Completed {
			ans = "false"
		} else {
			ans = "true"
		}
		user_.Completed = ans
	})

	form.AddButton("save", func() {
		arg := []string{}
		arg = append(arg, user_.ID, user_.Item, user_.Time, user_.Completed)
		cmd.Update(arg)
		pages.SwitchToPage("Edit Schedule")
	})

	form.AddButton("cancel", func() {
		pages.SwitchToPage("Edit Schedule")
	})
	form.SetBorder(true).SetBorderColor(tcell.ColorDarkRed).SetTitle("E[red]D[yellow]I[green]T[darkcyan]O[blue]R[darkmagenta]").SetTitleAlign(tview.AlignCenter)
	return form
}
