# Run the application in dev mode with auto reload
dev:
	@docker-compose -f docker-compose.yaml up dev

# Run the application
run:
	@docker-compose -f docker-compose.yaml up run

# Run linter and tests
precommit:
	@docker run  --rm -v "`pwd`:/workspace:cached" -w "/workspace/." golangci/golangci-lint:latest golangci-lint run -c .golangci.yml
	@docker image build --target test .
