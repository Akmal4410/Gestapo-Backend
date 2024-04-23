To connect to Evans
evans --host localhost --port 8080 -r repl

To list all the service running on the secfic port
lsof -i :8080

To kill the PID
kill -9 pid

For showing logs inside any service
docker logs deploy-authentication-service-1
