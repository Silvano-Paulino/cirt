package cirt

import "errors"

type IrtTable struct {
	Escalao      int
	InitialValue float64
	FinalValue   float64
	ParcelaFixa  float64
	Tax          float64
	Excess       float64
}

func Table() []IrtTable {
	return []IrtTable{
		{Escalao: 1, InitialValue: 0, FinalValue: 100000, ParcelaFixa: 0, Tax: 0.0, Excess: 0},
		{Escalao: 2, InitialValue: 100001, FinalValue: 150000, ParcelaFixa: 0, Tax: 0.13, Excess: 100001},
		{Escalao: 3, InitialValue: 150001, FinalValue: 200000, ParcelaFixa: 12500, Tax: 0.16, Excess: 150001},
		{Escalao: 4, InitialValue: 200001, FinalValue: 300000, ParcelaFixa: 31250, Tax: 0.18, Excess: 200001},
		{Escalao: 5, InitialValue: 300001, FinalValue: 500000, ParcelaFixa: 49250, Tax: 0.19, Excess: 300001},
		{Escalao: 6, InitialValue: 500001, FinalValue: 1000000, ParcelaFixa: 87250, Tax: 0.20, Excess: 500001},
		{Escalao: 7, InitialValue: 1000001, FinalValue: 1500000, ParcelaFixa: 187249, Tax: 0.21, Excess: 1000001},
		{Escalao: 8, InitialValue: 1500001, FinalValue: 2000000, ParcelaFixa: 292249, Tax: 0.22, Excess: 1500001},
		{Escalao: 9, InitialValue: 2000001, FinalValue: 2500000, ParcelaFixa: 402249, Tax: 0.23, Excess: 2000001},
		{Escalao: 10, InitialValue: 2500001, FinalValue: 5000000, ParcelaFixa: 517249, Tax: 0.24, Excess: 2500001},
		{Escalao: 11, InitialValue: 5000001, FinalValue: 10000000, ParcelaFixa: 1117249, Tax: 0.245, Excess: 5000001},
		{Escalao: 12, InitialValue: 10000001, FinalValue: 0, ParcelaFixa: 2342248, Tax: 0.25, Excess: 10000001},
	}
}

func Get(colletableMaterial float64) (IrtTable, error) {
	for _, t := range Table() {
		if colletableMaterial >= float64(t.InitialValue) && colletableMaterial <= float64(t.FinalValue) {
			return t, nil
		} else if float64(t.InitialValue) >= 10000001 {
			return t, nil
		}
	}

	return IrtTable{}, errors.New("not found")
}
