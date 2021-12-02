# Operating System

|#|category|config|source|
|---|---|---|---|
|1|operating_system|change hostname|https://github.com/RPi-Distro/raspi-config|
|2|operating_system|add user|https://www.digitalocean.com/community/tutorials/how-to-add-and-delete-users-on-ubuntu-18-04|
|3|operating_system|delete user|https://www.digitalocean.com/community/tutorials/how-to-add-and-delete-users-on-ubuntu-18-04|
|4|operating_system|change password|https://github.com/RPi-Distro/raspi-config|

### 1) Change hostname
1. regex on the front-end to have a well formatted hostname

    > RFCs mandate that a hostname's labels may contain only the ASCII letters 'a' through 'z' (case-insensitive), the digits '0' through '9', and the hyphen. Hostname labels cannot begin or end with a hyphen.No other symbols, punctuation characters, or blank spaces are permitted.

2. change hostname: POST /configure/changehostname?hostname=**hostname**
3. reboot

### 2) Add User
1. check if user exists: GET /humanusers

    > The user name should follow the following regex: ^[a-z-_][a-z0-9-_]{1,31}$

2. add or delete depending on the result: POST /configure/adduser?username=**username**&password=**password**

### 3) Delete User
1. check if user exists: GET /humanusers
2. add or delete depending on the result: POST /configure/deleteuser?username=**username**

### 4) Change password
1. check if user exists: GET /humanusers
2. change password : POST /configure/changepassword?password=**password**&username=**username**
3. no reboot needed

# Package Management

|#|category|config|source|
|---|---|---|---|
|1|package_management|update|https://www.cyberciti.biz/faq/how-do-i-update-ubuntu-linux-softwares/|
|2|package_management|upgrade|https://www.cyberciti.biz/faq/how-do-i-update-ubuntu-linux-softwares/|
|3|package_management|update & upgrade|https://www.cyberciti.biz/faq/how-do-i-update-ubuntu-linux-softwares/|

### 1) Update
1. check lastUpdate in Firestore: GET <TBD>.
2. update depending on the result: POST /configure/update

### 2) Upgrade
1. check lastUpgrade in Firestore: GET <TBD>.
2. upgrade depending on the result: POST /configure/upgrade

### 3) Update & Upgrade
1. check lastUpgrade in Firestore: GET <TBD>.
2. update & upgrade depending on the result: POST /configure/updateupgrade

# Network

|#|category|config|source|
|---|---|---|---|
|1|network|SSH|https://github.com/RPi-Distro/raspi-config|
|2|network|wifi country|https://github.com/RPi-Distro/raspi-config|
|3|network|wait for network at boot|https://github.com/RPi-Distro/raspi-config|
|4|network|VNC|https://github.com/RPi-Distro/raspi-config|

### 1) SSH
1. check isSSHKeyGenerating: GET /rpinterfaces. 
2. if true, try later
3. check isSSH: GET /rpinterfaces
4. depending on the result: POST /configure/ssh?action=**[enable/disable]**

### 2) Wifi Country
1. check isWifiInterfaces and isWpaSupCom for a given **[iface]**: GET /rpinterfaces.
2. change wifi country for wireless network: POST /configure/wificountry?iface=**[iface]**&country=**[code]**
3. reboot

### 3) Wait for network at boot
1. check if isWaitForNetworkAtBoot: GET /boots 
    > no need to check GET /configfiles because GET /boots does that
2. depending on that enable or disable: POST /configure/waitfornetworkatboot?action=**[enable/disable]**
3. no reboot needed

### 4) VNC
1. check isVNCInstalled: GET /softwares. 
CAUTION: isVNCInstalled is verified by two commands in raspi-config.
Here I verified with the one option only.
2. if true, continue
3. check isVNC: GET /rpinterfaces
4. depending on the result: POST /configure/vnc?action=**[enable/disable]**

# Screen

|#|category|config|source|
|---|---|---|---|
|1|screen|overscan|https://github.com/RPi-Distro/raspi-config|
|2|screen|blanking|https://github.com/RPi-Distro/raspi-config|

### 1) Overscan
1. check if file exists: GET /configfiles
2. if file exists, check isOverscan: GET /displays
3. depending on that enable or disable: POST /configure/overscan?action=**[enable/disable]**
4. reboot

### 2) Blanking
1. check if file exists: GET /configfiles
2. if file exists, check isXscreenSaverInstalled and isBlanking: GET /displays
3. if isXscreenSaverInstalled = true, it will override the raspi-config blanking config: ABORT here
4. if isXscreenSaverInstalled = false, then continue
5. enable or disable depending on isBlanking: POST /configure/blanking?action=**[enable/disable]**
PS: it will fail it try to enable while it is already enabled - same with disable status & disable
6. reboot

# Interface

|#|category|config|source|
|---|---|---|---|
|1|interface|camera|https://github.com/RPi-Distro/raspi-config|
|2|interface|SPI|https://github.com/RPi-Distro/raspi-config|
|3|interface|I2C|https://github.com/RPi-Distro/raspi-config|
|4|interface|one wire|https://github.com/RPi-Distro/raspi-config|
|5|interface|remote gpio|https://github.com/RPi-Distro/raspi-config|

### 1) Camera
1. check isStartXElf: GET /rpinterfaces. 
2. if false, update firmware
3. check isCamera: GET /rpinterfaces
4. depending on the result: POST /configure/camera?action=**[enable/disable]**

### 2) SPI
1. check isSPI: GET /rpinterfaces. 
2. depending on the result: POST /configure/spi?action=**[enable/disable]**
3. reboot

### 3) I2C
1. check isSPI: GET /rpinterfaces. 
2. depending on the result: POST /configure/i2c?action=**[enable/disable]**
3. reboot

### 4) One Wire
1. check isOneWire: GET /rpinterfaces. 
2. depending on the result: POST /configure/onewire?action=**[enable/disable]**
3. reboot

### 5) Remote GPIO
1. check isRemoteGpio: GET /rpinterfaces. 
2. depending on the result: POST /configure/rgpio?action=**[enable/disable]**
3. reboot
