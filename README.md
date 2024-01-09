# **Project: friend-management**
### How to Run:
- Setup: make pg, make setup
- Run app: make run

### **Technology**:
- Using Go 1.19
- PostgreSQL
- DB migration
- sqlBoiler
- mockery

### Project structure
- Workflow: Request => Routes => Handler => Controller => Repository => Database

- Three layers model:
    + internal/handler: Get request from httpRequest, decode, validate, call controller, write httpResponse
    + internal/controller: Handle business logic, call repositories
    + internal/repository: Data access layer

### API ENDPOINTS(localhost:3000)

*** Using your preferred browser or api call tools such as (Postman or Swagger) to trigger api call

#### 1*. Create User:
   + Path: localhost:3000/users
   + Method: POST.
   + Body:
   {
        "name":"Test",
        "email":"test@example.com"
   }
#### 1. Add Friend:
+ Path: localhost:3000/friends
+ Method: POST.
+ Body: {
  "friends":
  [
  "test1@example.com",
  "test2@example.com"
  ]
  }
#### 2. Get List of Friends:
+ Path: localhost:3000/friends/list?email=test@example.com
+ Method: GET
#### 3. Get List of Common Friends:
+ Path: localhost:3000/friends/common?email1=test1@example.com&email2=test2@example.com
+ Method: GET
#### 4. Subscribe to User:
+ Path: localhost:3000/subscriptions
+ Method: POST
+ Body: {
"requester": "test1@example.com",
"target": "test2@example.com"
}
#### 5. Block User:
+ Path: localhost:3000/subscriptions/block
+ Method: POST
+ Body: {
  "requester": "charlie@example.com",
  "target": "bob@example.com"
  }
#### 6. Update User Topic and Get List of Users received update:
+ Path: localhost:3000/users/update
+ Method: POST
+ Body: {
"sender": "test@example.com",
"text": "Hello World! test2@example.com"
}