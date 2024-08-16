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
	arquivoCsv := getArquivoCsv()
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

func BuscarTarefas(isFiltrarTarefasPendentes bool) [][]string {
	arquivoCsv := getArquivoCsv()
	csvReader := csv.NewReader(arquivoCsv)

	todosOsRegistro, _ := csvReader.ReadAll()

	if(isFiltrarTarefasPendentes) {
		var registrosFiltrados [][]string

		for i := 0; i < len(todosOsRegistro); i++ {
			if todosOsRegistro[i][2] == "Não" {
				registrosFiltrados = append(registrosFiltrados, todosOsRegistro[i])
			}
		}
	
		return registrosFiltrados
	}
	return todosOsRegistro;
}

func CompletarTarefa(id string) {
	arquivoCsvAntigo := getArquivoCsv()
	csvReaderAntigo := csv.NewReader(arquivoCsvAntigo)

	todosOsRegistro, _ := csvReaderAntigo.ReadAll()

	atualizaRegistroCsv(id, todosOsRegistro)
	removeArquivoDadosAntigos(*arquivoCsvAntigo)
	escreveNovoArquivoAtualizado(todosOsRegistro)
}

func escreveNovoArquivoAtualizado(todosOsRegistros [][]string) {
	csvWriterNovo := csv.NewWriter(getArquivoCsv())

	for i := 0; i < len(todosOsRegistros); i++ {
		csvWriterNovo.Write(todosOsRegistros[i])
	}
	csvWriterNovo.Flush()
}

func removeArquivoDadosAntigos(arquivoCsvAntigo os.File ) {
	arquivoCsvAntigo.Close()
	os.Remove(caminho_arquivo_dados)
}

func atualizaRegistroCsv(idParaAtualizar string, todosOsRegistros [][]string) {
	for i := 0; i < len(todosOsRegistros); i++ {
		if todosOsRegistros[i][0] == idParaAtualizar {
			todosOsRegistros[i][2] = "Sim"
		}
	}
}

func buscarRegistroPeloId(id string) []string {
	arquivoCsv := getArquivoCsv()
	csvReader := csv.NewReader(arquivoCsv)

	todosOsRegistro, _ := csvReader.ReadAll()
	var registroBuscado []string = nil

	for i := 0; i < len(todosOsRegistro); i++ {
		if todosOsRegistro[i][0] == id {
			registroBuscado = todosOsRegistro[i]
		}
	}

	if registroBuscado == nil {
		log.Fatal("Não foi encontrado registro com o ID passado!")
	}
	return registroBuscado
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
