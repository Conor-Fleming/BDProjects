# Social Media Backend 
Social Media Backend is a project ive been working on as part of the curriculum at [Boot.Dev](https://boot.dev/tracks/computer-science). The motive behind this project was to become more familiar with programming in Go, as well as backend development as a whole.

This backend (in its current state) can take request to Create, Update, Get, and Delete users, as well as Creating, Deleting, and getting posts from a specified user.

## Usage

To use this backend yourself you would need to use a REST client(I used REST Client for VSCode extension) and send GET, DELETE, PUT, or POST requests with the required data in the body of the request.

```http
POST http://localhost:8080/users

{
   "Email":"test@example.com",
   "Password":"testpass",
   "Name":"John Doe",
   "Age": 55
 }

 POST http://localhost:8080/posts

{
   "UserEmail":"test@example.com",
   "Text":"This would create John Doe's first post!"
 }
```

Thanks for looking!