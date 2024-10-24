package service

import (
	"fmt"
	"sort"

	"github.com/jung-kurt/gofpdf"
	"github.com/zmskv/sales-app/internal/model"
	"github.com/zmskv/sales-app/internal/repository"
)

type ProductWithIndex struct {
	Index   int           `json:"index"`
	Product model.Product `json:"product"`
}

type SalesService struct {
	repos repository.SalesList
}

func NewSalesService(repos repository.SalesList) *SalesService {
	return &SalesService{repos: repos}
}

func (s *SalesService) CreateRecord(record model.Product) (int, error) {
	return s.repos.CreateRecord(record)
}

func (s *SalesService) GetRecord(id string) (model.Product, error) {
	return s.repos.GetRecord(id)
}

func (s *SalesService) DeleteRecord(id string) (string, error) {
	return s.repos.DeleteRecord(id)
}

func (s *SalesService) GetAllRecords() ([]model.Product, error) {
	return s.repos.GetAllRecords()
}

func (s *SalesService) UpdateRecord(record model.Product) (string, error) {
	return s.repos.UpdateRecord(record)
}

func (s *SalesService) ExportToPDF(sales []ProductWithIndex) (*gofpdf.Fpdf, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(40, 10, "Sales Report")
	pdf.Ln(20)

	pdf.SetFont("Arial", "B", 12)
	headers := []string{"ID", "Title", "Amount", "Price", "Revenue", "Username", "Date"}
	for _, header := range headers {
		pdf.CellFormat(25, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 10)

	for _, record := range sales {
		pdf.CellFormat(25, 10, fmt.Sprintf("%d", record.Product.Id), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 10, record.Product.Title, "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 10, fmt.Sprintf("%d", record.Product.Amount), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 10, fmt.Sprintf("%.2f", record.Product.Price), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 10, fmt.Sprintf("%.2f", float64(record.Product.Amount)*record.Product.Price), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 10, record.Product.Username, "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 10, record.Product.Date.Format("2006-01-02"), "1", 0, "C", false, 0, "")
		pdf.Ln(10)
	}
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 12)
	headers = []string{"Date", "Total Revenue"}
	for _, header := range headers {
		pdf.CellFormat(87.5, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 10)
	revenueByDay := make(map[string]float64)
	for _, record := range sales {
		day := record.Product.Date.Format("2006-01-02")
		if _, exists := revenueByDay[day]; !exists {
			revenueByDay[day] = 0.0
		}
		revenueByDay[day] += float64(record.Product.Amount) * record.Product.Price
	}

	for day, revenue := range revenueByDay {
		pdf.CellFormat(87.5, 10, day, "1", 0, "C", false, 0, "")
		pdf.CellFormat(87.5, 10, fmt.Sprintf("%.2f", revenue), "1", 0, "C", false, 0, "")
		pdf.Ln(10)
	}
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 12)
	headers = []string{"Month", "Total Revenue"}
	for _, header := range headers {
		pdf.CellFormat(87.5, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 10)

	var totalRevenue float64
	revenueByMonth := make(map[string]float64)
	for _, record := range sales {
		month := record.Product.Date.Format("2006-01")
		if _, exists := revenueByMonth[month]; !exists {
			revenueByMonth[month] = 0.0
		}
		revenueByMonth[month] += float64(record.Product.Amount) * record.Product.Price
		totalRevenue += float64(record.Product.Amount) * record.Product.Price
	}

	for month, revenue := range revenueByMonth {
		pdf.CellFormat(87.5, 10, month, "1", 0, "C", false, 0, "")
		pdf.CellFormat(87.5, 10, fmt.Sprintf("%.2f", revenue), "1", 0, "C", false, 0, "")
		pdf.Ln(10)
	}
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(20, 10, "Top Sellers")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 12)
	headers = []string{"Username", "Total Revenue"}
	for _, header := range headers {
		pdf.CellFormat(87.5, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 10)

	revenueByUsername := make(map[string]float64)
	for _, record := range sales {
		username := record.Product.Username
		if _, exists := revenueByUsername[username]; !exists {
			revenueByUsername[username] = 0.0
		}
		revenueByUsername[username] += float64(record.Product.Amount) * record.Product.Price
	}

	var revenueSlice []struct {
		Username string
		Revenue  float64
	}

	for k, v := range revenueByUsername {
		revenueSlice = append(revenueSlice, struct {
			Username string
			Revenue  float64
		}{k, v})
	}

	sort.Slice(revenueSlice, func(i, j int) bool {
		return revenueSlice[i].Revenue > revenueSlice[j].Revenue
	})

	for _, item := range revenueSlice {
		pdf.CellFormat(87.5, 10, item.Username, "1", 0, "C", false, 0, "")
		pdf.CellFormat(87.5, 10, fmt.Sprintf("%.2f", item.Revenue), "1", 0, "C", false, 0, "")
		pdf.Ln(10)
	}
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 14)

	pdf.Cell(40, 15, fmt.Sprintf("Total Revenue for all time: %.2f", totalRevenue))

	return pdf, pdf.Error()
}
