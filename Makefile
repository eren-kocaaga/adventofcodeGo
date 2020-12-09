NUMBERS = 1 2
targets = $(foreach var,$(NUMBERS), day$(var)/main)
runall = $(foreach var,$(NUMBERS), $(var)_run)

run: $(runall)
build: $(targets)

day%/run::day%/main/main.go
	go build -o $@ $^

%_run: day%/run
	@echo "running $< ####################################################"
	./$< < $(subst run,,$<)main/input.txt

.PHONY: clean
clean:
	rm -fr $(targets)
