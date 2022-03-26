package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Paginate implements pagination for query
func Paginate(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := c.ParamsInt("page", 1)
		if page == 0 {
			page = 1
		}

		pageSize, _ := c.ParamsInt("size", 10)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
