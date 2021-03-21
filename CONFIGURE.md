## Introduction
---
There are a variety of options when it comes to configuration.

Here are some of the sources for those configurations:

| config | source |
|---|---|
| change hostname | https://github.com/RPi-Distro/raspi-config |
| change password | https://github.com/RPi-Distro/raspi-config |
| wait for network at boot | https://github.com/RPi-Distro/raspi-config |
| overscan | https://github.com/RPi-Distro/raspi-config |
| blanking | https://github.com/RPi-Distro/raspi-config |
| camera | https://github.com/RPi-Distro/raspi-config |
| SSH | https://github.com/RPi-Distro/raspi-config |
| VNC | https://github.com/RPi-Distro/raspi-config |
| SPI | https://github.com/RPi-Distro/raspi-config |
| I2C | https://github.com/RPi-Distro/raspi-config |
| one wire | https://github.com/RPi-Distro/raspi-config |
| remote gpio | https://github.com/RPi-Distro/raspi-config |
| add user | https://www.digitalocean.com/community/tutorials/how-to-add-and-delete-users-on-ubuntu-18-04 |
| delete user | https://www.digitalocean.com/community/tutorials/how-to-add-and-delete-users-on-ubuntu-18-04 |
| update | https://www.cyberciti.biz/faq/how-do-i-update-ubuntu-linux-softwares/ |
| upgrade | https://www.cyberciti.biz/faq/how-do-i-update-ubuntu-linux-softwares/ |
| update & upgrade | https://www.cyberciti.biz/faq/how-do-i-update-ubuntu-linux-softwares/ |
| wifi country | https://github.com/RPi-Distro/raspi-config |

## Logical flow
---
### 1) Change hostname
1. regex on the front-end to have a well formatted hostname

    > RFCs mandate that a hostname's labels may contain only the ASCII letters 'a' through 'z' (case-insensitive), the digits '0' through '9', and the hyphen. Hostname labels cannot begin or end with a hyphen.No other symbols, punctuation characters, or blank spaces are permitted.

2. change hostname: POST /configure/changehostname?hostname=**hostname**
3. reboot

### 2) Change password
1. change password : POST /configure/changepassword?password=**password**&username=**username**
2. no reboot needed

### 3) Wait for network at boot
1. check if isWaitForNetworkAtBoot: GET /boots 
    > no need to check GET /configfiles because GET /boots does that
2. depending on that enable or disable: POST /configure/waitfornetworkatboot?action=**[enable/disable]**
3. no reboot needed

### 4) Overscan
1. check if file exists: GET /configfiles
2. if file exists, check isOverscan: GET /displays
3. depending on that enable or disable: POST /configure/overscan?action=**[enable/disable]**
4. reboot

### 5) Blanking
1. check if file exists: GET /configfiles
2. if file exists, check isXscreenSaverInstalled and isBlanking: GET /displays
3. if isXscreenSaverInstalled = true, it will override the raspi-config blanking config: ABORT here
4. if isXscreenSaverInstalled = false, then continue
5. enable or disable depending on isBlanking: POST /configure/blanking?action=**[enable/disable]**
PS: it will fail it try to enable while it is already enabled - same with disable status & disable
6. reboot

### 6) Add User
1. check if user exists: GET /humanusers
2. add or delete depending on the result: POST /configure/adduser?username=**username**&password=**password**

### 7) Delete User
1. check if user exists: GET /humanusers
2. add or delete depending on the result: POST /configure/deleteuser?username=**username**

### 8) Camera
1. check isStartXElf: GET /rpinterfaces. 
2. if false, update firmware
3. check isCamera: GET /rpinterfaces
4. depending on the result: POST /configure/camera?action=**[enable/disable]**

### 9) SSH
1. check isSSHKeyGenerating: GET /rpinterfaces. 
2. if true, try later
3. check isSSH: GET /rpinterfaces
4. depending on the result: POST /configure/ssh?action=**[enable/disable]**

### 10) VNC
1. check isVNCInstalled: GET /softwares. 
CAUTION: isVNCInstalled is verified by two commands in raspi-config.
Here I verified with the one option only.
2. if true, continue
3. check isVNC: GET /rpinterfaces
4. depending on the result: POST /configure/vnc?action=**[enable/disable]**

### 11) SPI
1. check isSPI: GET /rpinterfaces. 
2. depending on the result: POST /configure/spi?action=**[enable/disable]**
3. reboot

### 12) I2C
1. check isSPI: GET /rpinterfaces. 
2. depending on the result: POST /configure/i2c?action=**[enable/disable]**
3. reboot

### 13) One Wire
1. check isOneWire: GET /rpinterfaces. 
2. depending on the result: POST /configure/onewire?action=**[enable/disable]**
3. reboot

### 14) Remote GPIO
1. check isRemoteGpio: GET /rpinterfaces. 
2. depending on the result: POST /configure/rgpio?action=**[enable/disable]**
3. reboot

### 15) Update
1. check lastUpdate in Firestore: GET <TBD>.
2. update depending on the result: POST /configure/update

### 16) Upgrade
1. check lastUpgrade in Firestore: GET <TBD>.
2. upgrade depending on the result: POST /configure/upgrade

### 16) Update & Upgrade
1. check lastUpgrade in Firestore: GET <TBD>.
2. update & upgrade depending on the result: POST /configure/updateupgrade

### 17) Wifi Country
1. check isWifiInterfaces and isWpaSupCom for a given **[iface]**: GET /rpinterfaces.
2. change wifi country for wireless network: POST /configure/wificountry?iface=**[iface]**&country=**[code]**
3. reboot