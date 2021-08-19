package data

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
	"github.com/thetkpark/golang-todo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"testing"
)

func gormSetup(t *testing.T, dbPath string) (UserRepository, *gorm.DB) {
	gormDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	require.NoError(t, err)

	err = gormDB.AutoMigrate(&models.User{})
	require.NoError(t, err)

	logger := hclog.Default()

	userRepo := NewGormUserRepository(gormDB, logger)
	return userRepo, gormDB
}

func gormTeardown(t *testing.T, db *gorm.DB, dbPath string) {
	tx := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.User{})
	require.NoError(t, tx.Error)

	err := os.Remove(dbPath)
	require.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	dbPath := fmt.Sprintf("todo-test-%d.test.db", rand.Intn(1000))
	userRepo, db := gormSetup(t, dbPath)

	user, err := userRepo.Create("randomUsername1", "password")
	require.NoError(t, err)

	// Check if exist in db
	var userInDB models.User
	tx := db.Where(&models.User{ID: user.ID}).First(&userInDB)
	require.NoError(t, tx.Error)

	gormTeardown(t, db, dbPath)
}
