# Benchmarking ssh-agent performance

This project uses Go benchmarking framework to measure performance of
ssh-agent. Turns out it's pretty fast.


## Usage

```console
$ make
ssh -V
OpenSSH_8.4p1 Debian-5+deb11u1, OpenSSL 1.1.1n  15 Mar 2022
chmod go-rwx key_ed25519 key_rsa4096
ssh-add -t 150 key_ed25519 key_rsa4096
Identity added: key_ed25519 (INSECURE TEST KEY)
Lifetime set to 150 seconds
Identity added: key_rsa4096 (INSECURE TEST KEY)
Lifetime set to 150 seconds
go test -bench=. -benchtime=5s
goos: linux
goarch: amd64
pkg: ssh-agent-benchmark
cpu: Intel Xeon Processor (Icelake)
BenchmarkSshAgent/key_ed25519/32B/unique-4          3722      1707378 ns/op
BenchmarkSshAgent/key_ed25519/32B/same-4            3632      1744296 ns/op
BenchmarkSshAgent/key_rsa4096/32B/unique-4           817      7802090 ns/op
BenchmarkSshAgent/key_rsa4096/32B/same-4             739      7634488 ns/op
BenchmarkSshAgent/key_ed25519/64B/unique-4          3518      1789925 ns/op
BenchmarkSshAgent/key_ed25519/64B/same-4            3657      1744986 ns/op
BenchmarkSshAgent/key_rsa4096/64B/unique-4           712      7750078 ns/op
BenchmarkSshAgent/key_rsa4096/64B/same-4             780      7706054 ns/op
BenchmarkSshAgent/key_ed25519/256B/unique-4         3664      1776065 ns/op
BenchmarkSshAgent/key_ed25519/256B/same-4           3529      1717271 ns/op
BenchmarkSshAgent/key_rsa4096/256B/unique-4          805      7564995 ns/op
BenchmarkSshAgent/key_rsa4096/256B/same-4            818      7428230 ns/op
BenchmarkSshAgent/key_ed25519/1024B/unique-4        3668      1732772 ns/op
BenchmarkSshAgent/key_ed25519/1024B/same-4          3548      1834202 ns/op
BenchmarkSshAgent/key_rsa4096/1024B/unique-4         818      7520015 ns/op
BenchmarkSshAgent/key_rsa4096/1024B/same-4           811      7562604 ns/op
BenchmarkSshAgent/key_ed25519/16384B/unique-4       3057      1826782 ns/op
BenchmarkSshAgent/key_ed25519/16384B/same-4         3165      1829706 ns/op
BenchmarkSshAgent/key_rsa4096/16384B/unique-4        790      7693554 ns/op
BenchmarkSshAgent/key_rsa4096/16384B/same-4          793      7690758 ns/op
PASS
ok    ssh-agent-benchmark    131.814s
```


## License and copyright

Copyright 2023 Vitaly Potyarkin

```
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```

