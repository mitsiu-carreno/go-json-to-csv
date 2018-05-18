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

func (p Entry) toString() string{
	return toJson(p)
}

func toJson(p interface{}) string{
	bytes, err := json.Marshal(p)
	check(err)
	return string(bytes)
} 

func main(){
	files := getFiles()
	for _, p := range files{
		fmt.Println(p.toString())
	}

	fmt.Println(toJson(files))
}

func getFiles() []Entry {
	raw, err := ioutil.ReadFile("/Users/jorandradefig/Desktop/Projects/inai_solicitudes/scrapper/output/page_1.json")
	check(err)

	var c[]Entry
	json.Unmarshal(raw, &c)
	return c
}