migrate -database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations
migrate create -ext sql -dir db/migrations create_table_wrong_up
