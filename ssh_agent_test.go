package main

import (
	"testing"

	"fmt"
	"math/rand"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func BenchmarkSshAgent(b *testing.B) {
	benchmarks := []struct {
		keyfile string
		msgsize int
	}{
		{"key_ed25519", 32},
		{"key_rsa4096", 32},
		{"key_ed25519", 64},
		{"key_rsa4096", 64},
		{"key_ed25519", 256},
		{"key_rsa4096", 256},
		{"key_ed25519", 1024},
		{"key_rsa4096", 1024},
		{"key_ed25519", 16 * 1024},
		{"key_rsa4096", 16 * 1024},
	}
	socket := os.Getenv("SSH_AUTH_SOCK")
	if socket == "" {
		b.Fatalf("environment variable not set: SSH_AUTH_SOCK")
	}
	conn, err := net.Dial("unix", socket)
	if err != nil {
		b.Fatalf("failed to open SSH_AUTH_SOCK: %v", err)
	}
	defer conn.Close()
	sshAgent := agent.NewClient(conn)
	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("%s/%dB", bm.keyfile, bm.msgsize), func(b *testing.B) {
			pubKeyRaw, err := os.ReadFile(bm.keyfile + ".pub")
			if err != nil {
				b.Fatalf("failed to read public key: %v", err)
			}
			pubKey, _, _, _, err := ssh.ParseAuthorizedKey(pubKeyRaw)
			if err != nil {
				b.Fatalf("failed to parse public key: %v", err)
			}
			msg := make([]byte, bm.msgsize)
			n, err := rand.Read(msg)
			if err != nil {
				b.Fatalf("failed to generate random message: %v", err)
			}
			if n != len(msg) {
				b.Fatalf("random generator returned %d bytes instead of %d", n, len(msg))
			}
			for _, unique := range []bool{true, false} {
				b.Run(repr(unique), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						if unique {
							n, err := rand.Read(msg)
							if err != nil {
								b.Fatalf("failed to generate random message: %v", err)
							}
							if n != len(msg) {
								b.Fatalf("random generator returned %d bytes instead of %d", n, len(msg))
							}
						}
						_, err = sshAgent.Sign(pubKey, msg) // all allocations happen here (ed25519 = 31 allocs/op, rsa = 38 allocs/op)
						if err != nil {
							b.Fatalf("message signing failed: %v", err)
						}
					}
				})
			}
		})
	}
}

func repr(unique bool) string {
	if unique {
		return "unique"
	}
	return "same"
}
