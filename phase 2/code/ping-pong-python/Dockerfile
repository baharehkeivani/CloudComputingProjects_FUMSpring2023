# Use an official Python runtime as a parent image
FROM tiangolo/uvicorn-gunicorn-fastapi:python3.11

# Set the working directory to /app
WORKDIR /app

# Copy the requirements.txt file into the container at /app
COPY requirements.txt requirements.txt

# Install any needed packages specified in requirements.txt
RUN pip3 install --trusted-host pypi.python.org -r requirements.txt

# Copy the rest of the working directory contents into the container at /app
COPY . .

# Make port 8002 available to the world outside this container
EXPOSE 8002

# Run app.py when the container launches
CMD ["uvicorn", "PingPong:app", "--host", "0.0.0.0", "--port", "8002" , "--reload"]
