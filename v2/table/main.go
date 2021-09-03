package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"

	examples "github.com/micheam/fyne-examples/v2"
)

type account struct {
	id       string
	name     string
	phone    string
	birthday string
}

var (
	windowSize = fyne.NewSize(600, 300)

	headers  = []string{"id", "name", "phone", "birthday"}
	accounts = []account{
		{uuid.New().String(), "Robert L. Taylor", "706-767-8575", "December 25, 1970"},
		{uuid.New().String(), "Kento Agano", "936-347-9392", "May 22, 1943"},
		{uuid.New().String(), "Mahito Fujiwara", "574-282-4340", "July 27, 1983"},
		{uuid.New().String(), "Suzanne R. Gonzalezxxxxxxxxxxx", "603-736-6867", "June 12, 1994"},
		{uuid.New().String(), "Linda M. Bailey", "317-738-7776", "January 2, 1962"},
	}
)

func main() {
	a := app.New()
	a.Settings().SetTheme(new(examples.DefaultTheme))
	w := a.NewWindow("Account List")

	table := widget.NewTable(

		func() (int, int) {
			rows := len(accounts) + 1 // data rows with header
			cols := 4
			return rows, cols
		},

		// callback fn for Create each cell.
		func() fyne.CanvasObject {
			l := widget.NewLabel("placeholder")
			l.Wrapping = fyne.TextTruncate
			return l
		},

		// callback fn for Update each cell.
		// This may trigger on initial rendering process.
		// override result of second param in NewTable()
		func(id widget.TableCellID, c fyne.CanvasObject) {
			label := c.(*widget.Label)
			col, row := id.Col, id.Row
			if row == 0 { // Header Row
				label.Alignment = fyne.TextAlignCenter
				label.TextStyle = fyne.TextStyle{Bold: true}
				label.Text = headers[col]
				return
			}
			// Data row
			acc := accounts[row-1]
			var text string
			switch col {
			case 0:
				text = acc.id
			case 1:
				text = acc.name
			case 2:
				text = acc.phone
			case 3:
				text = acc.birthday
			default:
				text = "-"
			}
			label.SetText(text)
		})

	// NOTE: Set width for each columns...
	//
	// Columns for widget.Table is automatically determined from the template object
	// specified in CreateCell (second arg of function NewTable) by default.
	// Here, the size of each column is determined separately from a sample of the data.
	sample := accounts[0]
	table.SetColumnWidth(0, widget.NewLabel(sample.id).MinSize().Width)       // id
	table.SetColumnWidth(1, widget.NewLabel(sample.name).MinSize().Width)     // name
	table.SetColumnWidth(2, widget.NewLabel(sample.phone).MinSize().Width)    // phone
	table.SetColumnWidth(3, widget.NewLabel(sample.birthday).MinSize().Width) // birthday

	// App Start
	w.SetContent(container.NewMax(table))
	w.Resize(windowSize)
	w.ShowAndRun()
}
