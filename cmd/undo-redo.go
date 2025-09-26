package main

type Operation interface {
	Apply(e *Editor)
	Undo(e *Editor)
}

type InsertOp struct {
	Pos  int
	Text string
}

func (op InsertOp) Apply(e *Editor) {
	e.Buffer = e.Buffer.Insert(op.Pos, op.Text)
	cursorPos := op.Pos + len([]rune(op.Text))
	e.Cursor.pos = cursorPos
	e.Cursor.col = cursorPos % e.ColSize
	e.Cursor.row = cursorPos / e.ColSize
}

func (op InsertOp) Undo(e *Editor) {
	e.Buffer = e.Buffer.Delete(op.Pos, op.Pos+len([]rune(op.Text)))
	e.Cursor.pos = op.Pos
	e.Cursor.col = op.Pos % e.ColSize
	e.Cursor.row = op.Pos / e.ColSize
}

type DeleteOp struct {
	Pos  int
	Text string
}

func (op DeleteOp) Apply(e *Editor) {
	e.Buffer = e.Buffer.Delete(op.Pos, op.Pos+len([]rune(op.Text)))
	e.Cursor.pos = op.Pos
	e.Cursor.col = op.Pos % e.ColSize
	e.Cursor.row = op.Pos / e.ColSize
}

func (op DeleteOp) Undo(e *Editor) {
	e.Buffer = e.Buffer.Insert(op.Pos, op.Text)
	cursorPos := op.Pos + len([]rune(op.Text))
	e.Cursor.pos = cursorPos
	e.Cursor.col = cursorPos % e.ColSize
	e.Cursor.row = cursorPos / e.ColSize
}
