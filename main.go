package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type EnumFilter struct {
	Values []string `json:"enum"`
}

type NumberFilter struct {
	Value interface{} `json:"number"`
}

type FreeFormFilter struct {
	Value string `json:"free_form"`
}

type Filter struct {
	EnumFilter     *EnumFilter     `json:"enum"`
	NumberFilter   *NumberFilter   `json:"number"`
	FreeFormFilter *FreeFormFilter `json:"free_form"`
}

type Ordering struct {
	Column    string `json:"kolon adı"`
	Direction string `json:"direction"`
}

type RequestBody struct {
	Filters  map[string]Filter `json:"filters"`
	Ordering []Ordering        `json:"ordering"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/parsejson", func(c *fiber.Ctx) error {
		data := new(RequestBody)

		if err := c.BodyParser(data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON verileri çözümlenemedi"})
		}

		fmt.Println("Filters:")
		for columnName, filter := range data.Filters {
			fmt.Printf("Kolon Adı: %s\n", columnName)
			if filter.EnumFilter != nil {
				fmt.Printf("Enum Filter Değeri: %v\n", filter.EnumFilter.Values)
			} else if filter.NumberFilter != nil {
				fmt.Printf("Number Filter Değeri: %v\n", filter.NumberFilter.Value)
			} else if filter.FreeFormFilter != nil {
				fmt.Printf("Free Form Filter Değeri: %s\n", filter.FreeFormFilter.Value)
			}
		}

		fmt.Println("Ordering:")
		for _, order := range data.Ordering {
			fmt.Printf("Kolon Adı: %s, Yön: %s\n", order.Column, order.Direction)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	app.Listen(":8080")

	/*database.ConnectDb()

	app.Get("/assignment/query", routes.GetReportOutputs)


	*/
}
