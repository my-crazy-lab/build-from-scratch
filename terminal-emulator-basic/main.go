package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/creack/pty"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const MaxBufferSize = 16

func main() {
	a := app.New()
	w := a.NewWindow("Terminal by minhle")

	ui := widget.NewTextGrid()
	ui.SetText("Hello")

	c := exec.Command("/bin/bash")
	p, err := pty.Start(c)
	if err != nil {
		fyne.LogError("Fail to open PTY", err)
		os.Exit(1)
	}

	defer c.Process.Kill()

	onTypedKey := func(e *fyne.KeyEvent) {
		fmt.Println("key event ", e.Name)
		if e.Name == fyne.KeyEnter || e.Name == fyne.KeyReturn {
			_, _ = p.Write([]byte{'\r'})
		}
	}

	onTypedRune := func(r rune) {
		fmt.Println("key rune ", r)
		_, _ = p.WriteString(string(r))
	}

	w.Canvas().SetOnTypedKey(onTypedKey)
	w.Canvas().SetOnTypedRune(onTypedRune)

	buffer := [][]rune{}
	reader := bufio.NewReader(p)

	go func() {
		line := []rune{}
		buffer = append(buffer, line)
		for {
			r, _, err := reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					return
				}
				os.Exit(0)
			}

			fmt.Println("reader rune ", r)
			line = append(line, r)
			buffer[len(buffer)-1] = line
			if r == '\n' {
				if len(buffer) > MaxBufferSize { // If the buffer is at capacity...
					buffer = buffer[1:] // ...pop the first line in the buffer
				}

				line = []rune{}
				buffer = append(buffer, line)
			}
		}
	}()

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			ui.SetText("")
			var lines string
			for _, line := range buffer {
				// fmt.Println("line: ", line)
				lines = lines + string(line)
			}
			ui.SetText(string(lines))
		}
	}()

	w.SetContent(
		container.New(layout.NewGridWrapLayout(fyne.NewSize(800, 400)), ui),
	)
	w.ShowAndRun()
}
