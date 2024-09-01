package main

import "flag"

func cli(salaryBase, subSidioA, subSidioT, premeo float64, late, days int) {
	flag.Float64Var(&salaryBase, "s", 0, "salary base")
	flag.Float64Var(&subSidioA, "a", 0, "subsidio alimentacao")
	flag.Float64Var(&subSidioT, "t", 0, "subsidio transporte")
	flag.Float64Var(&premeo, "p", 0, "premeo")
	flag.IntVar(&late, "f", 0, "lates")
	flag.IntVar(&days, "d", 22, "days utils")
	flag.Parse()
}
