package core

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"example.com/csv/src/models"
)

func ReadCsv() []models.Person {
	//abre el documento
	file, err := os.Open("C:\\Users\\gmero\\Documents\\GO\\04-csv\\resources\\SRI_RUC_Azuay(1).csv")

	if err != nil {
		fmt.Println("Error al abrir archivo: ", err)
		return nil
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '|'

	//leer un lector
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("ERROR al leer archivo CSV:", err)
		return nil
	}

	//crear una variable de una array de tipo person
	var people []models.Person

	for i, row := range records {
		if i == 0 {
			continue
		}

		numero_est, _ := strconv.ParseInt(row[11], 10, 64)

		person := models.Person{
			Ruc:                       row[0],
			RazonSocial:               row[1],
			CodigoJuridiccion:         row[2],
			EstadoContribuyente:       row[3],
			ClaseContribuyente:        row[4],
			FechaInicioActividades:    row[5],
			FechaActualización:        row[6],
			FechaSuspencionDefinitiva: row[7],
			FechaReinicioActividades:  row[8],
			Obligado:                  row[9],
			TipoContribuyente:         row[10],
			NumeroEstablecimiento:     int(numero_est),
			NombreFantasiaComercial:   row[12],
			EstadoEstablecimiento:     row[13],
			DescripcionProvinciaEst:   row[14],
			DescripcionCantonEst:      row[15],
			DescripcionParroquiaEst:   row[16],
			CodigoCIIU:                row[17],
			ActividadEconomica:        row[18],
			AgenteRetencion:           row[19],
			Especial:                  row[20],
		}

		people = append(people, person)
	}

	return people

}
