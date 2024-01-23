# Dockerfile for terminalbackend.go
FROM debian:sid

# Install golang
RUN apt-get update && apt-get install -y golang ttyd curl \
    && curl -fsSL https://code-server.dev/install.sh | sh

# Establecer un directorio de trabajo
WORKDIR /app

# Copiar los archivos de la aplicación
COPY . .

# Exponer el puerto en el que se ejecutará la aplicación
EXPOSE 5000

# CMD ["/bin/sh"]
# Definir el comando para ejecutar la aplicación
CMD ["go", "run", "main.go"]
