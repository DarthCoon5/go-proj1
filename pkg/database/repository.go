package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	DSN string
}

type RepositoryConfig struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
	SslMode  string
}

//Repository ...
type Repository struct {
	DB *gorm.DB
}

//RepositoryInterface ...
type RepositoryInterface interface {
	Database() *gorm.DB
	MinLimit() *uint64
	MaxLimit() *uint64
}

// Init ...
func NewRepository(config *RepositoryConfig) (rep *Repository, err error) {
	dbConfig := config.postgresConfig()

	conn, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dbConfig.DSN,
			PreferSimpleProtocol: true, // отключает неявное использование подготовленных операторов
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	if err != nil {
		return nil, err
	}
	rep = new(Repository)
	rep.DB = conn
	return rep, nil
}

// Status ...
func (rep Repository) Status() bool {
	return rep.Database != nil
}

// Close ...
func (rep Repository) Close() {
	sqlDB, _ := rep.DB.DB()
	sqlDB.Close()
}

// CurrentConnectionsCount ...
func (rep Repository) CurrentConnectionsCount() (int, error) {
	var count int
	err := rep.DB.Raw("select count(datid) from pg_stat_activity").Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c RepositoryConfig) postgresConfig() DBConfig {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.DbName,
		c.Password,
		c.SslMode,
	)
	return DBConfig{DSN: dsn}
}

//MinLimit ...
func (rep Repository) Database() *gorm.DB {
	return rep.DB
}

//MinLimit ...
func (rep Repository) MinLimit() *uint64 {
	return nil
}

//MaxLimit ...
func (rep Repository) MaxLimit() *uint64 {
	return nil
}
