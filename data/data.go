package data

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"to-go-list/model"
)

const caminho_arquivo_dados = "C:/GO/to-go-list/data/dados.csv"

func AdicionarNovaTarefaNoArquivo(novaTarefa *model.Tarefa) {
	arquivoCsv := getArquivoCsv();
	csvWriter := csv.NewWriter(arquivoCsv)
	

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

func BuscarTarefasPendentes() [][]string{
	arquivoCsv := getArquivoCsv();
	csvReader := csv.NewReader(arquivoCsv)

	todosOsRegistro, _ := csvReader.ReadAll()
	var registrosFiltrados [][]string;

	for i := 0; i < len(todosOsRegistro); i++ {
		if (todosOsRegistro[i][2] == "Não") {
			registrosFiltrados = append(registrosFiltrados, todosOsRegistro[i])
		}
	}

	return registrosFiltrados
}

func getArquivoCsv() *os.File {
	file, err := os.OpenFile(caminho_arquivo_dados, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func GetProximoIdRegistro() string {
	csvReader := csv.NewReader(getArquivoCsv())

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