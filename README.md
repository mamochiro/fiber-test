

# Create a RESTful API with the following endpoints:

- GET /users: Retrieve a list of all users
- POST /users: Create a new user
- GET /users/{id}: Retrieve a specific user by ID
- PUT /users/{id}: Update a specific user by ID
- DELETE /users/{id}: Delete a specific user by ID
---
### Each user should have the following properties:
ID (unique identifier)  
Name  
Email  
Use an in-memory data store (e.g., a slice) to store the user data.

Implement basic validation for the request payloads:

The name field should be non-empty.  
The email field should be a valid email address.  
Write unit tests to validate the functionality of the API endpoints.

Note: You can use any third-party libraries or frameworks you prefer for routing and JSON handling.

Evaluation Criteria:

Correctness: Does the API implementation meet the specified requirements?  
Code quality: Is the code well-structured, readable, and maintainable?  
Error handling: Are errors handled appropriately and gracefully?  
Testing: Are there comprehensive unit tests covering the API endpoints?  
Best practices: Does the code adhere to Go best practices and conventions?  
Additional Guidelines:

Provide a README file with instructions on how to run and test the application.  
Encourage the candidate to write clean and self-explanatory code, along with appropriate comments if needed.  
Specify a reasonable time frame for the candidate to complete the assignment.  
By providing this assignment, you can assess the candidate's ability to develop a RESTful API using Golang, their understanding of HTTP request handling, data validation, error handling, and testing practices.