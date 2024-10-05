
# CLI Projects to Streamline Developer Experience 🚀

This repository contains a collection of **cool CLI tools** aimed at simplifying and improving common developer workflows. Each tool is built to address specific tasks, reducing the manual effort and making your development life a bit easier!

## Projects Overview

### 1. **gclear.go** 🧹

The `gclear` CLI tool helps you clean up your git branches by deleting all branches except for `main`. It’s a great way to keep your git workspace tidy when you’ve accumulated a lot of stale or unused branches.

- **How it works**: The tool fetches a list of all branches and automatically deletes any branch that is not named `main`.  
- **Command**:  
  ```bash
  gclear
  ```
  This command will remove all local git branches except for the main branch.

### 2. **sst.go** ⏰

The `sst` (System Set Time) CLI tool helps synchronize your system time with an external time API, based on your system's current timezone. It also ensures that your hardware clock is updated accordingly.

- **How it works**:
  - The tool fetches your system’s timezone from `/etc/timezone` or the `date` command as a fallback.
  - It then makes a request to the [timeapi.io](https://timeapi.io) service to get the current time.
  - The tool sets the system time and updates the hardware clock.
  
- **Command**:  
  ```bash
  sst
  ```
  This command will fetch the current time for your system's timezone, set your system clock, and update the hardware clock.

---

## Installation 🛠️

To use any of these CLI tools, clone this repository and build the respective tools from the source code.

```bash
git clone https://github.com/your-username/cli-tools.git
cd cli-tools
```

### Building a Tool
For example, to build `gclear.go`:

```bash
cd gclear
go build -o gclear gclear.go
```

For `sst.go`:

```bash
cd sst
go build -o sst sst.go
```

Once built, you can move the executables to a directory in your `$PATH` (e.g., `/usr/local/bin/`) for easy access.

---

### Happy Coding! 💻

If you enjoy using these tools, please consider giving the repository a star ⭐!