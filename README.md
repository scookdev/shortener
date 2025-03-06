# 🚀 URL Shortener CLI (Go)

A simple, lightweight CLI tool written in Go that shortens URLs using TinyURL.

---

## 📥 Installation

First, clone the repository:

```sh
git clone https://github.com/scookdev/shortener.git
cd shortener
```

### 🛠️ **Building the Binary**

To build the executable:

```sh
make build
```

This creates a `shortener` binary.

---

## 🏃 Usage

### 🔗 **Shorten a URL**

```sh
./shortener --url https://equitas.cc
```

Alternatively, if you haven't built it yet:

```sh
make run
```

This will **build the binary first** and then run the command.

---

## ⚙️ Development

### 🔄 **Auto-Compile & Run on File Changes (Linux/macOS)**

For automatic recompilation whenever a file is modified:

```sh
make watch
```

This watches for `.go` file changes and reruns the program.

### 🧹 **Cleaning Up**

To remove the compiled binary:

```sh
make clean
```

---

## 🔥 Features

- 📏 **Simple CLI**: Just pass a URL to shorten it.
- ⚡ **Fast & Lightweight**: No dependencies except Go.
- 🏗️ **Easy to Build**: Uses `make build` for convenience.
- 🔄 **Auto-Rebuild**: `make watch` for live development.

---

## 📜 License

This project is licensed under the **MIT License**.
