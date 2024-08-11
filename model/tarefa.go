package model

import "time"

type Tarefa struct {
	Id int
	Descricao string
	EstaFinalizada bool
	CriadaEm time.Time
}

func ConverterParaPortugues(statusTarefa bool) string {
	if(statusTarefa) {
		return "Sim"
	}

	return "Não"
}