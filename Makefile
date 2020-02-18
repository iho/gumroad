create_db:
	echo "CREATE DATABASE db;" | sudo -u postgres psql 
	echo "CREATE USER db_user WITH password 'password';" | sudo -u postgres psql 
	echo "GRANT ALL ON DATABASE db TO db_user;" | sudo -u postgres psql 
