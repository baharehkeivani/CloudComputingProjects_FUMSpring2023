# imports OpenJDK Java 20 image as the base image
FROM openjdk:8

# Sets the working directory for the image to /app
WORKDIR /app

# Copies all files from the current directory of the host machine to the /app directory in the image
COPY . /app

#compiles java file
RUN javac PingPong.java

#specifies the command to run when the container is started
CMD ["java", "PingPong"]

#Expose the port your app is running on
EXPOSE 8001

