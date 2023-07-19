KEYS=key_ed25519 key_rsa4096
GO?=go

.PHONY: bench
bench: load-keys
	$(GO) test -bench=. -benchtime=5s

.PHONY: load-keys
load-keys:
	ssh-add -t 60 $(KEYS)
