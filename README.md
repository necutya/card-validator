# Card validator
This is a simple service for credit card validation.  
**Note**: Card validation only means the corectness of the format of the card number (starts with a `4` digit and has 16 digits), not its validity as a real card number issued by a specific provider.

To run the application docker and docker-compose must be installed in your PC:
1. run in a simple mode on 8080 port: `make run`
2. run in a dev mode with debugger and autoreload on 8080 port: `make dev`
3. run linter abd tests: `make precommit`

Swagger documentation: http://localhost:8080/docs
