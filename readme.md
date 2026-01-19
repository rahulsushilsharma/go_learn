# Go Learn - A Journey Through Go Programming

A comprehensive learning project showcasing various Go programming concepts through interactive applications and algorithms.

## 🎯 Project Overview

This project represents a learning journey through Go programming, featuring multiple modules that demonstrate different aspects of the language including web development, game development, terminal applications, and algorithm implementations.

## 🚀 Features

### 🐍 Snake Game
- **Terminal-based snake game** with real-time controls
- **WebSocket streaming** for live browser gameplay
- **Collision detection** and food generation
- **Score tracking** and game state management
- **Cross-platform input handling** using tcell library

### 🌐 Web API Server
- **RESTful API** built with Gin framework
- **User management** endpoints (GET/POST users)
- **CORS support** for frontend integration
- **WebSocket connections** for real-time communication
- **Swagger documentation** with auto-generated API specs

### 🖼️ Image Processing
- **Base64 encoding** for image data
- **Terminal rendering** capabilities
- **File I/O operations** for image handling

### 📚 Algorithm Collection
- **LeetCode-style solutions** including:
  - Two Sum
  - Valid Parentheses
  - Longest Common Prefix
  - Plus One (array increment)
  - Merge Two Sorted Lists
  - Binary Search implementations
  - String manipulation algorithms

## 🛠️ Technologies Used

### Core Technologies
- **Go 1.24.3** - Primary programming language
- **Gin Framework** - HTTP web framework
- **tcell/v2** - Terminal UI library
- **Gorilla WebSocket** - Real-time communication

### Documentation & API
- **Swaggo** - Swagger documentation generator
- **gin-swagger** - Gin middleware for Swagger
- **gin-contrib/cors** - Cross-origin resource sharing

### Additional Libraries
- **Base64 encoding** for data transformation
- **JSON marshaling** for API communication
- **File I/O operations** for persistent storage

## 📁 Project Structure

```
go_learn/
├── main.go                    # Application entry point
├── go.mod                     # Go module definition
├── go.sum                     # Dependency checksums
├── readme.md                  # This documentation
├── internal/                  # Internal packages
│   ├── snake_game/           # Snake game implementation
│   │   └── game.go           # Game logic and rendering
│   ├── handler/              # HTTP handlers
│   │   ├── rest.go           # REST API endpoints
│   │   └── game_stream.go    # WebSocket handling
│   ├── router/               # HTTP routing
│   │   └── router.go         # Route configuration
│   ├── model/                # Data models
│   │   ├── game_model.go     # Game state structure
│   │   ├── user_model.go     # User data structure
│   │   └── errors.go         # Error handling
│   ├── image_to_terminal/    # Image processing
│   │   └── render.go         # Image rendering logic
│   └── tutorial1/            # Algorithm implementations
│       └── tutorial.go       # Practice problems
├── docs/                     # API documentation
│   ├── docs.go              # Generated Swagger docs
│   ├── swagger.json         # OpenAPI specification
│   └── swagger.yaml         # YAML API spec
└── wifey.png                # Sample image for processing
```

## 🎮 How to Run

### Prerequisites
- Go 1.24.3 or higher
- Git for cloning

### Installation
```bash
git clone <repository-url>
cd go_learn
go mod tidy
```

### Running Different Components

#### 1. Snake Game (Terminal)
```bash
go run .
# Uncomment snake_game.RunGame() in main.go
```

#### 2. Web Server with API
```bash
go run .
# Uncomment Serve() in main.go
# Server runs on :8000
```

#### 3. Image Processing
```bash
go run .
# Currently active - processes wifey.png
```

#### 4. Access API Documentation
- **Swagger UI**: `http://localhost:8000/swagger/index.html`
- **JSON Spec**: `http://localhost:8000/swagger/doc.json`

## 🌐 API Endpoints

### Users
- `GET /users` - Retrieve all users
- `POST /users` - Create a new user

### WebSocket
- `GET /ws` - WebSocket connection for live game streaming

### CORS Configuration
- **Allowed Origins**: `http://localhost:5173`
- **Allowed Methods**: GET, POST, PUT, DELETE
- **Allowed Headers**: Origin, Content-Type

## 🎮 Game Controls

### Terminal Controls
- **W/↑** - Move up
- **S/↓** - Move down  
- **A/←** - Move left
- **D/→** - Move right
- **R** - Restart game
- **ESC/Ctrl+C** - Quit game

### Browser Controls
- Same controls via WebSocket connection
- Real-time game state synchronization

## 📊 Game Features

- **Dynamic food generation** with collision avoidance
- **Snake growth** on food consumption
- **Boundary wrapping** for continuous gameplay
- **Score tracking** based on snake length
- **Game over detection** on self-collision
- **Real-time rendering** at 10 FPS

## 🔧 Development Notes

### Learning Objectives
- **Go fundamentals** (structs, interfaces, goroutines)
- **Web development** with REST APIs
- **Real-time communication** via WebSockets
- **Terminal programming** with tcell
- **Algorithm implementation** and optimization
- **Error handling** patterns in Go

### Code Organization
- **Modular design** with clear separation of concerns
- **Dependency injection** for testability
- **Channel-based communication** for concurrency
- **JSON serialization** for API responses

## 🚧 Future Enhancements

- [ ] **Frontend client** for browser-based gameplay
- [ ] **Multiplayer support** with room management
- [ ] **Database integration** for persistent user data
- [ ] **Authentication system** with JWT tokens
- [ ] **Docker containerization** for easy deployment
- [ ] **Unit tests** for all components
- [ ] **Performance optimization** and profiling

## 📝 Learning Journey

This project serves as a practical learning experience covering:

1. **Basic Go syntax** and project structure
2. **Web API development** with Gin framework
3. **Real-time applications** using WebSockets
4. **Terminal programming** and user interfaces
5. **Algorithm problem-solving** and optimization
6. **Error handling** and logging patterns
7. **Documentation generation** with Swagger

## 🤝 Contributing

This is a learning project. Feel free to:
- Study the code structure
- Suggest improvements
- Report issues
- Share similar learning projects

## 📄 License

This project is for educational purposes. Feel free to use and modify for learning Go programming.

---

**Built with ❤️ as part of a Go programming learning journey**