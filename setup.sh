wget https://github.com/fsphil/fswebcam/archive/refs/heads/master.zip
unzip -j master.zip
sed -i s/main\(/cmain\(/g fswebcam.c
./configure 
