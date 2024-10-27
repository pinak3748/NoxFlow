# Build the docker image for the python application
FROM python:3.9.20-bullseye

# Set the working directory
WORKDIR /app

# Copy the requirements file
COPY requirements.txt .

# Install the dependencies
RUN pip install -r requirements.txt

# Copy the source code
COPY . .

# Run the application
CMD ["python", "app.py"]