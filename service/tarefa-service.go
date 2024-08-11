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
	writer := getConfigTabWriter();
	registros := data.BuscarTarefasPendentes();

	for i := 0; i < len(registros); i++ {
		writer.Write(getLinhaOutput(registros[i]))
	}
	writer.Flush()
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
