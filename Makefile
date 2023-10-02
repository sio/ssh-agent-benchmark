KEYS=key_ed25519 key_rsa4096
GO?=go

.PHONY: bench
bench: load-keys
	$(GO) test -bench=. -benchtime=5s -benchmem

.PHONY: load-keys
load-keys:
	ssh -V
	chmod go-rwx $(KEYS)
	ssh-add -t 200 $(KEYS)
