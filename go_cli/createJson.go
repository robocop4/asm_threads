package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Node struct {
	ID       string   `json:"id"`
	Data     Data     `json:"data"`
	Position Position `json:"position"`
	Type     string   `json:"type"`
	// SourcePosition Position `json:"sourcePosition"`
	// TargetPosition Position `json:"targetPosition"`
}

type Data struct {
	Label string `json:"label"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

var x = 0
var y = 0

func CreateJson(id string, name string) Node {

	x = x + 250

	node := Node{
		ID: id,
		Data: Data{
			Label: name,
		},
		Position: Position{
			X: x,
			Y: x,
		},
		Type: "bidirectional",
		// SourcePosition: Position{
		// 	X: 0, // Здесь нужно указать нужные вам координаты
		// 	Y: 0, // Здесь нужно указать нужные вам координаты
		// },
		// TargetPosition: Position{
		// 	X: 0, // Здесь нужно указать нужные вам координаты
		// 	Y: 0, // Здесь нужно указать нужные вам координаты
		// },
	}

	return node
}

type Edge struct {
	ID           string    `json:"id"`
	Step         int       `json:"step"`
	Source       string    `json:"source"`
	Target       string    `json:"target"`
	Type         string    `json:"type"`
	Animated     bool      `json:"animated"`
	SourceHandle string    `json:"sourceHandle"`
	TargetHandle string    `json:"targetHandle"`
	MarkerEnd    MarkerEnd `json:"markerEnd"`
	Style        Style     `json:"style"`
}

type MarkerEnd struct {
	Type string `json:"type"`
}

type Style struct {
	StrokeWidth int    `json:"strokeWidth"`
	Stroke      string `json:"stroke"`
}

func createEdges(from string, to string, step int, color string) Edge {
	uniq, _ := uuid.NewRandom()
	// Создаем новый элемент Edge
	newEdge := Edge{
		ID:           uniq.String(),
		Step:         step,
		Source:       from,
		Target:       to,
		Type:         "bidirectional",
		Animated:     false,
		SourceHandle: "left",
		TargetHandle: "right",
		// MarkerEnd: MarkerEnd{
		// 	Type: "MarkerType.Arrow",
		// },
		Style: Style{
			StrokeWidth: 2,
			Stroke:      color,
		},
	}

	return newEdge
}

func saveJson(name string, jsonF interface{}) {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(jsonF)
	if err != nil {
		panic(err)
	}

}

// Создаем структуры данных, которые соответствуют структуре JSON.

type CallRef struct {
	Addr uint64 `json:"addr"`
	Type string `json:"type"`
	At   uint64 `json:"at"`
}

type Function struct {
	Offset uint64 `json:"offset"`
	Name   string `json:"name"`
	Size   uint64 `json:"size"`
	IsPure string `json:"is-pure"`
	// Добавьте остальные поля из вашего JSON
	CallRefs []CallRef `json:"callrefs"`
	DataRefs []uint64  `json:"datarefs"`
}

func parseJsonFunctionsList(jsonString string) []Function {

	// Создаем слайс для хранения данных.
	var functions []Function

	// Распарсиваем JSON.
	err := json.Unmarshal([]byte(jsonString), &functions)
	if err != nil {
		fmt.Println("Ошибка при распарсивании JSON:", err)
	}

	return functions

}

type IEStruct struct {
	Vaddr  int64  `json:"vaddr"`
	Paddr  int64  `json:"paddr"`
	Baddr  int64  `json:"baddr"`
	Laddr  int64  `json:"laddr"`
	Hvaddr int64  `json:"hvaddr"`
	Haddr  int64  `json:"haddr"`
	Type   string `json:"type"`
}

func jsonIE(jsonString string) []IEStruct {

	var data []IEStruct

	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
	}

	return data

}
