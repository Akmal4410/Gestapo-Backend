FROM alpine:latest

# Create a directory in the container for the application
RUN mkdir /app
WORKDIR /app

# Copy the binary into the container
COPY ./build/authenticationServiceApp /app

# # Copy the config.yaml file into the container
COPY ./configs/config.yaml /app/configs/config.yaml

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the application
CMD ["./authenticationServiceApp"]
