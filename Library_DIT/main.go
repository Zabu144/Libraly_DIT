package main

import (
	"fmt"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "log"
)

// =====================================================================================
// DECLARAÇÃO DAS VARIAVEIS GERAIS
var (
	TEXTmenu    string
	TEXTmenuGen string
	TEXTwelcome string
)
var dados []TextVariable
var cart []Books
var choiceMenu, choiceGen, choiceBook, rentBook, manyShop, manyRent, finalChoice int

//======================================================================================

// =====================================================================================
// FUNÇÃO PRINCIPAL
func main() {
	cadastro() //FUNÇÃO CADASTRO

	for true {
		menu() //VARIÁVEL DE TEXTO MENU DE OPÇÕES PARA AS FUNÇÕES
		fmt.Print("Sua escolha: ")
		_, _ = fmt.Scan(&choiceMenu) //INPUT PARA OS SWITCHS

		//=========================================================================
		switch choiceMenu {
		case 1: //CASO A ESCOLHA SEJA COMPRAR
			clearTerm()
			fmt.Print("-=-= COMPRANDO LIVROS =-=-")
			jumpLine()
			shop()
		//==========================================================================

		//==========================================================================
		case 2: //CASO A ESCOLHA SEJA ADICONAR LIVROS
			clearTerm()
			fmt.Print("-=-= ADICIONANDO LIVROS =-=-")
			jumpLine()
			insert()
		//==========================================================================

		//==========================================================================
		case 3: //LISTAR LIVROS
			clearTerm()
			jumpLine()
			toList()
		//==========================================================================

		//==========================================================================
		case 4: //REMOVER LIVROS
			clearTerm()
			fmt.Print("-=-= REMOVENDO LIVROS =-=-")
			jumpLine()
			delete()
		}
		//==========================================================================

		//==========================================================================
		//CONDIÇÃO PARA CHAMAR O MENU OU ENCERRAR
		jumpLine()
		fmt.Print("[1] VOLTAR MENU  [2] ENCERRAR PROGRAMA ")
		_, _ = fmt.Scan(&finalChoice)

		if finalChoice == 1 {
			clearTerm()
			continue
		}
		if finalChoice == 2 {
			clearTerm()
			fmt.Printf("TOTAL A PAGAR: R$%d,00", cont2*25+sum*10)
			jumpLine()
			fmt.Print("Boa leitura!")
			break
		}
		//==========================================================================
	}
}
