package main

import "fmt"

func main() {
	idUsuario := solution(4, [][]int{
		{3, 5},
		{2, 5},
		{3, 4},
		{1, 4},
		{2, 3},
	})

	fmt.Println(idUsuario)
}

func solution(N int, ratings [][]int) int {
	userMax := make(map[int]int) // Mapa: ID do usuário → maior nota dada

	// Passo 1: Atualizar a maior nota de cada usuário
	for _, review := range ratings {
		userID := review[0]
		rating := review[1]

		// Se a nota atual for maior que a armazenada, atualiza
		if currentMax, exists := userMax[userID]; !exists || rating > currentMax {
			userMax[userID] = rating
		}
	}

	// Passo 2: Encontrar a maior nota global e o menor ID em caso de empate
	maxRating := -1
	minID := -1

	for userID, rating := range userMax {
		if rating > maxRating || (rating == maxRating && userID < minID) {
			maxRating = rating
			minID = userID
		}
	}

	return minID
}
