package main

type Editor struct {
	Buffer *Rope
	Cursor *Cursor
}

type Cursor struct {
	row int
	col int
}

func NewEditor() *Editor {
	return &Editor{
		Buffer: NewRope(""),
		Cursor: &Cursor{
			row: 0,
			col: 0,
		},
	}
}

func (e *Editor) Insert(word string) {
	e.Buffer = e.Buffer.Insert(e.Cursor.col, word)
	e.Cursor.col += len([]rune(word))
}

func (e *Editor) MoveCursorToLast() {
	e.Cursor.col = e.Buffer.Length()
}

func (e *Editor) Backspace(chars int) {
	if e.Cursor.col == 0 {
		return
	}

	chars = min(chars, e.Cursor.col)

	e.Buffer = e.Buffer.Delete(e.Cursor.col-chars, e.Cursor.col)
	e.Cursor.col -= chars
}

func (e *Editor) MoveCursorToLeft(chars int) {
	if e.Cursor.col > 0 {
		e.Cursor.col = max(e.Cursor.col-chars, 0)
	}
}

func (e *Editor) MoveCursorToRight(chars int) {
	if e.Cursor.col < e.Buffer.Length() {
		e.Cursor.col = min(e.Cursor.col+chars, e.Buffer.Length())
	}
}

func (e *Editor) Content(showCursor bool) string {
	if !showCursor {
		return e.Buffer.String()
	}

	runes := []rune(e.Buffer.String())
	cursorPos := e.Cursor.col

	runes = append(runes[:cursorPos], append([]rune{'|'}, runes[cursorPos:]...)...)
	return string(runes)
}
