datanow=`date "+%Y-%m-%d %H:%M:%S"`
commitStr='[deploy] auto deploy '$datanow  
echo 'wait for building' && 
yarn build &&
echo 'build done wait for git push' &&
cp -r ./catdogs ../../web-build &&
cd ../../web-build &&
git pull &&
git add . &&
git commit -m '$commitStr' &&
git push &&
curl 'http://118.24.146.34:3001/'

