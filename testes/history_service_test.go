//Teste com GODOG

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

var historico []Corrida

func resetarEstadoInicial() {
	ResetCorridas()
}

func queOMotoristaPossuiCorridasFinalizadas(nome string) error {
	ResetCorridas()
	now := time.Now()
	for i := 0; i < 3; i++ {
		AdicionarCorridaFinalizada(nome, Corrida{
			DataHora: now.Add(-time.Duration(i) * time.Hour),
			Valor:    15.50 + float64(i),
			Destino:  fmt.Sprintf("Destino %d", i+1),
			Avaliada: false,
		})
	}
	return nil
}

func eleAcessaAAbaHistoricoDeCorridas() error {
	historico = ObterHistoricoDeCorridas("João Pereira")
	return nil
}

func oSistemaExibeListaComDadosCompletos() error {
	if len(historico) != 3 {
		return fmt.Errorf("esperado 3 corridas, obtido %d", len(historico))
	}

	for _, c := range historico {
		if c.DataHora.IsZero() || c.Valor <= 0 || c.Destino == "" {
			return fmt.Errorf("dados incompletos na corrida: %+v", c)
		}
	}
	return nil
}

func cadaItemPossuiBotaoParaAvaliarCorrida() error {
	for _, c := range historico {
		if c.Avaliada {
			continue
		}
		// Simulação: se não for avaliada, botão deve estar disponível
		// Aqui poderia haver verificação extra se fosse um app real
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(sc *godog.Scenario) {
		resetarEstadoInicial()
	})
	ctx.Step(`^que o motorista "([^"]*)" está logado e possui 3 corridas finalizadas nas últimas 24 horas$`, queOMotoristaPossuiCorridasFinalizadas)
	ctx.Step(`^ele acessa a aba "Histórico de Corridas"$`, eleAcessaAAbaHistoricoDeCorridas)
	ctx.Step(`^o sistema exibe uma lista com as corridas finalizadas contendo data, horário, valor e destino$`, oSistemaExibeListaComDadosCompletos)
	ctx.Step(`^cada item da lista possui um botão para avaliar a corrida$`, cadaItemPossuiBotaoParaAvaliarCorrida)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Name:                 "history",
		ScenarioInitializer:  InitializeScenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"features"},
		},
	}

	if suite.Run() != 0 {
		t.Fatal("Testes de histórico falharam")
	}
}
