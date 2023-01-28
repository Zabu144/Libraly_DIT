package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// ===================================================
// DECLARAÇÃO DE VARIÁVEIS AUXILIARES
var cont2 = 0
var sum = 0

// ===================================================
// BANCO DE DADOS
func getConnection() (*sqlx.DB, error) {
	conn, err := sqlx.Open("postgres", "user=postgres password=140404 dbname=LibralyDB host=localhost sslmode=disable")
	return conn, err
}

// ===================================================
// FUNÇÕES PRINCIPAIS

// ===================================================
// CADASTRO DE DADOS
func cadastro() {
	Text := TextVariable{}

	Text.variable()
	fmt.Print("Digite o nome: ")
	_, _ = fmt.Scan(&Text.Name)

	fmt.Print("Digite o email: ")
	_, _ = fmt.Scan(&Text.Email)

	fmt.Print("Digite a senha: ")
	_, _ = fmt.Scan(&Text.Pass)

	dados = append(dados, Text)
	fmt.Print(dados)
	clearTerm()
}

// ===================================================

// ===================================================
// COMPRAR LIVROS
func shop() {
	fmt.Print("Quantos livros você quer comprar? ")
	_, _ = fmt.Scan(&manyShop)

	for i := 1; i <= manyShop; i++ {
		clearTerm()

		fmt.Print("Livros disponíveis: ")
		jumpLine()
		toList()
		jumpLine()

		var id int
		fmt.Print("Digite o ID do livro que deseja comprar: ")
		_, _ = fmt.Scan(&id)

		conn, err := getConnection()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		var book Books
		err = conn.QueryRow("SELECT idbook, namebook, genbook, pricebook, qnt FROM dados WHERE idbook = $1",
			id).Scan(&book.Id, &book.NameB, &book.GenB, &book.Price, &book.QntB)
		if err != nil {
			fmt.Println("Erro ao recuperar livro do banco de dados", err)
			return
		}

		cart = append(cart, book)
		fmt.Println("Livro adicionado ao carrinho com sucesso!")
	}

	fmt.Println("=======================================================")
	fmt.Println("Livros no carrinho: ", cart)
}

// ===================================================

// ===================================================
// ALUGAR LIVROS
func rent() {
	fmt.Print("Quantos livros você quer alugar? ")
	_, _ = fmt.Scan(&manyRent)

	for i := 1; i <= manyRent; i++ {
		clearTerm()
		menuGen()
		jumpLine()
		fmt.Print("Escolha um Gênero: ")
		_, _ = fmt.Scan(&choiceGen)

		switch choiceGen { // SWITCH PARA GÊNERO
		// =================================================================================================
		// AÇÃO
		case 1:
			jumpLine()
			cont := 0

			fmt.Println("=======================================================")
			fmt.Print("Preço por aluguel de livro: R$10,00 P/DIA")
			jumpLine()
			for i := 1; i <= 5; i++ {
				cont += 1
				fmt.Printf("[%d] %s", cont, Actionlist[i])
				jumpLine()

			}
			fmt.Print("Escolha um livro: ")
			_, _ = fmt.Scan(&choiceBook)
			fmt.Printf("Por quantos dias você quer alugar esse livro? ")
			_, _ = fmt.Scan(&rentBook)
			sum += rentBook

			BooksRentlist = append(BooksRentlist, Actionlist[choiceBook])
			cont2 += 1
			// =================================================================================================
		// ROMANCE
		case 2:
			jumpLine()
			cont := 0

			fmt.Println("=======================================================")
			fmt.Print("Preço por aluguel de livro: R$10,00 P/DIA")
			jumpLine()
			for i := 1; i <= 5; i++ {
				cont += 1
				fmt.Printf("[%d] %s", cont, Romancelist[i])
				jumpLine()

			}
			fmt.Print("Escolha um livro: ")
			_, _ = fmt.Scan(&choiceBook)
			fmt.Printf("Por quantos dias você quer alugar esse livro? ")
			_, _ = fmt.Scan(&rentBook)
			sum += rentBook

			BooksRentlist = append(BooksRentlist, Romancelist[choiceBook])
			// =================================================================================================
		// FICÇÃO
		case 3:
			jumpLine()
			cont := 0

			fmt.Println("=======================================================")
			fmt.Print("Preço por aluguel de livro: R$10,00 P/DIA")
			jumpLine()
			for i := 1; i <= 5; i++ {
				cont += 1
				fmt.Printf("[%d] %s", cont, Fictionlist[i])
				jumpLine()

			}
			fmt.Print("Escolha um livro: ")
			_, _ = fmt.Scan(&choiceBook)
			fmt.Printf("Por quantos dias você quer alugar esse livro? ")
			_, _ = fmt.Scan(&rentBook)
			sum += rentBook

			BooksRentlist = append(BooksRentlist, Fictionlist[choiceBook])
			// =================================================================================================
		// TERROR
		case 4:
			jumpLine()
			cont := 0

			fmt.Println("=======================================================")
			fmt.Print("Preço por aluguel de livro: R$10,00 P/DIA")
			jumpLine()
			for i := 1; i <= 5; i++ {
				cont += 1
				fmt.Printf("[%d] %s", cont, Horrorlist[i])
				jumpLine()

			}
			fmt.Print("Escolha um livro: ")
			_, _ = fmt.Scan(&choiceBook)
			fmt.Printf("Por quantos dias você quer alugar esse livro? ")
			_, _ = fmt.Scan(&rentBook)
			sum += rentBook

			BooksRentlist = append(BooksRentlist, Horrorlist[choiceBook])
			// =================================================================================================
		}
	}

	fmt.Print("==========================================================")
	jumpLine()
	fmt.Printf("Você alugou %d livros", cont2)
	jumpLine()
	fmt.Print(BooksRentlist)
	jumpLine()
	fmt.Printf("Total a pagar pelos livros ALUGADOS: R$%d,00", sum*10)
}

// ===================================================

// ===================================================
// ADICIONAR LIVROS
func insert() {
	myBook := Books{}

	fmt.Print("Nome do livro: ")
	_, _ = fmt.Scan(&myBook.NameB)

	if myBook.NameB == "" {
		fmt.Print("Dado Inválido!")
		return
	}

	fmt.Print("Gênero do livro: ")
	_, _ = fmt.Scan(&myBook.GenB)

	if myBook.GenB == "" {
		fmt.Print("Dado Inválido!")
		return
	}

	fmt.Print("Preço do livro: ")
	_, _ = fmt.Scan(&myBook.Price)

	if myBook.Price <= 0 {
		fmt.Print("Preço inválido!")
		return
	}

	fmt.Print("Quantidade de livros: ")
	_, _ = fmt.Scan(&myBook.QntB)

	if myBook.QntB <= 0 {
		fmt.Print("Você não pode adcionar 0 livros!")
		return
	}

	conn, err := getConnection()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	insert := `
		INSERT INTO dados (nameBook, genBook, priceBook, qnt)
		VALUES ($1, $2, $3, $4)`

	_, err = conn.Exec(insert, myBook.NameB, myBook.GenB, myBook.Price, myBook.QntB)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Livro Adicionado a biblioteca!")
	}
}

// ===================================================

// ===================================================
// LISTAR
func toList() {

	book := searching2()

	for _, Books := range book {
		fmt.Println("=======================================")
		fmt.Println("ID: ", Books.Id)
		fmt.Println("Nome do livro: ", Books.NameB)
		fmt.Println("Gênero do livro: ", Books.GenB)
		fmt.Println("Preço do livro: R$", Books.Price)
		fmt.Println("Quantidade de livros: ", Books.QntB)
		fmt.Println("=======================================")
		jumpLine()
	}
}

func searching2() []Books {
	conn, err := getConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var book []Books
	err = conn.Select(&book, "SELECT idbook, namebook, genbook, pricebook, qnt from dados")
	if err != nil {
		log.Fatal(err)
	}

	return book
}

// ===================================================

// ===================================================
// DELETE
func delete() {
	toList()

	var id int
	fmt.Print("Digite o id: ")
	fmt.Scan(&id)

	conn, err := getConnection()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	query := "DELETE FROM dados WHERE idbook = $1"
	r, err := conn.Exec(query, id)
	if err != nil {
		panic(err)
	}
	linhasDeletadas, _ := r.RowsAffected()
	if linhasDeletadas > 0 {
		fmt.Println("O livro foi retirado da biblioteca")
	} else {
		fmt.Println("Nenhum livro foi alterado")
	}
}

// ===================================================

// ===================================================
// FUNÇÕES SECUNDÁRIAS
func jumpLine() {
	fmt.Print("\n")
} //Pula linha

func execTerm(name string, arg ...string) {
	terminal := exec.Command(name, arg...)
	terminal.Stdout = os.Stdout
	_ = terminal.Run()
}

func clearTerm() {
	switch runtime.GOOS {
	case "windows":
		execTerm("cmd", "/c", "cls")
	default:
		execTerm("clear")
	}
} //Limpa o terminal
//===================================================
