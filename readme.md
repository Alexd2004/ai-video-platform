# AI Video Intelligence Platform

Backend system for uploading, processing, and searching videos with AI.

Tech Stack:
- Go
- PostgreSQL
- Redis
- Amazon S3
- Kubernetes



# Database Setup Importance
Please create a .env file following the structure of the .env.example file. Use your detials 

This is only for a local host, later on we will use a deployed database version. So use your own local host values for the env.


1. brew install postgresql (this is the download for mac)

The name is important 
2. createdb video_platform

3. Create the env based off of the env example and use your own password 

4. go test ./... do this to verify you made it,  should get this as a result
- ok      video-platform/internal/database        0.291s

5. Later when we build migrations in the database files i will update the next steps



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