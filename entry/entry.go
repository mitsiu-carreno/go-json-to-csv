package entry

type Entry struct{
	_id						string `json:"_id"`
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