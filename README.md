# Vue-Go-Redis-PostgreSQL Stack Template

*Template for a full-stack application built on Vue.js, Go net/http, Redis, and PostgreSQL.*

## Architecture

### Vue.js
- Vue 3 - composition API
- (Potentially) Pinia for state management
- Vuetify or shadcn/ui component libraries

### Go
- net/http webserver
- gorilla/websocket for real time communication via WebSockets
- CRUD scaffolding
- user management, authentication, authorization
- modular for decomposition into microservices if needed

### Redis
- KV caching for common requests
- Pub/Sub for messaging support

### PostgreSQL
- relational DB for data persistence
