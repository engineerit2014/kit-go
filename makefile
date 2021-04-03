run-test: test-up test-down

mod:
	# This make rule requires Go 1.11+
	GO111MODULE=on go mod tidy

# Test Rules
test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes

test-up:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit

test-down:
	docker-compose -f docker-compose.test.yml down --volumes
