# Tools
GO_BIN := $(shell which go)
BUN_BIN := $(shell which bun)
FSWATCH_BIN := $(shell which fswatch)
TEMPL_BIN := $(shell which templ)

# Project info
PROJECT_NAME := $(shell basename $(CURDIR))
STATIC_DIR := static

# Build configuration
BUILD_DIR := build
MAIN_GO := cmd/main.go
TEMPL_FILES := $(shell find . -type f -name "*.templ")
TEMPL_GO_FILES := $(TEMPL_FILES:.templ=_templ.go)

PROD_BINARY := $(BUILD_DIR)/$(PROJECT_NAME)

# Simple color scheme
CYAN := \033[36m
DIM := \033[2m
RESET := \033[0m

# Basic symbols
CHECK := ✓
ARROW := →

# Simplified Tailwind config focused on templ files
# TAILWIND_CONFIG := 'module.exports = {\n\
#   content: ["./views/**/*.templ"],\n\
#   theme: {\n\
#     extend: {}\n\
#   },\n\
#   plugins: []\n\
# }'

# Start Base Layout
BASELAYOUT := 'package layouts\n\
templ BaseLayout() {\n\
    <!doctype html>\n\
    <html lang="de">\n\
    <head>\n\
        <meta charset="UTF-8">\n\
        <meta name="viewport" content="width=device-width, initial-scale=1.0">\n\
        <title>Document</title>\n\
        // TailwindCSS\n\
        <link rel="stylesheet" href="/static/css/styles.css">\n\
        // AlpineJS\n\
        <script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"></script>\n\
    </head>\n\
    <body>\n\
        <h1 class="text-3xl font-bold text-[\#007D9C]">Hello World</h1>\n\
    </body>\n\
    </html>\n\
}'

MAIN_GO_FILE := 'package main\n\
import (\n\
	"fmt"\n\
	"net/http"\n\
	"github.com/a-h/templ"\n\
	"$(PROJECT_NAME)/views/layouts"\n\
)\n\
func main() {\n\
	fs := http.FileServer(http.Dir("./static"))\n\
    http.Handle("/static/", http.StripPrefix("/static/", fs))\n\
	component := layouts.BaseLayout()\n\
	http.Handle("/", templ.Handler(component))\n\
	fmt.Println("Listening on :3000")\n\
	http.ListenAndServe(":3000", nil)\n\
}\n\
'

AIR_TOML := '\
root = "."\n\
tmp_dir = "build"\n\
\n\
[build]\n\
\tbin = "./build/main"\n\
\tcmd = "templ generate && go build -o ./build/main cmd/main.go"\n\
\tdelay = 1000\n\
\texclude_dir = ["static", "node_modules", "build"]\n\
\texclude_regex = [".*_templ.go"]\n\
\texclude_unchanged = false\n\
\tfollow_symlink = false\n\
\tinclude_ext = ["go", "tpl", "tmpl", "templ", "html"]\n\
\tkill_delay = "0s"\n\
\tlog = "build-errors.log"\n\
\tsend_interrupt = false\n\
\tstop_on_error = true\n\
\n\
[color]\n\
\tbuild = "yellow"\n\
\tmain = "magenta"\n\
\trunner = "green"\n\
\twatcher = "cyan"\n\
\n\
[log]\n\
\ttime = false\n\
\n\
[misc]\n\
\tclean_on_exit = true\n\
'


.PHONY: check-deps init create-dirs setup-go setup-tailwind build serve watch clean help templ deploy

init: check-deps create-dirs setup-go setup-tailwind
	@printf "\n$(CHECK) Project setup complete!\n"

check-deps:
	@printf "\n$(CYAN)Checking dependencies$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@test -n "$(GO_BIN)" || (printf "✗ Go not installed\n" && exit 1)
	@test -n "$(BUN_BIN)" || (printf "✗ Bun not installed\n" && exit 1)
	@test -n "$(FSWATCH_BIN)" || (printf "✗ fswatch not installed\n" && exit 1)
	@test -n "$(TEMPL_BIN)" || (printf "✗ templ not installed\n" && exit 1)
	@printf "$(CHECK) All dependencies ready\n"

create-dirs:
	@printf "\n$(CYAN)Creating project structure$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@mkdir -p $(STATIC_DIR)/js
	@mkdir -p $(STATIC_DIR)/css
	@mkdir -p views/{components,layouts,pages}
	@mkdir -p internal/{handler,middleware,model,service,util}
	@mkdir -p cmd/
	@mkdir -p $(BUILD_DIR)
	@printf "%b" $(MAIN_GO_FILE) > $(MAIN_GO)
	@$(GO_BIN) fmt $(MAIN_GO)
	@echo "package handler" > internal/handler/handler.go
	@echo "package middleware" > internal/middleware/middleware.go
	@echo "package model" > internal/model/model.go
	@echo "package service" > internal/service/service.go
	@echo "package util" > internal/util/util.go
	@printf "%b" $(BASELAYOUT) > views/layouts/BaseLayout.templ
	@templ fmt ./views/layouts/BaseLayout.templ   
	@echo "// Navigation component" > views/components/Navbar.templ
	@echo "// Home page component" > views/pages/HomePage.templ
	@printf "%b" $(AIR_TOML) > .air.toml
	@printf "$(CHECK) Project structure created\n"

setup-go:
	@printf "\n$(CYAN)Initializing Go module$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(GO_BIN) mod init $(PROJECT_NAME) 2>/dev/null || true
	@$(GO_BIN) mod tidy >/dev/null 2>&1
	@$(GO_BIN) get -u github.com/a-h/templ@latest
	@printf "$(CHECK) Go modules ready\n"

setup-tailwind:
	@printf "\n$(CYAN)Setting up Tailwind CSS$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(BUN_BIN) init -y
	@$(BUN_BIN) install -D tailwindcss@latest @tailwindcss/cli
	@echo '@import "tailwindcss";' > $(STATIC_DIR)/css/input.css
	@printf "$(CHECK) Tailwind CSS ready\n"

build: check-deps
	@printf "\n$(CYAN)Building project$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(GO_BIN) build -o $(BUILD_DIR) $(MAIN_GO)
	@printf "$(CHECK) Build complete\n"

templ:
	@printf "\n$(CYAN)Generating templ files$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(TEMPL_BIN) generate
	@printf "$(CHECK) Templ files generated\n"

css:
	@printf "\n$(CYAN)Generating CSS$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@bunx @tailwindcss/cli -i $(STATIC_DIR)/css/input.css -o $(STATIC_DIR)/css/styles.css --minify
	@printf "$(CHECK) CSS generated\n"

watch:
	@printf "\n$(CYAN)Watching for changes$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@printf "\n$(CYAN)Rebuilding...$(RESET)\n"
	@bunx @tailwindcss/cli -i $(STATIC_DIR)/css/input.css -o $(STATIC_DIR)/css/styles.css --watch=always &
	@air

clean:
	@printf "\n$(CYAN)Cleaning build files$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@rm -rf $(BUILD_DIR)
	@find . -type f -name "*_templ.go" -delete
	@printf "$(CHECK) Clean complete\n"

serve: templ css
	@printf "\n$(CYAN)Starting server$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(GO_BIN) run $(MAIN_GO) serve

deploy: clean templ css
	@printf "\n$(CYAN)Building production binary$(RESET)\n"
	@GOOS=linux GOARCH=amd64 $(GO_BIN) build -ldflags="-s -w" -o $(PROD_BINARY) $(MAIN_GO)
	@printf "$(CHECK) Production build complete: $(PROD_BINARY)\n"

help:
	@printf "\n$(CYAN)Available commands$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@printf "  make init$(RESET)          $(ARROW) Initialize project\n"
	@printf "  make build$(RESET)         $(ARROW) Build project\n"
	@printf "  make watch$(RESET)         $(ARROW) Watch for changes\n"
	@printf "  make serve$(RESET)         $(ARROW) Start server\n"
