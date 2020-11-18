package software

import (
	"testing"

	"fyne.io/fyne/container"
	"fyne.io/fyne/test"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func TestRender(t *testing.T) {
	obj := widget.NewLabel("Hi")
	test.AssertImageMatches(t, "label_dark.png", Render(obj, theme.DarkTheme()))
	test.AssertImageMatches(t, "label_light.png", Render(obj, theme.LightTheme()))
}

func TestRender_State(t *testing.T) {
	obj := widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {})
	test.AssertImageMatches(t, "button.png", Render(obj, theme.DarkTheme()))

	obj.Importance = widget.HighImportance
	obj.Refresh()
	test.AssertImageMatches(t, "button_important.png", Render(obj, theme.DarkTheme()))
}

func TestRender_Focus(t *testing.T) {
	obj := widget.NewEntry()
	test.AssertImageMatches(t, "entry.png", Render(obj, theme.DarkTheme()))

	test.Tap(obj)
	test.AssertImageMatches(t, "entry_focus.png", Render(obj, theme.DarkTheme()))
}

func TestRenderCanvas(t *testing.T) {
	obj := container.NewAppTabs(
		container.NewTabItem("Tab 1", container.NewVBox(
			widget.NewLabel("Label"),
			widget.NewButton("Button", func() {}),
		)))

	c := NewCanvas()
	c.SetContent(obj)
	test.AssertImageMatches(t, "canvas.png", RenderCanvas(c, theme.LightTheme()))
}
