package main

import "fmt"

const (
	cachorro  = "cachorro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	funcCachorro := Animal("cachorro")
	funcGato := Animal("gato")
	funcHamster := Animal("hamster")
	funcTarantula := Animal("tarantula")

	racaoCachorro := funcCachorro(2)
	racaoGato := funcGato(2)
	racaoHamster := funcHamster(3)
	racaoTarantula := funcTarantula(3)

	fmt.Println("A ração necessária é de: ", racaoCachorro, " gramas")
	fmt.Println("A ração necessária é de: ", racaoGato, " gramas")
	fmt.Println("A ração necessária é de: ", racaoHamster, " gramas")
	fmt.Println("A ração necessária é de: ", racaoTarantula, " gramas")

}

func Animal(animal string) func(quantidade int) int {
	switch animal {
	case cachorro:
		return calCachorro
	case gato:
		return calGato
	case hamster:
		return calHamster
	case tarantula:
		return calTarantula
	}
	return nil
}

func calCachorro(quantidade int) int {
	return quantidade * 10000
}

func calGato(quantidade int) int {
	return quantidade * 5000
}

func calHamster(quantidade int) int {
	return quantidade * 250
}

func calTarantula(quantidade int) int {
	return quantidade * 150
}
