NAME         =health-service
MAIN_FILE    =main.go
BIN_DIR      =bin
LINT_DIR_LIST = $(shell ls -d */ | grep -v -E scripts\|vendor\|log\|bin\|docs\|mocks/)

build:
	@echo "STEP: BUILD"
	@echo "   1. create dir: $(BIN_DIR)" \
		&& mkdir -p $(BIN_DIR)\
		&& echo "   ==> ok"
	@echo "   2. build: $(MAIN_FILE)" \
		&& go build -o $(BIN_DIR)/$(NAME) $(MAIN_FILE) \
		&& echo "   ==> ok: SERVICE=$(BIN_DIR)/$(NAME)"

run:
	$(BIN_DIR)/$(NAME) api

clean:
	@echo "STEP: CLEAN"
	@echo "   1. remove dir: $(BIN_DIR)"
	@rm -rf bin \
	 	&& echo "   ==> ok"

swagger:
	swag init -g $(MAIN_FILE)

test:
	go test ./...