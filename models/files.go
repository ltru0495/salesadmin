package models

import (
	"fmt"
	"math"

	"github.com/tealeg/xlsx"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/barcode"

	"strings"
)

var salesTableHeaders = [...]string{"FECHA DE VENTA", "MARCA", "TALLA", "MODELO", "PERTENECE", "VENDEDOR",
	"LUGAR DE VENTA", "PRECIO DE VENTA", "PRECIO DE COMPRA", "GANANCIA", "COMENTARIO", "COMENTARIO PRODUCTO"}
var salesTableUserHeaders = [...]string{"FECHA DE VENTA", "MARCA", "TALLA", "MODELO",
	"PERTENECE", "VENDEDOR", "PRECIO DE VENTA", "COMENTARIO"}
var sellersTableHeaders = [...]string{"NOMBRE DE VENDEDOR", "CANTIDAD VENDIDA"}
var totalTableHeaders = [...]string{"TOTAL VENDIDO", "GANANCIA"}

var inventoryTableHeaders = [...]string{"Nº", "FECHA DE REGISTRO", "COD PROD", "MARCA", "SERIE", "TALLA",
	"MODELO", "PRECIO DE COMPRA", "CANTIDAD", "UBICACION", "NOTA"}

const CODEFONTSIZE = 7
const SIZEFONTSIZE = 36
const MODELX = 26
const WIDTH = 62.0
const HEIGHT = 18.0

func BarcodesFile(products []Product) *gofpdf.Fpdf {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: "mm",
		Size:    gofpdf.SizeType{Wd: WIDTH, Ht: 40.0},
	})

	for _, product := range products {
		LargePage(pdf, product)
	}
	return pdf
}

func WithSizePage(pdf *gofpdf.Fpdf, product Product) {
	pdf.AddPage()
	marginX := 4.0
	marginY := 3.0
	// x, y := marginX, marginY

	barcodeX, barcodeY := marginX, marginY
	barcodeW, barcodeH := 36.0, 10.0

	// w, h := 30.0, 12.0
	pdf.SetFont("Times", "B", CODEFONTSIZE)

	key := barcode.RegisterCode128(pdf, product.Code)
	barcode.Barcode(pdf, key, barcodeX, barcodeY, barcodeW, barcodeH, false)

	codeX, codeY := marginX, marginY+barcodeH+2.0
	pdf.SetFont("Times", "B", CODEFONTSIZE-1)

	pdf.Text(codeX, codeY, product.Code)

	modelX, modelY := barcodeW/2.0+marginX+3.0, codeY
	if len(product.Model) > 10 {
		pdf.SetFont("Times", "B", CODEFONTSIZE-3)
	} else {
		pdf.SetFont("Times", "B", CODEFONTSIZE)
	}
	model := strings.Replace(product.Model, "ñ", "n", -1)
	pdf.Text(modelX, modelY, model)

	sizeX, sizeY := 44.0, 13.0
	pdf.SetFont("Times", "B", SIZEFONTSIZE+3)
	pdf.Text(sizeX, sizeY, fmt.Sprintf("%d", product.Size))
}

func NormalPage(pdf *gofpdf.Fpdf, product Product) {
	pdf.AddPage()
	marginX := 4.0
	marginY := 3.0
	// x, y := marginX, marginY

	barcodeX, barcodeY := marginX, marginY
	barcodeW, barcodeH := 35.0, 10.0
	key := barcode.RegisterCode128(pdf, product.Code)
	barcode.Barcode(pdf, key, barcodeX, barcodeY, barcodeW, barcodeH, false)
	pdf.SetFont("Times", "B", 8)

	codeX, codeY := marginX, marginY+barcodeH+2.5
	pdf.Text(codeX, codeY, product.Code)

	model := strings.Replace(product.Model, "ñ", "n", -1)
	pdf.SetFont("Times", "B", 8)
	x2 := 40.0
	w := 28
	pdf.Text(x2, marginY+2.0, format(w, product.Brand))
	pdf.Text(x2, marginY+12.0, format(w, fmt.Sprintf("%d", product.Size)))

	pdf.SetFont("Times", "B", 6)

	if product.SPrice != 0.00 {
		x := math.Round(product.SPrice)
		pdf.Text(barcodeX+25, barcodeY+barcodeH+2.0, format(20, fmt.Sprintf("%.d", int(x))))
	}

	w = 45
	x2 = 40.0
	modelWords := split(model)
	switch len(modelWords) {
	case 1:
		w = 30
		pdf.SetFont("Times", "B", 8)
		pdf.Text(x2, marginY+7.0, format(w, model))
	case 2:
		w = 38
		pdf.SetFont("Times", "B", 6)

		pdf.Text(x2, marginY+6.0, format(w, modelWords[0]))
		pdf.Text(x2, marginY+8.0, format(w, modelWords[1]))

	default:
		pdf.SetFont("Times", "B", 5)

		pdf.Text(x2, marginY+4.0, format(w, modelWords[0]))
		pdf.Text(x2, marginY+6.5, format(w, modelWords[1]))
		pdf.Text(x2, marginY+9.0, format(w, modelWords[2]))
	}

}

func format(w int, val string) string {
	return fmt.Sprintf(fmt.Sprintf("%%-%ds", w/2), fmt.Sprintf(fmt.Sprintf("%%%ds", w/2), val))
}

func split(str string) []string {
	model := strings.Replace(str, "ñ", "n", -1)

	var words []string
	modelWords := strings.Split(model, " ")

	for _, word := range modelWords {
		if strings.Contains(word, "/") && len(word) > 5 {
			moreWords := strings.Split(word, "/")
			words = append(words, moreWords...)
		} else {
			words = append(words, word)
		}
	}
	return words

}
func LargePage(pdf *gofpdf.Fpdf, product Product) {
	pdf.AddPage()
	marginX := 3.8
	marginY := 3.0
	// x, y := marginX, marginY

	x2 := marginX

	pdf.SetFont("Times", "B", 8)
	brandX := x2
	brandY := marginY + 7.0
	pdf.Text(brandX, brandY, product.Brand)

	barcodeX, barcodeY := 29.0, marginY+2.0
	barcodeW, barcodeH := 28.0, 6.0
	key := barcode.RegisterCode128(pdf, product.Code)
	barcode.Barcode(pdf, key, barcodeX, barcodeY, barcodeW, barcodeH, false)

	pdf.SetFont("Times", "B", 6)
	fCode := format(20, product.Code)
	pdf.Text(barcodeX, barcodeY+barcodeH+2.0, fCode)

	if product.SPrice != 0.00 {
		x := math.Round(product.SPrice)
		pdf.Text(barcodeX+25, barcodeY+barcodeH+2.0, format(20, fmt.Sprintf("%.d", int(x))))
	}

	pdf.SetFont("Times", "B", 65)
	fSize := format(5, fmt.Sprintf("%d", product.Size))
	pdf.Text(32.0, 33.0, fSize)

	pdf.SetFont("Times", "B", 6)
	// codeX, codeY := x2, marginY+30.5
	// pdf.Text(codeX, codeY, product.PFC)
	modelWords := split(product.Model)
	model := product.Model
	marginY = 13.5
	fs := 15.0
	switch len(modelWords) {

	case 1:
		pdf.SetFont("Times", "B", fs)
		pdf.Text(x2, marginY+11.0, model)
	case 2:
		pdf.SetFont("Times", "B", fs)
		pdf.Text(x2, marginY+6.0, modelWords[0])
		pdf.Text(x2, marginY+12.0, modelWords[1])

	default:
		pdf.SetFont("Times", "B", fs)
		pdf.Text(x2, marginY+4.0, modelWords[0])
		pdf.Text(x2, marginY+10.0, modelWords[1])
		pdf.Text(x2, marginY+16.0, modelWords[2])
	}

}

func spanishD(date string) string {
	date = strings.Replace(date, "Jan", "Ene", 1)
	date = strings.Replace(date, "Apr", "Abr", 1)
	date = strings.Replace(date, "Aug", "Ago", 1)
	date = strings.Replace(date, "Dec", "Dic", 1)
	return date
}
func GroupedBarcodesFile(products []Product) *gofpdf.Fpdf {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: "mm",
		Size:    gofpdf.SizeType{Wd: WIDTH, Ht: HEIGHT},
	})

	WithSizePage(pdf, products[0])

	for _, product := range products {
		NormalPage(pdf, product)
		// NormalPage(pdf, product)
		WithSizePage(pdf, product)
	}
	return pdf
}

func BarcodesWithSizeFile(products []Product) *gofpdf.Fpdf {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: "mm",
		Size:    gofpdf.SizeType{Wd: WIDTH, Ht: HEIGHT},
	})
	for _, product := range products {
		NormalPage(pdf, product)

		// WithSizePage(pdf, product)
	}
	return pdf

}

func LastBarcodesFile(products []Product) *gofpdf.Fpdf {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: "mm",
		Size:    gofpdf.SizeType{Wd: WIDTH, Ht: HEIGHT},
	})

	for _, product := range products {
		NormalPage(pdf, product)

	}
	return pdf

}
func GetReportFile(total float64, totalSales float64,
	sellers []Seller, sales []Sale, date string, salePlace string) *xlsx.File {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	hStyle := headersStyle()

	file = xlsx.NewFile()
	sheet, err := file.AddSheet("Ventas")
	if err != nil {
		fmt.Printf(err.Error())
	}
	sheet.SetColWidth(0, 1, 5)
	sheet.SetColWidth(1, 20, 20)

	row = sheet.AddRow() //Empty Row

	if salePlace == "all" {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell = row.AddCell()
		cell.SetStyle(extraStyle())
		cell.Value = "Ventas " + date
		row = sheet.AddRow()
	} else {
		cell = row.AddCell()
		cell = row.AddCell()
		cell.Value = "TIENDA: " + salePlace
		cell.SetStyle(extraStyle())
		cell = row.AddCell()
		cell = row.AddCell()
		cell = row.AddCell()

		cell.SetStyle(extraStyle())
		cell.Value = "FECHA: " + date
		row = sheet.AddRow()

	}

	// **************** TOTAL TABLE ********************
	row = sheet.AddRow()
	row.SetHeight(20.0)

	cell = row.AddCell()
	for _, tth := range totalTableHeaders {
		cell = row.AddCell()
		cell.SetStyle(hStyle)
		cell.Value = tth
	}
	row = sheet.AddRow()

	flag := false
	cell = row.AddCell() //Empty cell
	cell = row.AddCell()

	cStyle := contentStyle(flag)
	cell.SetStyle(cStyle)
	cell.SetFloatWithFormat(totalSales, "###0.00")
	// cell.Value = "S/ " + fmt.Sprintf("%.2f", totalSales)

	cell = row.AddCell()
	cell.SetStyle(cStyle)
	cell.SetFloatWithFormat(total, "###0.00")
	// cell.Value = "S/ " + fmt.Sprintf("%.2f", total)

	row = sheet.AddRow() //Empty Row

	// ********************* SELLERS TABLE **********************
	row = sheet.AddRow()
	row.SetHeight(20.0)
	cell = row.AddCell()
	for _, sth := range sellersTableHeaders {
		cell = row.AddCell()
		cell.SetStyle(hStyle)
		cell.Value = sth
	}
	flag = false
	for _, seller := range sellers {
		row = sheet.AddRow()
		cell = row.AddCell() //Empty cell
		cStyle = contentStyle(flag)

		cell = row.AddCell()
		cell.Value = seller.Name
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", seller.Amount)
		cell.SetStyle(cStyle)
		flag = !flag
	}

	//******************* SALES TABLE ***********************
	row = sheet.AddRow() //Empty row
	// Headings
	row = sheet.AddRow()
	row.SetHeight(20.0)
	cell = row.AddCell()

	for _, sth := range salesTableHeaders {
		cell = row.AddCell()
		cell.SetStyle(hStyle)
		cell.Value = sth
	}
	flag = false

	for k, sale := range sales {
		row = sheet.AddRow()
		cStyle = contentStyle(flag)

		cell = row.AddCell()

		cell.Value = fmt.Sprintf("%d", (k + 1))
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Timestamp.Format("02/01/2006")
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Brand
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", sale.Size)
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Model
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Location
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Seller
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Place
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.SetFloatWithFormat(sale.Price, "###0.00")
		// cell.Value = "S/ " + fmt.Sprintf("%.2f", sale.Price)
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.SetFloatWithFormat(sale.PriceBuy, "###0.00")
		// cell.Value = "S/ " + fmt.Sprintf("%.2f", sale.PriceBuy)
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.SetFloatWithFormat(sale.Earning, "###0.00")
		// cell.Value = "S/ " + fmt.Sprintf("%.2f", sale.Earning)
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Comment
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.PNote
		cell.SetStyle(cStyle)

		flag = !flag
	}

	return file
}

func GetReportFileForUser(sellers []Seller, sales []Sale, date string, place string) *xlsx.File {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	flag := false

	hStyle := headersStyle()
	cStyle := contentStyle(flag)

	file = xlsx.NewFile()
	sheet, err := file.AddSheet("Ventas")
	if err != nil {
		fmt.Printf(err.Error())
	}
	sheet.SetColWidth(0, 1, 5)
	sheet.SetColWidth(1, 20, 20)

	row = sheet.AddRow() //Empty Row

	row = sheet.AddRow()
	cell = row.AddCell()
	cell = row.AddCell()
	cell.Value = "TIENDA: " + place
	cell.SetStyle(extraStyle())
	cell = row.AddCell()
	cell = row.AddCell()
	cell = row.AddCell()

	cell.SetStyle(extraStyle())
	cell.Value = "FECHA " + date
	row = sheet.AddRow()

	// ********************* SELLERS TABLE **********************
	row = sheet.AddRow()
	row.SetHeight(20.0)
	cell = row.AddCell()
	for _, sth := range sellersTableHeaders {
		cell = row.AddCell()
		cell.SetStyle(hStyle)
		cell.Value = sth
	}
	for _, seller := range sellers {
		row = sheet.AddRow()
		cell = row.AddCell() //Empty cell
		cStyle = contentStyle(flag)

		cell = row.AddCell()
		cell.Value = seller.Name
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", seller.Amount)
		cell.SetStyle(cStyle)
		flag = !flag
	}

	//******************* SALES TABLE ***********************
	row = sheet.AddRow() //Empty row
	// Headings
	row = sheet.AddRow()
	row.SetHeight(20.0)
	cell = row.AddCell()
	for _, sth := range salesTableUserHeaders {
		cell = row.AddCell()
		cell.SetStyle(hStyle)
		cell.Value = sth
	}
	flag = false

	for k, sale := range sales {
		row = sheet.AddRow()
		cStyle = contentStyle(flag)
		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", (k + 1))
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Timestamp.Format("02/01/2006")
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Brand
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", sale.Size)
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Model
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Location
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Seller
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.SetFloatWithFormat(sale.Price, "###0.00")
		// cell.Value = "S/ " + fmt.Sprintf("%.2f", sale.Price)
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = sale.Comment
		cell.SetStyle(cStyle)

		flag = !flag
	}

	return file
}

func GetReportPDFFile(total float64, totalSales float64,
	sellers []Seller, sales []Sale, date string, salePlace string) *gofpdf.Fpdf {

	pdf := gofpdf.New("L", "mm", "Letter", "")

	pdf.AddPage()
	pdf.SetFont("Times", "B", 28)
	pdf.Cell(40, 10, "Reporte de Ventas "+date)
	pdf.Ln(12)

	pdf.SetFont("Times", "B", 14)
	pdf.SetFillColor(240, 240, 240)

	for _, sth := range salesTableHeaders {
		pdf.CellFormat(40, 7, sth, "1", 0, "C", true, 0, "")
	}

	pdf.SetFont("Times", "", 12)
	pdf.SetFillColor(255, 255, 255)
	align := []string{"C", "C", "C", "C", "C", "C"}

	for i, sale := range sales {
		pdf.Ln(-1)

		pdf.CellFormat(40, 7, sale.Brand, "1", 0, align[i], false, 0, "")

		pdf.CellFormat(40, 7, fmt.Sprintf("%d", sale.Size), "1", 0, align[i], false, 0, "")

		pdf.CellFormat(40, 7, sale.Model, "1", 0, align[i], false, 0, "")

		pdf.CellFormat(40, 7, sale.Location, "1", 0, align[i], false, 0, "")

		pdf.CellFormat(40, 7, fmt.Sprintf("%f", sale.Price), "1", 0, align[i], false, 0, "")

		pdf.CellFormat(40, 7, fmt.Sprintf("%f", sale.PriceBuy), "1", 0, align[i], false, 0, "")

	}

	return pdf

}

func GetInventoryFile(products []Product) *xlsx.File {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	hStyle := headersStyle()

	file = xlsx.NewFile()
	sheet, err := file.AddSheet("INVENTARIO")
	if err != nil {
		fmt.Printf(err.Error())
	}
	sheet.SetColWidth(1, 2, 10)
	sheet.SetColWidth(0, 0, 5)

	// **************** TOTAL TABLE ********************
	row = sheet.AddRow()
	row.SetHeight(25.0)

	for _, ith := range inventoryTableHeaders {
		cell = row.AddCell()
		cell.SetStyle(hStyle)
		cell.Value = ith
	}
	flag := false
	for k, product := range products {
		if k == 0 {
			fmt.Println(product.ToMap())
		}
		row = sheet.AddRow()
		cStyle := contentStyle(flag)

		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", (k + 1))
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.SetStyle(cStyle)
		cell.Value = product.RegDate.Format("02/01/2006")

		cell = row.AddCell()
		cell.Value = product.Code

		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = product.Brand
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = product.Serie
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", product.Size)
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = product.Model
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.SetFloatWithFormat(product.Price, "###0.00")
		// cell.Value = fmt.Sprintf("%.2f", product.Price)

		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", product.Quantity)
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = product.Location
		cell.SetStyle(cStyle)

		cell = row.AddCell()
		cell.Value = product.Note
		cell.SetStyle(cStyle)

		flag = !flag
	}

	return file
}

func headersStyle() *xlsx.Style {
	style := xlsx.NewStyle()

	fill := xlsx.NewFill("solid", "3f9a9b", "a0c11d")
	style.ApplyFill = true
	style.Fill = *fill

	font := xlsx.NewFont(13, "Consolas")
	font.Bold = true
	style.ApplyFont = true
	style.Font = *font

	alignment := xlsx.DefaultAlignment()
	alignment.WrapText = true
	alignment.Horizontal = "center"
	alignment.Vertical = "center"
	style.ApplyAlignment = true
	style.Alignment = *alignment

	border := xlsx.NewBorder("thin", "thin", "thin", "thin")
	border.LeftColor = "000000"
	border.RightColor = "000000"
	border.TopColor = "000000"
	border.BottomColor = "000000"
	style.ApplyBorder = true
	style.Border = *border

	return style
}

func highlight(style *xlsx.Style) {
	fill := xlsx.NewFill("solid", "ffff00", "ffff00")
	style.ApplyFill = true
	style.Fill = *fill
}

func contentStyle(flag bool) *xlsx.Style {
	style := xlsx.NewStyle()
	fill := xlsx.DefaultFill()

	font := xlsx.NewFont(11, "Consolas")
	style.ApplyFont = true
	style.Font = *font

	if flag {
		fill = xlsx.NewFill("solid", "ebeef4", "ebeef4")
	}
	style.ApplyFill = true
	style.Fill = *fill

	alignment := xlsx.DefaultAlignment()
	alignment.WrapText = true
	alignment.Horizontal = "center"
	alignment.Vertical = "center"
	style.ApplyAlignment = true
	style.Alignment = *alignment

	border := xlsx.NewBorder("thin", "thin", "thin", "thin")
	border.LeftColor = "000000"
	border.RightColor = "000000"
	border.TopColor = "000000"
	border.BottomColor = "000000"
	style.ApplyBorder = true
	style.Border = *border

	return style
}

func extraStyle() *xlsx.Style {
	style := xlsx.NewStyle()

	font := xlsx.NewFont(12, "Consolas")
	style.ApplyFont = true
	style.Font = *font

	alignment := xlsx.DefaultAlignment()
	alignment.WrapText = true
	alignment.Horizontal = "center"
	alignment.Vertical = "center"
	style.ApplyAlignment = true
	style.Alignment = *alignment

	return style
}
