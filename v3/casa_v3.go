package v3

import (
	"fmt"
)

//START OMIT

func NewCasa(materialParede string) *casa {
	return &casa{materialParede: materialParede}
} // HLnewCasa1

type casa struct {
	paredeCor           string
	alturaPeDireito     int
	quantidadeQuartos   int
	quantidadeBanheiros int
	materialParede      string
}

type option func(*casa)

func (t *casa) Configurar(options ...option) {
	for _, opFunc := range options {
		opFunc(t)
	}
}

func opParedeAzul(c *casa) {
	c.paredeCor = "azul"
}

func opParedeCor(cor string) option {
	return func(c *casa) {
		c.paredeCor = cor
	}
}

func opQuantidadeQuartos(qtd int) option {
	return func(c *casa) {
		c.quantidadeQuartos = qtd
	}
}

func opCasaPequena(c *casa) {
	c.alturaPeDireito = 2
	c.materialParede = "alvenaria"
	c.paredeCor = "amarela"
}

func exemploMain() {
	casa := NewCasa("madeira")                            // HL10START
	casa.Configurar(opParedeAzul, opQuantidadeQuartos(2)) // HL10
	// HL10END

	fmt.Printf("Casa %v \n", casa)
}

//END OMIT
