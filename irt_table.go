package main

import "errors"

type IrtTable struct {
	Escalao      int
	ValorInicial float64
	ValorFinal   float64
	ParcelaFixa  float64
	Taxa         float64
	Excesso      float64
}

func Data() []IrtTable {
	return []IrtTable{
		{Escalao: 1, ValorInicial: 0, ValorFinal: 100000, ParcelaFixa: 0, Taxa: 0.0, Excesso: 0},
		{Escalao: 2, ValorInicial: 100001, ValorFinal: 150000, ParcelaFixa: 0, Taxa: 0.13, Excesso: 100001},
		{Escalao: 3, ValorInicial: 150001, ValorFinal: 200000, ParcelaFixa: 12500, Taxa: 0.16, Excesso: 150001},
		{Escalao: 4, ValorInicial: 200001, ValorFinal: 300000, ParcelaFixa: 31250, Taxa: 0.18, Excesso: 200001},
		{Escalao: 5, ValorInicial: 300001, ValorFinal: 500000, ParcelaFixa: 49250, Taxa: 0.19, Excesso: 300001},
		{Escalao: 6, ValorInicial: 500001, ValorFinal: 1000000, ParcelaFixa: 87250, Taxa: 0.20, Excesso: 500001},
		{Escalao: 7, ValorInicial: 1000001, ValorFinal: 1500000, ParcelaFixa: 187249, Taxa: 0.21, Excesso: 1000001},
		{Escalao: 8, ValorInicial: 1500001, ValorFinal: 2000000, ParcelaFixa: 292249, Taxa: 0.22, Excesso: 1500001},
		{Escalao: 9, ValorInicial: 2000001, ValorFinal: 2500000, ParcelaFixa: 402249, Taxa: 0.23, Excesso: 2000001},
		{Escalao: 10, ValorInicial: 2500001, ValorFinal: 5000000, ParcelaFixa: 517249, Taxa: 0.24, Excesso: 2500001},
		{Escalao: 11, ValorInicial: 5000001, ValorFinal: 10000000, ParcelaFixa: 1117249, Taxa: 0.245, Excesso: 5000001},
		{Escalao: 12, ValorInicial: 10000001, ValorFinal: 0, ParcelaFixa: 2342248, Taxa: 0.25, Excesso: 10000001},
	}
}

func Get(materiaColectavel float64) (IrtTable, error) {
	for _, t := range Data() {
		if materiaColectavel >= float64(t.ValorInicial) && materiaColectavel <= float64(t.ValorFinal) {
			return t, nil
		} else if float64(t.ValorInicial) >= 10000001 {
			return t, nil
		}
	}

	return IrtTable{}, errors.New("not found")
}
