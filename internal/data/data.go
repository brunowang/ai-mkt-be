package data

import (
	"ai-mkt-be/internal/biz"
	"ai-mkt-be/internal/conf"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPlanRepo)

// Data .
type Data struct {
	lg *log.Helper
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	d := &Data{lg: log.NewHelper(logger)}
	cleanup := func() {
		d.lg.Info("closing the data resources")
	}
	if err := d.InitSqlDB(c); err != nil {
		return nil, cleanup, err
	}
	return d, cleanup, nil
}

func (d *Data) InitSqlDB(c *conf.Data) error {
	dsn := c.Database.Source
	var err error
	d.db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: NewGormLogger(d.lg),
	})
	if err != nil {
		return err
	}

	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	d.ensureTable()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				d.lg.Errorf("ensureTable.Panic, r: %+v", r)
				msg := fmt.Sprintf("### 数据库自动分表失败\n ```%s```", "InitSqlDB")
				d.lg.Errorf(msg)
			}
		}()

		ticker := time.NewTicker(time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			d.ensureTable()
		}
	}()
	return nil
}

func (d *Data) ensureTable() {
	ctx := context.Background()
	mig := d.db.WithContext(ctx).Migrator()
	if err := mig.AutoMigrate(
		&biz.Plan{},
	); err != nil {
		d.lg.WithContext(ctx).Errorf("auto migrate table failed, err: %v", errors.WithStack(err))
	}
}
