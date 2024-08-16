package data

import (
	"testing"
	"to-go-list/model"
	"time"
	"strconv"
)

func TestCompletarTarefa (t *testing.T) {

	t.Run("Completar uma tarefa existente", func(t *testing.T) {
		CompletarTarefa("1")
		tarefaCompletadaComSucesso := false

		conteudoArquivo := BuscarTarefas(false);

		for _, linha := range conteudoArquivo {
			if(linha[0] == "1" && linha[2] == "Sim") {
				tarefaCompletadaComSucesso = true;
			}
		}
		
		if(!tarefaCompletadaComSucesso) {
			t.Log("A tarefa não foi completada")
			t.Fail()
		}
	})
}

func TestDeletarRegistro (t *testing.T) {

	t.Run("Deletar registro existente", func(t *testing.T) {
		DeletarRegistro("1")
		registroDeletadoComSucesso := true

		conteudoArquivo := BuscarTarefas(false);

		for _, linha := range conteudoArquivo {
			if(linha[0] == "1") {
				registroDeletadoComSucesso = false;
			}
		}
		
		if(!registroDeletadoComSucesso) {
			t.Log("O registro não foi deletado")
			t.Fail()
		}
	})
}

func TestAdicionarNovaTarefaNoArquivo(t *testing.T) {

	t.Run("Adiciona novo registro", func(t *testing.T) {

		agora := time.Now()
		registroInseridoComSucesso := false

		novoRegistro := model.Tarefa {
			Id: 1,
			Descricao: "Descrição teste",
			EstaFinalizada: false,
			CriadaEm: agora.Format("02/01/2006"),
		}

		AdicionarNovaTarefaNoArquivo(&novoRegistro)
		conteudoArquivo := BuscarTarefas(false);

		for _, linha := range conteudoArquivo {
			if(linha[0] == strconv.Itoa(novoRegistro.Id) && linha[1] == novoRegistro.Descricao) {
				registroInseridoComSucesso = true;
			}
		}
		
		if(!registroInseridoComSucesso) {
			t.Log("Não foi encontrado o registro com as informações passadas")
			t.Fail()
		}
	})
}

func TestBuscarTarefas(t *testing.T) {

	t.Run("Execução sem filtro", func(t *testing.T) {
		conteudoArquivo := BuscarTarefas(true);
		apresentaTarefasFinalizadas := false
		for _, linha := range conteudoArquivo {
			if(linha[2] == "Sim") {
				apresentaTarefasFinalizadas = true;
			}
		}

		if(apresentaTarefasFinalizadas) {
			t.Fail()
			t.Log("Execução sem Flag mostra todos os registros")
		}
	})

	t.Run("Execução com a flag -t", func(t *testing.T) {
		conteudoArquivo := BuscarTarefas(false);
		apresentaTarefasFinalizadas := false
		for _, linha := range conteudoArquivo {
			if(linha[2] == "Sim") {
				apresentaTarefasFinalizadas = true;
			}
		}

		if(!apresentaTarefasFinalizadas) {
			t.Fail()
			t.Log("Execução com Flag não mostra todos os registros")
		}
	})

}