FROM postgres:latest

# Set environment variables
ENV POSTGRES_USER=your_username
ENV POSTGRES_PASSWORD=your_password
ENV POSTGRES_DB=your_database_name

# Expose the PostgreSQL port
EXPOSE 5432

# Add initialization scripts
COPY ./scripts/ /docker-entrypoint-initdb.d/