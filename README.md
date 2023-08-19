# Simple Go CRUD API

This is a basic Go HTTP server that implements CRUD operations for managing notes. It uses the `httprouter` package for routing and includes middleware for logging.

## Prerequisites

- Go (1.13 or later)
- [httprouter](https://github.com/julienschmidt/httprouter)
- [go-chi middleware](https://github.com/go-chi/chi)

## Getting Started

1. Clone this repository:

```bash
git clone https://github.com/yourusername/your-repo.git
cd your-repo
```

2. Install dependencies:

```bash
go mod download
```

3. Run the application:

```bash
go run main.go
```

The server will start and listen on port 4000 by default. You can access the server by visiting `http://localhost:4000`.

## Endpoints

<!-- List of endpoints and their descriptions here -->

## Middleware

The application uses the `go-chi` middleware for logging. Middleware enhances the server with additional functionality.

## Future Work

The current version of the application provides a foundation for managing notes through CRUD operations. Here are some areas of improvement for future development:

- **Authentication and Authorization**: Enhance the security of the application by implementing user authentication and authorization mechanisms. Ensure that only authorized users can perform CRUD operations on notes.

- **JWT (JSON Web Tokens)**: Implement JWT-based authentication to securely manage user sessions. This can help in maintaining user sessions and ensuring the integrity of user data.

- **Rate Limiting**: Introduce rate limiting to prevent abuse of the server's resources. This can prevent malicious users or automated bots from overwhelming the server with requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
