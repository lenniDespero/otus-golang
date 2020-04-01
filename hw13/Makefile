PROJECTNAME="hw13"

help: Makefile
	@echo "Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## build: Build application
build:
	@for type in "api" "scheduler" "notifier" ; do \
		CGO_ENABLED=0 go build -o ./bin/$$type ./cmd/$$type/ ; \
	done
