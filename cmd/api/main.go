package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	infra_adapters "github.com/mariojuniortrab/hauling-api/internal/infra/adapters"
	web_assembler "github.com/mariojuniortrab/hauling-api/internal/presentation/web/assembler"
	user_routes "github.com/mariojuniortrab/hauling-api/internal/presentation/web/route/user"
)

func main() {
	mysqlDb, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/hauling")
	if err != nil {
		panic(err)
	}
	defer mysqlDb.Close()

	validator := infra_adapters.NewValidator()
	encrypter := infra_adapters.NewBcryptAdapter()
	tokenizer := infra_adapters.NewJwtAdapter()
	urlParser := infra_adapters.NewChiUrlParserAdapter()

	userAssembler := web_assembler.NewUserAssembler(mysqlDb, validator, encrypter, tokenizer, urlParser)
	middlewareAssembler := web_assembler.NewMiddleWareAssembler(validator, urlParser, tokenizer)

	//Routes
	userRouter := user_routes.NewRouter(userAssembler, middlewareAssembler)

	//Using chi with an adapter to manage routes
	r := infra_adapters.NewChiRouteAdapter()

	//routing
	userRouter.Route(r)

	//starting server
	fmt.Println("Server has started")
	http.ListenAndServe(":8000", r)

}
