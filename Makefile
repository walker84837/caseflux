operatingsystem := $(shell go env GOOS)
EXEC ?= cflux

.PHONY: all clean release debug

all:
	@echo "Building release for '$(operatingsystem)'..."
	GOOS=$(operatingsystem) go build -ldflags="-s -w -linkmode external" -o $(EXEC) src/caseflux.go 

clean:
	@echo "Cleaning up..."
	@rm -rf ./*.a ./*.o ./core ./*.log $(EXEC) || true
