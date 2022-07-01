BINARY=engine

COVERPKG := -coverpkg=./...
CMD_COVERAGE := go test -count=1 -coverprofile=coverage.out -coverpkg=./... ./core/...
CMD_TEST := go test -count=1 -failfast
EXCLUDE_FILE := mv ./coverage.out ./coverage.out.tmp; cat ./coverage.out.tmp | egrep -v "mocks" | egrep -v "util" > ./coverage.out && rm ./coverage.out.tmp

# you can include only spesific folder such as mode:|util|need_to_show_in_cover_show|another_path_that_important_to_see_in_cover_show
# then you need to run command like this ```make cover-show inc=1```
INCLUDE_FILE := mv ./coverage.out ./coverage.out.tmp; cat ./coverage.out.tmp | egrep -v "mocks" | egrep "mode:|util" > ./coverage.out && rm ./coverage.out.tmp

dep:
	go mod tidy

test:
	${CMD_TEST} ./...

test-verbose:
	${CMD_TEST} -v ./...

cover:
	${CMD_COVERAGE}
	# ${EXCLUDE_FILE}
	go tool cover -func=coverage.out

coverhtml:
	${CMD_COVERAGE}
	${EXCLUDE_FILE}
	go tool cover -html=coverage.out -o cover.html

cover-show:
	${CMD_COVERAGE}
ifeq ($(inc),)
	${EXCLUDE_FILE}
else
	${INCLUDE_FILE}
endif
	go tool cover -html=coverage.out

run:
	go run app/main/main.go

build:
	go build -o pos_lite.service app/main/main.go
