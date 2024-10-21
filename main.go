package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
)

// Definimos la estructura XML con las etiquetas correspondientes
type RDE struct {
	XMLName xml.Name `xml:"rDE"`
	DE      DE       `xml:"DE"`
}

type DE struct {
	GOpeDE     GOpeDE     `xml:"gOpeDE"`
	GTimb      GTimb      `xml:"gTimb"`
	GDatGralOpe GDatGralOpe `xml:"gDatGralOpe"`
	GDtipDE    GDtipDE    `xml:"gDtipDE"`
}

type GOpeDE struct {
	ITipEmi int `xml:"iTipEmi"`
}

type GTimb struct {
	ITiDE   int    `xml:"iTiDE"`
	DNumTim int    `xml:"dNumTim"`
	DEst    string `xml:"dEst"`
	DPunExp string `xml:"dPunExp"`
	DNumDoc int    `xml:"dNumDoc"`
}

type GDatGralOpe struct {
	DFeEmiDE string `xml:"dFeEmiDE"`
	GOpeCom  GOpeCom `xml:"gOpeCom"`
	GEmis    GEmis   `xml:"gEmis"`
	GDatRec  GDatRec `xml:"gDatRec"`
}

type GOpeCom struct {
	ITipTra  int    `xml:"iTipTra"`
	ITImp    int    `xml:"iTImp"`
	CMoneOpe string `xml:"cMoneOpe"`
}

type GEmis struct {
	DRucEm   int    `xml:"dRucEm"`
	DDVEmi   int    `xml:"dDVEmi"`
	ITipCont int    `xml:"iTipCont"`
}

type GDatRec struct {
	INatRec   int    `xml:"iNatRec"`
	ITiOpe    int    `xml:"iTiOpe"`
	CPaisRec  string `xml:"cPaisRec"`
	DNomRec   string `xml:"dNomRec"`
	ITipIDRec int    `xml:"iTipIDRec"`
	DNumIDRec int    `xml:"dNumIDRec"`
}

type GDtipDE struct {
	GCamFE   GCamFE   `xml:"gCamFE"`
	GCamCond GCamCond `xml:"gCamCond"`
	GCamItem GCamItem `xml:"gCamItem"`
}

type GCamFE struct {
	IIndPres int `xml:"iIndPres"`
}

type GCamCond struct {
	ICondOpe int       `xml:"iCondOpe"`
	GPaConEIni GPaConEIni `xml:"gPaConEIni"`
}

type GPaConEIni struct {
	ITiPago    int     `xml:"iTiPago"`
	DMonTiPag  float64 `xml:"dMonTiPag"`
	CMoneTiPag string  `xml:"cMoneTiPag"`
}

type GCamItem struct {
	DCodInt       string  `xml:"dCodInt"`
	DDesProSer    string  `xml:"dDesProSer"`
	DCantProSer   int     `xml:"dCantProSer"`
	GCamIVA       GCamIVA `xml:"gCamIVA"`
	GValorItem    GValorItem `xml:"gValorItem"`
}

type GCamIVA struct {
	IAfecIVA    int     `xml:"iAfecIVA"`
	DTasaIVA    int     `xml:"dTasaIVA"`
	DBasGravIVA int     `xml:"dBasGravIVA"`
	DLiqIVAItem int     `xml:"dLiqIVAItem"`
}

type GValorItem struct {
	DPUniProSer     float64 `xml:"dPUniProSer"`
	DTotBruOpeItem  float64 `xml:"dTotBruOpeItem"`
	GValorRestaItem GValorRestaItem `xml:"gValorRestaItem"`
}

type GValorRestaItem struct {
	DDescItem         int     `xml:"dDescItem"`
	DAntGloPreUniIt   int     `xml:"dAntGloPreUniIt"`
	DTotOpeItem       float64 `xml:"dTotOpeItem"`
}

func main() {
	// Abrir archivo CSV
	csvFile, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo CSV:", err)
		return
	}
	defer csvFile.Close()

	// Crear un nuevo lector CSV
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return
	}

	// Crear un nuevo archivo CSV de salida
	outputFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error al crear el archivo CSV de salida:", err)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Escribir encabezado (agregando la columna para el XML)
	header := append(records[0], "xml")
	if err := writer.Write(header); err != nil {
		fmt.Println("Error al escribir el encabezado en CSV:", err)
		return
	}

	// Procesar registros (omitimos la primera fila que es el encabezado)
	for _, record := range records[1:] {
		_, _ = strconv.Atoi(record[0])
		fecha := record[1]
		nombre := record[2]
		numID, _ := strconv.Atoi(record[3])
		total, _ := strconv.ParseFloat(record[4], 64)

		// Crear una estructura con los valores extraídos
		rde := RDE{
			DE: DE{
				GOpeDE: GOpeDE{
					ITipEmi: 1,
				},
				GTimb: GTimb{
					ITiDE:   1,
					DNumTim: 15674904,
					DEst:    "001",
					DPunExp: "001",
					DNumDoc: 777780,
				},
				GDatGralOpe: GDatGralOpe{
					DFeEmiDE: fecha,
					GOpeCom: GOpeCom{
						ITipTra:  2,
						ITImp:    1,
						CMoneOpe: "PYG",
					},
					GEmis: GEmis{
						DRucEm:   80021477,
						DDVEmi:   3,
						ITipCont: 2,
					},
					GDatRec: GDatRec{
						INatRec:   2,
						ITiOpe:    2,
						CPaisRec:  "PRY",
						DNomRec:   nombre,
						ITipIDRec: 1,
						DNumIDRec: numID,
					},
				},
				GDtipDE: GDtipDE{
					GCamFE: GCamFE{
						IIndPres: 2,
					},
					GCamCond: GCamCond{
						ICondOpe: 1,
						GPaConEIni: GPaConEIni{
							ITiPago:    1,
							DMonTiPag:  total,
							CMoneTiPag: "PYG",
						},
					},
					GCamItem: GCamItem{
						DCodInt:     "REV0002",
						DDesProSer:  "Intereses de prestamo activo",
						DCantProSer: 1,
						GCamIVA: GCamIVA{
							IAfecIVA:    1,
							DTasaIVA:    10,
							DBasGravIVA: 2535,
							DLiqIVAItem: 254,
						},
						GValorItem: GValorItem{
							DPUniProSer:    total,
							DTotBruOpeItem: total,
							GValorRestaItem: GValorRestaItem{
								DDescItem:       0,
								DAntGloPreUniIt: 0,
								DTotOpeItem:     total,
							},
						},
					},
				},
			},
		}

		// Convertir la estructura a XML
		xmlData, err := xml.MarshalIndent(rde, "", "  ")
		if err != nil {
			fmt.Println("Error al generar el XML:", err)
			return
		}

		// Agregar el XML como una columna
		newRecord := append(record, string(xmlData))

		// Escribir el nuevo registro en el archivo de salida
		if err := writer.Write(newRecord); err != nil {
			fmt.Println("Error al escribir en CSV:", err)
			return
		}
	}

	fmt.Println("Archivo CSV con columna XML generado con éxito.")
}
