#!/bin/sh

source /usr/sbin/helper.sh

# Does the firmware support addons?
nvram get rc_support | grep -q am_addons
if [ $? != 0 ]
then
    echo "This firmware does not support addons!"
    exit 5
fi

# Obtain the first available mount point in $am_webui_page
am_get_webui_page /tmp/mnt/liao/shadow/index.asp

if [ "$am_webui_page" = "none" ]
then
    echo "Unable to install shadow"
    exit 5
fi
echo "Mounting shadow as $am_webui_page"

# Copy custom page
ln -s /tmp/mnt/liao/shadow/index.asp /www/user/$am_webui_page
ln -s /tmp/mnt/liao/shadow/index.js /www/user/shadowApp.js

# Copy menuTree (if no other script has done it yet) so we can modify it
if [ ! -f /tmp/menuTree.js ]
then
    echo "Copying menuTree.js..." 
    cp /www/require/modules/menuTree.js /tmp/
    mount -o bind /tmp/menuTree.js /www/require/modules/menuTree.js
fi

# Insert link at the end of the Tools menu.  Match partial string, since tabname can change between builds (if using an AS tag)
echo "Inserting shadow menu..."
sed -i "/url: \"Tools_OtherSettings.asp\", tabName:/a {url: \"$am_webui_page\", tabName: \"Shadow\"}," /tmp/menuTree.js

# sed and binding mounts don't work well together, so remount modified file
umount /www/require/modules/menuTree.js && mount -o bind /tmp/menuTree.js /www/require/modules/menuTree.js

cd /tmp/mnt/liao/shadow

# Starting v2ctl deamon...
echo "Starting v2ctl deamon..."
./v2ctl server 1>>./v2ctl.log 2>&1 &

