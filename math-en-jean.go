package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

var limit int

func init() {
	flag.IntVar(&limit, "limit", 100, "Entrez la limite")
}

func main() {
	start := time.Now()

	flag.Parse()

	A := 0.0

	for i := 1.0; i <= float64(limit); i++ {
		A = A + 1/i
	}

	result := strconv.FormatFloat(A, 'f', 12, 64)
	elapsed := time.Since(start)

	fmt.Printf("Limite : %v\nRésultat : %s\nTemps d'exécution : %s\n", limit, result, elapsed)
}
