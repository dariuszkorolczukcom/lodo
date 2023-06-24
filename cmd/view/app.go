package view

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/dariuszkorolczukcom/lodo/internal/file"
)

var w fyne.Window

func Create(checkFiles func(string) []file.File) {
	a := app.New()
	w = a.NewWindow("LODO")
	w.Resize(fyne.NewSize(800, 600))

	//content := widget.NewButtonWithIcon("Check", theme.HomeIcon(), func() {
	//})

	func() {
		fmt.Println("button tapped")
		onChosen := func(f fyne.ListableURI, err error) {
			if err != nil {
				fmt.Println(err)
				return
			}
			if f == nil {
				return
			}
			fmt.Printf("chosen: %v\n\n", f.String())

			button1 := widget.NewButton("Search",
				func() {

					files := checkFiles(f.String())
					createTable(files)
				})
			button1.Resize(fyne.NewSize(100, 20))

			label1 := widget.NewLabel(f.String())
			label1.Resize(fyne.NewSize(100, 20))

			content1 := container.NewVBox(
				label1,
				button1,
			)

			w.SetContent(content1)
		}
		dialog.ShowFolderOpen(onChosen, w)
	}()
	w.SetContent(widget.NewLabel("L.O.D.O."))
	w.ShowAndRun()
}

func createTable(files []file.File) {
	progress := createProgressBar()
	w.SetContent(container.NewVBox(progress))

	var data = [][]string{{"file name", "modified"}}
	for _, file := range files {
		data = append(data, []string{file.Name, file.Modified.Format(time.Stamp)})
	}
	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			item.Resize(fyne.Size{
				Width:  400,
				Height: 20,
			})
			return item
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
			o.(*widget.Label).Resize(fyne.Size{
				Width:  400,
				Height: 20,
			})
		})
	table.Resize(fyne.NewSize(800, 500))
	content := container.NewVBox(
		widget.NewLabel("Files"),
		table,
	)
	content.Resize(fyne.NewSize(800, 520))

	w.SetContent(content)
	// w.SetContent(table)

}

func createProgressBar() *widget.ProgressBar {

	progress := widget.NewProgressBar()
	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 400)
			progress.SetValue(i)
		}
	}()
	return progress
}
