package v1

//START OMIT

// Na v1 temos poucos parâmetros em uso HLnewCasa1
func NewCasa(corParede string, alturaPeDireito int) *casa {
	return &casa{corParede, alturaPeDireito}
} // HLnewCasa1

type casa struct {
	paredeCor       string
	alturaPeDireito int
}

//END OMIT
