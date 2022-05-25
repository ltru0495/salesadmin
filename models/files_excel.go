package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tealeg/xlsx"
)

func InventoryFile(products []Product) *xlsx.File {
	headers := []string{"Nº", "FECHA DE REGISTRO", "CODIGO", "MARCA", "TALLA", "MODELO", "PRECIO DE COSTO", "UBICACION", "COMENTARIO"}
	fields := []string{"regdate", "code", "brand", "size", "model", "price", "location", "note"}

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, err := file.AddSheet("INVENTARIO")
	if err != nil {
		log.Println(err)
	}
	sheet.SetColWidth(0, 0, 3)
	sheet.SetColWidth(1, 1, 13)
	sheet.SetColWidth(2, 2, 18)
	sheet.SetColWidth(3, 3, 13)
	sheet.SetColWidth(4, 4, 8)
	sheet.SetColWidth(5, 8, 15)

	// HEADER ROW
	row = sheet.AddRow()
	row.SetHeight(32.0)
	for _, v := range headers {
		hStyle := headersStyle()
		cell = row.AddCell()
		if v == "FECHA DE REGISTRO" {
			highlight(hStyle)
		}
		cell.SetStyle(hStyle)
		cell.Value = v
	}

	flag := false
	for k, product := range products {
		p := product.ToMap()
		row = sheet.AddRow()

		cStyle := contentStyle(flag)
		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", (k + 1))
		cell.SetStyle(cStyle)

		for _, val := range fields {
			cStyle := contentStyle(flag)
			if val == "regdate" {
				highlight(cStyle)
			}
			cell = row.AddCell()
			cell.SetStyle(cStyle)
			cell.Value = p[val]
		}

		flag = !flag
	}

	return file
}

// N°	FECHA DE REGISTRO	FECHA DE VENTA	HORA DE VENTA	CODIGO	MARCA	TALLA	MODELO	PRECIO DE COSTO	UBICACION	VENDEDOR	LUGAR DE VENTA	PRECIO DE VENTA	FORMA DE PAGO	COMENTARIO
func SaleFile(sales []Sale) *xlsx.File {
	headers := []string{"Nº", "FECHA DE REGISTRO", "FECHA DE VENTA", "HORA DE VENTA", "CODIGO", "MARCA", "TALLA", "MODELO", "PRECIO DE COSTO", "UBICACION", "VENDEDOR", "LUGAR DE VENTA", "PRECIO DE VENTA", "FORMA DE PAGO", "COMENTARIO"}
	fields := []string{"regdate", "timestamp", "time", "code", "brand", "size", "model", "pricebuy", "location", "seller", "place", "price", "payment_method", "note"}

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, err := file.AddSheet("INVENTARIO")
	if err != nil {
		log.Println(err)
	}
	sheet.SetColWidth(0, 0, 5)
	sheet.SetColWidth(1, 1, 13)
	sheet.SetColWidth(2, 2, 18)
	sheet.SetColWidth(3, 3, 13)
	sheet.SetColWidth(4, 5, 15)
	sheet.SetColWidth(5, 5, 10)
	sheet.SetColWidth(5, 13, 13)
	sheet.SetColWidth(13, 14, 15)

	// HEADER ROW
	row = sheet.AddRow()
	row.SetHeight(32.0)
	for _, v := range headers {
		hStyle := headersStyle()
		cell = row.AddCell()

		cell.SetStyle(hStyle)
		cell.Value = v
	}

	flag := false
	for k, sale := range sales {

		if sale.Price < 0 {
			continue
		}
		p := sale.ToMap()
		row = sheet.AddRow()

		cStyle := contentStyle(flag)
		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", (k + 1))
		cell.SetStyle(cStyle)

		for _, field := range fields {
			cStyle := contentStyle(flag)
			cell = row.AddCell()
			cell.SetStyle(cStyle)
			if field == "price" || field == "pricebuy" {
				floatV, err := strconv.ParseFloat(p[field], 64)
				if err != nil {
					log.Println(err)
				}
				cell.SetFloatWithFormat(floatV, "###.00")
			} else {
				cell.Value = p[field]
			}
		}

		flag = !flag
	}

	return file
}
