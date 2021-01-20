package v2

//START OMIT

//Na v2 temos novos parametros que "quebram" seu utilizadores HLnewCasa1
func NewCasa(corParede string, alturaPeDireito int, materialParede string,
	qtdQuartos, qtdBanheiros int) *casa {
	return &casa{corParede, alturaPeDireito,
		materialParede, qtdQuartos,
		qtdBanheiros}
} // HLnewCasa1

type casa struct {
	paredeCor           string
	alturaPeDireito     int
	materialParede      string
	quantidadeQuartos   int
	quantidadeBanheiros int
}

//END OMIT
