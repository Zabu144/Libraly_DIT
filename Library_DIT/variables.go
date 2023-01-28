package main

import "fmt"

//================================================================================================

type TextVariable struct {
	Name  string
	Email string
	Pass  string
}

func (t TextVariable) variable() {
	TEXTwelcome = `
-=-=-= BIBLIOTECA VIRTUAL =-=-=-
    -=-=-= Cadastro =-=-=-
`
	fmt.Print(TEXTwelcome)
}

type Books struct {
	Id    int     `db:"idbook"`
	NameB string  `db:"namebook"`
	GenB  string  `db:"genbook"`
	Price float64 `db:"pricebook"`
	QntB  int     `db:"qnt"`
}

func menu() {
	TEXTmenu = `
-=-=-= MENU DE OPÇÕES =-=-=-

[1] COMPRAR LIVROS
[2] ADICONAR LIVROS
[3] MOSTRAR LIVROS
[4] REMOVER LIVROS
`
	fmt.Print(TEXTmenu)
}

func menuGen() {
	TEXTmenuGen = `
-=-=-= MENU DE OPÇÕES =-=-=-

[1] AÇÃO
[2] ROMANCE
[3] FICÇÃO
[4] TERROR
`
	fmt.Print(TEXTmenuGen)
}

var BooksShoplist = []string{}
var BooksRentlist = []string{}
