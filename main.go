/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"database/sql"
	"ethereum-explorer/logger"

	"github.com/go-sql-driver/mysql"
)

func main() {
//  cmd.Execute()

  cfg := mysql.Config{
    User:   "test_user",
    Passwd: "1234",
    Net:    "tcp",
    Addr:   "127.0.0.1:3306",
    DBName: "testdb",
  }

  db, err := sql.Open("mysql", cfg.FormatDSN())
  
	if err != nil {
		logger.Logger.Error("db error")
		
	}
  
	result, err := db.Query("SELECT aa, bb FROM test")
  
  for result.Next() {
    
  }
}

