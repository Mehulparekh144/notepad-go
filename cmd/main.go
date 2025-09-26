package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadRawString(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimRight(line, "\r\n"), nil
}

const HELPMENU = `insert <word> : Add text at the end of the document
cursor reset : Move cursor to the end of the document
cursor left <num> : Move cursor left by <num> characters
cursor right <num> : Move cursor right by <num> characters
undo : Undoes the operation
redo : Redoes the operation
test : Test
save : Saves File
backspace <num> : Delete <num> characters to the left of cursor
quit : Exit the text editor`

func main() {
	editor := NewEditor(10)
	reader := bufio.NewReader(os.Stdin)
	editor.LoadFile()

	println("My Editor")
	fmt.Println("--------Loaded Contents-------")
	fmt.Println(editor.Content(true))
	fmt.Println("-----------------------------")
	println(HELPMENU)
	for {

		fmt.Print(">")
		text, err := ReadRawString(reader)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if text == "quit" {
			break
		}

		parts := strings.Fields(text)
		if len(parts) == 0 {
			fmt.Println(HELPMENU)
			continue
		}

		switch parts[0] {
		case "insert":
			if len(parts) < 2 {
				fmt.Println("Invalid Command \n" + HELPMENU)
				continue
			}

			word := parts[1]
			editor.Insert(word)

			fmt.Println(editor.Content(true))

		case "cursor":
			if len(parts) < 2 {
				fmt.Println("Invalid Command \n" + HELPMENU)
				continue
			}

			switch parts[1] {
			case "reset":
				editor.MoveCursorToLast()
			case "left":
				chars := 1
				if len(parts) > 2 {
					chars, _ = strconv.Atoi(parts[2])
				}
				editor.MoveCursorToLeft(chars)
			case "right":
				chars := 1
				if len(parts) > 2 {
					chars, _ = strconv.Atoi(parts[2])
				}
				editor.MoveCursorToRight(chars)

			default:
				fmt.Println("Invalid Command \n" + HELPMENU)
			}

			fmt.Println(editor.Content(true))

		case "backspace":
			if len(parts) < 2 {
				editor.Backspace(1)
			} else {
				chars, _ := strconv.Atoi(parts[1])
				editor.Backspace(chars)
			}
			fmt.Println(editor.Content(true))
		case "undo":
			editor.Undo()
			fmt.Println(editor.Content(true))
		case "redo":
			editor.Redo()
			fmt.Println(editor.Content(true))
		case "test":
			fmt.Println("Col Size", editor.ColSize)
			fmt.Println("Row Size", editor.RowSize)
			editor.Insert("hello")
			fmt.Println(editor.Content(true))
			fmt.Println("Col", editor.Cursor.col)
			fmt.Println("Rol", editor.Cursor.row)
			editor.Insert("world")
			fmt.Println(editor.Content(true))
			fmt.Println("Col", editor.Cursor.col)
			fmt.Println("Rol", editor.Cursor.row)
		case "save":
			editor.SaveBuffer()
		default:
			fmt.Println("Invalid Command \n" + HELPMENU)
			continue
		}
	}

}
