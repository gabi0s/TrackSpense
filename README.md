# TrackSpense is an expense & budget tracker tool which aims to help people managin their money by creating budget limits to track their expense in a handy and playful way !

# To launch the project :

## 1. clone the project

```bash
git clone <URL_REPO_GITLAB>
```

## 2. installer les packages
build the images & containers with docker cli

```bash
docker-compose build
```

then launch it :

```bash
docker-compose up
```


## Access the frontend

```bash
http://localhost:4200/login
```

# Specifications

All services are hosted on a Golang Fiber server and interacts with an isolated dedicated PostgreSQL database using the GORM ORM.

USER SERVICE :

this service is running on port 8080 so be carefull it does not overlaps with other programs you may already have on your machine.

to use /auth route :

exemple :

curl -X POST http://localhost:8080/auth -H "Content-Type: application/json" -d "{"Nom":"Alice", "Password":"123"}"

should return :

{"user_id":"eb8b27f7-0f1d-4356-8615-d3edd8059f30","message":"Connexion succeed"}

to use /solde/:userid route :

exemple :

curl -X GET http://localhost:8080/get-solde/fc5671b0-3625-4272-8b1c-60556f131a0a (you will need to change the UUID (big string) with one for the user_id you'll have once database is created)

should return :

{"message":"Solde gathered with success","solde":1500

to use /solde/update route :

exemple :

curl -X POST http://localhost:8080/update-solde -H "Content-Type: application/json" -d "{"user_id": "fc5671b0-3625-4272-8b1c-60556f131a0a", "amount": 500}

should return :

{"message":"Solde successfully updated","solde":2000}

and curl -X POST http://localhost:8080/update-solde -H "Content-Type: application/json" -d "{"user_id": "fc5671b0-3625-4272-8b1c-60556f131a0a", "amount": -500}

should return :

{"message":"Solde successfully updated","solde":1500}

if needed :

to see the users table :

run docker exec -it containername bash (opens a bash in the db container)

then in the container run : psql -U user_admin -d users_db

then in the psql prompt do : SELECT * FROM users;

if needed :
to see the users table :
run docker exec -it containername bash (opens a bash in the db container)
then in the container run : psql -U user_admin -d users_db
then in the psql prompt do : SELECT * FROM users;


BUDGET SERVICE : 

This service is running on port 8082.

Routes : 

POST :
  
  -/create-budget (create budget for a specific user and category)

  exemple : 
  
  curl -X POST http://localhost:8082/create-budget -H "Content-Type: application/json" -d "{\"user_id\":\"e5a33e48-23a6-4da5-89d3-7bdd40ded66b\", \"category\":\"Food\", \"budgets_limit\":200}"}

  should return : 

  {"budgets":{"user_id":"e5a33e48-23a6-4da5-89d3-7bdd40ded66b","category":"Food","budgets_limit":200,"current_amount":0},"message":"Budget successfully created !"}

GET :

  -/get-budget/:userid (gets budgets for a selected user)

  exemple : 

  curl -X GET http://localhost:8082/get-budget/eb8b27f7-0f1d-4356-8615-d3edd8059f30 (some user_id)

  should return : 

  [{"budgets_id":"c6cdf862-4abb-4be1-b4fc-fd6d02a19533","user_id":"e5a33e48-23a6-4da5-89d3-7bdd40ded66b","category":"Food","budgets_limit":200,"current_amount":0},{"budgets_id":"66693455-8846-49b0-bd0f-7a6fb3fd20bb","user_id":"e5a33e48-23a6-4da5-89d3-7bdd40ded66b","category":"Food","budgets_limit":200,"current_amount":0}]

PUT :

  -/update-budget/:userid/:budgetid (updates a user budget)

  exemple : 

  curl -X PUT http://localhost:8082/update-budget/e5a33e48-23a6-4da5-89d3-7bdd40ded66b/94807550-e16a-450b-891c-21a8f4328ec1 -H "Content-Type: application/json" -d "{\"amount\": 100.50}"

  should return : 

  {"message":"Amount successfully added to budget","new_amount":703.9}

DELETE :

  -/budget/:userid/:budgetid (deletes a specific budget) 

  exemple : 

  curl -X DELETE http://localhost:8082/delete-budget/e5a33e48-23a6-4da5-89d3-7bdd40ded66b/d79b8a80-4b69-4bd7-b5b2-709757760875

  should return : 

  {"message": "Budget successfully deleted"}

EXPENSE SERVICE :

This service is running on port 8081.

POST :

/create-expense (create a new expense for a specific user and budget)

Example :

curl -X POST http://localhost:8081/create-expense -H "Content-Type: application/json" -d \
'{
  "user_id": "fdab61fe-117c-4697-aabf-4eecd54ace9c",
  "budget_id": "03d027c0-095d-4f8f-ab38-fd3f6a9356b0",
  "budget_categorie": "Alimentation",
  "price": 50.75
}'

Expected response :

{"expense":{"user_id":"e5a33e48-23a6-4da5-89d3-7bdd40ded66b","budget_id":"49ae878e-a5ab-4ab8-ac17-5ed69e137f5b","budget_categorie":"MMMMM","price":50.75},"message":"Expense added successfully"

GET :

/get-expenses/:user_id (get all expenses for a user)

Example :

curl -X GET http://localhost:8081/expenses/fdab61fe-117c-4697-aabf-4eecd54ace9c

Expected response :

[
  {
    "expense_id": "c6ac3d7a-f417-42df-b6dd-497a5765e9c0",
    "user_id": "fdab61fe-117c-4697-aabf-4eecd54ace9c",
    "budget_id": "03d027c0-095d-4f8f-ab38-fd3f6a9356b0",
    "budget_categorie": "Alimentation",
    "price": 50.75
  }
]