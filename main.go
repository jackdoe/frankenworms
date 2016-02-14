package main

import (
	"encoding/csv"
	"fmt"
	gv "github.com/awalterschulze/gographviz"
	parser "github.com/awalterschulze/gographviz/parser"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var ID_TO_NEURON = map[string]*neuron{}
var BODY = newBody()
var DEBUG = false

func rcsv(name string, cb func([]string)) {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	r.Comma = ';'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		cb(record)
	}
}

func main() {
	rcsv("data/neurons.csv", func(records []string) {
		newEmptyNeuron(records[0], records, 5, BODY)
	})

	rcsv("data/connectome.csv", func(records []string) {
		weight, err := strconv.ParseInt(records[3], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		ID_TO_NEURON[records[0]].connect(ID_TO_NEURON[records[1]], int32(weight))
	})

	rcsv("data/fake_motor.csv", func(records []string) {
		weight, err := strconv.ParseInt(records[3], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		if records[1] == "LEFT" {
			ID_TO_NEURON[records[0]].connect(BODY.left, int32(weight))
		} else if records[1] == "RIGHT" {
			ID_TO_NEURON[records[0]].connect(BODY.right, int32(weight))
		}
	})

	for _, v := range ID_TO_NEURON {
		v.ticker()
	}
	BODY.ticker()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.RawQuery
		if n, ok := ID_TO_NEURON[r.URL.RawQuery]; ok {
			n.ping(30)
			fmt.Fprintf(w, "nReceived: %d", n.nReceived)
		} else {
			http.Error(w, fmt.Sprintf("missing neuron: %s", q), 404)
		}
	})

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		DEBUG = !DEBUG
		fmt.Fprintf(w, "current: %#v", DEBUG)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		graphAst, _ := parser.ParseString(`digraph G {}`)
		graph := gv.NewGraph()
		graph.SetDir(true)
		for _, v := range ID_TO_NEURON {

			if v.activity() > 0 {
				graph.AddNode("G", v.id(), nil)
				for _, c := range v.connections {
					graph.AddNode("G", c.to.id(), nil)
					graph.AddEdge(v.id(), c.to.id(), true, nil)
				}
			}
		}
		gv.Analyse(graphAst, graph)
		output := graph.String()
		fmt.Fprintf(w, output)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
