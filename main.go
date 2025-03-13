package main

import (
	"fmt"
	"os"
	"strings"
	"image/png"
	"bytes"
	"encoding/base64"


	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "robotgo-cli",
		Short: "A CLI wrapper for robotgo functionality",
	}

	// ======================
	// Mouse Commands Group
	// ======================
	var mouseCmd = &cobra.Command{
		Use:   "mouse",
		Short: "Mouse related commands",
	}

	// Mouse move command
	var mouseMoveCmd = &cobra.Command{
		Use:   "move",
		Short: "Move the mouse to specified coordinates",
		Run: func(cmd *cobra.Command, args []string) {
			x, _ := cmd.Flags().GetInt("x")
			y, _ := cmd.Flags().GetInt("y")
			robotgo.Move(x, y)
			fmt.Printf("Moved mouse to (%d, %d)\n", x, y)
		},
	}
	mouseMoveCmd.Flags().Int("x", 0, "X coordinate")
	mouseMoveCmd.Flags().Int("y", 0, "Y coordinate")

	// Mouse click command
	var mouseClickCmd = &cobra.Command{
		Use:   "click",
		Short: "Click a mouse button",
		Run: func(cmd *cobra.Command, args []string) {
			button, _ := cmd.Flags().GetString("button")
			double, _ := cmd.Flags().GetBool("double")
			robotgo.Click(button, double)
			fmt.Printf("Clicked %s button (double=%v)\n", button, double)
		},
	}
	mouseClickCmd.Flags().String("button", "left", "Mouse button (left, right, wheelLeft, wheelRight, etc.)")
	mouseClickCmd.Flags().Bool("double", false, "Double click flag")

	// Mouse scroll command
	var mouseScrollCmd = &cobra.Command{
		Use:   "scroll",
		Short: "Scroll the mouse in a direction",
		Run: func(cmd *cobra.Command, args []string) {
			direction, _ := cmd.Flags().GetString("direction")
			steps, _ := cmd.Flags().GetInt("steps")
			robotgo.ScrollDir(steps, direction)
			fmt.Printf("Scrolled %s for %d steps\n", direction, steps)
		},
	}
	mouseScrollCmd.Flags().String("direction", "up", "Scroll direction: up, down, left, right")
	mouseScrollCmd.Flags().Int("steps", 10, "Number of steps to scroll")

	// Mouse toggle command
	var mouseToggleCmd = &cobra.Command{
		Use:   "toggle",
		Short: "Toggle mouse button state (down/up)",
		Run: func(cmd *cobra.Command, args []string) {
			button, _ := cmd.Flags().GetString("button")
			state, _ := cmd.Flags().GetString("state")
			err := robotgo.Toggle(button, state)
			if err != nil {
				fmt.Printf("Error toggling mouse button: %v\n", err)
			} else {
				fmt.Printf("Toggled %s button to state %s\n", button, state)
			}
		},
	}
	mouseToggleCmd.Flags().String("button", "left", "Mouse button to toggle")
	mouseToggleCmd.Flags().String("state", "down", "State: down or up")

	mouseCmd.AddCommand(mouseMoveCmd, mouseClickCmd, mouseScrollCmd, mouseToggleCmd)
	rootCmd.AddCommand(mouseCmd)

	// ===========================
	// Keyboard Commands Group
	// ===========================
	var keyboardCmd = &cobra.Command{
		Use:   "keyboard",
		Short: "Keyboard related commands",
	}

	// Keyboard type command
	var keyboardTypeCmd = &cobra.Command{
		Use:   "type",
		Short: "Type a string using the keyboard",
		Run: func(cmd *cobra.Command, args []string) {
			text, _ := cmd.Flags().GetString("text")
			robotgo.TypeStr(text)
			fmt.Printf("Typed string: %s\n", text)
		},
	}
	keyboardTypeCmd.Flags().String("text", "", "Text to type")
	keyboardTypeCmd.MarkFlagRequired("text")

	// Keyboard tap command
	var keyboardTapCmd = &cobra.Command{
		Use:   "tap",
		Short: "Tap a key with optional modifiers",
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cmd.Flags().GetString("key")
			mods, _ := cmd.Flags().GetString("mods")
			var modArray []string
			if mods != "" {
				modArray = strings.Split(mods, ",")
			}
			err := robotgo.KeyTap(key, modArray)
			if err != nil {
				fmt.Printf("Error tapping key: %v\n", err)
			} else {
				fmt.Printf("Tapped key: %s with modifiers: %v\n", key, modArray)
			}
		},
	}
	keyboardTapCmd.Flags().String("key", "", "Key to tap")
	keyboardTapCmd.Flags().String("mods", "", "Comma-separated list of modifiers (e.g., ctrl,shift)")
	keyboardTapCmd.MarkFlagRequired("key")

	// Keyboard toggle command
	var keyboardToggleCmd = &cobra.Command{
		Use:   "toggle",
		Short: "Toggle a key state (down/up)",
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cmd.Flags().GetString("key")
			state, _ := cmd.Flags().GetString("state")
			err := robotgo.KeyToggle(key, state)
			if err != nil {
				fmt.Printf("Error toggling key: %v\n", err)
			} else {
				fmt.Printf("Toggled key %s to state %s\n", key, state)
			}
		},
	}
	keyboardToggleCmd.Flags().String("key", "", "Key to toggle")
	keyboardToggleCmd.Flags().String("state", "down", "State: down or up")
	keyboardToggleCmd.MarkFlagRequired("key")

	keyboardCmd.AddCommand(keyboardTypeCmd, keyboardTapCmd, keyboardToggleCmd)
	rootCmd.AddCommand(keyboardCmd)

	// ===========================
	// Screen Commands Group
	// ===========================
	var screenCmd = &cobra.Command{
		Use:   "screen",
		Short: "Screen related commands",
	}

	var screenCaptureCmd = &cobra.Command{
		Use:   "capture",
		Short: "Capture a portion of the screen and return a Base64 string",
		Run: func(cmd *cobra.Command, args []string) {
			// Get flags
			full, _ := cmd.Flags().GetBool("full")
			x, _ := cmd.Flags().GetInt("x")
			y, _ := cmd.Flags().GetInt("y")
			width, _ := cmd.Flags().GetInt("width")
			height, _ := cmd.Flags().GetInt("height")
			
			if full {
				// Override values with full screen dimensions
				x = 0
				y = 0
				width, height = robotgo.GetScreenSize()
			}

			// Capture screen area
			bit := robotgo.CaptureScreen(x, y, width, height)
			if bit == nil {
				fmt.Println("Failed to capture screen")
				return
			}
			defer robotgo.FreeBitmap(bit)

			// Convert captured bitmap to image.Image
			img := robotgo.ToImage(bit)

			// Encode the image to PNG into a buffer
			var buf bytes.Buffer
			err := png.Encode(&buf, img)
			if err != nil {
				fmt.Printf("Error encoding screenshot: %v\n", err)
				return
			}

			// Encode the PNG bytes to a Base64 string and print it
			base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
			fmt.Println(base64Str)
		},
	}
	screenCaptureCmd.Flags().Bool("full", false, "Capture the full screen")
	screenCaptureCmd.Flags().Int("x", 0, "X coordinate of capture")
	screenCaptureCmd.Flags().Int("y", 0, "Y coordinate of capture")
	screenCaptureCmd.Flags().Int("width", 100, "Width of capture")
	screenCaptureCmd.Flags().Int("height", 100, "Height of capture")

	// Screen getpixel command
	var screenGetPixelCmd = &cobra.Command{
		Use:   "getpixel",
		Short: "Get the color of the pixel at given coordinates",
		Run: func(cmd *cobra.Command, args []string) {
			x, _ := cmd.Flags().GetInt("x")
			y, _ := cmd.Flags().GetInt("y")
			color := robotgo.GetPixelColor(x, y)
			fmt.Printf("Pixel color at (%d, %d): %s\n", x, y, color)
		},
	}
	screenGetPixelCmd.Flags().Int("x", 0, "X coordinate")
	screenGetPixelCmd.Flags().Int("y", 0, "Y coordinate")

	// Screen size command
	var screenSizeCmd = &cobra.Command{
		Use:   "size",
		Short: "Get the screen size",
		Run: func(cmd *cobra.Command, args []string) {
			width, height := robotgo.GetScreenSize()
			fmt.Printf("Screen size: width=%d, height=%d\n", width, height)
		},
	}

	screenCmd.AddCommand(screenCaptureCmd, screenGetPixelCmd, screenSizeCmd)
	rootCmd.AddCommand(screenCmd)

	// ===========================
	// Window Commands Group
	// ===========================
	var windowCmd = &cobra.Command{
		Use:   "window",
		Short: "Window related commands",
	}

	// Window activate command
	var windowActivateCmd = &cobra.Command{
		Use:   "activate",
		Short: "Activate a window by name or pid",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			pid, _ := cmd.Flags().GetInt("pid")
			if name != "" {
				err := robotgo.ActiveName(name)
				if err != nil {
					fmt.Printf("Error activating window by name: %v\n", err)
				} else {
					fmt.Printf("Activated window with name: %s\n", name)
				}
			} else if pid != 0 {
				err := robotgo.ActivePid(pid)
				if err != nil {
					fmt.Printf("Error activating window by pid: %v\n", err)
				} else {
					fmt.Printf("Activated window with pid: %d\n", pid)
				}
			} else {
				fmt.Println("Please provide either a window name or a pid")
			}
		},
	}
	windowActivateCmd.Flags().String("name", "", "Window name to activate")
	windowActivateCmd.Flags().Int("pid", 0, "Process ID of the window to activate")

	// Window kill command
	var windowKillCmd = &cobra.Command{
		Use:   "kill",
		Short: "Kill a process by pid",
		Run: func(cmd *cobra.Command, args []string) {
			pid, _ := cmd.Flags().GetInt("pid")
			if pid == 0 {
				fmt.Println("Please provide a pid")
				return
			}
			err := robotgo.Kill(pid)
			if err != nil {
				fmt.Printf("Error killing process: %v\n", err)
			} else {
				fmt.Printf("Killed process with pid: %d\n", pid)
			}
		},
	}
	windowKillCmd.Flags().Int("pid", 0, "Process ID to kill")

	// Window title command
	var windowTitleCmd = &cobra.Command{
		Use:   "title",
		Short: "Get the title of the active window",
		Run: func(cmd *cobra.Command, args []string) {
			title := robotgo.GetTitle()
			fmt.Printf("Active window title: %s\n", title)
		},
	}

	windowCmd.AddCommand(windowActivateCmd, windowKillCmd, windowTitleCmd)
	rootCmd.AddCommand(windowCmd)

	// ===========================
	// Clipboard Commands Group
	// ===========================
	var clipboardCmd = &cobra.Command{
		Use:   "clipboard",
		Short: "Clipboard read/write commands",
	}

	// Clipboard read command
	var clipboardReadCmd = &cobra.Command{
		Use:   "read",
		Short: "Read text from the clipboard",
		Run: func(cmd *cobra.Command, args []string) {
			text, err := robotgo.ReadAll()
			if err != nil {
				fmt.Printf("Error reading clipboard: %v\n", err)
			} else {
				fmt.Printf("Clipboard contents: %s\n", text)
			}
		},
	}

	// Clipboard write command
	var clipboardWriteCmd = &cobra.Command{
		Use:   "write",
		Short: "Write text to the clipboard",
		Run: func(cmd *cobra.Command, args []string) {
			text, _ := cmd.Flags().GetString("text")
			if text == "" {
				fmt.Println("Please provide text to write to the clipboard")
				return
			}
			err := robotgo.WriteAll(text)
			if err != nil {
				fmt.Printf("Error writing to clipboard: %v\n", err)
			} else {
				fmt.Println("Text written to clipboard")
			}
		},
	}
	clipboardWriteCmd.Flags().String("text", "", "Text to write to clipboard")
	clipboardWriteCmd.MarkFlagRequired("text")

	clipboardCmd.AddCommand(clipboardReadCmd, clipboardWriteCmd)
	rootCmd.AddCommand(clipboardCmd)

	// ===========================
	// Process Commands Group
	// ===========================
	var processCmd = &cobra.Command{
		Use:   "process",
		Short: "Process related commands",
	}

	// Process list command
	var processListCmd = &cobra.Command{
		Use:   "list",
		Short: "List processes (PID and Name)",
		Run: func(cmd *cobra.Command, args []string) {
			processes, err := robotgo.Process()
			if err != nil {
				fmt.Printf("Error listing processes: %v\n", err)
				return
			}
			fmt.Println("Processes:")
			for _, proc := range processes {
				fmt.Printf("PID: %d, Name: %s\n", proc.Pid, proc.Name)
			}
		},
	}

	processCmd.AddCommand(processListCmd)
	rootCmd.AddCommand(processCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
