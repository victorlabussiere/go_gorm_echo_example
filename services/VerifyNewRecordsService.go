package services

import (
	"fmt"
	"log"
	"time"

	"github.com/victorlabussiere/go_gorm_echo_postgres_example/initializer"
	"gorm.io/gorm"
)

func listTables(db *gorm.DB) ([]string, error) {
	var tables []string
	query := `
		SELECT table_name 
		FROM information_schema.tables 
		WHERE table_schema = 'public' AND table_type = 'BASE TABLE';
	`

	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}

	return tables, nil
}

func checkNewRecords(table string, lastCheck time.Time) {
	var count int64
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE created_at > ?", table)

	if err := initializer.DB.Raw(query, lastCheck).Scan(&count).Error; err != nil {
		log.Printf("Erro ao consultar tabela %s: %v", table, err)
		return
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	if count > 0 {
		fmt.Printf("[%s] :: Encontrados %d novos registros na tabela %s\n", currentTime, count, table)
	}
}

func VerifyNewRecords() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	lastCheck := time.Now()

	for {
		select {
		case <-ticker.C:
			tables, err := listTables(initializer.DB)
			if err != nil {
				log.Printf("Erro ao listar tabelas: %v", err)
				continue
			}

			for _, table := range tables {
				checkNewRecords(table, lastCheck)
			}

			lastCheck = time.Now()
		}
	}
}
