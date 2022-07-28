git submodule init
git submodule update
for i in fswebcam/*
do ln -s $i
done
sed -i s/main\(/cmain\(/g fswebcam.c
./configure 
