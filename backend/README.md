## Folder Descriptions

### `/cmd/server`

- **Purpose**: Contains the main entry point for the backend application.
- **Files**:
  - `main.go`: Initializes the server, sets up routes, and starts the HTTP and WebSocket servers.

### `/internal`

- **Purpose**: Houses internal packages that contain the core business logic and application functionalities.
- **Subfolders**:

  #### `/handlers`

  - **Purpose**: Contains HTTP and WebSocket handlers for managing requests and real-time communication.
  - **Files**:
    - `http.go`: Handles standard HTTP requests.
    - `websocket.go`: Manages WebSocket connections and message handling.

  #### `/services`

  - **Purpose**: Contains business logic related to session management, message handling, and other core functionalities.
  - **Files**:
    - `session_service.go`: Manages session creation, expiration, and deletion.
    - `message_service.go`: Handles the logic for processing messages sent through WebSocket.

  #### `/models`

  - **Purpose**: Defines data structures (models) used throughout the application.
  - **Files**:
    - `session.go`: Represents a session with fields like ID, last active time, and user details.
    - `message.go`: Represents messages sent in the chat, including sender information and content.

### `/websocket`

- **Purpose**: Contains functionality specific to WebSocket operations, such as connection management and broadcasting messages.
- **Files**:
  - `connection.go`: Manages individual WebSocket connections, including reading from and writing to clients.
  - `broadcast.go`: Handles broadcasting messages to all connected clients.

### `/config`

- **Purpose**: Manages application configuration, including loading environment variables and application settings.
- **Files**:
  - `config.go`: Contains functions for loading and accessing configuration values from the `.env` file.
  - `constants.go`: Holds any constants used throughout the application.
