create_db:
	echo "CREATE DATABASE db;" | sudo -u postgres psql 
	echo "CREATE USER db_user WITH password 'password';" | sudo -u postgres psql 
	echo "GRANT ALL ON DATABASE db TO db_user;" | sudo -u postgres psql 

watch:
	watchman watch ./
	watchman -- trigger ./ sqlc 'migrations/*.sql' -- sqlc generate
	watchman -- trigger ./ gqlgen 'schema.graphqls' -- gqlgen generate
	nodemon --exec "go run ./server.go"  --watch  "graph/*.go"
