GenModels:
	goctl model mysql ddl -src="./userAccount/models/*.sql" -dir="./userAccount/models"
GenNewApi:
	goctl api new -home "." ./userAccount/api
GenUpdateApi:
	goctl api go -api ./userAccount/api/api.api -dir ./userAccount/api
GenGoRun:
	go run ./userAccount/api/api.go -f ./userAccount/api/etc/api-api.yaml