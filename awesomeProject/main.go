package main

import (
	"fmt"
)

func main() {

	// Exercicio 1 - Lista de Nomes

	names := []string{
		"Benjamin",
		"Fernando",
		"Brenda",
		"Marcos",
		"Pedro",
		"Evandro",
		"Alex",
		"Maria",
		"Frederico",
		"Juan",
		"Leandro",
		"Eduardo",
		"Camila",
	}

	names = append(names, "Gabriela")

	fmt.Println("A nova lista depois da inserção é: ", names)

	//Exercicio 2 = Quantos anos você tem...

	var employees = map[string]int{"Benjamin": 20, "Fernando": 28, "Brenda": 19, "Marcos": 44, "Pedro": 30}

	fmt.Println("A idade é", employees["Benjamin"])
	number := 0

	for _, valor := range employees {
		if valor > 21 {
			number++
		}
	}

	//Saber quantos dos seus funcionários têm mais de 21 anos.

	fmt.Println("A quantidade de funcionários com idade superior a 21 anos é: ", number)

	//Adicionar um funcionário novo a lista, chamado Igor, que tem 25 anos.

	employees["Igor"] = 25

	//Remover o Pedro do map.

	delete(employees, "Pedro")
	
}
