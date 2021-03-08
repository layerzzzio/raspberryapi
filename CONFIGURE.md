## 1) Change hostname
1. regex on the front-end to have a well formatted hostname

    > RFCs mandate that a hostname's labels may contain only the ASCII letters 'a' through 'z' (case-insensitive), the digits '0' through '9', and the hyphen. Hostname labels cannot begin or end with a hyphen.No other symbols, punctuation characters, or blank spaces are permitted.

2. change hostname: POST /configure/changehostname?hostname=**hostname**
3. reboot

## 2) Change password
1. change password : POST /configure/changepassword?password=**password**&username=**username**
2. no reboot needed

## 3) Wait for network at boot
1. check if isWaitForNetworkAtBoot: GET /boots 
    > no need to check GET /configfiles because GET /boots does that
2. depending on that enable or disable: POST /configure/waitfornetworkatboot?action=**[enable/disable]**
3. no reboot needed

## 4) Overscan
1. check if file exists: GET /configfiles
2. if file exists, check isOverscan: GET /displays
3. depending on that enable or disable: POST /configure/overscan?action=**[enable/disable]**
4. reboot

## 5) Blanking
1. check if file exists: GET /configfiles
2. if file exists, check isXscreenSaverInstalled and isBlanking: GET /displays
3. if isXscreenSaverInstalled = true, it will override the raspi-config blanking config: ABORT here
4. if isXscreenSaverInstalled = false, then continue
5. enable or disable depending on isBlanking: POST /configure/blanking?action=**[enable/disable]**
PS: it will fail it try to enable while it is already enabled - same with disable status & disable
6. reboot

## 6) Add User
1. check if user exists: GET /humanusers
2. add or delete depending on the result: POST /configure/adduser?username=**username**&password=**password**

## 7) Delete User
1. check if user exists: GET /humanusers
2. add or delete depending on the result: POST /configure/deleteuser?username=**username**

## 8) Camera
1. check isStartXElf: GET /rpinterfaces. 
2. if false, update firmware
3. check isCamera: GET /rpinterfaces
4. depending on the result: POST /configure/camera?action=**[enable/disable]**

## 9) SSH
1. check isSSHKeyGenerating: GET /rpinterfaces. 
2. if true, try later
3. check isSSH: GET /rpinterfaces
4. depending on the result: POST /configure/ssh?action=**[enable/disable]**

## 10) VNC
1. check isVNCInstalled: GET /rpinterfaces. 
CAUTION: isVNCInstalled is verified by two commands in raspi-config.
Here I verified with the one option only.
2. if true, continue
3. check isVNC: GET /rpinterfaces
4. depending on the result: POST /configure/vnc?action=**[enable/disable]**


## 11) SPI
1. check isSPI: GET /rpinterfaces. 
2. depending on the result: POST /configure/spi?action=**[enable/disable]**
3. reboot