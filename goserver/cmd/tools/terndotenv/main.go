package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Diretorio de trabalho atual:", dir)

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	fmt.Println("Executando comando:", cmd.String())

	//if err := cmd.Run(); err != nil {
	//	panic(err)
	//}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error ao executat comando: %s\n", err)
		fmt.Printf("Saida do comando: %s\n", output)
		panic(err)
	}

	fmt.Println("Saida do comando:", string(output))
}
