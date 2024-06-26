# Dating App API

This project is a RESTful API for a dating application built using the Go programming language, Gin web framework, and JWT for authentication. The service includes user authentication and profile management functionalities.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/doc/install) 1.16+
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [PostGis] (https://postgis.net/documentation/getting_started/)
- [PostgreSQL](https://www.postgresql.org/download/) (or any other preferred database)

### Installing

1. **Clone the repository:**

    ```sh
    git clone https://github.com/wayydev/dating-apps-api.git
    cd dating-apps
    ```

2. **Setup Environment Variables:**

    Create a `.env` file in the root directory and add your environment-specific variables. Here is an example:

    ```sh
    APP_DB_HOST=
    APP_DB_PORT=
    APP_DB_NAME=
    APP_DB_USER=
    APP_DB_PASS=
    APP_JWT_KEY=
    ```

    OR Using development server
    ```sh
    APP_DB_HOST=103.134.154.190
    APP_DB_PORT=5432
    APP_DB_NAME=datingapps
    APP_DB_USER=datingapps
    APP_DB_PASS="datingapps#321"

    APP_JWT_KEY="5F3WeYsVsbSO4FcRWdxoKZzrrVjB1RcZ"
    ```

3. **Install Dependencies:**

    ```sh
    go mod tidy
    ```

4. **Database Configuration**
    Following this step:
    - Install Postgres and Postgis
    - Create A New Database
    - Restore the sql file on databases folders

    ```sh
    # Step 1: Install Postgres and Postgis
    sudo apt update
    sudo apt install postgresql postgis postgis-doc

    # Step 2: Create a New Database
    sudo -u postgres psql -c "CREATE DATABASE mydatabase;"

    # Step 3: Install PostGIS Extension
    sudo -u postgres psql -d mydatabase -c "CREATE EXTENSION postgis;"

    # Step 4: Restore the SQL file to the database
    sudo -u postgres psql -d mydatabase -f /project/path/pkg/databases/databases.sql
    ```

The service will be running at `http://localhost:8000`.

5. **Build and Run the Application:**

    ```sh
    go run main.go
    ```

    The service will be running at `http://localhost:8000`.

### Running the Tests

To run the tests for this project using postman:

https://documenter.getpostman.com/view/26168305/2sA3QqhYnC

### Deployment URL

To online access the api `http://103.134.154.190:8080`.
