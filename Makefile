NAME = ffk # this is because the go compiler is stupid
ifeq ($(OS), Windows_NT)
	NAME = ffk.exe
endif

ffk:
	go build -o $(NAME) src/tokens.go src/main.go

windows: # for the github action
	go build -o ffk.exe src/tokens.go src/main.go

unix:
	go build -o ffk src/tokens.go src/main.go

push:
	git add --all
	git commit -m %1
	git push
