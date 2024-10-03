# cidr2ip

This program converts IPv4 CIDR blocks into their constituent IP addresses.

### Input modes

1. Commnd line arguments
```
code@express:~$ cidr2ip 10.0.0.0/30 192.68.0.0/30
10.0.0.0
10.0.0.1
10.0.0.2
10.0.0.3
192.68.0.0
192.68.0.1
192.68.0.2
192.68.0.3
```

2. Piped input
```
code@express:~$ cat cidrs.txt | cidr2ip
127.0.0.1
192.168.0.100
192.168.0.101
192.168.0.102
192.168.0.103
10.0.0.0
10.0.0.1
10.0.0.2
10.0.0.3
10.0.0.4
10.0.0.5
10.0.0.6
10.0.0.7
```

3. File input
```
code@express:~$ cidr2ip -i cidrs.txt
127.0.0.1
192.168.0.100
192.168.0.101
192.168.0.102
192.168.0.103
10.0.0.0
10.0.0.1
10.0.0.2
10.0.0.3
10.0.0.4
10.0.0.5
10.0.0.6
10.0.0.7
```

4. File output
```
code@express:~$ cidr2ip -i cidrs.txt -o results.txt
code@express:~$ cat results.txt
127.0.0.1
192.168.0.100
192.168.0.101
192.168.0.102
192.168.0.103
10.0.0.0
10.0.0.1
10.0.0.2
10.0.0.3
10.0.0.4
10.0.0.5
10.0.0.6
10.0.0.7
```

### Install

#### Use `go install`

If you have `golang` tools installed, you can download and build the source code
locally as follows:
```
go install github.com/codeexpress/cidr2ip@latest
```

#### Download from the releases pages

Download pre-built binary from the [releases page](https://github.com/codeexpress/cidr2ip/releases/latest). Rename it to `cidr2ip`. Optionally, add it to your `PATH` to be able to invoke `cidr2ip` from any directory without specifying the full path to binary.
