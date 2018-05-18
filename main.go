package main

import(
	"fmt"
	"io/ioutil"
	"encoding/json"
	// models "github.com/mitsiu-carreno/go-json-to-csv/entry"
)

func check(e error){
	if e != nil{
		fmt.Println(e)
		panic(e)
	}
}

type Entry struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Url   string `json:"url"`
}

type Entry3 struct {
    ID    					string `json:"_id"`
	FOLIO					string `json:"FOLIO"`
	FECHASOLICITUD			string `json:"FECHASOLICITUD"`
	DEPENDENCIA				string `json:"DEPENDENCIA"`
	ESTATUS					string `json:"ESTATUS"`
	MEDIOENTRADA			string `json:"MEDIOENTRADA"`
	TIPOSOLICITUD			string `json:"TIPOSOLICITUD"`
	DESCRIPCIONSOLICITUD	string `json:"DESCRIPCIONSOLICITUD"`
	OTROSDATOS				string `json:"OTROSDATOS"`
	ARCHIVOADJUNTOSOLICITUD	string `json:"ARCHIVOADJUNTOSOLICITUD"`
	MEDIOENTREGA			string `json:"MEDIOENTREGA"`
	FECHALIMITE				string `json:"FECHALIMITE"`
	RESPUESTA				string `json:"RESPUESTA"`
	TEXTORESPUESTA			string `json:"TEXTORESPUESTA"`
	ARCHIVORESPUESTA		string `json:"ARCHIVORESPUESTA"`
	FECHARESPUESTA			string `json:"FECHARESPUESTA"`
	PAIS					string `json:"PAIS"`
	ESTADO					string `json:"ESTADO"`
	MUNICIPIO				string `json:"MUNICIPIO"`
	CODIGOPOSTAL			string `json:"CODIGOPOSTAL"`
	SECTOR					string `json:"SECTOR"`
}

func (n Entry) toString() string {
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

func getEntries() []Entry{
	raw, err := ioutil.ReadFile("./input/pages.json")
	check(err)

	var c []Entry
	json.Unmarshal(raw, &c)
	return c
}