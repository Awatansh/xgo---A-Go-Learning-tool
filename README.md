# xgo 🧭

**`xgo`** is a terminal-based CLI tool to track your progress through a 60-day Go learning roadmap — including fundamentals, web dev, microservices, DevOps, and real-time systems.

No external files needed after installation. Your progress is saved locally and can be reset anytime.

---
** Built on MANJARO-OS **

## 📦 Features

- ✅ View current step with `xgo -now`
- 🎯 Mark step as done with `xgo -done`
- 🔁 Reset progress with `xgo -reset`
- 📂 Stores progress in `~/.config/xgo/progress.json`
- 🔒 Self-contained — no external roadmap file needed after build

---

## 🚀 Installation

```bash
git clone https://github.com/Awatansh/xgo.git
cd xgo
go build -o xgo main.go
sudo mv xgo /usr/local/bin
