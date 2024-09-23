package login

import (
	"encoding/json"

	"io/ioutil"
	"log"
	sql_basic "main/src/sql"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Main(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("PANIC: Erro Interno")
			http.Error(response, "Erro interno", http.StatusInternalServerError)
		}
	}()
	startHour := time.Now().UTC()

	if request.Method == "GET" {
		// Pegando pararms
		userMail := request.URL.Query().Get("mail")
		userPassword := request.URL.Query().Get("password")
		if userMail == "" {
			http.Error(response, "Email não enviado", http.StatusBadRequest)
			return
		} else if userPassword == "" {
			http.Error(response, "Senha não enviada", http.StatusBadRequest)
			return
		}

		resp := getLogin(userMail, userPassword)

		response.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Panicln("Erro ao converter em JSON")
		}
		if resp.Error.Status == 0 {
			_, err = response.Write(jsonResp)
			if err != nil {
				log.Panicln("Erro ao retornar reqeust")
			}
			log.Println(request.Method, "- Terminated Request time", time.Since(startHour), " - ", http.StatusOK)
		} else {
			http.Error(response, string(jsonResp), resp.Error.Status)
			log.Println(request.Method, "- Terminated Request time", time.Since(startHour), " - ", resp.Error.Status)
		}

		return
	} else {
		http.Error(response, "Metodo inválido", http.StatusMethodNotAllowed)
		return
	}

}

func getLogin(userMail string, userPassword string) responseLogin {
	List, err := listPermissions("CALL usermanagementdb.GetUserPermissions(?, ?);", "usermanagementdb", userMail, userPassword)
	if err != nil {
		return responseLogin{Error: typeError{Status: http.StatusInternalServerError, Error: err.Error()}}
	}
	if len(List) == 0 {
		return responseLogin{Error: typeError{Status: http.StatusUnauthorized, Error: "Senha e email incorreto "}}
	}
	return responseLogin{Token: createTokenJWT(userMail, List)}

}

func listPermissions(query string, dataBase string, args ...interface{}) ([]string, error) {

	var mapResp []map[string]interface{}
	resp, err := sql_basic.RunQuery(query, dataBase, true, true, args...)
	if err != nil {
		return nil, err
	}
	var listValue []string
	err = json.Unmarshal(resp.([]byte), &mapResp)
	if err != nil {
		return nil, err
	}

	for _, datamap := range mapResp {
		for _, value := range datamap {
			listValue = append(listValue, value.(string))
		}
	}

	return listValue, nil
}

func createTokenJWT(userMail string, PemissionsList []string) (token string) {
	// Carregar a chave privada (private.pem)
	privateKeyData, err := ioutil.ReadFile("private.pem")
	if err != nil {
		log.Panicln("Erro ao ler a chave privada:", err)
		return
	}

	// Parse da chave privada
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		log.Panicln("Erro ao parsear a chave privada:", err)
		return
	}

	// Definindo a expiração do token (aqui ele expira em 12 horas)
	expirationTime := time.Now().Add(12 * time.Hour)

	// Criando os body do token, que incluem o username, permissões e a expiração
	bodyToken := &BodyJWT{
		Email:       userMail,
		Permissions: PemissionsList,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Gerando token assinado e passando para string
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, bodyToken).SignedString(privateKey)
	if err != nil {
		log.Panicln("Erro ao gerar o token JWT:", err)
		return
	}

	return tokenString

}
