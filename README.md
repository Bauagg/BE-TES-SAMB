# BE-TES-SAMB

This is the backend for the **TES-SAMB** project. It is built using Go and connects to a MySQL database.

## Prerequisites

Before you get started, make sure you have the following installed:

- [Go](https://golang.org/dl/)
- [MySQL](https://www.mysql.com/)

## Setup Instructions

### 1. Clone the Repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/Bauagg/BE-TES-SAMB.git
cd BE-TES-SAMB

### 1. Install Dependencies

Use the following command to install all necessary dependencies:

go mod tidy

This will download all required Go packages for the project.

### 3. Create a .env File

In the root directory of the project, create a .env file and add the following configuration settings:

# KONFIGURASI DATABASE
DB_HOST=localhost
DB_PORT=3306
DB_NAME=tes-golang-samb
DB_USER=root
DB_PASSWORD=root

# Application settings
APP_PORT=:8080
SECRETKEY_TOKEN=jhgker90c9i09wrkkpo0
URL_HOST_SERVER=http://localhost:8080

Note: Replace DB_USER and DB_PASSWORD with your actual MySQL credentials.

### 4. Set Up the Database

Ensure that your MySQL server is running and that the specified database (tes-golang-samb) exists. If it doesn’t exist, create it by running the following SQL command:

CREATE DATABASE tes-golang-samb;

You can run this command in the MySQL command line or using any SQL client.

### 5. Run the Application

To start the application, use the following command:

go run main.go

This will start the server on http://localhost:8080 (or on the port specified in your .env file).

### 6. Verify the Application

To verify that the application is running, open your browser or a tool like Postman and go to:

http://localhost:8080

You should see a response indicating that the application is active.

# Building for Production

To build the application for production, run:

go build -o be-tes-samb

This command will create an executable file named be-tes-samb in the project directory. To run it:

./be-tes-samb

## Additional Notes

- Make sure that the .env file and database configuration match your development or production environment.
- Keep the SECRETKEY_TOKEN safe, as it’s used for secure operations within the app.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

This `README.md` covers all the essential setup and running steps, ensuring the project is easy to set up from scratch. Let me know if you need any further assistance!
