package model


type Tarefa struct {
	Id int
	Descricao string
	EstaFinalizada bool
	CriadaEm string
}

func ConverterParaPortugues(statusTarefa bool) string {
	if(statusTarefa) {
		return "Sim"
	}

	return "Não"
}