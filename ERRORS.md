# Known errors

Below are a list of errors observed while testing Raspibuddy against the two Raspberry Zero W1 used for testing purposes.

## A) VPN
### A.1) when clicking on a VPN client
**Description**
In /etc/openvpn/wov_nordvpn for example, if vpnconfigs.zip exists but not vpnconfigs, the function VPNCountries fails. It results in the NordVPN Client to load forever.

**Don't**
Never delete vpnconfigs manually only.

**Do**
Instead delete the entire nordvpn folder: /etc/openvpn/wov_nordvpn.

**Correction**
N/A

**Status**
N/A

## B) Package management
### B.1) dpkg issue when upgrading
**Description**
The error message is as followed when running `sudo apt-get upgrade -y` manually in the shell:
`error dpkg was interrupted, you must manually run 'sudo dpkg --configure -a' to correct the problem`

**Don't**
N/A

**Do**
N/A

**Correction**
To add `sudo dpkg --configure -a` in the upgrade process.

**Status**
Correction done.