package main

import (
	"time"

	"golang.org/x/sys/windows"
)

const VK_ESCAPE = 0x1B
const VK_DOWN = 0x28
const VK_UP = 0x26
const VK_RIGHT = 0x27
const VK_LEFT = 0x25
const VK_F11 = 0x7A
const VK_F12 = 0x7B
const VK_VOLUME_DOWN = 0xAE
const VK_VOLUME_UP = 0xAF
const VK_MEDIA_NEXT_TRACK = 0xB0
const VK_MEDIA_PREV_TRACK = 0xB1
const VK_VOLUME_MUTE = 0xAD
const VK_MEDIA_PLAY_PAUSE = 0xB3

type MediaKeysEmulator struct {
	modUser32            *windows.LazyDLL
	procGetAsyncKeyState *windows.LazyProc
	procKeybdEvent       *windows.LazyProc
}

func (mke *MediaKeysEmulator) Init() {
	mke.modUser32 = windows.NewLazyDLL("user32.dll")
	mke.procGetAsyncKeyState = mke.modUser32.NewProc("GetAsyncKeyState")
	mke.procKeybdEvent = mke.modUser32.NewProc("keybd_event")
}

func (mke *MediaKeysEmulator) EmulateInLoop() {
	for {
		time.Sleep(100 * time.Millisecond)

		if escapePressed := mke.CallGetAsyncKeyState(VK_ESCAPE); escapePressed == 0 {
			// ESCAPE not pressed
			continue
		}

		downPressed := mke.CallGetAsyncKeyState(VK_DOWN)
		upPressed := mke.CallGetAsyncKeyState(VK_UP)
		rightPressed := mke.CallGetAsyncKeyState(VK_RIGHT)
		leftPressed := mke.CallGetAsyncKeyState(VK_LEFT)
		f11Pressed := mke.CallGetAsyncKeyState(VK_F11)
		f12Pressed := mke.CallGetAsyncKeyState(VK_F12)

		if downPressed != 0 {
			// volume down
			mke.procKeybdEvent.Call(
				uintptr(VK_VOLUME_DOWN),
				uintptr(0),
				uintptr(0),
				uintptr(0))
		} else if upPressed != 0 {
			// volume up
			mke.procKeybdEvent.Call(
				uintptr(VK_VOLUME_UP),
				uintptr(0),
				uintptr(0),
				uintptr(0))
		} else if rightPressed != 0 {
			// media next
			mke.procKeybdEvent.Call(
				uintptr(VK_MEDIA_NEXT_TRACK),
				uintptr(0),
				uintptr(0),
				uintptr(0))
		} else if leftPressed != 0 {
			// media prev
			mke.procKeybdEvent.Call(
				uintptr(VK_MEDIA_PREV_TRACK),
				uintptr(0),
				uintptr(0),
				uintptr(0))
		} else if f11Pressed != 0 {
			// mute/unmute
			mke.procKeybdEvent.Call(
				uintptr(VK_VOLUME_MUTE),
				uintptr(0),
				uintptr(0),
				uintptr(0))
		} else if f12Pressed != 0 {
			// play/pause
			mke.procKeybdEvent.Call(
				uintptr(VK_MEDIA_PLAY_PAUSE),
				uintptr(0),
				uintptr(0),
				uintptr(0))
		}
	}
}

func (mke *MediaKeysEmulator) CallGetAsyncKeyState(key int) uintptr {
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
