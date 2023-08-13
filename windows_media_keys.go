package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2/app"
	"github.com/skazanyNaGlany/windows_media_keys/windows/main_window"
	"github.com/sqweek/dialog"
)

func main() {
	checkPlatform()
	changeCwd()

	if isRunCommand() {
		run()
	} else {
		userInterface()
	}
}

func checkPlatform() {
	if runtime.GOOS != "windows" {
		msg := "This program can be run only on Windows."

		LogDialog{}.Panicln(APP_NAME, msg)
	}
}

func changeCwd() {
	os.Chdir(
		filepath.Dir(os.Args[0]))
}

func isRunCommand() bool {
	return len(os.Args) == 2 && os.Args[1] == RUN_COMMAND
}

func sleepCallback() bool {
	return dialog.Message("%s", "Do you want to suspend the computer ?").Title(APP_NAME).YesNo()
}

func restartCallback() bool {
	return dialog.Message("%s", "Do you want to restart the computer ?").Title(APP_NAME).YesNo()
}

func shutdownCallback() bool {
	return dialog.Message("%s", "Do you want to shutdown the computer ?").Title(APP_NAME).YesNo()
}

func run() {
	emulator := MediaKeysEmulator{}
	emulator.Init()

	emulator.sleepHook = sleepCallback
	emulator.restartHook = restartCallback
	emulator.shutdownHook = shutdownCallback

	emulator.EmulateInLoop()
}

func userInterface() {
	app := app.New()

	installer := AutostartInstaller{}
	installer.Init()

	instance := Instance{}
	instance.Init()

	mainWindow := main_window.MainWindow{}
	mainWindow.Init(app)

	running := instance.IsRunnnig()

	mainWindow.SetRunOrTestButtonState(!running)
	mainWindow.SetKillButtonState(running)

	mainWindow.SetToggleAutostartButtonState(
		!installer.IsAutostartEnabled())

	// if installer.IsAutostartEnabled() {
	// }

	mainWindow.GetRunOrTestButton().OnTapped = func() {
		if instance.IsRunnnig() {
			LogDialog{}.Panicln(APP_NAME, "Already running.")
		}

		exec.Command(os.Args[0], RUN_COMMAND).Start()

		mainWindow.SetRunOrTestButtonState(false)
		mainWindow.SetKillButtonState(true)
	}

	mainWindow.GetKillButton().OnTapped = func() {
		instance.IsRunnnig()
		instance.Kill()

		mainWindow.SetRunOrTestButtonState(true)
		mainWindow.SetKillButtonState(false)
	}

	mainWindow.GetExitButton().OnTapped = func() {
		app.Quit()
	}

	mainWindow.GetToggleAutostartButton().OnTapped = func() {
		if err := installer.EnableAutostart(
			!installer.IsAutostartEnabled()); err != nil {
			LogDialog{}.Panicln(APP_NAME, err.Error())
		}

		mainWindow.SetToggleAutostartButtonState(!installer.IsAutostartEnabled())

		// if installer.IsAutostartEnabled() {
		// } else {
		// }
	}

	mainWindow.GetWindow().SetOnClosed(func() {
	})

	mainWindow.GetWindow().SetMaster()
	mainWindow.GetWindow().CenterOnScreen()
	mainWindow.GetWindow().ShowAndRun()
}
