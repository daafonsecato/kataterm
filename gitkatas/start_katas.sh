
ttyd -p 7681 -i 0.0.0.0 --writable -w /home/git-katas-user/exercise bash &
code-server --auth none --bind-addr 0.0.0.0:8080 /home/git-katas-user/exercise &
# Copy or link files from the mounted volume to /app
cp -r /mnt/gitkatas/* /app/ || ln -s /mnt/gitkatas/* /app/
# Ensure the tmp directory exists and has the right permissions
mkdir -p /app/tmp
chmod -R 777 /app/tmp
cd /app
/usr/bin/air