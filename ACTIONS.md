## Change hostname
1. regex on the front-end to have a well formatted hostname

    > RFCs mandate that a hostname's labels may contain only the ASCII letters 'a' through 'z' (case-insensitive), the digits '0' through '9', and the hyphen. Hostname labels cannot begin or end with a hyphen.No other symbols, punctuation characters, or blank spaces are permitted.

2. change hostname: POST /configure/changehostname?hostname=**hostname**
3. reboot

## Change password
1. change password : POST /configure/changepassword?password=**password**&username=**username**
2. no reboot needed

## Wait for network at boot
1. check if isWaitForNetworkAtBoot: GET /boots 
    > no need to check GET /configfiles because GET /boots does that
2. depending on that enable or disable: POST /configure/waitfornetworkatboot?action=**[enable/disable]**
3. no reboot needed

## Overscan
1. check if file exists: GET /configfiles
2. if file exists, check isOverscan: GET /displays
3. depending on that enable or disable: POST /configure/overscan?action=**[enable/disable]**
4. reboot

## Blanking
1. check if file exists: GET /configfiles
2. if file exists, check isXscreenSaverInstalled and isBlanking: GET /displays
3. if isXscreenSaverInstalled = true, it will override the raspi-config blanking config: ABORT here
4. if isXscreenSaverInstalled = false, then continue
5. enable or disable depending on isBlanking: POST /configure/overscan?action=**[enable/disable]**
PS: it will fail it try to enable while it is already enabled - same with disable status & disable
6. reboot