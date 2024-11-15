# Use the official Node.js image as a base image
FROM node:18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json (if available)
COPY frontend/package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the frontend code
COPY frontend/ ./

# Build the frontend application
RUN npm run build

# Expose the port that Nuxt3 will run on
EXPOSE 3000

# Command to start the application
CMD ["npm", "run", "start"]
