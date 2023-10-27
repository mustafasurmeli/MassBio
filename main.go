package main

import (
	"WebAPI1/database"
	"WebAPI1/models"
	"WebAPI1/routes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"math"
	"strconv"
)

type Filter interface {
	Type() string
}

type EnumFilter struct {
	Values []string `json:"values"`
}

func (e EnumFilter) Type() string {
	return "enum"
}

type NumberFilter struct {
	Value interface{} `json:"value"`
}

func (n NumberFilter) Type() string {
	return "number"
}

type FreeFormFilter struct {
	Value string `json:"value"`
}

func (f FreeFormFilter) Type() string {
	return "free_form"
}

type Ordering struct {
	Column    string `json:"kolon_adi"`
	Direction string `json:"direction"`
}

type RequestBody struct {
	Filters  map[string]json.RawMessage `json:"filters"`
	Ordering []Ordering                 `json:"ordering"`
}

var filteredData []map[string]interface{}

func main() {
	app := fiber.New()
	baseURL := `https://api-dev.massbio.info/assignment/query`
	app.Use(cors.New())
	app.Get(baseURL, func(c *fiber.Ctx) error {
		data := new(RequestBody)

		if err := c.BodyParser(data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON verileri çözümlenemedi"})
		}

		fmt.Println("Filters:")
		for columnName, filter := range data.Filters {
			var f Filter

			var enumFilter EnumFilter
			if err := json.Unmarshal(filter, &enumFilter); err == nil && len(enumFilter.Values) > 0 {
				f = enumFilter
			} else {
				var numberFilter NumberFilter
				if err := json.Unmarshal(filter, &numberFilter); err == nil {
					f = numberFilter

				} else {
					var freeFormFilter FreeFormFilter
					if err := json.Unmarshal(filter, &freeFormFilter); err == nil {
						f = freeFormFilter
					}
				}
			}
			fmt.Println(data)
			if f != nil {
				fmt.Printf("Kolon Adı: %s, Filter Type: %s\n", columnName, f.Type())
				filteredData = append(filteredData, map[string]interface{}{
					"Kolon Adı":   columnName,
					"Filter Type": f.Type(),
				})
			}
		}

		fmt.Println("Ordering:")
		for _, order := range data.Ordering {
			fmt.Printf("Kolon Adı: %s, Yön: %s\n", order.Column, order.Direction)
			if order.Direction != "" {
				return routes.Sort(c, order.Column, order.Direction)
			}
		}
		page, _ := strconv.Atoi(c.Query("page", "0"))
		perPage, _ := strconv.Atoi(c.Query("page_size", "0"))
		if page != 0 || perPage != 0 {
			return pagination(c, page, perPage)
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Listen(":8080")
}
func pagination(c *fiber.Ctx, page int, page_size int) error {
	reportOutputs := []models.ReportOutput{}
	sql := "SELECT * from report_outputs"
	var total int64

	database.Database.Db.Raw(sql).Count(&total)
	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, page_size, (page-1)*page_size)
	database.Database.Db.Raw(sql).Scan(&reportOutputs)
	return c.JSON(fiber.Map{
		"data":      reportOutputs,
		"page":      page,
		"last_page": math.Ceil(float64(total / int64(page_size))),
	})
}

//{"filters": {"links.pheno pubmed": "https://www.ebi.ac.uk/ols/ontologies/mondo/terms?iri=http%3A%2F%2Fpurl.obolibrary.org%2Fobo%2FMONDO_0060596"}}
//{"filters": {"main.af_vcf": 0.2601, "main.dp": 229}}
//{"filters": {"MainUploadedVariation": "AGT", "MainSymbol": "HA", "Details2Provean": "mondo"}}
//{"ordering": [{"MainUploadedVariation": "ASC"}]
