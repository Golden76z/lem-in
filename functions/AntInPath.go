package functions

func AntInPath(paths [][]string, numberOfAnts int) []int {
	// Initialiser le tableau résultat avec un élément pour chaque chemin
	result := make([]int, len(paths))

	// Boucle tant qu'il reste des fourmis à répartir
	for numberOfAnts > 0 {
		minLen := 0 // Réinitialiser minLen pour chaque nouvelle fourmi

		// Trouver le chemin qui prendra le moins de temps pour la fourmi suivante
		for i := 1; i < len(paths); i++ {
			// Comparer les longueurs des chemins avec le nombre actuel de fourmis
			if len(paths[i])+result[i] < len(paths[minLen])+result[minLen] {
				minLen = i
			}
		}

		// Ajouter une fourmi au chemin trouvé
		result[minLen]++

		// Réduire le nombre de fourmis restantes
		numberOfAnts--
	}

	return result
}
