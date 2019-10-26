package solve

import (
	"fmt"
	"sync"
)

const MaxTileNumber = 13

// UnstructuredBoard represents a collection of tiles without dividing into Groups or Runs
type UnstructuredBoard [MaxTileNumber]TileStack

type GroupAndRemainder struct {
	GroupingOption GroupCombo
	Remainder      TileStack
}

func AllGroupingOptions(board UnstructuredBoard) [MaxTileNumber][]GroupAndRemainder {
	options := [MaxTileNumber][]GroupAndRemainder{}
	for i := 0; i < MaxTileNumber; i++ {
		asMap := FindGroupings(board[i])
		asSlice := make([]GroupAndRemainder, 0, len(asMap))
		for gc, rem := range asMap {
			asSlice = append(asSlice, GroupAndRemainder{gc, rem})
		}
		options[i] = asSlice
	}
	return options
}

func FindAllValidOptions(options [MaxTileNumber][]GroupAndRemainder) [][MaxTileNumber]GroupAndRemainder {
	resultsSlice := [][MaxTileNumber]GroupAndRemainder{}
	resultsChan := make(chan []GroupAndRemainder)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for res := range resultsChan {
			ar := [MaxTileNumber]GroupAndRemainder{}
			copy(ar[:], res)
			resultsSlice = append(resultsSlice, ar)
		}
		wg.Done()
	}()

	choices := make([]GroupAndRemainder, 0, MaxTileNumber)

	recurse(choices, options[:], resultsChan)

	close(resultsChan)
	wg.Wait()
	return resultsSlice
}

func recurse(
	choices []GroupAndRemainder,
	optionsAhead [][]GroupAndRemainder,
	results chan<- []GroupAndRemainder) {

	if len(optionsAhead) == 0 { // no more choices
		results <- choices
		return
	}

	fmt.Printf("recurse depth %02d\n", len(choices))

	index := len(choices)
	if index >= cap(choices) {
		panic("choices capacity insufficient, would allocate")
	}
	choices = append(choices, GroupAndRemainder{}) // one more slot
	nextOptions := optionsAhead[1:]                // one less slot

	for _, option := range optionsAhead[0] {
		// TODO: form the struct in FindGroupings instead?
		choices[index] = option

		if !runnable(choices) {
			continue
		}

		recurse(choices, nextOptions, results)
	}
}

func runnable(choices []GroupAndRemainder) bool {
	var runs runsState
	nChoices := len(choices)
	for position := 0; position < nChoices; position++ {
		if !runs.attemptUpdate(position, choices[position].Remainder) {
			return false
		}
	}
	return true
}

const NumColors = 4

type runsState [NumColors]runsStateForColor

type runsStateForColor [2]runCursor

func (r *runsStateForColor) numActive() int {
	n := 0
	if r[0].Active {
		n += 1
	}
	if r[1].Active {
		n += 1
	}
	return n
}

func (r *runsStateForColor) startNew(position int) {
	if r[0].tryStartNew(position) {
		return
	}
	if r[1].tryStartNew(position) {
		return
	}
	panic("cannot start new run")
}

type runCursor struct {
	Active     bool
	StartIndex int
}

func (c *runCursor) tryStartNew(position int) bool {
	if c.Active {
		return false // cannot start a new one, since already active
	}
	*c = runCursor{true, position}
	return true
}

const MinRunLength = 3

func (c *runCursor) canStopBefore(position int) bool {
	return position-c.StartIndex >= MinRunLength
}

func (r *runsStateForColor) tryStopRun(position int) bool {
	if r[0].Active && r[0].canStopBefore(position) {
		r[0] = runCursor{false, 0}
		return true
	}
	if r[1].Active && r[1].canStopBefore(position) {
		r[1] = runCursor{false, 0}
		return true
	}
	return false // no active runs, or cannot stop them
}

func (r *runsStateForColor) tryStopBothRuns(position int) bool {
	if !r.tryStopRun(position) {
		return false
	}
	return r.tryStopRun(position)
}

// update runs with next set of tiles at a given position
// returns true if the new tiles can be accomodated by valid Runs
// returns false otherwise
func (r *runsState) attemptUpdate(position int, t TileStack) bool {
	for colorIndex := 0; colorIndex < NumColors; colorIndex++ {
		if !r[colorIndex].attemptUpdate(position, CountOfColor(t, colorIndex)) {
			return false
		}
	}
	return true
}

// update runsState for this color with the quantity of tiles at the given position
// returns true if new tiles of the given quantity can be accomodated by valid Runs
// returns false otherwise
func (r *runsStateForColor) attemptUpdate(position int, quantity int) bool {
	switch r.numActive() {
	case 0: // no active runs
		switch quantity {
		case 0: // no tiles of this color
			return true
		case 1: // need to start a new run
			r.startNew(position) // would panic if 2 active runs
			return true
		case 2:
			r.startNew(position) // would panic if we already had 2 active runs
			r.startNew(position) // would panic if we already had 1 active run
			return true
		}
	case 1:
		switch quantity {
		case 0:
			return r.tryStopRun(position)
		case 1: // run continues, no change
			return true
		case 2:
			r.startNew(position)
			return true
		}
	case 2:
		switch quantity {
		case 0:
			return r.tryStopBothRuns(position)
		case 1:
			return r.tryStopRun(position)
		case 2:
			return true
		}
	}
	return false
}
