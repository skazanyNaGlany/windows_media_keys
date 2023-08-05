package main_window

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type MainWindow struct {
	window                     fyne.Window
	vContainer                 *fyne.Container
	runOrTestButton            *widget.Button
	killButton                 *widget.Button
	aboutButton                *widget.Button
	exitButton                 *widget.Button
	toggleAutostartButton      *widget.Button
	toggleAutostartButtonState bool
}

func (mw *MainWindow) Init(app fyne.App) {
	mw.window = app.NewWindow(fmt.Sprintf(TITLE, APP_VERSION))
	mw.vContainer = container.NewVBox()
	mw.runOrTestButton = widget.NewButton(RUN_TEST_BUTTON, func() {})
	mw.killButton = widget.NewButton(KILL_BUTTON, func() {})
	mw.aboutButton = widget.NewButton(ABOUT_BUTTON, func() {})
	mw.exitButton = widget.NewButton(EXIT_BUTTON, func() {})
	mw.toggleAutostartButton = widget.NewButton(TOGGLE_AUTOSTART_BUTTON_DISABLE, func() {})

	mw.vContainer.Add(widget.NewSeparator())
	mw.vContainer.Add(mw.runOrTestButton)
	mw.vContainer.Add(mw.killButton)
	mw.vContainer.Add(widget.NewSeparator())
	mw.vContainer.Add(mw.toggleAutostartButton)
	mw.vContainer.Add(mw.aboutButton)
	mw.vContainer.Add(mw.exitButton)

	mw.window.SetContent(mw.vContainer)
	mw.window.SetPadded(true)

	mw.window.Resize(fyne.NewSize(
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
	))

	mw.aboutButton.OnTapped = func() {
		dialog.ShowInformation(
			ABOUT_TITLE,
			fmt.Sprintf(
				strings.ReplaceAll(ABOUT_MESSAGE, `\n`, "\n"),
				APP_VERSION),
			mw.GetWindow())
	}
}

func (mw *MainWindow) GetWindow() fyne.Window {
	return mw.window
}

func (mw *MainWindow) GetRunOrTestButton() *widget.Button {
	return mw.runOrTestButton
}

func (mw *MainWindow) GetKillButton() *widget.Button {
	return mw.killButton
}

func (mw *MainWindow) SetToggleAutostartButtonState(enable bool) {
	if enable {
		mw.toggleAutostartButton.SetText(TOGGLE_AUTOSTART_BUTTON_ENABLE)
	} else {
		mw.toggleAutostartButton.SetText(TOGGLE_AUTOSTART_BUTTON_DISABLE)
	}

	mw.toggleAutostartButtonState = enable
}

func (mw *MainWindow) GetToggleAutostartButtonState() bool {
	return mw.toggleAutostartButtonState
}

func (mw *MainWindow) GetToggleAutostartButton() *widget.Button {
	return mw.toggleAutostartButton
}

func (mw *MainWindow) GetExitButton() *widget.Button {
	return mw.exitButton
}

func (mw *MainWindow) SetRunOrTestButtonState(enable bool) {
	if enable {
		mw.runOrTestButton.Enable()
	} else {
		mw.runOrTestButton.Disable()
	}
}

func (mw *MainWindow) SetKillButtonState(enable bool) {
	if enable {
		mw.killButton.Enable()
	} else {
		mw.killButton.Disable()
	}
}
