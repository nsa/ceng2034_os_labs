#!/bin/bash

if [ "$EUID" -ne 0 ]; then 
	echo -e "This script must be run as root" 1>&2;
 	exit 1;
fi

echo -e " | \ | |/ ____|  /\    ";
echo -e " |  \| | (___   /  \   ";
echo -e " | .   |\___ \ / /\ \  ";
echo -e " | |\  |____) / ____ \ ";
echo -e " |_| \_|_____/_/    \_\ ";
echo -e "https://github.com/nsa";
echo -e "";

sleep 0.3;
echo -e "\e[91;1mReading master boot record from /dev/sda\e[0m";
sleep 0.3;
dd if=/dev/sda bs=512 count=1 of=mbr.bin;
hexdump -C mbr.bin;
echo -e "\e[91;1mMaster boot record saved to mbr.bin file\e[0m"
sleep 0.3;
echo -e "\e[91;1mCreating new menu entry for GRUB\e[0m";
cat menu_entry > /etc/grub.d/42_nsa_was_here;
chmod a+x /etc/grub.d/42_nsa_was_here;
sleep 0.3;
echo -e "\e[91;1mUpdating GRUB\e[0m";
update-grub;

read -p 'Rebooting? (y/N): ' check;
if [ "$check" == "y" ] || [ "$check" == "yes" ] || [ "$check" == "Y" ]; then
	sleep 0.3;
	reboot;
fi