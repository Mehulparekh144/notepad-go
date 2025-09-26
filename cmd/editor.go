package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Editor struct {
	ColSize   int
	RowSize   int
	Buffer    *Rope
	Cursor    *Cursor
	UndoStack []Operation
	RedoStack []Operation
}

type Cursor struct {
	row int
	col int
	pos int
}

const BUFFERFILE = "buffer.txt"

func NewEditor(size int) *Editor {
	return &Editor{
		ColSize: size,
		RowSize: 10,
		Buffer:  NewRope(""),
		Cursor: &Cursor{
			row: 0,
			col: 0,
			pos: 0,
		},
	}
}

func (e *Editor) Insert(word string) {
	e.Cursor.pos = e.Cursor.col + e.Cursor.row*e.ColSize
	if e.Cursor.pos == e.RowSize {
		e.RowSize++
	}
	e.Execute(InsertOp{Pos: e.Cursor.pos, Text: word})
}

func (e *Editor) MoveCursorToLast() {
	length := e.Buffer.Length()
	e.Cursor.col = length % e.ColSize
	e.Cursor.row = length / e.ColSize
}

func (e *Editor) Backspace(chars int) {
	if e.Cursor.pos == 0 {
		return
	}

	end := e.Cursor.pos
	start := max(0, end-chars)
	deleted := e.Buffer.Substring(start, end)
	e.Execute(DeleteOp{Pos: start, Text: deleted})
}

func (e *Editor) MoveCursorToLeft(chars int) {
	if e.Cursor.pos > 0 {
		length := e.Cursor.pos - chars
		length = max(length, 0)
		e.Cursor.pos = length
		e.Cursor.col = e.Cursor.pos % e.ColSize
		e.Cursor.row = e.Cursor.pos / e.ColSize
	}
}

func (e *Editor) MoveCursorToRight(chars int) {
	if e.Cursor.pos < e.Buffer.Length() {
		length := e.Cursor.pos + chars
		length = min(length, e.Buffer.Length())
		e.Cursor.pos = length
		e.Cursor.col = e.Cursor.pos % e.ColSize
		e.Cursor.row = e.Cursor.pos / e.ColSize
	}
}

func (e *Editor) Content(showCursor bool) string {
	runes := []rune(e.Buffer.String())
	matrix := make([][]rune, e.RowSize)

	for i := range matrix {
		matrix[i] = make([]rune, e.ColSize)
	}

	for i, c := range runes {
		row := i / e.ColSize
		col := i % e.ColSize

		matrix[row][col] = c
	}

	rows := []string{}

	for i, c := range matrix {
		empty := true
		for _, r := range c {
			if r != 0 {
				empty = false
				break
			}
		}
		if empty {
			continue
		}
		rowStr := string(c)

		if showCursor && i == e.Cursor.row {
			if e.Cursor.col < len(rowStr) {
				rowStr = rowStr[:e.Cursor.col] + "|" + rowStr[e.Cursor.col:]
			} else {
				rowStr = rowStr + "|"
			}
		}
		rows = append(rows, rowStr)
	}

	return strings.Join(rows, "\n")
}

func (e *Editor) Execute(op Operation) {
	op.Apply(e)
	e.UndoStack = append(e.UndoStack, op)
	e.RedoStack = nil
}

func (e *Editor) Undo() {
	if len(e.UndoStack) == 0 {
		return
	}

	op := e.UndoStack[len(e.UndoStack)-1]
	op.Undo(e)
	e.UndoStack = e.UndoStack[:len(e.UndoStack)-1]
	e.RedoStack = append(e.RedoStack, op)
}
func (e *Editor) Redo() {
	if len(e.RedoStack) == 0 {
		return
	}
	op := e.RedoStack[len(e.RedoStack)-1]
	op.Apply(e)
	e.RedoStack = e.RedoStack[:len(e.RedoStack)-1]
	e.UndoStack = append(e.UndoStack, op)
}

func (e *Editor) LoadFile() {

	file, err := os.Open(BUFFERFILE)
	if err != nil {
		fmt.Println("Error opening the file, Try again")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()

		clean := bytes.ReplaceAll(line, []byte{0}, nil)

		if len(clean) > 0 {
			e.Insert(string(clean))
		}
	}

}

func (e *Editor) SaveBuffer() {
	file, err := os.OpenFile(BUFFERFILE, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if os.IsNotExist(err) {
		os.Create(BUFFERFILE)
	}

	defer file.Close()

	_, err = file.WriteString(e.Content(false))
	if err != nil {
		fmt.Println("Error writing to the file, Try again")
		return
	}

	fmt.Println(BUFFERFILE + " saved successfully !")
}
