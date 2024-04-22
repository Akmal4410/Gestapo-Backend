FROM alpine:latest 

RUN mkdir /app

WORKDIR /app


COPY ./build/authenticationServiceApp /app


EXPOSE 8080

CMD [ "./authenticationServiceApp" ]