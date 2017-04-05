PORT = 8008
SRC  = web.go
OBJ  = rbrian.bit

run:
	PORT=$(PORT) go run $(SRC)

heroku:
	git push herouk master

clean:
	@rm $(OBJ)

.PHONY: run heroku clean
