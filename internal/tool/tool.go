package tool

type Tool int

const (
	Pencil Tool = iota
	Fill
	Drag
	Undo
	Redo
	Open
	Save
	Info
)

func (t Tool) String() string {
	switch t {
	case Pencil:
		return "Pencil"
	case Fill:
		return "Fill"
	case Drag:
		return "Drag"
	case Undo:
		return "Undo"
	case Redo:
		return "Redo"
	case Open:
		return "Open"
	case Save:
		return "Save"
	case Info:
		return "Info"
	}

	return "unknown"
}
