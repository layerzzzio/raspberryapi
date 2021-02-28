## Change hostname
1. regex on the front-end to have a well formatted hostname

    > RFCs mandate that a hostname's labels may contain only the ASCII letters 'a' through 'z' (case-insensitive), the digits '0' through '9', and the hyphen. Hostname labels cannot begin or end with a hyphen.No other symbols, punctuation characters, or blank spaces are permitted.

2. change hostname: POST /configure/changehostname?hostname=**hostname**
3. reboot

## Change password
1. change password : POST /configure/changepassword?password=**password**&username=**username**
2. no reboot needed

## Wait for network at boot
1. check isWaitForNetworkAtBoot: GET /boots
2. depending on that enable or disable: POST /configure/waitfornetworkatboot?action=**[enable/disable]**
3. no reboot needed

## Overscan
1. check isOverscan: GET /displays (if not the file might not exist)
2. depending on that enable or disable: POST /configure/overscan?action=**[enable/disable]**
3. reboot