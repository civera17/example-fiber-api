package routes

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func TestAllBooks(t *testing.T) {
	// mockDb := database.ConnectMockDb()
	mockCtx := fiber.New().AcquireCtx(&fasthttp.RequestCtx{})

	type args struct {
		c *fiber.Ctx
	}

	tests := []struct {
		name          string
		args          args
		wantErr       bool
		queryToExpect string
	}{
		{
			name:    "test all books",
			args:    args{mockCtx},
			wantErr: false,
			queryToExpect: `SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL`,
		},
		{
			name:    "expect error test all books",
			args:    args{mockCtx},
			wantErr: true,
			queryToExpect: "invalid query",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AllBooks(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AllBooks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
