# Robot Apocalypse

An implementation of robot_apocalypse and postgres

The project is meant to run on local machine. Docker is not implemented yet.

Steps to run and test the service.
## DB Connect 
1. Connect PostgresSQL or run PgAdmin in your machine.
2. Create database 'robo', script given below-
    CREATE DATABASE robo;
3. All the database details are hardcoded in db_connect.go file (robot_apocalypse/pkg/dbg/db_connect.go)
Update the below database configurations in db_connect.go file according to your database-
    Name     = "robo"
	User     = "postgres"
	Password = "******"
	Host     = "localhost"
	Port     = "5432"

## Run application
1. Clone the repository.
2. Go inside the robot_apocalypse folder (Run 'cd robot_apocalypse').
3. Go inside the cmd folder (Run cd cmd)  
3. Run 'go run main.go' command.
4. The server will be up in the local machine.

## API calls
NOTE: All the request and response id captured in api.md file (robot_apocalypse/documention/api.md)
1. The endpoints can be accessed on http://localhost:8080/
2. POST call on /survivors endpoint will insert data in survivors table.
3. UPDATE call on /survivors endpoint will update the location and survivor infected flag in survivors table. Pass the reporter field(assumptions: reporter field stores the number of survivors reporting contamination) as query param when need to update the infected flag.
4. GET call on /survivors endpoint will return all the infected and non-infected survivors based on infected flag. Pass query param 'infect=true' to fetch all the infected survivors and 'infect=false' for non-infected survivors.
5. GET call on /survivor_percentage will return the percentage of infected and non-infected survivors based on infected flag. Pass query param 'infect=true' to fetch the percentage of infected survivors and 'infect=false' to fetch the percentage of non-infected survivors.
6. GET call on /robo will return all the robo information. 