package repository_test

import (
	"log"
	"testing"

	"github.com/LeoTwins/go-clean-architecture/internal/infrastructure/db"
	"github.com/LeoTwins/go-clean-architecture/pkg/config"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var _ = BeforeSuite(func() {
	var err error
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("環境変数の読み込みに失敗しました: %v", err)
	}
	database, err = db.NewDB(cfg.DBConfig)
	database.Logger = database.Logger.LogMode(logger.Info)

	if err != nil {
		log.Fatalf("データベースとの接続に失敗しました: %v", err)
	}

	flushRecords()
})

var _ = BeforeEach(func() {
	if database == nil {
		Skip("repository test skipped because no database connection")
	}
})

func flushRecords() {
	database.Exec("TRUNCATE shopping_items CASCADE")
}
