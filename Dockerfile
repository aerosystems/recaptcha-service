FROM alpine:latest
RUN mkdir /app
RUN mkdir /app/logs

COPY ./recaptcha-service/recaptcha-service.bin /app

# Run the server executable
CMD [ "/app/recaptcha-service.bin" ]