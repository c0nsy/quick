# Quick Brainstorm

**Quick Brainstorm** is a drop-in, drop-out collaborative workspace for quick brainstorming sessions. It allows users to create unique sessions with real-time communication via instant messaging and a collaborative whiteboard. Sessions are automatically deleted when everyone leaves or after 15 minutes of inactivity.

## Features

- **Drop-in, Drop-out Communication**: Easily create or join brainstorming sessions with a unique URL.
- **Real-time Collaboration**: Instant chat messaging and whiteboard drawing synced across all connected users.
- **Session Expiry**: Sessions are deleted automatically when all users leave or after 15 minutes of inactivity.

## Tech Stack

- **Backend**: Go (Golang) for WebSocket-based real-time communication and session management.
- **Frontend**: Vue.js for handling the user interface and WebSocket communication.
- **Optional**: Redis for session management and tracking active sessions.


### Frontend

- Vue.js is used for the user interface, split into:
  - **Left half**: Whiteboard for drawing.
  - **Right half**: Chat window for instant messaging.
- WebSocket communication to the Go backend for real-time syncing of drawings and messages.

### Backend

- Go (Golang) is used to manage sessions, WebSocket connections, and session expiry (after 15 minutes or when everyone leaves).
- WebSocket routes manage real-time communication for both the whiteboard and chat.

## Getting Started

### Prerequisites

- **Docker**: Ensure that Docker is installed for containerization of the backend and frontend.
- **Node.js & npm**: For local development of the Vue frontend.
- **Go**: For local development of the Go backend.

### Running the Project

1. **Clone the repository**:
   ```bash
   git clone https://github.com/c0nsy/quick.git
   cd quick 
2. **Run with Docker:**
   ```bash
   docker-compose up --build
   ```
3. Access the app: After the build, the application will be available at:

    Frontend (Vue.js): http://localhost:8080
    Backend (Go): http://localhost:8081
