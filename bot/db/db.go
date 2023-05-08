package db

import (
	"Keeper/models"
	"context"
	"errors"
	"github.com/go-pg/pg/v10"
	"os"
)

// DB - это обертка над *pg.DB
type DB struct {
	*pg.DB
}

// GetDB - возвращает новый экземпляр DB с настройками из переменных окружения
func GetDB() *DB {
	addr := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")

	pgOptions := &pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: dbname,
		PoolSize: 10,
		OnConnect: func(ctx context.Context, conn *pg.Conn) error {
			_, err := conn.Exec("SELECT 1")
			return err
		},
	}

	db := pg.Connect(pgOptions)

	return &DB{db}
}

// Migrate - выполняет миграции из папки db/migrations
func (db *DB) Migrate() error {
	readFile, err := os.ReadFile("db/migrations/20230507130000_create_users_table.sql")
	_, err = db.Exec(string(readFile))
	if err != nil {
		return err
	}
	readFile, err = os.ReadFile("db/migrations/20230507130200_create_services_table.sql")
	_, err = db.Exec(string(readFile))
	if err != nil {
		return err
	}
	return nil
}

// Close - закрывает соединение с БД
func (db *DB) Close() error {
	return db.Close()
}

// AddUser - добавляет нового пользователя в БД
func (db *DB) AddUser(user *models.User) error {

	_, err := db.Model(user).OnConflict("DO NOTHING").Insert()
	return err
}

// AddService - добавляет новый сервис в БД
func (db *DB) AddService(service *models.Service) error {
	tmp := &models.Service{}
	err := db.Model(tmp).Where("id_tg = ? and service_name = ?", service.IdTg, service.ServiceName).Select()
	if err == nil {
		return errors.New("service already exists")
	}
	_, err = db.Model(service).Insert()
	return err
}

// GetService - возвращает сервис по idTg и serviceName
func (db *DB) GetService(idTg int, serviceName string) (*models.Service, error) {
	service := &models.Service{
		IdTg:        idTg,
		ServiceName: serviceName,
	}
	err := db.Model(service).Where("id_tg = ? and service_name = ?", idTg, serviceName).Select()
	if err != nil {
		return nil, err
	}
	return service, nil
}

// GetServices - возвращает все сервисы пользователя по idTg
func (db *DB) GetServices(idTg int) ([]models.Service, error) {
	var services []models.Service
	err := db.Model(&services).Where("id_tg = ?", idTg).Select()
	if err != nil {
		return nil, err
	}
	return services, nil
}

// DeleteService - удаляет сервис по idTg и serviceName
func (db *DB) DeleteService(idTg int, serviceName string) (*models.Service, error) {
	service := &models.Service{
		IdTg:        idTg,
		ServiceName: serviceName,
	}
	err := db.Model(service).Where("id_tg = ? and service_name = ?", idTg, serviceName).Select()
	if err != nil {
		return nil, err
	}
	_, err = db.Model(service).Where("id_tg = ? and service_name = ?", idTg, serviceName).Delete()
	if err != nil {
		return nil, err
	}
	return service, nil
}
