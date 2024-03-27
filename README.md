## Login App

This is a simple login application with a Golang backend and a React frontend. The backend uses MongoDB for storage and provides a REST API for user authentication.

### Prerequisites:

Before running the application, make sure you have the following installed on your system:

Golang<br>
Node.js<br>
npm<br>
Docker (optional)<br>
Running the application<br>

### How-To set up and run the application:

Clone the repository:
```bash
git clone git@github.com:Cerebrovinny/login-app.git
cd login-app
```
Run the MongoDB Docker container (optional):

```bash
make run-docker
```

Alternatively, you can use a local MongoDB installation or a remote MongoDB instance. Just make sure to set the MONGO_CONNECTION_STRING environment variable accordingly.

Migrate the server and create an admin user:
```bash
make migrate-server
```

Run the backend server:
```bash
make run-server
```

The backend server will be available at http://localhost:8080.

Run the React development server:
```bash
make run-frontend
```

The frontend will be available at http://localhost:3000.

### Running tests

To run the tests for the different packages, use the following commands:

Test the handlers package:
```bash
make test-handlers
```

Test the config package:
```bash
make test-config
```

Test the registerAdmin function:
```bash
make test-register-admin
```

### Assumptions

When developing the Login App, several assumptions have been made:

Rate limiting: There is a rate limiting in the application. This is made to protect against brute force attacks.

Environment variables: It is assumed that the user will correctly set the required environment variables, such as MONGO_CONNECTION_STRING, MONGO_DATABASE_NAME, and JWT_KEY, for the backend to function properly.

MongoDB setup: The application assumes that MongoDB is set up and accessible, either through a local installation, a Docker container, or a remote instance.

Single admin user: The application assumes that there will only be one admin user created during the server migration process. This admin user is responsible for managing the application and its users, although user management functionality is not provided in the current version of the app.

### Limitations

The Login App is a simple application for user authentication, and as such, it has some limitations:

No user registration: This application only provides a login functionality, without any registration process for new users. In its current form, the admin user is created during the server migration process.

No password recovery: There is no functionality for password recovery or password reset.

Limited frontend: The frontend is minimalistic, providing only a simple login form without any additional features or user interface elements.