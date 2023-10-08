package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"

	r2pipe "github.com/radareorg/r2pipe-go"
)

type Item struct {
	Offset int64  `json:"offset"`
	Size   int    `json:"size"`
	Bytes  string `json:"bytes"`
	Type   string `json:"type"`
}

func checkInvalid(currentAddress string) bool {
	var items []Item
	err := json.Unmarshal([]byte(currentAddress), &items)
	if err != nil {
		fmt.Println(err.Error())
	}
	if items[0].Type == "invalid" {
		return false
	}
	return true
}

// func containsString(arr []string, target string) bool {
// 	for _, s := range arr {
// 		if s == target {
// 			return true
// 		}
// 	}
// 	return false
// }

func userInput(r2p *r2pipe.Pipe) {

	var userInput string

	fmt.Print("Enter something: ")
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	fmt.Println("You entered:", userInput)

	currentAddress, err := r2p.Cmd(userInput)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("[+]: %s\n", currentAddress)

}

func calcOffset(offset int64, baseadr string) string {

	result, _ := strconv.ParseInt(baseadr, 16, 64)
	hexstring := fmt.Sprintf("0x%X", result+offset)

	return hexstring
}

func calcOffset2(addr string, baseaddr string) int64 {
	// Конвертируем шестнадцатеричные строки в целые числа
	num1, _ := strconv.ParseInt(addr, 0, 64)
	num2, _ := strconv.ParseInt(baseaddr, 0, 64)
	result := num2 - num1
	return result
}

func main() {

	// Определение флагов с помощью функций flag
	// Первый аргумент - имя флага, второй - значение по умолчанию, третий - описание
	// Примеры:
	// -name=John
	// -age=30
	// -verbose

	// Определение флагов
	pFlag := flag.String("p", "", "Continue execution if set")
	argFlag := flag.String("arg", "", "Optional argument")
	secondThread := flag.Bool("st", false, "Optional argument")

	// Парсинг аргументов командной строки
	flag.Parse()

	// Если флаг -p установлен, продолжаем выполнение
	if *pFlag == "" {
		// Если флаг -p не установлен, выводим справку
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1) // Выход с кодом ошибки
	}

	//check arg program
	if *argFlag != "" {
		fmt.Printf("Optional argument -arg: %s\n", *argFlag)
	}

	colorKey := "first"
	fileName := "edges1.json"
	if *secondThread {
		colorKey = "second"
		fileName = "edges2.json"
	}
	color := map[string]string{
		"first":  "rgba(255, 0, 0, 0.5)",
		"second": "rgba(10, 0, 255, 0.5)",
	}

	jsonArray := []Node{}
	breackpoints := []string{}
	trace := []string{}
	edges := []Edge{}
	r2p, err := r2pipe.NewPipe(*pFlag)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	defer r2p.Close()

	// Передайте аргумент программе

	// Запускаем команду Radare2 для анализа функций
	r2p.Cmd("aaaa")

	//base_static, _ := r2p.Cmd("s")

	functions, err := r2p.Cmd("aflj")
	if err != nil {
		fmt.Println(err.Error())
	}

	functionsList := parseJsonFunctionsList(functions)

	//add breackpoint
	for _, substr := range functionsList {

		//offset is base addr in dec
		offset := substr.Offset
		base_addr_func := fmt.Sprintf("0x%X", offset)
		//fmt.Println(base_addr_func)

		// add BAF to breackpoints array
		breackpoints = append(breackpoints, base_addr_func)

		//dbte <addr> Enable Breakpoint Trace
		//fmt.Println(calcOffset(int64(offset), base))
		_, errr := r2p.Cmd("db " + base_addr_func)
		if errr != nil {
			fmt.Println(errr.Error())
		}

		jsonArray = append(jsonArray, CreateJson(base_addr_func, substr.Name))

		//fmt.Println(id)

	}

	saveJson("base.json", jsonArray)

	combinedStr := fmt.Sprintf("ood %s", *argFlag)
	_, err = r2p.Cmd(combinedStr)
	if err != nil {
		fmt.Println(err.Error())
	}

	base_dinamic, _ := r2p.Cmd("iej")
	iej := jsonIE(base_dinamic)

	//fmt.Println(iej[0].Baddr)

	// jsonData, err := json.MarshalIndent(jsonArray, "", "  ")
	// if err != nil {
	// 	fmt.Println("Ошибка при маршалинге в JSON:", err)
	// 	return
	// }
	// fmt.Println(string(jsonData))
	// fmt.Println(len(substrings))
	// dblist, errdblist := r2p.Cmd("db")
	// if errdblist != nil {
	// 	fmt.Println(errdblist.Error())
	// }
	// fmt.Println(dblist)

	for {

		currentAddress, _ := r2p.Cmd("pdj 1")
		// 	current, _ := r2p.Cmd("pd 1")

		// 	//current, _ := r2p.Cmd("dcr")
		if !checkInvalid(currentAddress) {
			break
		}

		// 	fmt.Println(current)
		_, err := r2p.Cmd("dc")
		if err != nil {
			fmt.Println(err.Error())
		}
		jSonPD, _ := r2p.Cmd("pdj 1")

		var items []Item
		err = json.Unmarshal([]byte(jSonPD), &items)
		if err != nil {
			fmt.Println(err.Error())
		}
		//fmt.Println(items[0].Offset - iej[0].Baddr)
		base_addr_func := fmt.Sprintf("0x%X", items[0].Offset-iej[0].Baddr)
		trace = append(trace, base_addr_func)
		//trace = append(trace, calcOffset(items[0].Offset, base))

	}

	//fmt.Println(trace)

	// Разбиваем массив на пары элементов
	for i := 0; i < len(trace)-1; i++ {
		//pair := []string{trace[i], trace[i+1]}
		//fmt.Println(pair)
		newEdge := createEdges(trace[i], trace[i+1], i, color[colorKey])
		edges = append(edges, newEdge)

	}

	//    save json

	saveJson(fileName, edges)

	// join two files
	if *secondThread {
		// Чтение содержимого первого JSON-файла
		edges1, err := os.ReadFile("edges1.json")
		if err != nil {
			fmt.Println("Error reading edges1.json:", err)
			return
		}

		// Чтение содержимого второго JSON-файла
		edges2, err := os.ReadFile("edges2.json")
		if err != nil {
			fmt.Println("Error reading edges2.json:", err)
			return
		}

		// Распаковка JSON-данных в массивы
		var items1 []Edge
		if err := json.Unmarshal(edges1, &items1); err != nil {
			fmt.Println("Error unmarshalling data from file1.json:", err)
			return
		}

		var items2 []Edge
		if err := json.Unmarshal(edges2, &items2); err != nil {
			fmt.Println("Error unmarshalling data from file2.json:", err)
			return
		}

		combinedItems := append(items1, items2...)
		saveJson(fileName, combinedItems)

	}

	//jsonEdges, _ := json.MarshalIndent(edges, "", "  ")

	//fmt.Println(string(jsonEdges))

	// for {
	// 	userInput(r2p)
	// }
	//userInput(r2p)

}
