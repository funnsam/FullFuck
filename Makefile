GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)

ifeq ($(GOOS), windows)
	NAME = ffk.exe
else ifeq ($(GOOS), darwin)
	NAME = ffkd
else ifeq ($(GOOS), linux)
	NAME = ffk
else
	NAME = ffku
endif

ifeq ($(GOARCH), arm64)
	ifeq ($(GOOS), windows)
		NAME = ffkr.exe
	else
		NAME := $(NAME)r
	endif
endif

all:
	go build -o $(NAME) ./src

push:
	git add --all
	git commit -m %1
	git push -f -u origin HEAD:main