NAME = ffk # this is because the go compiler is stupid
ifeq ($(OS), Windows_NT)
	NAME = ffk.exe
endif

ffk:
	go build -o $(NAME) ./src

windows: # for the github action
	go build -o ffk.exe ./src

unix:
	go build -o ffk ./src

push:
	git add --all
	git commit -m %1
	git push
