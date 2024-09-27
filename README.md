# Task Management API

This is a task management API built with Go and MongoDB. The API supports user registration, task creation, updating, deletion, and other task management operations.

## Project Structure

The project is organized into several key directories:

├── controllers/ 
# Contains all the logic for handling API requests 

├── middlewares/ 

# Authentication, authorization, and other middleware 
├── models/ 

# Contains struct definitions for the data models (e.g., Task, User) 
├── routes/ 

# Defines API routes 
├── utils/ 

# Contains helper functions (e.g., database connection)


## Requirements

To run this project, you need to have the following installed:

- **Go**: Version 1.16 or later
- **MongoDB**: A MongoDB instance or cluster
- **Git** (optional, for cloning the project)

## Environment Variables

You need to define the following environment variables for the application to run properly:

- `MONGO_URI`: The MongoDB connection string.
- `JWT_SECRET`: The secret key used for JWT token signing.

You can place these variables in a `.env` file at the root of the project.

### Example `.env` file:

```bash
MONGO_URI=mongodb+srv://<username>:<password>@cluster.mongodb.net/<database>?retryWrites=true&w=majority
JWT_SECRET=your_jwt_secret_key
```
Installation & Running the Project
1. Clone the Project
git clone https://github.com/your-username/task-management-api.git
cd task-management-api

2. Install Dependencies
The project uses Go modules. To install the required dependencies, run:

go mod tidy


3. Run the Application
You can now run the application by executing:
go run main.go
