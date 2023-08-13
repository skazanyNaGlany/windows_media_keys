package main

import (
	"os/exec"
	"time"

	"golang.org/x/sys/windows"
)

const VK_ESCAPE = 0x1B
const VK_DOWN = 0x28
const VK_UP = 0x26
const VK_RIGHT = 0x27
const VK_LEFT = 0x25
const VK_F4 = 0x73
const VK_F7 = 0x76
const VK_F8 = 0x77
const VK_F11 = 0x7A
const VK_F12 = 0x7B
const VK_VOLUME_DOWN = 0xAE
const VK_VOLUME_UP = 0xAF
const VK_MEDIA_NEXT_TRACK = 0xB0
const VK_MEDIA_PREV_TRACK = 0xB1
const VK_VOLUME_MUTE = 0xAD
const VK_MEDIA_PLAY_PAUSE = 0xB3
const VK_SLEEP = 0x5F
const VK_OEM_RESET = 0xE9

type OnActionHook func() bool

type MediaKeysEmulator struct {
	modUser32            *windows.LazyDLL
	procGetAsyncKeyState *windows.LazyProc
	procKeybdEvent       *windows.LazyProc
	volumeDownHook       OnActionHook
	volumeUpHook         OnActionHook
	mediaNextHook        OnActionHook
	mediaPrevHook        OnActionHook
	sleepHook            OnActionHook
	restartHook          OnActionHook
	shutdownHook         OnActionHook
	muteHook             OnActionHook
	playHook             OnActionHook
}

func (mke *MediaKeysEmulator) Init() {
	mke.modUser32 = windows.NewLazyDLL("user32.dll")
	mke.procGetAsyncKeyState = mke.modUser32.NewProc("GetAsyncKeyState")
	mke.procKeybdEvent = mke.modUser32.NewProc("keybd_event")
}

func (mke *MediaKeysEmulator) EmulateInLoop() {
	for {
		time.Sleep(100 * time.Millisecond)

		if escapePressed := mke.callGetAsyncKeyState(VK_ESCAPE); escapePressed == 0 {
			// ESCAPE not pressed
			continue
		}

		downPressed := mke.callGetAsyncKeyState(VK_DOWN)
		upPressed := mke.callGetAsyncKeyState(VK_UP)
		rightPressed := mke.callGetAsyncKeyState(VK_RIGHT)
		leftPressed := mke.callGetAsyncKeyState(VK_LEFT)
		f4Pressed := mke.callGetAsyncKeyState(VK_F4)
		f7Pressed := mke.callGetAsyncKeyState(VK_F7)
		f8Pressed := mke.callGetAsyncKeyState(VK_F8)
		f11Pressed := mke.callGetAsyncKeyState(VK_F11)
		f12Pressed := mke.callGetAsyncKeyState(VK_F12)

		if downPressed != 0 {
			// volume down
			if mke.volumeDownHook != nil && !mke.volumeDownHook() {
				continue
			}

			mke.procKeybdEvent.Call(
				uintptr(VK_VOLUME_DOWN),
				uintptr(0),
				uintptr(0),
				uintptr(0))
		} else if upPressed != 0 {
			// volume up
			if mke.volumeUpHook != nil && !mke.volumeUpHook() {
				continue
			}

			mke.procKeybdEvent.Call(
				uintptr(VK_VOLUME_UP),
				uintptr(0),
				uintptr(0),
				uintptr(0))
		} else if rightPressed != 0 {
			// media next
			if mke.mediaNextHook != nil && !mke.mediaNextHook() {
				continue
			}

			mke.procKeybdEvent.Call(
				uintptr(VK_MEDIA_NEXT_TRACK),
				uintptr(0),
				uintptr(0),
				uintptr(0))
		} else if leftPressed != 0 {
			// media prev
			if mke.mediaPrevHook != nil && !mke.mediaPrevHook() {
				continue
			}

			mke.procKeybdEvent.Call(
				uintptr(VK_MEDIA_PREV_TRACK),
				uintptr(0),
				uintptr(0),
				uintptr(0))
		} else if f4Pressed != 0 {
			// sleep
			if mke.sleepHook != nil && !mke.sleepHook() {
				continue
			}

			exec.Command("rundll32.exe", "powrprof.dll,SetSuspendState", "0,1,0").Start()
		} else if f7Pressed != 0 {
			// restart
			if mke.restartHook != nil && !mke.restartHook() {
				continue
			}

			exec.Command("shutdown", "/r", "/t", "0").Start()
		} else if f8Pressed != 0 {
			// shutdown
			if mke.shutdownHook != nil && !mke.shutdownHook() {
				continue
			}

			exec.Command("shutdown", "/s", "/t", "0").Start()
		} else if f11Pressed != 0 {
			// mute/unmute
			if mke.muteHook != nil && !mke.muteHook() {
				continue
			}

			mke.procKeybdEvent.Call(
				uintptr(VK_VOLUME_MUTE),
				uintptr(0),
				uintptr(0),
				uintptr(0))
		} else if f12Pressed != 0 {
			// play/pause
			if mke.playHook != nil && !mke.playHook() {
				continue
			}

			mke.procKeybdEvent.Call(
				uintptr(VK_MEDIA_PLAY_PAUSE),
				uintptr(0),
				uintptr(0),
				uintptr(0))
		}
	}
}

func (mke *MediaKeysEmulator) callGetAsyncKeyState(key int) uintptr {
	state, _, _ := mke.procGetAsyncKeyState.Call(uintptr(key))

	if state&0x1 != 0 {
		// clear the least significant bit
		// as it's written in the documentation:
		// if the least significant bit is set, the key was pressed
		// after the previous call to GetAsyncKeyState.
		state ^= 0x1
	}

	return state
}
