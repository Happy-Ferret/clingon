package specs

import (
	"testing"
	pt "spectrum/prettytest"
	"cli"
)

func should_receive_commands(t *pt.T) {
	t.True(Interact([]Interactor{NewEnterCommand(console, "Hello gopher!", 2e8)}))
}

func should_browse_history(t *pt.T) {
	t.True(Interact([]Interactor{NewEnterCommand(console, "foo", 2e8)}))
	t.True(Interact([]Interactor{NewEnterCommand(console, "bar", 2e8)}))
	t.True(Interact([]Interactor{NewEnterCommand(console, "biz", 2e8)}))
	t.True(Interact([]Interactor{NewBrowseHistory(console, 
			[]int{
			cli.HISTORY_PREV, 
			cli.HISTORY_PREV, 
			cli.HISTORY_PREV, 
			cli.HISTORY_NEXT,
			cli.HISTORY_NEXT,
			cli.HISTORY_NEXT,
		}, 2e8)}))
}

func should_move_cursor(t *pt.T) {
	t.True(Interact([]Interactor{NewEnterCommand(console, "fooooooooooooooooo", 2e8)}))
	t.True(Interact([]Interactor{NewBrowseHistory(console, []int{cli.HISTORY_PREV}, 2e8)}))
	t.True(Interact([]Interactor{NewMoveCursor(console, []int{
			cli.CURSOR_LEFT,
			cli.CURSOR_LEFT,
			cli.CURSOR_LEFT,
			cli.CURSOR_LEFT,
			cli.CURSOR_LEFT,
			cli.CURSOR_LEFT,
			cli.CURSOR_LEFT,
			cli.CURSOR_RIGHT,
			cli.CURSOR_RIGHT,
			cli.CURSOR_RIGHT,
			cli.CURSOR_RIGHT,
			cli.CURSOR_RIGHT,
		}, 2e8)}))
}

func TestConsoleSpecs(t *testing.T) {
	pt.Describe(
		t,
		"Console",
		should_receive_commands,
		should_browse_history,
		should_move_cursor,

		beforeAll,
		afterAll,
	)
}