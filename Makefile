ARG := $(word 2, $(MAKECMDGOALS))

.PHONY: day
day:
	go run main.go --day=$(ARG)

%:
	@: