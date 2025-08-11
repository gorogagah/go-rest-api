# 📚 Go REST API with Gin and MongoDB

This is a simple RESTful API built with [Golang](https://go.dev/), [Gin](https://github.com/gin-gonic/gin), and [MongoDB](https://www.mongodb.com/). It's designed as a beginner-friendly project to help you get started with building APIs in Go.

📝 **Based on this article:**  
1. 👉 [Getting Started with REST APIs in Golang: A Practical Guide](https://medium.com/@gorogagah/get-started-with-go-golang-rest-apis-dd58c8a468ed)
2. 👉 [JWT Auth in Golang REST API — A Beginner’s Guide](https://medium.com/@gorogagah/jwt-auth-in-golang-rest-api-a-beginners-guide-5ce5c0c1d3d0)  

---

## 🚀 Features

- Create, Read, Update, Delete (CRUD) operations on books
- JWT-based authentication
- Session management with MongoDB
- Clean project structure (models, routes, controllers)
- MongoDB connection using official Go driver
- Simple and easy to follow

---

## 📁 Project Structure

```
go-rest-api/
├── main.go
├── controllers/
│ ├── book_controller.go
│ └── auth_controller.go
├── models/
│ ├── book.go
│ ├── user.go
│ └── session.go
├── routes/
│ ├── book_routes.go
│ └── user_routes.go
├── middleware/
│ └── auth_middleware.go
├── config/
│ └── db.go
```

---

## ⚙️ Requirements

- Go 1.20+
- MongoDB (local or Atlas)
- Git

---

## 📦 Installation

```bash
# Clone the repo
git clone https://github.com/gorogagah/go-rest-api.git
cd go-rest-api

# Initialize Go modules
go mod tidy

# Run the app
go run main.go
```

---

## 🔌 API Endpoints

### Books

| Method | Endpoint    | Description       |
| ------ | ----------- | ----------------- |
| GET    | /books      | Get all books     |
| GET    | /books/\:id | Get book by ID    |
| POST   | /books      | Create a new book |
| PUT    | /books/\:id | Update a book     |
| DELETE | /books/\:id | Delete a book     |

### Authentication

| Method | Endpoint       | Description                 |
| ------ | -------------- | --------------------------- |
| POST   | /auth/register | Register a user             |
| POST   | /auth/login    | Login and get JWT           |
| GET    | /auth/logout   | Logout (invalidate session) |

---

## 🛠 Sample Payload

```json
{
  "title": "Harry Potter and the Philosopher's Stone",
  "author": "J. K. Rowling"
}
```

---

## 🧑‍💻 Author

Made with ❤️ by [gorogagah](https://github.com/gorogagah)

Follow me on [Medium](https://medium.com/@gorogagah) or [LinkedIn](https://www.linkedin.com/in/anggoro-gagah/) for more beginner Go tutorials.

---

## 📜 License
This project is open-source and available under the MIT [License](LICENSE).
