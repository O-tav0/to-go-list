package service

import (
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
	"to-go-list/data"
	"to-go-list/model"
)

func AdicionarTarefa(descricao string) {
	novaTarefa := preencherNovaTarefaStruct(descricao)
	data.AdicionarNovaTarefaNoArquivo(novaTarefa)
}

func ListarTarefas(isFiltrarTarefas bool) {
	writer := getConfigTabWriter();
	registros := data.BuscarTarefas(isFiltrarTarefas);

	for i := 0; i < len(registros); i++ {
		writer.Write(getLinhaOutput(registros[i]))
	}
	writer.Flush()
}

func CompletarTarefa(id string) {
	data.CompletarTarefa(id)
}

func getConfigTabWriter() *tabwriter.Writer {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	w.Write([]byte("\tID\tDescrição\tEstá Completa\tCriada Em\t\n"))
	return w
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
