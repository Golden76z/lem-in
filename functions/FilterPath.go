package functions

func FilterPath(AllPaths [][]string, start string, end string) [][]string {
	BestSolution := [][]string{}
	CurrentSolution := [][]string{}
	duplicate := false
	visitedrooms := ""
	for i := 0; i < len(AllPaths); i++ {
		for j := 0; j < len(AllPaths); j++ {
			for k := 0; k < len(visitedrooms); k++ {
				for l := 0; l < len(AllPaths[j]); l++ {
					if string(visitedrooms[k]) == AllPaths[j][l] {
						duplicate = true
					}
				}
			}
			if !duplicate {
				CurrentSolution = append(CurrentSolution, AllPaths[j])
				for m := 0; m < len(AllPaths[j]); m++ {

				}
			}
			duplicate = false
		}
		if len(CurrentSolution) > len(BestSolution) {
			BestSolution = CurrentSolution
		}
		CurrentSolution = [][]string{}
	}
	return BestSolution
}

func Duplicate(letter string, s string) bool {

	return true
}
