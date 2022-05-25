# Create survivors
Request: 
curl --location --request POST 'http://localhost:8080/survivors' \
--header 'Content-Type: application/json' \
--data-raw '[
    {
        "name": "Ahana",
        "age": 15,
        "gender": "F",
        "infected": false
    }
]'

Response:
[
    {
        "id": 8,
        "name": "Ahana",
        "age": 15,
        "gender": "F",
        "infected": false
    }
]

# Create survivors with inventory and location details
Request:
curl --location --request POST 'http://localhost:8080/survivors' \
--header 'Content-Type: application/json' \
--data-raw '[
    {
        "name": "Vami",
        "age": 19,
        "gender": "F",
        "infected": false,
        "last_location": {
            "latitude": "27.2046° N",
            "longitude": "77.4977° E"
        },
        "inventory": {
            "items": "Water, Food, and Ammunition"
        }
    }
]'

Response:
[
    {
        "id": 9,
        "name": "Vami",
        "age": 19,
        "gender": "F",
        "infected": false,
        "last_location": {
            "id": 7,
            "survivor_id": 9,
            "latitude": "27.2046° N",
            "longitude": "77.4977° E"
        },
        "inventory": {
            "id": 7,
            "survivor_id": 9,
            "items": "Water, Food, and Ammunition"
        }
    }
]

# Update survivor flag as infected if reporters are at least three
Request: 
curl --location --request PUT 'http://localhost:8080/survivors' \
--header 'Content-Type: application/json' \
--data-raw '
     {
        "id": 4,
        "name": "Shreya",
        "age": 30,
        "gender": "F",
        "infected": false,
        "reporter":3
    }
'

Response:
{
    "id": 4,
    "name": "Shreya",
    "age": 30,
    "gender": "F",
    "infected": true,
    "reporter":{
            "Int64": 3,
            "Valid": true
    }
}

# Update survivor flag as infected when reporters count is less
Request: curl --location --request PUT 'http://localhost:8080/survivors' \
--header 'Content-Type: application/json' \
--data-raw '
     {
        "id": 4,
        "name": "Shreya",
        "age": 30,
        "gender": "F",
        "infected": false,
        "reporter":{
            "Int64": 2,
            "Valid": true
    }
    }
'
Response:
{
    "error": "cannot update survivor to infected as reporter count is less"
}

# Update survivor
Request: curl --location --request PUT 'http://localhost:8080/survivors' \
--header 'Content-Type: application/json' \
--data-raw '
     {
        "id": 4,
        "name": "Shreya",
        "age": 30,
        "gender": "F",
        "infected": false
    }
'
Response:
{
    "id": 4,
    "name": "Shreya",
    "age": 30,
    "gender": "F",
    "infected": false
}

# Update survivor location
curl --location --request PUT 'http://localhost:8080/survivors' \
--header 'Content-Type: application/json' \
--data-raw '
    {
        "id": 7,
        "last_location": {
            "id": 6,
            "survivor_id": 7,
            "latitude": "27.2046° N",
            "longitude": "140.4977° E"
        }
    }
'
Request:
{
    "id": 7,
    "name": "",
    "infected": false,
    "last_location": {
        "id": 6,
        "survivor_id": 7,
        "latitude": "27.2046° N",
        "longitude": "140.4977° E"
    }
}

# Percentage of non-infected survivors
Request: curl --location --request GET 'http://localhost:8080/survivor_percentage?infect=false'
Response:
{
    "percentage": 62.5
}

# Percentage of infected survivors
Request: curl --location --request GET 'http://localhost:8080/survivor_percentage?infect=true'
Response:
{
    "percentage": 37.5
}

# Fetch all non-infected survivors
Request: curl --location --request GET 'http://localhost:8080/survivors?infect=false'
Response:
[
    {
        "id": 5,
        "name": "Heer",
        "age": 22,
        "gender": "F",
        "infected": false
    },
    {
        "id": 6,
        "name": "Ameya",
        "age": 33,
        "gender": "M",
        "infected": false
    }
]

# Fetch all infected survivors
Request: curl --location --request GET 'http://localhost:8080/survivors?infect=true'
Response:
[
    {
        "id": 2,
        "name": "Priya",
        "age": 26,
        "gender": "F",
        "infected": true
    },
    {
        "id": 3,
        "name": "Shantanu",
        "age": 28,
        "gender": "M",
        "infected": true
    },
    {
        "id": 4,
        "name": "Shreya",
        "age": 30,
        "gender": "F",
        "infected": true
    }
]

# Get robots details
Request: curl --location --request GET 'http://localhost:8080/robo'
Response:
[
    {
        "model": "0LKCW",
        "serialNumber": "LTT3J1KLVYPON2J",
        "manufacturedDate": "2022-07-13T01:15:33.9486736+00:00",
        "category": "Land"
    },
    {
        "model": "0UQT4",
        "serialNumber": "141ZY6HWKQ26LWP",
        "manufacturedDate": "2022-05-30T01:15:33.9486256+00:00",
        "category": "Land"
    },
    ...
    {
        "model": "26AMY",
        "serialNumber": "GEFZMXDKP6TJQ2X",
        "manufacturedDate": "2022-06-02T01:15:33.9486283+00:00",
        "category": "Land"
    }
]