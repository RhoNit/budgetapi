<div align="center">
  <h1 style="font-family: cursive;">𝓑𝓾𝓭𝓰𝓮𝓽𝓐𝓟𝓘💰💳📊</h1>
</div>

## Description
BudgetAPI is a simple API that helps manage and track budgets for personal finance. It allows users to create, view, update, and delete budget entries and track expenses. Built using `golang`, `labstack/echo/v4`, `golang-jwt/jwt/v5`, `gomail.v2`, `gorm/Postgres`, `Docker`, `Terraform` etc

## Features/User Journey
- [x] Users can register their details in the application
- [x] Once registered, users would get a welcome note via email
- [x] Users can log into the application using their login credentials i.e. `email_id` and `password`
- [x] Authenticated user can create category, list all categories and delete a specific category by `id`
- [x] We can associate few existing categories to an authenticated user
- [x] Users can create a custom category during the process of associating categories to themselves
- [x] Users can list out all the user-categories associations
- [x] Users should be able to input their `budget` for each category of expenses, e.g. `Food 70.00 Ranit` 
- [x] Users can view their budgets, update and delete a specific budget
- [ ] Users should be able to record expenses or income as a transaction
- [ ] Users should be able to see a detail page for their expenses for selected dates or month
- [ ] Users should get notified at the end of the day to provide their expenses input
- [ ] Users should be notified at the start of the month to prepare their expenses for the new month
- [ ] Users should be allowed to invite their others to their account 

## Project Structure
Project structure:
```plaintext
budgetapi/
  ├── cmd/api/
  │   ├── handlers/
  │   │   ├── app_handler.go
  │   │   ├── auth_handler.go
  │   │   ├── budget_handler.go
  │   │   ├── category_handler.go
  │   │   ├── handler.go
  │   │   └── validate_request_handler.go
  │   │
  │   ├── middlewares/
  │   │   └── auth_middleware.go
  │   │
  │   ├── requests/
  │   │   ├── budget_request.go
  │   │   ├── category_request.go
  │   │   ├── param_request.go
  │   │   └── user_request.go
  │   │
  │   ├── routes/
  │   │   └── endpoints.go
  │   │
  │   ├── services/
  │   │   ├── budget_service.go
  │   │   ├── category_service.go
  │   │   └── user_service.go
  │   │
  │   ├── validation/
  │   │   └── validation.go
  │   │
  │   └── main.go
  |
  ├── common/
  │   ├── custom_errors/
  │   │   └── not_found_error.go
  │   │
  │   ├── api_response.go
  │   ├── db_connection.go
  │   ├── jwt.go
  │   ├── pagination.go
  │   ├── password.go
  │   └── scopes.go
  |
  ├── internal/
  │   ├── mailer/
  │   │   ├── templates/
  │   │   │   ├── hello.html
  │   │   │   └── welcome.html
  │   │   │
  │   │   └── mailer.go
  │   │
  │   ├── migration/
  │   │   ├── seeders/
  │   │   │   └── category_seeders.go
  │   │   │
  │   │   └── migrate_up.go
  │   │
  │   └── models/
  │       ├── base_model.go
  │       ├── budget.go
  │       ├── category.go
  │       ├── user_category.go
  │       └── user.go
  |
  ├── terraform-database-stack/
  │   ├── db_ec2.tf
  │   ├── provider.tf
  │   ├── variables.tf
  │   └── user-data-db.sh
  |
  ├── terraform-application-stack/
  │   ├── app_ec2.tf
  │   ├── provider.tf
  │   ├── variables.tf
  │   └── user-data-app.sh
  |
  ├── .env
  ├── .gitignore
  ├── docker-compose.yml
  ├── Dockerfile
  ├── go.mod
  ├── go.sum
  └── README.md
```
---
## Architecture
![BudgetAPI Logo](schematic_architecture.png "Architecture")
---

## Project Setup & Configuration
**1. Clone the project**
```bash
git clone --depth 1 -b v1 https://github.com/RhoNit/budgetapi.git
cd budgetapi/
```
---
**2. Create a .env file and set it up with the values (you can go through the `.env.sample` for ease of reference)**

---
**3. Download all the required libraries/modules mentioned in go.mod file**
```go
go get .                        // or you can download all the dependencies mentioned in the go.mod by executing
go mod download                 // use any one command to fetch all modules/dependencies into your local
```
---
**4. Database Setup/configuration**
- You can setup and connect the database using any locally installed postgres database client tool i.e. `PgAdmin 4`, `TablePlus`, `DBeaver` etc
- Or you can create a postgres container using the database configs mentioned in the `.env` file
```bash
docker run -itd \
  --name <postgres_container_name> \
  -p 5432:5432 \
  -e POSTGRES_USER=<user_mentioned_in_.env>
  -e POSTGRES_PASSWORD=<password_mentioned_in_.env>
  -e POSTGRES_DB=<database_mentioned_in_.env> \
  postgres:alpine
```
- If you have `docker-compose` installed in your local, you can spin up a postgres container in localhost using the `db` service mentioned in the `docker-compose.yml` file
```bash
docker-compose up -d db         # here `db` is the database service name from `docker-compose.yml` file
```
- Verify if your database container has been started or not
```bash
docker ps                       # list of all running conatainers in the host
docker logs <container_name>    # helpgful while troubleshooting an issue 
```
---
**5. Run the `./internal/migration/seeders/category_seeders.go` file to seed a list of static data (or categories) into `categories` table (OPTIONAL)**
```go
go run ./internal/migration/seeders/category_seeders.go
```
---
**6. Build+Run or Run the application**
```golang
go build -o ./mainbin ./cmd/api/main.go
./mainbin                       // either generate the binary after building the app and then execute the binary

go run ./cmd/api/main.go        // or directly run the application
```
---

## API Documentation
---

### Endpoints
***I already have deployed the application in an EC2 instance. To test the APIs, you can use the `public_ip` of `ec2_instance` i.e. `18.118.156.190`***
 ```json
  {
    "base_url": "http://18.118.156.190:3000/api"
  }
 ```
#### 1. User Registration
- **Endpoint:** `{{ base_url }}/auth/register`
- **Method:** `POST`
- **Description:** Registers a user into the Database.
---
#### Request
- **Headers:**  
  - `Content-Type: application/json`

- **Body:**
  ```json
  {
    "first_name": "Rishit",
    "last_name": "Awasthi",
    "email": "rishit.awasthi@gmail.com",
    "hashed_password": "LOL123"
  }
  ```
---
#### Response
- **Body:**
```json
  {
    "success": true,
    "message": "User Registration successful!",
    "data": {
      "id": 2,
      "created_at": "2025-01-26T22:58:51.699666076Z",
      "updated_at": "2025-01-26T22:58:51.699666076Z",
      "first_name": "Rishit",
      "last_name": "Awasthi",
      "gender": null,
      "email": "rishit.awasthi@gmail.com",
      "categories": null
    }
  }
```
---

#### 2. User Login
- **Endpoint:** `{{ base_url }}/auth/login`
- **Method:** `POST`
- **Description:** Authenticates a user and provides a JWT token upon successful login.
---
#### Request
- **Headers:**  
  - `Content-Type: application/json`

- **Body:**
  ```json
  {
    "email": "rishit.awasthi@gmail.com",
    "hashed_password": "LOL123"
  }
  ```
---
#### Response
- **Body:**
  ```json
  {
    "success": true,
    "message": "user login successful",
    "data": {
      "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNzM3OTY0NzcxLCJpYXQiOjE3Mzc5NjM4NzF9.Q2kISH2Vjd0wcbM6aLyNHyBTDsXRodDN8ecTfeux45s",
      "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNzM3OTc0NjcxLCJpYXQiOjE3Mzc5NjM4NzF9.BwRvA8Q3HVASbDkk4EhvIHcjygWI_H1lz3AyuBxjKlE",
      "user": {
        "id": 2,
        "created_at": "2025-01-26T22:58:51.699666Z",
        "updated_at": "2025-01-26T22:58:51.699666Z",
        "first_name": "Rishit",
        "last_name": "Awasthi",
        "gender": null,
        "email": "rishit.awasthi@gmail.com",
        "categories": null
      }
    }
  }
  ```
---

#### 3. Get Authenticated User
- **Endpoint:** `{{ base_url }}/profile/authenticated-user`
- **Method:** `GET`
- **Description:** Get the user details of current authenticated user 
---
#### Request
- **Headers:** 

  - `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNzM3OTMzMjY1LCJpYXQiOjE3Mzc5MzIzNjV9.dcE9Q5hBIQ4iOHazhZvlxGu7l2V9NU_n4FMiE548hPQ`
---
#### Response
- **Body:**
```json
  {
    "success": true,
    "message": "Authenticated user",
    "data": {
      "id": 1,
      "created_at": "2025-01-26T14:14:58.252422Z",
      "updated_at": "2025-01-26T14:14:58.252422Z",
      "first_name": "Sushant",
      "last_name": "Shahi",
      "gender": null,
      "email": "sushant1@gmail.com",
      "categories": null
    }
  }
```
---

#### 4. Get Category List
- **Endpoint:** `{{ base_url }}/categories/all?page=5&page_size=3`
- **Method:** `GET`
- **Description:** Get the paginated categories list seeded in the application
---
#### Request
- **Headers:** 
  - `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNzM3OTMzMjY1LCJpYXQiOjE3Mzc5MzIzNjV9.dcE9Q5hBIQ4iOHazhZvlxGu7l2V9NU_n4FMiE548hPQ`

- **Query Parameters:** 
  - `page` (integer, required): The page number to fetch.  
    Example: `5`
  - `page_size` (integer, required): The number of items per page.  
    Example: `3`
---
#### Response
- **Body:**
```json
  {
    "success": true,
    "message": "categories retrieved",
    "data": {
      "page_size": 3,
      "page": 5,
      "Sort": "",
      "total_rows": 15,
      "total_pages": 5,
      "items": [
        {
          "id": 13,
          "created_at": "2025-01-26T19:13:36.072884Z",
          "updated_at": "2025-01-26T19:13:36.072884Z",
          "name": "Insurance",
          "slug": "insurance",
          "is_custom": false
        },
        {
          "id": 14,
          "created_at": "2025-01-26T19:13:37.018597Z",
          "updated_at": "2025-01-26T19:13:37.018597Z",
          "name": "Investments",
          "slug": "investments",
          "is_custom": false
        },
        {
          "id": 15,
          "created_at": "2025-01-26T19:13:37.965668Z",
          "updated_at": "2025-01-26T19:13:37.965668Z",
          "name": "Savings",
          "slug": "savings",
          "is_custom": false
        }
      ]
    }
  }
```
---

#### 5. Create Category
- **Endpoint:** `{{ base_url }}/categories/create`
- **Method:** `POST`
- **Description:** Create a category
---
#### Request
- **Headers:** 
  - `Content-Type: application/json`

  - `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNzM3OTMzMjY1LCJpYXQiOjE3Mzc5MzIzNjV9.dcE9Q5hBIQ4iOHazhZvlxGu7l2V9NU_n4FMiE548hPQ`
---
- **Body:** 
```json
  {
	  "name": "test category 1",
	  "is_custom": true
  }
```
---
#### Response
- **Body:**
```json
  {
    "success": true,
    "message": "category created",
    "data": {
      "id": 16,
      "created_at": "2025-01-27T09:35:41.824037364Z",
      "updated_at": "2025-01-27T09:35:41.824037364Z",
      "name": "test category 01",
      "slug": "test_category_01",
      "is_custom": true
    }
  }
```
---

#### 6. Delete a Category
- **Endpoint:** `{{ base_url }}/categories/delete/:id`
- **Method:** `DELETE`
- **Description:** Delete a category by ID
---
#### Request
- **Headers:** 
  - `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNzM3OTMzMjY1LCJpYXQiOjE3Mzc5MzIzNjV9.dcE9Q5hBIQ4iOHazhZvlxGu7l2V9NU_n4FMiE548hPQ`
---
#### Response
- **Body:**
```json
  {
    "success": true,
    "message": "category deleted",
    "data": null
  }
```
---

#### 7. List All Associated User-Categories
- **Endpoint:** `{{ base_url }}/users/categories/all`
- **Method:** `GET`
- **Description:** Get all category lists which are associated with currently authenticated user
---
#### Request
- **Headers:** 

  - `Content-Type: application/json`

  - `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNzM3OTMzMjY1LCJpYXQiOjE3Mzc5MzIzNjV9.dcE9Q5hBIQ4iOHazhZvlxGu7l2V9NU_n4FMiE548hPQ`
---
#### Response
- **Body:** 
```json
  {
    "success": true,
    "message": "categories retrieved for user: 2 | rishit.awasthi@gmail.com",
    "data": {
      "page_size": 10,
      "page": 1,
      "Sort": "",
      "total_rows": 15,
      "total_pages": 2,
      "items": [
        {
          "id": 6,
          "created_at": "2025-01-26T19:13:29.446575Z",
          "updated_at": "2025-01-26T19:13:29.446575Z",
          "name": "House Rent",
          "slug": "house_rent",
          "is_custom": false
        },
        {
          "id": 10,
          "created_at": "2025-01-26T19:13:33.231589Z",
          "updated_at": "2025-01-26T19:13:33.231589Z",
          "name": "Travel & Vacations",
          "slug": "travel_&_vacations",
          "is_custom": false
        },
        {
          "id": 14,
          "created_at": "2025-01-26T19:13:37.018597Z",
          "updated_at": "2025-01-26T19:13:37.018597Z",
          "name": "Investments",
          "slug": "investments",
          "is_custom": false
        }
      ]
    }
  }
```
---

#### 8. Associate User to Existing Categories
- **Endpoint:** `{{ base_url }}/users/categories/associate`
- **Method:** `POST`
- **Description:** Create User-Categories Association (Attach existing categories to user profile)
---
#### Request
- **Headers:** 

  - `Content-Type: application/json`

  - `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNzM3OTMzMjY1LCJpYXQiOjE3Mzc5MzIzNjV9.dcE9Q5hBIQ4iOHazhZvlxGu7l2V9NU_n4FMiE548hPQ`
---
- **Body:** 
```json
  {
	  "category_ids": [10, 6, 14]
  }
```
---
#### Response
- **Body:** 
```json
  {
    "success": true,
    "message": "3 categories are associated with user: 2 | rishit.awasthi@gmail.com",
    "data": null
  }
```
---

#### 9. Create Custom Category During User-Categories Association
- **Endpoint:** `{{ base_url }}/users/categories/create`
- **Method:** `POST`
- **Description:** Create custom categories and associate it to currently authenticated user (Attach custom-created category to current user profile)
---
#### Request
- **Headers:** 

  - `Content-Type: application/json`

  - `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNzM3OTMzMjY1LCJpYXQiOjE3Mzc5MzIzNjV9.dcE9Q5hBIQ4iOHazhZvlxGu7l2V9NU_n4FMiE548hPQ`
---
- **Body:** 
```json
  {
	  "name": "test category"
  }
```
---
#### Response
- **Body:** 
```json
  {
    "success": true,
    "message": "custom category created",
    "data": {
      "id": 16,
      "created_at": "2025-01-23T16:39:25.982295981+05:30",
      "updated_at": "2025-01-23T16:39:25.982295981+05:30",
      "name": "test category",
      "slug": "test_category",
      "is_custom": false
    }
  }
```
---

#### 10. Create Budget for Current User
- **Endpoint:** `{{ base_url }}/budgets/create`
- **Method:** `POST`
- **Description:** Budget creation for a user
---
#### Request
- **Headers:** 

  - `Content-Type: application/json`

  - `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNzM3OTMzMjY1LCJpYXQiOjE3Mzc5MzIzNjV9.dcE9Q5hBIQ4iOHazhZvlxGu7l2V9NU_n4FMiE548hPQ`
---
- **Body:** 
```json
  {
    "title": "Bijlee Bill @At Home PG/Jan, 2025 + Netflix Subscription",
    "category_ids": [7,11],
    "amount": 310
  }
```
---
#### Response
- **Body:** 
```json
  {
    "success": false,
    "message": "budget with user_id 2, title bijlee_bill_@at_home_pg/jan,_2025_+_netflix_subscription, year 2025 and month 1 is already there in the database"
  }
```
---

#### 11. Get All Budgets for Current User
- **Endpoint:** `{{ base_url }}/budgets/all`
- **Method:** `GET`
- **Description:** List all the created budgets of currently authenticated user
---
#### Request
- **Headers:** 

  - `Content-Type: application/json`

  - `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNzM3OTMzMjY1LCJpYXQiOjE3Mzc5MzIzNjV9.dcE9Q5hBIQ4iOHazhZvlxGu7l2V9NU_n4FMiE548hPQ`
---
#### Response
- **Body:** 
```json
  {
    "success": true,
    "message": "budgets retrieved",
    "data": {
      "page_size": 10,
      "page": 1,
      "Sort": "",
      "total_rows": 1,
      "total_pages": 1,
      "items": [
        {
          "id": 1,
          "created_at": "2025-01-26T23:04:31.083471Z",
          "updated_at": "2025-01-26T23:04:31.092958Z",
          "title": "Bijlee Bill @At Home PG/Jan, 2025 + Netflix Subscription",
          "slug": "bijlee_bill_@at_home_pg/jan,_2025_+_netflix_subscription",
          "description": null,
          "user_id": 2,
          "amount": 310,
          "categories": [
            {
              "id": 7,
              "created_at": "2025-01-26T19:13:30.393429Z",
              "updated_at": "2025-01-26T19:13:30.393429Z",
              "name": "Electricity Bill",
              "slug": "electricity_bill",
              "is_custom": false
            },
            {
              "id": 11,
              "created_at": "2025-01-26T19:13:34.179714Z",
              "updated_at": "2025-01-26T19:13:34.179714Z",
              "name": "OTT Subscription",
              "slug": "ott_subscription",
              "is_custom": false
            }
          ],
          "date": "2025-01-26T23:04:31.077758Z",
          "month": 1,
          "year": 2025
        }
      ]
    }
  }
```
---

## Infra Setup & Deployment
### Database Deployment
- [x] Created VPC stack along with provisioning subnet, igw, route table associations for Postgres container
- [x] Then using Terraform script spinned up `postgres:14-alpine` image based container inside an EC2 instance provisioned in the above VPC stack
---
### Application Deployment
- [x] Using the DB based EC2's public IP, setup the DB connection inside the application
- [x] Created a docker image of application using the Dockerfile
```bash
docker build -t <app_img_name> .
```
- [ ] Tried to create a multi-stage docker image to reduce the image size.. but encountered some issues when a container was created using that image. Will try to devote some time on that in future.
- [x] Tagged and pushed the docker image in AWS ECR repository
```bash
aws configure          # first configure with your aws creds then login to the Elasctic Container Registry
aws ecr get-login-password --region <region_name> | docker login --username AWS --password-stdin <ecr_uri>
docker tag <app_img_name>:<tag> <ecr_uri>/<ecr_repo>:<new_tag>
docker push <ecr_uri>/<ecr_repo>:<new_tag>
```
- [x] Create another VPC stack like previous to provision the second EC2 machine where the application's container is gonna be hosted
- [x] Used the ECR repo's image to spin up the application_container
```bash
docker run --name <application_container> \
  -p <host_port>:<container_port> \
  --restart unless-stopped
  -d <ecr_uri>/<ecr_repo>:<new_tag>
```
---