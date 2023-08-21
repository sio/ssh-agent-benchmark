# Benchmarking ssh-agent performance

This project uses Go benchmarking framework to measure performance of
ssh-agent. Turns out it's pretty fast.

See the [blog post] for more information.

[blog post]: https://potyarkin.com/posts/2023/ssh-agent-benchmark/


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
BenchmarkSshAgent/key_ed25519/32B/unique-4          3553      1763363 ns/op
BenchmarkSshAgent/key_ed25519/32B/same-4            3568      1708270 ns/op
BenchmarkSshAgent/key_rsa4096/32B/unique-4           778      7824780 ns/op
BenchmarkSshAgent/key_rsa4096/32B/same-4             763      7657785 ns/op
BenchmarkSshAgent/key_ed25519/64B/unique-4          3456      1752457 ns/op
BenchmarkSshAgent/key_ed25519/64B/same-4            3598      1733750 ns/op
BenchmarkSshAgent/key_rsa4096/64B/unique-4           781      7639828 ns/op
BenchmarkSshAgent/key_rsa4096/64B/same-4             798      7720210 ns/op
BenchmarkSshAgent/key_ed25519/256B/unique-4         3549      1735906 ns/op
BenchmarkSshAgent/key_ed25519/256B/same-4           3417      1722301 ns/op
BenchmarkSshAgent/key_rsa4096/256B/unique-4          698      7738767 ns/op
BenchmarkSshAgent/key_rsa4096/256B/same-4            787      7625366 ns/op
BenchmarkSshAgent/key_ed25519/1024B/unique-4        3555      1703601 ns/op
BenchmarkSshAgent/key_ed25519/1024B/same-4          3651      1633226 ns/op
BenchmarkSshAgent/key_rsa4096/1024B/unique-4         805      7542115 ns/op
BenchmarkSshAgent/key_rsa4096/1024B/same-4           810      7437307 ns/op
BenchmarkSshAgent/key_ed25519/16384B/unique-4       3205      1935190 ns/op
BenchmarkSshAgent/key_ed25519/16384B/same-4         3296      1907921 ns/op
BenchmarkSshAgent/key_rsa4096/16384B/unique-4        783      7624776 ns/op
BenchmarkSshAgent/key_rsa4096/16384B/same-4          759      7620548 ns/op
PASS
ok    ssh-agent-benchmark    130.115s
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
