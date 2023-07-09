# Use the official Node.js image as the base image
FROM node:slim

# Set the working directory inside the container
WORKDIR /js-app

# Copy package.json and package-lock.json (if available) to the working directory
COPY package*.json ./

# Copy the rest of the application code to the working directory
COPY . .

# Install the app dependencies
RUN npm install

# Define the command to run our app when the container is started
CMD ["node", "PingPong.js"]


# Expose the port your app is running on
EXPOSE 8003


