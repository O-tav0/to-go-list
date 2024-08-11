package service

import (
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
	"to-go-list/model"
	"to-go-list/data"
)

func AdicionarTarefa(descricao string) {
	novaTarefa := preencherNovaTarefaStruct(descricao)
	data.AdicionarNovaTarefaNoArquivo(novaTarefa)
}

func ListarTarefasPendentes() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)

	registros := data.BuscarTarefasPendentes();

	w.Write([]byte("\tID\tDescrição\tEstá Completa\tCriada Em\t\n"))
	for i := 0; i < len(registros); i++ {
		w.Write(getLinhaOutput(registros[i]))
	}

	w.Flush()
}

func getLinhaOutput(registro []string) []byte {
	linhaOutput := "\t" + registro[0] + "\t" + registro[1] + "\t" + registro[2] + "\t" + registro[3] + "\t" + "\n"
	return []byte(linhaOutput)
}

func preencherNovaTarefaStruct(descricao string) *model.Tarefa {
	id, err := strconv.Atoi(data.GetProximoIdRegistro())
	if err != nil {
		log.Fatal("Falha ao converter o número doúltimo Id")
	}

	agora := time.Now()

	novaTarefa := &model.Tarefa{
		Id:             id,
		Descricao:      descricao,
		EstaFinalizada: false,
		CriadaEm:       agora.Format("02/01/2006"),
	}

	return novaTarefa
}
