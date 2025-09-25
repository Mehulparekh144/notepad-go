package main

type Rope struct {
	left   *Rope
	right  *Rope
	weight int // Left rope weight
	data   string
}

func NewRope(data string) *Rope {
	return &Rope{
		data: data,
	}
}

func (r *Rope) Insert(index int, data string) *Rope {
	leftPart, rightPart := r.split(index)
	newLeftPart := concat(leftPart, NewRope(data))
	return concat(newLeftPart, rightPart)
}

func (r *Rope) Length() int {
	if r == nil {
		return 0
	}
	if r.left == nil && r.right == nil {
		return len([]rune(r.data))
	}

	return r.weight + r.right.Length()
}

func (r *Rope) Delete(start, end int) *Rope {
	if start >= end {
		return r
	}

	left, rest := r.split(start)
	_, right := rest.split(end - start)
	return concat(left, right)
}

func (r *Rope) String() string {
	if r.left == nil && r.right == nil {
		return r.data
	}

	return r.left.String() + r.right.String()
}

func (r *Rope) split(index int) (*Rope, *Rope) {
	if r.left == nil && r.right == nil {
		arr := []rune(r.data)
		left := &Rope{
			data: string(arr[:index]),
		}

		right := &Rope{
			data: string(arr[index:]),
		}

		return left, right
	}

	if index < r.weight {
		left, right := r.left.split(index)
		return left, concat(right, r.right)
	}

	left, right := r.right.split(index - r.weight)
	return concat(r.left, left), right
}

func concat(left, right *Rope) *Rope {
	return &Rope{
		left:   left,
		right:  right,
		weight: left.Length(),
	}
}
