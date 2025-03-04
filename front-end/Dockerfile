# Stage 1: Build the React app
FROM node:18-alpine AS builder

WORKDIR /app

# Copy package.json and package-lock.json, then install dependencies
COPY package*.json ./
RUN npm install

# Copy the rest of the application files and build
COPY . .
RUN npm run build

# Stage 2: Serve with Nginx
FROM nginx:alpine

# Copy the built files to Nginx's default directory
COPY --from=builder /app/build /usr/share/nginx/html

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
