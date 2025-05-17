package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	// Подключаемся к тестовой БД
	var err error
	dsn := "host=localhost port=5433 user=test password=test dbname=testdb sslmode=disable"
	for i := 0; i < 10; i++ {
		testDB, err = sql.Open("postgres", dsn)
		if err == nil && testDB.Ping() == nil {
			break
		}
		log.Println("Waiting for DB...")
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("could not connect to test db: %v", err)
	}

	// Создаём таблицу
	_, err = testDB.Exec(`DROP TABLE IF EXISTS users; CREATE TABLE users (id INT PRIMARY KEY, name TEXT)`)
	if err != nil {
		log.Fatalf("could not create table: %v", err)
	}

	code := m.Run()
	testDB.Close()
	os.Exit(code)
}

func TestUserRepository_CreateAndFind(t *testing.T) {
	repo := NewUserRepository(testDB)

	// Добавляем пользователя
	user := User{ID: 1, Name: "Alice"}
	err := repo.Create(user)
	if err != nil {
		t.Fatalf("failed to insert user: %v", err)
	}

	// Извлекаем пользователя
	got, err := repo.Find(1)
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}

	if got.Name != "Alice" {
		t.Errorf("expected name 'Alice', got '%s'", got.Name)
	}
}
