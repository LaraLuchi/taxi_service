//Serviço simulado

package main

import (
	"time"
)

type Corrida struct {
	DataHora time.Time
	Valor    float64
	Destino  string
	Avaliada bool
}

var corridasPorMotorista = make(map[string][]Corrida)

func ResetCorridas() {
	corridasPorMotorista = make(map[string][]Corrida)
}

func AdicionarCorridaFinalizada(motorista string, corrida Corrida) {
	corridasPorMotorista[motorista] = append(corridasPorMotorista[motorista], corrida)
}

func ObterHistoricoDeCorridas(motorista string) []Corrida {
	return corridasPorMotorista[motorista]
}
