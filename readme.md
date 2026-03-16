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

run the database migration
- psql video_platform < migrations/001_create_users_table.sql

6. Verify it 
- psql video_platform
- \dt (checks the tables)
- \d users (to insepect the table)
- users
- to exit write \q 

7. You can do some tests inside psql if you want to play around with inserting users there





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
