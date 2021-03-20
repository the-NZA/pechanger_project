APP = pechanger 

FLGS = -v
.PHONY: build_backend
build_backend:
	go build $(FLGS) ./backend/cmd/$(APP)

.PHONY: run_backend
run_backend:
	go run $(FLGS) ./backend/cmd/$(APP)

.PHONY: buildnrace_backend
buildnrace_backend:
	go build $(FLGS) -race ./backend/cmd/$(APP)

.PHONY: test_backend
test_backend:
	go test ./...

.PHONY: run_frontend
run_frontend:
	cd frontend && npm run serve

.PHONY: build_frontend
build_frontend:
	cd frontend && npm run build 

.PHONY: clean
clean:
	rm ./$(APP)

.DEFAULT_GOAL := build_backend
