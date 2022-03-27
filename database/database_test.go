package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectMockDB(t *testing.T) {
	sqlMock := ConnectMockDB()
	assert.NotNil(t, sqlMock)

	db := DB.Db
	assert.NotNil(t, db)
}