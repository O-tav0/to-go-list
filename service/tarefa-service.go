package service

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
	"to-go-list/model"
	// "fmt"
)

const caminho_arquivo_dados = "C:/GO/to-go-list/data/dados.csv"

func AdicionarTarefa(descricao string) {
	dados := getArquivoCsv()
	csvWriter := csv.NewWriter(dados)
	novaTarefa := preencherNovaTarefaStruct(descricao, dados)

	csvWriter.Write([]string{
		strconv.Itoa(novaTarefa.Id),
		novaTarefa.Descricao,
		model.ConverterParaPortugues(novaTarefa.EstaFinalizada),
		novaTarefa.CriadaEm,
	})
	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		log.Fatalf("Erro ao realizar o flush: %v", err)
	}
}

func ListarTarefasPendentes() {
	dados := getArquivoCsv()
	csvReader := csv.NewReader(dados)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)

	registros, _ := csvReader.ReadAll()

	w.Write([]byte("\tID\tDescrição\tEstá Completa\tCriada Em\t\n"))
	for i := 0; i < len(registros); i++ {
		if(registros[i][2] == "Não") {
			w.Write(getLinhaOutput(registros[i]))
		}
	}

	w.Flush()
}

func getLinhaOutput(registro []string) []byte {
	linhaOutput := "\t" + registro[0] + "\t" + registro[1] + "\t" + registro[2] + "\t" + registro[3] + "\t"+"\n"
	return []byte(linhaOutput)
}

func getArquivoCsv() *os.File {
	file, err := os.OpenFile(caminho_arquivo_dados, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func getProximoIdRegistro(arquivoDados *os.File) string {
	csvReader := csv.NewReader(arquivoDados)

	registros, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Erro ao ler o arquivo de dados para recuperar último ID!")
	}

	if len(registros) == 0 {
		return "1"
	}

	ultimoRegistro := registros[len(registros)-1]
	ultimoIdRegistrado, _ := strconv.Atoi(ultimoRegistro[0])

	return strconv.Itoa(ultimoIdRegistrado + 1)
}

func preencherNovaTarefaStruct(descricao string, arquivoDados *os.File) *model.Tarefa {
	id, err := strconv.Atoi(getProximoIdRegistro(arquivoDados))
	if err != nil {
		log.Fatal("Falha ao converter o número doúltimo Id")
	}

	agora := time.Now()

	novaTarefa := &model.Tarefa{
		Id:             id,
		Descricao:      descricao,
		EstaFinalizada: false,
		CriadaEm:      agora.Format("02/01/2006"),
	}

	return novaTarefa
}
