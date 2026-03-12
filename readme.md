# AI Video Intelligence Platform

Backend system for uploading, processing, and searching videos with AI.

Tech Stack:
- Go
- PostgreSQL
- Redis
- Amazon S3
- Kubernetes



Here is an AI summary on how the current project structre will look

cmd/
Contains application entrypoints. The api folder starts the HTTP server for the backend.

internal/
Core application logic that is private to the project. This is where most backend code lives.

auth/ handles user authentication (register, login, JWT)

video/ manages video metadata and video-related APIs

task/ manages processing jobs and task tracking

database/ handles database connections and setup

middleware/ contains API middleware like authentication and logging

config/ loads environment variables and configuration

pkg/
Reusable utilities or helper functions used across the project.

workers/
Background processing services that run tasks such as video transcoding, transcription, and AI processing.

deployments/
Infrastructure configuration for running the project in containers or Kubernetes.

docker/ contains Dockerfiles and Docker Compose configs

k8s/ contains Kubernetes deployment manifests

scripts/
Automation scripts for development tasks such as database setup or running workers.

migrations/
Database schema migration files for creating and updating database tables.

.env.example
Example environment variables needed to run the project.

go.mod
Defines the Go module and tracks project dependencies.
