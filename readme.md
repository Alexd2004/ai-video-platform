# AI Video Intelligence Platform

Backend system for uploading, processing, and searching videos with AI.

Tech Stack:
- Go
- PostgreSQL
- Redis
- Amazon S3
- Kubernetes



Here is an AI summary on how the current project structre will look
video-platform
│
├── cmd/                # Application entrypoints (API server, workers)
│   └── api/            # Starts the HTTP server
│
├── internal/           # Core application logic (private to this project)
│   ├── auth/           # User authentication (register, login, JWT)
│   ├── video/          # Video management (CRUD, metadata, playback)
│   ├── task/           # Task/job management for video processing
│   ├── database/       # Database connection and setup
│   ├── middleware/     # API middleware (auth, logging, etc.)
│   └── config/         # Environment variables and configuration
│
├── pkg/                # Reusable utilities (helpers, validators, etc.)
│
├── workers/            # Background processing services
│                       # (transcoding, transcription, AI tasks)
│
├── deployments/        # Infrastructure configuration
│   ├── docker/         # Dockerfiles and docker-compose
│   └── k8s/            # Kubernetes deployment configs
│
├── scripts/            # Automation scripts (setup, seeding, dev tools)
│
├── migrations/         # Database schema migrations
│
├── .env.example        # Example environment variables
├── go.mod              # Go module and dependencies
└── README.md           # Project documentation