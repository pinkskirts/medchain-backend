package sql_basic

import (
	"database/sql"
	"encoding/json"
	"fmt"

	//"os"

	_ "github.com/go-sql-driver/mysql"
)

// var username string = os.Getenv("DB_USERNAME")
// var password string = os.Getenv("DB_PASSWORD")
// var host string = os.Getenv("DB_URL")
// var port string = os.Getenv("DB_PORT")

var username string = "root"
var password string = "root"
var host string = "localhost"
var port string = "3306"

func RunQuery(query string, data_base string, has_return bool, inBytes bool, args ...interface{}) (interface{}, error) {

	db, err := ConexaoMySQL(data_base)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if has_return {

		rows, err := db.Query(query, args...)
		if err != nil {
			return nil, fmt.Errorf("\nFalha ao executar a query: %s", err.Error())
		}
		defer rows.Close()
		json_data, err := Rows_in_json(rows, inBytes)
		if err != nil {
			return nil, err
		}

		return json_data, nil

	} else {
		result, err := db.Exec(query, args...)
		if err != nil {

			return nil, err
		}
		last_id, err := result.LastInsertId()
		if err != nil {

			return nil, err

		}
		return last_id, nil
	}
}

func Rows_in_json(rows *sql.Rows, inBytes bool) (interface{}, error) {

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("\nFalha ao pegar as colunas da tabela: %s", err.Error())
	}

	var result []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("\nFalha ao escanear a tabela: %s", err.Error())
		}

		row := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			if val == nil {
				row[col] = nil
			} else {

				switch val.(type) {
				case []byte:

					encodedVal := string(val.([]byte))
					/*
						_, err := base64.StdEncoding.DecodeString(encodedVal)

						if err != nil {
							fmt.Println(err.Error())
							// o valor não está codificado em base64, então é tratado como uma string normal
							row[col] = encodedVal
						} else {
							decoded, err := base64.StdEncoding.DecodeString(encodedVal)
							if err != nil {
								return nil, fmt.Errorf("\nFalha ao decodificar valor em base64: %s", err.Error())
							}
							row[col] = string(decoded)
						}*/
					row[col] = encodedVal

				default:
					row[col] = val
				}
			}
		}

		result = append(result, row)
	}
	if inBytes {
		json_data, err := json.Marshal(result)
		if err != nil {
			return nil, fmt.Errorf("\nFalha ao converter o arquivo em json: %s", err.Error())
		}
		return json_data, nil
	}
	return result, nil
}

func ConexaoMySQL(dataBase string) (*sql.DB, error) {
	// Constrói a string de conexão do MySQL.
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dataBase)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("\nFalha ao conectar com o banco de dados: %s", err.Error())
	}
	return db, nil
}
