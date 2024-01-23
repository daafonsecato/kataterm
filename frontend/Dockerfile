FROM node:21-alpine

LABEL maintainer="David Fonseca <8128david@gmail.com>"

# Step 2: Set the working directory in the container
WORKDIR /usr/src/app

# Step 3: Copy package.json and package-lock.json
COPY package*.json ./

# Step 4: Install app dependencies
RUN npm install

# Step 5: Bundle app source inside Docker image
COPY . .

# Step 6: Build the app for production
RUN npm run build

# Step 7: Install serve package globally to serve the app
RUN npm install -g serve

# Step 8: Define environment variables
# The port number to bind on
ENV PORT=3000
# The hostname of your backend
ENV REACT_APP_BACKEND_URL=http://backend:3001
# The hostname of your backend
ENV REACT_APP_FRONTEND_URL=http://192.168.1.4:3000

# Step 9: Expose the port the app runs on
EXPOSE $PORT

# Step 10: Define the command to run the app
CMD ["serve", "-s", "build", "-l", "tcp://0.0.0.0:$PORT"]