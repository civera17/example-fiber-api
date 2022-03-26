package routes

import (
	"sort"
	"strings"

	"github.com/civera17/fintech-test/database"
	"github.com/civera17/fintech-test/models"
	"github.com/gofiber/fiber/v2"
)

//AddBook
func AddBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	
	database.DB.Db.Create(&book)

	return c.Status(200).JSON(book)
}

//AllBooks
func AllBooks(c *fiber.Ctx) error {
	books := []models.Book{}
	database.DB.Db.Find(&books)

	return c.Status(200).JSON(books)
}

//SlowestQueries
func SlowestQueries(c *fiber.Ctx) error {
	queries := []models.QueryInfo{}

	queryType := strings.ToUpper(c.Params("type"))

	database.DB.Db.Where("type = ?", queryType).Scopes(Paginate(c)).Find(&queries)
	sort.Slice(queries, func(i, j int) bool {
		return queries[i].CostSeconds > queries[j].CostSeconds
	})

	return c.Status(200).JSON(queries)
}

//Update
func Update(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Model(&book).Where("title = ?", title.Title).Update("author", title.Author)

	return c.Status(400).JSON("updated")
}

//Delete
func Delete(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Where("title = ?", title.Title).Delete(&book)

	return c.Status(200).JSON("deleted")
}