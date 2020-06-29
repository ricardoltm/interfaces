package main

import (
	"encoding/json"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	rg := RG{
		Numero:             "376817884",
		OrgaoEmissor:       "SSP",
		UF:                 "SP",
		DataExpedicao:      time.Now().AddDate(-1, 0, 0),
		NaturalidadeEstado: "SP",
		NaturalidadeCidade: "São Paulo",
		Filiacao1:          "Margie Simpson",
		Filiacao2:          "Homer Simpson",
		Comprovantes: []Voucher{
			{
				Type: "frente",
				File: "/img/teste.jpg",
			},
			{
				Type: "verso",
				File: "/img/teste.jpg",
			},
		},
	}

	cpf := CPF{
		Numero: "25036156005",
		Comprovantes: []Voucher{
			{
				Type: "frente",
				File: "/img/teste.jpg",
			},
		},
	}

	store := MemoryStore{}

	err := store.Save(rg)
	if err != nil {
		panic(err)
	}

	err = store.Save(cpf)
	if err != nil {
		panic(err)
	}

	docs := store.List()

	bytes, err := json.Marshal(docs)
	if err != nil {
		panic(err)
	}

	specs := []DocumentSpec{}

	err = json.Unmarshal(bytes, &specs)
	if err != nil {
		panic(err)
	}

	spew.Dump(specs)
}
