
###############################################################################
# Initial
###############################################################################
goget:
	go get -u -d -v ./...


###############################################################################
# Build
###############################################################################
bld:
	go build -o wallet ./cmd/wallet/main.go
	go build -o coldwallet ./cmd/coldwallet/main.go

run: bld
	./wallet -f 1

###############################################################################
# Test
###############################################################################
gotest:
	go test -v ./...


###############################################################################
# Docker compose
###############################################################################



###############################################################################
# Utility
###############################################################################
.PHONY: clean
clean:
	rm -rf detect