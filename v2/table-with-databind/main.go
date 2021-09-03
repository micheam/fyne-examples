package main

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/brianvoe/gofakeit/v6"

	examples "github.com/micheam/fyne-examples/v2"
)

type account struct {
	ID      string `fake:"{uuid}"`
	Name    string `fake:"{name}"`
	Country string `fake:"{country}"`
	Phone   string `fake:"{phone}"`
}

var (
	windowSize = fyne.NewSize(600, 300)

	headers  = []string{"id", "name", "country", "phone"}
	datasize = 100
	accounts = []binding.DataMap{}
)

func init() {
	gofakeit.Seed(time.Now().Unix())
	for i := 0; i < datasize; i++ {
		acc := new(account)
		if err := gofakeit.Struct(acc); err != nil {
			fyne.LogError("failed to generate fake data", err)
			panic(err)
		}
		accounts = append(accounts, binding.BindStruct(acc))
	}
}

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
			col, row := id.Col, id.Row
			if row == 0 { // Header Row
				label := c.(*widget.Label)
				label.Alignment = fyne.TextAlignCenter
				label.TextStyle = fyne.TextStyle{Bold: true}
				label.Text = headers[col]
				return
			}
			// Data row
			acc := accounts[row-1]
			switch col {
			case 0:
				item, _ := acc.GetItem("ID")
				c.(*widget.Label).Bind(item.(binding.String))
			case 1:
				item, _ := acc.GetItem("Name")
				c.(*widget.Label).Bind(item.(binding.String))
			case 2:
				item, _ := acc.GetItem("Country")
				c.(*widget.Label).Bind(item.(binding.String))
			case 3:
				item, _ := acc.GetItem("Phone")
				c.(*widget.Label).Bind(item.(binding.String))
			default:
				c.(*widget.Label).SetText("-")
			}
		})

	// NOTE: Set width for each columns...
	//
	// Columns for widget.Table is automatically determined from the template object
	// specified in CreateCell (second arg of function NewTable) by default.
	// Here, the size of each column is determined separately from a sample of the data.
	sample := accounts[0]
	fmt.Printf("%+v\n", sample.Keys())
	table.SetColumnWidth(0, widget.NewLabel(mustStringVal(sample, "ID")).MinSize().Width)
	table.SetColumnWidth(1, widget.NewLabel(mustStringVal(sample, "Name")).MinSize().Width)
	table.SetColumnWidth(2, widget.NewLabel(mustStringVal(sample, "Country")).MinSize().Width)
	table.SetColumnWidth(3, widget.NewLabel(mustStringVal(sample, "Phone")).MinSize().Width)

	// data controll
	var ctrl *fyne.Container
	{
		shuffle := widget.NewButtonWithIcon("Shuffle", theme.ViewRefreshIcon(), func() {
			rand.Shuffle(len(accounts), func(i, j int) {
				accounts[i], accounts[j] = accounts[j], accounts[i]
			})
			table.Refresh()
		})
		shuffle.Importance = widget.HighImportance
		ctrl = container.NewHBox(
			widget.NewLabel(fmt.Sprintf("data count: %d", datasize)),
			layout.NewSpacer(),
			shuffle)
	}

	// App Start
	w.SetContent(container.New(
		layout.NewBorderLayout(nil, ctrl, nil, nil),
		ctrl,
		table,
	))
	w.Resize(windowSize)
	w.ShowAndRun()
}

// ----------------------------
// misc

const dateFormat = "2006-01-02"

func mustParseTime(s string) time.Time {
	t, err := time.Parse(dateFormat, s)
	if err != nil {
		panic(err)
	}
	return t
}

func mustStringVal(m binding.DataMap, key string) string {
	bs := mustGetItem(m, key)
	v, err := bs.Get()
	if err != nil {
		panic(err)
	}
	return v
}

func mustGetItem(m binding.DataMap, key string) binding.String {
	item, err := m.GetItem(key)
	if err != nil {
		panic(err)
	}
	return item.(binding.String)
}
