# robotgo-cli

**robotgo-cli** is a cross-platform command-line tool that wraps the [robotgo](https://github.com/go-vgo/robotgo) library's functionality. It allows you to automate desktop tasks from the terminal such as controlling the mouse, keyboard, capturing screenshots, managing windows, handling clipboard operations, and more.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
  - [Mouse Commands](#mouse-commands)
  - [Keyboard Commands](#keyboard-commands)
  - [Screen Commands](#screen-commands)
  - [Window Commands](#window-commands)
  - [Event Commands](#event-commands)
  - [Clipboard Commands](#clipboard-commands)
  - [Process Commands](#process-commands)
- [Cross-Compilation](#cross-compilation)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

## Features

- **Mouse Commands:**  
  Move, click, scroll, and toggle mouse buttons.
- **Keyboard Commands:**  
  Type strings, tap keys (with optional modifiers), and toggle key states.
- **Screen Commands:**  
  Capture a portion of the screen or the entire screen, get pixel color, and retrieve screen size.
- **Window Commands:**  
  Activate a window by name or process ID, kill processes, and fetch the active window title.
- **Event Commands:**  
  Listen for specific key events or print low-level system events.
- **Clipboard Commands:**  
  Read from and write text to the clipboard.
- **Process Commands:**  
  List running processes (PID and name).

## Requirements

- **Go:** Version 1.11 or later (with module support)
- **Dependencies:**
  - [robotgo](https://github.com/go-vgo/robotgo)
  - [Cobra](https://github.com/spf13/cobra)
  - [gohook](https://github.com/robotn/gohook)

**Platform-specific requirements:**

- **Linux:**  
  You may need the following development packages:
  ```bash
  sudo apt-get install libx11-xcb-dev libx11-dev libxtst-dev
  ```

- **Windows:**  
  When cross-compiling on Linux, ensure you have a Windows C compiler like gcc-mingw-w64-x86-64.

## Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/sampagon/robotgo-cli.git
   cd robotgo-cli
   ```

2. **Initialize a Go Module:**
   ```bash
   go mod init robotgo-cli
   ```

3. **Install Dependencies:**
   ```bash
   go get github.com/go-vgo/robotgo
   go get github.com/spf13/cobra
   go get github.com/robotn/gohook
   ```

4. **Build the Tool:**
   ```bash
   go build -o robotgo-cli main.go
   ```

## Usage

After building, run the executable. You can see all commands by executing:

```bash
./robotgo-cli --help
```

### Mouse Commands

**Move the Mouse:**
```bash
./robotgo-cli mouse move --x 100 --y 200
```
Moves the mouse pointer to the coordinates (100, 200).

**Click the Mouse:**
```bash
./robotgo-cli mouse click --button left --double
```
Clicks the left mouse button (double click if --double is set).

**Scroll the Mouse:**
```bash
./robotgo-cli mouse scroll --direction up --steps 10
```
Scrolls up 10 steps. Change the --direction flag as needed.

**Toggle Mouse Button State:**
```bash
./robotgo-cli mouse toggle --button left --state down
```
Toggles the specified mouse button state to either down or up.

### Keyboard Commands

**Type a String:**
```bash
./robotgo-cli keyboard type --text "Hello, RobotGo!"
```
Simulates typing the provided string.

**Tap a Key (with Modifiers):**
```bash
./robotgo-cli keyboard tap --key a --mods ctrl,shift
```
Taps the a key with ctrl and shift modifiers. Use a comma-separated list for modifiers.

**Toggle a Key State:**
```bash
./robotgo-cli keyboard toggle --key a --state up
```
Toggles the key a to the specified state.

### Screen Commands

**Capture a Portion of the Screen:**
```bash
./robotgo-cli screen capture --x 10 --y 10 --width 300 --height 300 --output shot.png
```
Captures a 300x300 region starting at (10, 10) and saves it as shot.png.

**Capture the Full Screen:**

The tool supports a --full flag to capture the entire screen. When this flag is used, the provided coordinates and dimensions are ignored.

```bash
./robotgo-cli screen capture --full --output full_shot.png
```

**Get Pixel Color:**
```bash
./robotgo-cli screen getpixel --x 150 --y 150
```
Returns the color of the pixel at (150, 150).

**Get Screen Size:**
```bash
./robotgo-cli screen size
```
Prints the current screen width and height.

### Window Commands

**Activate a Window:**

Activate by name:
```bash
./robotgo-cli window activate --name chrome
```

Or by process ID:
```bash
./robotgo-cli window activate --pid 12345
```

**Kill a Window Process:**
```bash
./robotgo-cli window kill --pid 12345
```
Terminates the process with the given PID.

**Get Active Window Title:**
```bash
./robotgo-cli window title
```
Prints the title of the currently active window.

### Event Commands

**Listen for a Specific Key Combination:**
```bash
./robotgo-cli event listen --keys ctrl,shift,q
```
Listens for the ctrl+shift+q combination and prints event details when triggered (listener stops after the first event).

**Print Low-Level Events:**
```bash
./robotgo-cli event low
```
Continuously prints low-level events until interrupted (e.g., with Ctrl+C).

### Clipboard Commands

**Read from Clipboard:**
```bash
./robotgo-cli clipboard read
```
Prints the current contents of the clipboard.

**Write to Clipboard:**
```bash
./robotgo-cli clipboard write --text "Clipboard text"
```
Writes the provided text to the clipboard.

### Process Commands

**List Running Processes:**
```bash
./robotgo-cli process list
```
Displays a list of running processes along with their PID and name.

## Cross-Compilation

### Building for Windows on Linux

1. **Install the Windows Cross-Compiler:**
   ```bash
   sudo apt-get install gcc-mingw-w64-x86-64
   ```

2. **Set the CC Environment Variable:**
   ```bash
   export CC=x86_64-w64-mingw32-gcc
   ```

3. **Build for Windows:**
   ```bash
   GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o robotgo-cli.exe main.go
   ```
   This will produce robotgo-cli.exe, which can be run on Windows.

## Contributing

Contributions are welcome! If you'd like to add new features, improve documentation, or fix bugs, please fork the repository and submit a pull request. For major changes, open an issue first to discuss your ideas.

## Acknowledgements

- [robotgo](https://github.com/go-vgo/robotgo): Provides the underlying automation library.
- [Cobra](https://github.com/spf13/cobra): The CLI framework used in this project.
- [gohook](https://github.com/robotn/gohook): Used for low-level event handling.
