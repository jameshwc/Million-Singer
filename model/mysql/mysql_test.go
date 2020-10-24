package mysql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest"
)

var db *sql.DB

type fields struct {
	db *sql.DB
}

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mysql", "8.0", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	err = initSQL()
	if err != nil {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}

	code := m.Run()
	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

func initSQL() error {
	commands := strings.Split(readFile("sql", "init.sql"), ";")
	if err := executeCommands(commands); err != nil {
		return err
	}
	commands = strings.Split(readFile("sql", "test_data.sql"), ";")
	if err := executeCommands(commands); err != nil {
		return err
	}
	return nil
}

func executeCommands(commands []string) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	for _, command := range commands {
		if len(strings.TrimSpace(command)) == 0 {
			continue
		}
		if _, err := tx.Exec(command); err != nil {
			fmt.Println("hi: ", command)
			log.Fatal(err)
		}
	}
	if err := tx.Commit(); err != nil {
		fmt.Println("hello")
		return err
	}
	return nil
}
func readFile(dirname, filename string) string {
	wd, _ := os.Getwd()
	var dat []byte
	var err error
	for {
		if findDir(wd, dirname) == true {
			dat, err = ioutil.ReadFile(filepath.Join(wd, dirname, filename))
			if err == nil {
				break
			} else {
				log.Fatal(err)
			}
		}
		wd = filepath.Dir(wd)
	}
	return string(dat)
}

func findDir(wd, dirname string) bool {
	files, _ := ioutil.ReadDir(wd)
	for _, f := range files {
		if f.Name() == dirname {
			return true
		}
	}
	return false
}
