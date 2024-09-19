package functions

func FilterPath(AllPaths [][]string, start string, end string) [][]string {
	BestSolution := [][]string{}

	// Parcourir tous les chemins comme point de départ potentiel
	for i := 0; i < len(AllPaths); i++ {
		CurrentSolution := [][]string{AllPaths[i]} // Commence avec le premier chemin

		// Essayer de combiner ce chemin avec d'autres
		for j := 0; j < len(AllPaths); j++ {
			if i != j && CheckPath(CurrentSolution, AllPaths[j], start, end) {
				CurrentSolution = append(CurrentSolution, AllPaths[j])
			}
		}

		// Mettre à jour la meilleure solution si la solution courante est meilleure
		if len(CurrentSolution) > len(BestSolution) {
			BestSolution = CurrentSolution
		}
	}

	return BestSolution
}

// CheckPath vérifie si le chemin "current" peut être ajouté à la solution courante "path"
// sans partager de pièces autres que start et end
func CheckPath(path [][]string, current []string, start string, end string) bool {
	// Vérifier chaque chemin déjà dans la solution
	for i := 0; i < len(path); i++ {
		for _, room := range path[i] {
			// Ignorer les pièces de départ et d'arrivée
			if room == start || room == end {
				continue
			}

			// Si une pièce du chemin actuel existe déjà dans le chemin en cours, on retourne false
			for _, curRoom := range current {
				if curRoom == room && curRoom != start && curRoom != end {
					return false
				}
			}
		}
	}
	return true
}
