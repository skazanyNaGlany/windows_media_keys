package main_window

const TITLE = "Windows Media Keys v%v"
const EXIT_BUTTON = "Exit"
const ABOUT_BUTTON = "About"
const WINDOW_WIDTH = 800
const WINDOW_HEIGHT = 600
const TOGGLE_AUTOSTART_BUTTON_ENABLE = "Enable autostart"
const TOGGLE_AUTOSTART_BUTTON_DISABLE = "Disable autostart"
const KILL_BUTTON = "Kill running instance"
const RUN_TEST_BUTTON = "Run (test)"
const ABOUT_TITLE = "Windows Media Keys"
const ABOUT_MESSAGE = `Emulate Windows media keys

Version %v

Esc-Up - VOLUME UP
Esc-Down - VOLUME DOWN
Esc-Left - PREVIOUS MEDIA
Esc-Right - NEXT MEDIA

Esc-F4 - SUSPEND COMPUTER
Esc-F7 - RESTART COMPUTER
Esc-F8 - SHUTDOWN COMPUTER

Esc-F11 - TOGGLE MUTE
Esc-F12 - TOGGLE PLAY

If the suspend feature is not working properly, you
may need to disable hibernation:

powercfg -hibernate off
`
const APP_VERSION = "0.2"
