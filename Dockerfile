# Dockerfile
FROM postgres:17-alpine

# Set environment variables for the default database and user
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=CoursePlatform

# Copy the schema.sql file to the container to initialize the database
COPY ./schema.sql /docker-entrypoint-initdb.d/