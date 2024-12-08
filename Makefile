# Define the number as the second word in the command
DAY := $(word 2, $(MAKECMDGOALS))
PADDED_DAY := $(shell printf "%02d" $(DAY)) # Ensure 01, 02, etc.

# Base directories
SOURCE_DIR := ./template/dayXX
SOLUTIONS_DIR := ./solutions
TARGET_DIR := $(SOLUTIONS_DIR)/day$(PADDED_DAY)

.PHONY: run
run:
	@echo "Running Advent Day $(DAY)"
	@echo "-------------------------"
	@go run main.go --day=$(DAY)

# Ensure the input is valid
check:
	@if ! echo "$(DAY)" | grep -Eq '^[0-9]+$$'; then \
	    echo "Error: You must provide a valid day number (e.g., 1-99)."; \
	    exit 1; \
	fi

new-day: check
	@cp -r $(SOURCE_DIR) $(TARGET_DIR) && \
	find $(TARGET_DIR) -type f -exec sed -i '' -e "s/dayXX/day$(PADDED_DAY)/g" -e "s/DayXX/Day$(PADDED_DAY)/g" {} \; && \
	find $(TARGET_DIR) -type f -name "*XX*" -exec bash -c 'mv "$$0" "$${0/XX/$(PADDED_DAY)}"' {} \; && \
	echo "Created new day $(PADDED_DAY)"
	@echo "Make sure the update ./solutions/day_factory/day_factory.go"

# Pattern rule to prevent make from misinterpreting the number
%:
	@: