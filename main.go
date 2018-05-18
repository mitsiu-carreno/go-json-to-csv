package main

import(
	"fmt"
	"io/ioutil"
	"encoding/json"
	models "github.com/mitsiu-carreno/go-json-to-csv/entry"
)

func check(e error){
	if e != nil{
		fmt.Println(e)
		panic(e)
	}
}

type MyEntry struct{
	models.Entry
}

func (n MyEntry) toString() string {
	return toJson(n)
}

func toJson(n interface{}) string {
	bytes, err := json.Marshal(n)
	check(err)

	return string(bytes)
}

func main(){
	entries := getEntries()
	for _, entry := range entries{
		fmt.Println(entry.toString())
	}

	fmt.Println(toJson(entries))
}

func getEntries() []MyEntry{
	raw, err := ioutil.ReadFile("./input/page_test.json")
	check(err)

	var c []MyEntry
	json.Unmarshal(raw, &c)
	return c
}