## Introduction
---
There are a variety of options when it comes to installing apps.

Here are some of the sources for those configurations:

| type | config | source |
|---|---|---|
| vpn | open vpn | https://www.ovpn.com/en/guides/raspberry-pi-raspbian |
| vpn | nord vpn | https://pimylifeup.com/raspberry-pi-nordvpn/ |
| vpn | surfshark | https://pimylifeup.com/raspberry-pi-surfshark/ |
| web browser | vivaldi | https://pimylifeup.com/raspberry-pi-vivaldi/ |
| web browser | chromium | https://pimylifeup.com/raspberry-pi-chromium-browser/ |
| web browser | firefox esr | https://pimylifeup.com/raspberry-pi-firefox/ |
| system | docker | https://phoenixnap.com/kb/docker-on-raspberry-pi |
| system | portainer | https://pimylifeup.com/raspberry-pi-portainer/ |
| system | syncthing | https://pimylifeup.com/raspberry-pi-syncthing/ |
| system | grafana | https://pimylifeup.com/raspberry-pi-grafana/ |
| system | influxdb | https://pimylifeup.com/raspberry-pi-influxdb/ |
| media | plex | https://pimylifeup.com/raspberry-pi-plex-server/ |
| media | jellyfin | https://pimylifeup.com/raspberry-pi-jellyfin/ |


## Logical flow
---
## VPN
---
### 1) Open VPN
#### 1.A) Install
1. check isOpenVPN & isUnzip: GET /softwares
2. install package: POST /appinstall/aptget?action=**[install/uninstall]**&pkg=openvpn

### 2) Nord VPN
#### 2.A) Install
1. check isOpenVPN & isNordVPN & isUnzip: GET /softwares
2. install NordVPN: POST /appinstall/vpnwithovpn?action=**[install/uninstall]**&vpnName=**[vpnName]**&url=**[url]

#### 2.B) Connect/Disconnect NordVPN
1. check isVyprVPN : GET /softwares
2. connect/disconnect NordVPN: POST /appaction/vpnwithovpn?action=**[connect/disconnect]**&vpnName=nordvpn&relativeConfigPath=**[relativeConfigPath]**&country=**[country]**&username**=**[username]**&password=**[password]**

### 3) Surfshark
#### 3.A) Install
1. check isOpenVPN & isSurfShark & isUnzip: GET /softwares
2. install Surfshark: POST /appinstall/vpnwithovpn?action=**[install/uninstall]**&vpnName=**[vpnName]**&url=**[url]

#### 3.B) Connect/Disconnect Surfshark
cf. 2.B or 5.B

### 4) IPVanish
#### 4.A) Install
1. check isOpenVPN & isIPVanish & isUnzip: GET /softwares
2. install IPVanish: POST /appinstall/vpnwithovpn?action=**[install/uninstall]**&vpnName=**[vpnName]**&url=**[url]

#### 4.B) Connect/Disconnect IPVanish
cf. 2.B or 5.B

### 5) VyprVPN
#### 5.A) Install
1. check isOpenVPN & isVyprVPN & isUnzip: GET /softwares
2. install VyprVPN: POST /appinstall/vpnwithovpn?action=**[install/uninstall]**&vpnName=**[vpnName]**&url=**[url]

#### 5.B) Connect/Disconnect VyprVPN
1. check isVyprVPN : GET /softwares
2. connect/disconnect VyprVPN: POST /appaction/vpnwithovpn?action=**[connect/disconnect]**&vpnName=**[vpnName]**&relativeConfigPath=**[relativeConfigPath]**&country=**[country]**&username**=**[username]**&password=**[password]**

Example: 
POST http://10.0.0.143:3333/v1/appaction/vpnwithovpn?action=connect&vpnName=vyprvpn&relativeConfigPath=GF_OpenVPN_20200320/OpenVPN256&country=France&username=josharchibal@gmail.com&password=H?j6ft4J9NJCq?JS

---
## Web Browser
---
### 1) Vivaldi
1. 

### 2) Chromium
1. 

### 3) Firefox ESR
1. 

---
## System
---
### 1) Docker
1. 

### 2) Portainer
1. 

### 3) Syncthing
1.

### 4) Grafana
1.

### 5) InfluxDB
1.

---
## Media
---
### 1) Plex
1. 

### 2) Jellyfin
1. 