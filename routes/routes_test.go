package routes

import (
	"regexp"
	"testing"

	"github.com/civera17/fintech-test/database"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestAllBooks(t *testing.T) {
	mockDb := database.ConnectMockDB()

	mockCtx := fiber.New().AcquireCtx(&fasthttp.RequestCtx{})

	queryToExpect := `SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL`

	mockDb.ExpectQuery(regexp.QuoteMeta(queryToExpect))

	err := AllBooks(mockCtx)
	assert.NoError(t, err)

	err = mockDb.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAddBook(t *testing.T) {
	mockDb := database.ConnectMockDB()

	mockCtx := fiber.New().AcquireCtx(&fasthttp.RequestCtx{})

	queryToExpect := `INSERT INTO "books" ("created_at","updated_at","deleted_at","title","author") VALUES ('2022-03-27 12:42:31.983','2022-03-27 12:42:31.983',NULL,'WarA','DostoverskiiA')`

	mockDb.ExpectQuery(regexp.QuoteMeta(queryToExpect))

	err := AddBook(mockCtx)
	assert.NoError(t, err)

}


func TestUpdateBook(t *testing.T) {
	mockDb := database.ConnectMockDB()

	mockCtx := fiber.New().AcquireCtx(&fasthttp.RequestCtx{})

	queryToExpect := `UPDATE "books" SET "author"='Bulgakov',"updated_at"='2022-03-27 13:01:25.804' WHERE title = 'War' AND "books"."deleted_at" IS NULL`

	mockDb.ExpectQuery(regexp.QuoteMeta(queryToExpect))

	err := Update(mockCtx)
	assert.NoError(t, err)
}

func TestDeleteBook(t *testing.T) {
	mockDb := database.ConnectMockDB()

	mockCtx := fiber.New().AcquireCtx(&fasthttp.RequestCtx{})

	queryToExpect := `UPDATE "books" SET "deleted_at"='2022-03-27 13:06:10.539' WHERE title = 'War' AND "books"."deleted_at" IS NULL`

	mockDb.ExpectQuery(regexp.QuoteMeta(queryToExpect))

	err := Delete(mockCtx)
	assert.NoError(t, err)
}