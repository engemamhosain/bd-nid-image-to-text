# Mlkit
Overview
Mlkit is a Go-based application designed to run as a background service on Linux systems. It is configured to be managed by systemd, allowing for easy control, monitoring, and automatic startup at boot.

# Table of Contents
- Prerequisites
- Installation
    - Clone the Repository
    - Build the Application
    - Create a Systemd Service File
    - Reload systemd and Start the Service
    - Enable the Service on Boot
-Usage
  - Logs
  - Configuration
  - Uninstallation
  - Contributing
  - License
-Prerequisites
   - bee go framwork: Make sure Go is installed on your system (version 1.16 or newer is recommended).
  - Go: Make sure Go is installed on your system (version 1.16 or newer is recommended).
  - Linux: This guide assumes a Linux environment with systemd for service management.
  - Git: Required for cloning the repository (optional if downloading the source code directly).
 
## Installation
1. Clone the Repository
Clone the repository to your local machine using Git.
```bash
git clone https://github.com/engemamhosain/bd-nid-image-to-text.git
cd bd-nid-image-to-text
```
## Build the Application
```bash
go build -o myapp
```
## Move the Executable
```bash
sudo mv myapp /usr/local/bin/
```
## Create a Systemd Service File
```bash
sudo nano /etc/systemd/system/mlkit.service 
```
## mlkit json generate
[link](https://firebase.google.com/docs/ml-kit)

## Add the following
    [Service]
    WorkingDirectory=/home/MLKit/
    ExecStart=/home/MLKit//MLKit
    Environment="GOOGLE_APPLICATION_CREDENTIALS=/home/service-account-file.json"
    ExecReload=/bin/kill -HUP $MAINPID
    LimitNOFILE=65536
    Restart=always
    RestartSec=5
    
    
    [Install]
    WantedBy=multi-user.target



## Start and Enable the Service
```bash
    sudo systemctl daemon-reload
    sudo systemctl start mlkit.service 
    sudo systemctl enable mlkit.service 
```

## Contributing
- Fork the repository and submit a pull request. For major changes, discuss first.



