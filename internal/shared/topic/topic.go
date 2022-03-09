package topic

type Topic int

const (
	SET_OFFSET Topic = iota
	SET_CURRENT_COLOR
	SET_CURRENT_TOOL
	SET_CURRENT_SPRITE
	GET_SELECTED_SPRITE
	SET_CANVAS
	GET_PIXELS
	SET_ELEMENT
	SET_PIXEL
	PUSH_PIXELS
	RENDER
	DRAW
	UPDATE
	SAVE
	HANDLE_CLICK
	COPY
	PASTE
	UNDO
	REDO
	LEFT_CLICK
	RIGHT_CLICK
	PUSH_TO_CLIPBOARD
)

func (t Topic) String() string {
	switch t {
	case SET_OFFSET:
		return "SET_OFFSET"
	case SET_CURRENT_COLOR:
		return "SET_CURRENT_COLOR"
	case SET_CURRENT_TOOL:
		return "SET_CURRENT_TOOL"
	case SET_CURRENT_SPRITE:
		return "SET_CURRENT_SPRITE"
	case GET_SELECTED_SPRITE:
		return "GET_SELECTED_SPRITE"
	case SET_CANVAS:
		return "SET_CANVAS"
	case GET_PIXELS:
		return "GET_PIXELS"
	case SET_ELEMENT:
		return "SET_ELEMENT"
	case SET_PIXEL:
		return "SET_PIXEL"
	case PUSH_PIXELS:
		return "PUSH_PIXELS"
	case RENDER:
		return "RENDER"
	case DRAW:
		return "DRAW"
	case UPDATE:
		return "UPDATE"
	case SAVE:
		return "SAVE"
	case HANDLE_CLICK:
		return "HANDLE_CLICK"
	case PASTE:
		return "PASTE"
	case COPY:
		return "COPY"
	case UNDO:
		return "UNDO"
	case REDO:
		return "REDO"
	case LEFT_CLICK:
		return "LEFT_CLICK"
	case RIGHT_CLICK:
		return "RIGHT_CLICK"
	case PUSH_TO_CLIPBOARD:
		return "PUSH_TO_CLIPBOARD"
	}
	return "unkown"
}
