# Hl7 Listener

This is a golang microservice designed to listen to HL7 messages and forward them onto an outgoing address.

## Development

You will need to have golang 1.18 installed locally. This can be done either via the website or your package manager. apt, dnf, yum etc.

## Deployment

This microservice is deployed as a container on a server that is expecting incoming traffic.

## Expected files

The service requires some environmental variables to be set in order to initialise properly.

Here is an example of what an ENV file would look like. docker-compose will automatically pick this up.

```

PORT=8080
FORWARDING_ADDRESS=http://localhost:5001

````