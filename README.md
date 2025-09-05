# 🕵️ go-sniffer

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)
![Go Version](https://img.shields.io/badge/Language-Go-blue)
[![CLI](https://img.shields.io/badge/type-CLI-blueviolet)](#)
[![pcap](https://img.shields.io/badge/pcap-supported-orange)](#)
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macos%20%7C%20windows-blue)


A lightweight network packet sniffer written in Go.
This tool captures network traffic in real-time, analyzes packets, and provides readable statistics, including IP-to-domain resolution and traffic details per host or connection.

### 🔍 What is a Sniffer?
A network sniffer is a tool that monitors network traffic by capturing packets on a given network interface. It helps developers, network administrators, and security engineers analyze communication between hosts, identify issues, and debug applications.
This project is a simple and educational implementation of a sniffer that uses Go’s concurrency model and libraries to provide a clean, minimal, and extensible codebase.

---

### ✨ Features
📡 Capture packets from a specified network interface

🎯 Support for BPF filters to capture only relevant packets

⏱️ Optional capture duration (or run indefinitely)

🌍 Reverse DNS lookups (resolve IP addresses to domain names)

📊 Display per-IP statistics: Number of packets sent/received Total bytes (formatted in KB/MB/GB)

🔗 Show source-to-destination traffic pairs for better network visibility

⚙️ Clean modular design with a focus on extensibility

---

### 🔎 How It Works
The sniffer follows a pipeline-based design where each component focuses on a single responsibility:
```bash
┌─────────────┐      ┌──────────────┐      ┌──────────────┐      ┌──────────────┐
│ Packet Data │ ---> │  Sniffer     │ ---> │  Stats       │ ---> │   Printer    │
│ (from NIC)  │      │ (Capture &   │      │ (Aggregate   │      │ (Pretty      │
│             │      │  Parse)      │      │  Traffic)    │      │  Output)     │
└─────────────┘      └──────────────┘      └──────────────┘      └──────────────┘
                           │
                           ▼
                     ┌──────────────┐
                     │  Resolver    │
                     │ (IP → Domain)│
                     └──────────────┘
```

Sniffer – Captures raw packets from the selected interface using gopacket.

Resolver – Resolves IP addresses to domain names via reverse DNS lookups.

Stats – Aggregates packet counts and total bytes.

Printer – Displays sorted and human-readable network stats.

---

### 🚀 Getting Started

if you have macos on intel as me, you can download binary in releases, because cross compile don't work with pcap and cgo

#### 1 Clon repo

```cmd
git clone https://github.com/alexcfv/go-sniffer.git
```

#### 2 Go into the project and install dependencies:

```cmd
go mod tidy
```
#### 3 Build the sniffer
(linux and macos)

```cmd
go build -o go-sniffer .
```

(windows MinGW or MSY32)
```cmd
go build -o go-sniffer.exe .
```
#### 4 Run the sniffer

```cmd
sudo ./go-sniffer -i <interface> [options]
```

#### 5 Options

| Flag        | Description                                          |
| ----------- | ---------------------------------------------------- |
| `-i string` | Network interface to capture packets (**required**)  |
| `-f string` | Optional BPF filter (e.g. `tcp port 80`)         |
| `-t int`    | Capture duration in seconds (`0` = run indefinitely, CTR+C for break) |
| `-resolve`  | Enable IP → domain resolution                        |

---

### 🌐 Choosing the Right Network Interface
Network interface names vary by operating system. Below are common examples:

| OS          | Example Interfaces  | Description                                                                     |
| ----------- | ------------------- | ------------------------------------------------------------------------------- |
| **Linux**   | `eth0`, `eth1`      | Wired Ethernet connections                                                      |
|             | `wlan0`, `wlp3s0`   | Wireless/Wi-Fi connections                                                      |
|             | `lo`                | Loopback (localhost, used for testing only)                                     |
| **macOS**   | `en0`, `en1`        | Typically `en0` is Wi-Fi; `en1` may be wired or virtual                         |
|             | `lo0`               | Loopback interface                                                              |
| **Windows** | `Ethernet`, `Wi-Fi` | Interface names match adapter names (quotes may be required, e.g. `-i "Wi-Fi"`) |

### 📌 Tip:
On Linux/macOS, run ifconfig or ip link to list interfaces.
On Windows, use Get-NetAdapter in PowerShell or ipconfig /all.
Choose the interface connected to the network you want to monitor.

### 🔐 Permissions
Sniffing network traffic requires elevated privileges on all operating systems:

| OS              | How to Run                                                        |
| --------------- | ----------------------------------------------------------------- |
| **Linux/macOS** | Run with `sudo`: <br> `sudo ./go-sniffer -i eth0`                 |
| **Windows**     | Run the terminal as **Administrator** before starting the sniffer |

Without proper permissions, the sniffer will not be able to capture traffic.

---

### 🧩 Example
input
```cmd
sudo ./go-sniffer -i eth0 -t 30 -resolve
```
This captures traffic for 30 seconds on eth0 and shows domain names instead of raw IPs.

output
```cmd
Capture stopped. Summary:

=== Traffic Summary ===
Total Packets: 53
Total Bytes:   17651
TCP:           27
UDP:           26
ICMP:          0
Other:         0

=== IP/Domain Stats ===
IP/Domain                                         Packets     Bytes     
-------------------------------------------------------------------------
[PUBLIC] 2a06:63c1:110a:6c00:985c:a6e5:f994:2423  29          14.0 KB   
[PUBLIC] demetra1.isp.telekom.rs.                 13          1.4 KB    
[PUBLIC] 2a00:ff0:1234:fe80::1                    11          1.9 KB    

=== Traffic by IP pairs ===
SRC:PORT -> DST:PORT                                                        PACKETS
------------------------------------------------------------------------------------
2a00:ff0:1234:fe80::1:443 -> 2a06:63c1:110a:6c00:985c:a6e5:f994:2423:60227  16
2a06:63c1:110a:6c00:985c:a6e5:f994:2423:60227 -> 2a00:ff0:1234:fe80::1:443  11
2a06:63c1:110a:6c00:985c:a6e5:f994:2423:36653 -> 2a00:e90:0:3:3:3:3:3:53    1
2a06:63c1:110a:6c00:985c:a6e5:f994:2423:6725 -> 2a00:e90:0:3:3:3:3:3:53     1
2a06:63c1:110a:6c00:985c:a6e5:f994:2423:50456 -> 2a00:e90:0:3:3:3:3:3:53    1
2a00:e90:0:3:3:3:3:3:53 -> 2a06:63c1:110a:6c00:985c:a6e5:f994:2423:50456    1
2a06:63c1:110a:6c00:985c:a6e5:f994:2423:6029 -> 2a00:e90:0:3:3:3:3:3:53     1
```

### 🛠️ Technologies Used

Go – core language

gopacket – for packet capturing and decoding

net – for reverse DNS resolution

map-based stats – simple in-memory aggregation
