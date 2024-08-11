package service

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"to-go-list/model"
	"time"
)

const caminho_arquivo_dados = "C:/GO/to-go-list/data/dados.csv";

func AdicionarTarefa(descricao string) {
	dados := getArquivoCsv();
	csvWriter := csv.NewWriter(dados)
	novaTarefa := preencherNovaTarefaStruct(descricao)

	csvWriter.Write([]string {
		strconv.Itoa(novaTarefa.Id),
		novaTarefa.Descricao,
		model.ConverterParaPortugues(novaTarefa.EstaFinalizada),
		novaTarefa.CriadaEm.String(),
	})
	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		log.Fatalf("Erro ao realizar o flush: %v", err)
	}
}	

func getArquivoCsv() *os.File {
	file, err := os.OpenFile(caminho_arquivo_dados, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func getProximoIdRegistro() string {
	arquivoDados := getArquivoCsv();
	csvReader := csv.NewReader(arquivoDados)

	registros, err := csvReader.ReadAll()
	if(err != nil) {
		log.Fatal("Erro ao ler o arquivo de dados para recuperar último ID!")
	}

	if(len(registros) == 0) {
		return "1"
	}

	ultimoRegistro := registros[len(registros)-1]
	ultimoIdRegistrado, _ := strconv.Atoi(ultimoRegistro[0])

	return strconv.Itoa(ultimoIdRegistrado + 1)
}

func preencherNovaTarefaStruct(descricao string) *model.Tarefa {
	id, err := strconv.Atoi(getProximoIdRegistro())
	if(err != nil) {
		log.Fatal("Falha ao converter o número doúltimo Id")
	}

	novaTarefa := &model.Tarefa {
		Id: id ,
		Descricao: descricao,
		EstaFinalizada: false,
		CriadaEm: time.Now(),
	}

	return novaTarefa
}