# Define the output binaries
LG_BINARY := bin\lg.exe
LV_BINARY := bin\lv.exe

# Define the Go compiler and build flags
GO := go
BUILD_FLAGS := -v

.PHONY: all lg lv install install-lg install-lv clean

all: lg lv

lg: $(LG_BINARY)

lv: $(LV_BINARY)

# Build the lg binary
$(LG_BINARY):
	@echo "Building lg..."
	$(GO) build $(BUILD_FLAGS) -o $(LG_BINARY) .\cmd\lg

# Build the lv binary
$(LV_BINARY):
	@echo "Building lv..."
	$(GO) build $(BUILD_FLAGS) -o $(LV_BINARY) .\cmd\lv

# Install the lg binary
install-lg: lg
	@echo "Installing lg binary..."
	$(GO) install .\cmd\lg

# Install the lv binary
install-lv: lv
	@echo "Installing lv binary..."
	$(GO) install .\cmd\lv

# Install both binaries
install: install-lg install-lv

clean:
	@echo "Cleaning up..."
	del /q $(LG_BINARY) $(LV_BINARY)
