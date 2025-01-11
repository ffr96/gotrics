package api

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"gotrics.cl/internal/models"
)

type metrics struct {
	brandCategory       *models.BrandCategoryModel
	brand               *models.BrandModel
	browserClient       *models.BrowserClientModel
	browserRecord       *models.BrowserRecordModel
	browserSession      *models.BrowserSessionModel
	container           *models.ContainerModel
	edition             *models.EditionModel
	participantCategory *models.ParticipantCategoryModel
	resource            *models.ResourceModel
	timeDimension       *models.TimeDimensionModel
}

type application struct {
	logger  *slog.Logger
	metrics *metrics
}

func newMetrics(db *sql.DB) *metrics {
	return &metrics{
		brandCategory:       &models.BrandCategoryModel{DB: db},
		brand:               &models.BrandModel{DB: db},
		browserClient:       &models.BrowserClientModel{DB: db},
		browserRecord:       &models.BrowserRecordModel{DB: db},
		browserSession:      &models.BrowserSessionModel{DB: db},
		container:           &models.ContainerModel{DB: db},
		edition:             &models.EditionModel{DB: db},
		participantCategory: &models.ParticipantCategoryModel{DB: db},
		resource:            &models.ResourceModel{DB: db},
		timeDimension:       &models.TimeDimensionModel{DB: db},
	}
}

func Main(addr *string, dsn *string) {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger:  logger,
		metrics: newMetrics(db),
	}

	logger.Info("Server started at", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
