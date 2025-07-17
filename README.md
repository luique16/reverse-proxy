<h1 align="center"> Reverse Proxy </h1>

<p align="center"> A reverse proxy written in Go. </p>

## 📦 Installation

You can install the reverse-proxy cli using the following command:
```bash
go install github.com/luique16/reverse-proxy@latest
```

Or you can download the code and build it yourself:

```bash
git clone https://github.com/luique16/reverse-proxy.git
cd reverse-proxy
go build
```

## 🚀 Usage

```bash
reverse-proxy -h

# Or, if you have built the binary

./reverse-proxy -h
```

## 📝 TODO

- [ ] Check health of each instance
- [ ] Block DDoS attacks
- [ ] Transform to a library

