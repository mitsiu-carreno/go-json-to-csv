package main

import(
	"os"
	"io"
	"fmt"
	"html"
	"strconv"
	"io/ioutil"
	"archive/zip"
	"encoding/csv"
	"encoding/json"
	models "github.com/mitsiu-carreno/go-json-to-csv/entry"
)

func check(e error){
	if e != nil{
		fmt.Println(e)
		panic(e)
	}
}

type Entry struct{
	models.Entry
}

type NewPage struct{
	Results 	[]Entry `json:"results"`
	Pagination 	string `json:"-"` 
}

/*
func (n NewPage) toString() string {
	return toJson(n)
}

func toJson(n interface{}) string {
	bytes, err := json.Marshal(n)
	check(err)

	return string(bytes)
}
*/

func main(){

	var outputPath = "./output/"
	var outfileCSVName ="INAI-Solicitudes.csv"
	var outfileZIPName = "INAI-Solicitudes.zip"
	
	// Check if output directory exists
	_, err := os.Stat(outputPath)
	if os.IsNotExist(err){
		os.MkdirAll(outputPath, os.ModePerm)
	}else {
		check(err)
	}

	// Create new file
	outfile, err := os.Create(outputPath + outfileCSVName)
	check(err)
	defer outfile.Close()

	writer := csv.NewWriter(outfile)
	

	// Write headers in new file
	err = writer.Write([]string{"_id","FOLIO","FECHASOLICITUD","DEPENDENCIA","ESTATUS","MEDIOENTRADA","TIPOSOLICITUD","DESCRIPCIONSOLICITUD","OTROSDATOS","ARCHIVOADJUNTOSOLICITUD","MEDIOENTREGA","FECHALIMITE","RESPUESTA","TEXTORESPUESTA","ARCHIVORESPUESTA","FECHARESPUESTA","PAIS","ESTADO","MUNICIPIO","CODIGOPOSTAL","SECTOR"})
	check(err)
	writer.Flush()
	for p:=1; p<=9620; p++{
		page := getPage(p)
		if len(page.Results) != 100{
			fmt.Println("not 100 " + strconv.Itoa(p))
		}
		for _, entry := range page.Results{
			writeCSV(entry, outfile)
		}
	}

}

func getPage(pageNum int) NewPage{
	var inputfile = "/Users/jorandradefig/Desktop/Projects/inai_solicitudes/scrapper/output/page_" + strconv.Itoa(pageNum) +".json"
	raw, err := ioutil.ReadFile(inputfile)
	check(err)

	var c NewPage
	json.Unmarshal(raw, &c)
	return c
}

func writeCSV(entry Entry, file *os.File){
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err := writer.Write([]string{entry.ID, entry.FOLIO, entry.FECHASOLICITUD, html.UnescapeString(entry.DEPENDENCIA), entry.ESTATUS, entry.MEDIOENTRADA, entry.TIPOSOLICITUD, entry.DESCRIPCIONSOLICITUD, entry.OTROSDATOS, entry.ARCHIVOADJUNTOSOLICITUD, entry.MEDIOENTREGA, entry.FECHALIMITE, entry.RESPUESTA, html.UnescapeString(entry.TEXTORESPUESTA), entry.ARCHIVORESPUESTA, entry.FECHARESPUESTA, entry.PAIS, entry.ESTADO, entry.MUNICIPIO, entry.CODIGOPOSTAL, entry.SECTOR})
	check(err)
	err =  writer.Error()
	check(err);

}
