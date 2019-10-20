package solve

const MaxTileNumber = 13

// UnstructuredBoard represents a collection of tiles without dividing into Groups or Runs
type UnstructuredBoard [MaxTileNumber]TileStack

// GroupingOptions collects all possible ways of forming Groups, but does not account for Runs
type groupingOptions [MaxTileNumber]GroupingOptions

func GenerateGroupingOptions(board UnstructuredBoard) groupingOptions {
	options := [MaxTileNumber]GroupingOptions{}
	for i := 0; i < MaxTileNumber; i++ {
		options[i] = FindGroupings(board[i])
	}

	pruneOptions(options[0], options[1], options[2]) // left-most

	for i := 1; i < 12; i++ {
		pruneOptions(options[i], options[i-1], options[i+1])
	}

	pruneOptions(options[12], options[10], options[11]) // right-most
	return options
}

// remove any options where the remainder can't be covered by Runs with adjacent tiles
func pruneOptions(toPrune GroupingOptions, adjacent1, adjacent2 GroupingOptions) {
	for mapKey, remainder := range toPrune {
		if !remainderCouldBeInRuns(remainder, adjacent1, adjacent2) {
			delete(toPrune, mapKey)
		}
	}
}

// determine if the remainder pattern could be covered by runs across the adjacent positions
func remainderCouldBeInRuns(remainder TileStack, adjacent1, adjacent2 GroupingOptions) bool {
	return mapContainsValue(adjacent1, remainder) && mapContainsValue(adjacent2, remainder)
}

func mapContainsValue(m map[GroupCombo]TileStack, testValue TileStack) bool {
	for _, v := range m {
		if v == testValue {
			return true
		}
	}
	return false
}
