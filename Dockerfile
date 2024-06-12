FROM alpine:3.20.0
RUN mkdir /app
RUN mkdir /app/logs

COPY ./recaptcha-service.bin /app

# Run the server executable
CMD [ "/app/recaptcha-service.bin" ]