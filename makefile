IMPORT_PATH=github.com/bluemir/zumo
BIN_NAME=$(notdir $(IMPORT_PATH))

default: $(BIN_NAME)

GIT_COMMIT_ID = $(shell git rev-parse --short HEAD)

# if gopath not set, make inside current dir
ifeq ($(GOPATH),)
	GOPATH=$(PWD)/.GOPATH
endif

GO_SOURCES = $(shell find . -name ".GOPATH" -prune -o -type f -name '*.go' -print)
JS_SOURCES = $(shell find app/js -type f -name '*.js' -print)
HTML_SOURCES = $(shell find app/html -type f -name '*.html' -print)
CSS_SOURCES = $(shell find app/less -type f -name "*.less" -print)
WEB_LIBS = $(shell find app/lib -type f -type f -print)

DISTS += $(HTML_SOURCES:app/html/%=dist/html/%)
DISTS += $(JS_SOURCES:app/js/%=dist/js/%)
DISTS += dist/css/common.css dist/css/custom-element.css
DISTS += $(WEB_LIBS:app/lib/%=dist/lib/%)

# Automatic runner
DIRS = $(shell find . -name dist -prune -o -name ".git" -prune -o -type d -print)

.sources:
	@echo $(DIRS) makefile \
		$(GO_SOURCES) \
		$(JS_SOURCES) \
		$(HTML_SOURCES) \
		$(CSS_SOURCES) \
		$(WEB_LIBS)| tr " " "\n"
run: $(BIN_NAME)
	go test ./backend/...
	./$(BIN_NAME) #--config ../zumo-config.yaml
auto-run:
	while true; do \
		make .sources | entr -rd make run ;  \
		echo "hit ^C again to quit" && sleep 1  \
	; done
reset:
	ps -e | grep make | grep -v grep | awk '{print $$1}' | xargs kill

## Binary build
$(BIN_NAME).bin: $(GO_SOURCES) $(GOPATH)/src/$(IMPORT_PATH)
	go get -v -d $(IMPORT_PATH)            # can replace with glide
	go build -v \
		-ldflags "-X main.VERSION=$(GIT_COMMIT_ID)-$(shell date +"%Y%m%d.%H%M%S")" \
		-o $(BIN_NAME).bin .
	@echo Build DONE

$(BIN_NAME): $(BIN_NAME).bin $(DISTS)
	cp $(BIN_NAME).bin $(BIN_NAME).tmp
	rice append -v --exec $(BIN_NAME).tmp \
		-i $(IMPORT_PATH)/server
	mv $(BIN_NAME).tmp $(BIN_NAME)
	@echo Embed resources DONE

## Web dist
dist/css/common.css: $(CSS_SOURCES)
	lessc app/less/main.less $@
dist/css/custom-element.css: $(CSS_SOURCES)
	lessc app/less/custom-element.less $@
dist/%: app/%
	@mkdir -p $(basename $@)
	cp $< $@

tools:
	npm install -g less
	go get github.com/GeertJohan/go.rice/rice
clean:
	rm -rf dist/ vendor/ $(BIN_NAME) $(BIN_NAME).bin $(BIN_NAME).tmp
	go clean

$(GOPATH)/src/$(IMPORT_PATH):
	@echo "make symbolic link on $(GOPATH)/src/$(IMPORT_PATH)..."
	@mkdir -p $(dir $(GOPATH)/src/$(IMPORT_PATH))
	ln -s $(PWD) $(GOPATH)/src/$(IMPORT_PATH)

.PHONY: .sources run auto-run reset tools clean
