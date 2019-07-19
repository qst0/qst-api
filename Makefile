default: build

src = qst-api.go handler.go logger.go error.go router.go routes.go

build:
	go build $(src)

clean:
	@rm
