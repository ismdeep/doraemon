package pkg

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectToMySQL connect to mysql
func ConnectToMySQL(dsn string, timeout int) {
	startTime := time.Now().Second()
	for {
		if _, err := gorm.Open(mysql.Open(dsn)); err != nil {
			if time.Now().Second()-startTime >= timeout {
				fmt.Println("Error: timeout")
				os.Exit(1)
			}
			time.Sleep(1 * time.Second)
			continue
		}

		fmt.Println("OK")
		return
	}
}

// CreateDBOnMySQL create db on mysql
func CreateDBOnMySQL(dsn string, dbName string, additionAuths []string) error {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	/*
		create database `hello` default character set utf8mb4 collate utf8mb4_unicode_ci;
		CREATE USER 'hello'@'%' IDENTIFIED WITH mysql_native_password BY '4HtD427cfb2kF0n9';
		grant all privileges on `hello`.* to 'hello'@'%';
		flush privileges;
	*/

	if err := db.Exec(fmt.Sprintf("create database `%v` default character set utf8mb4 collate utf8mb4_unicode_ci;", dbName)).Error; err != nil {
		return err
	}

	for _, auth := range additionAuths {
		t := strings.Split(auth, ":")
		if len(t) != 2 {
			return errors.New("addition auth format invalid")
		}

		dbUser := t[0]
		dbPass := t[1]

		if err := db.Exec(fmt.Sprintf("CREATE USER '%v'@'%%' IDENTIFIED WITH mysql_native_password BY '%v';", dbUser, dbPass)).Error; err != nil {
			return err
		}
		if err := db.Exec(fmt.Sprintf("grant all privileges on `%v`.* to '%v'@'%%';", dbName, dbUser)).Error; err != nil {
			return err
		}
	}

	if err := db.Exec("flush privileges;").Error; err != nil {
		return err
	}

	return nil
}
