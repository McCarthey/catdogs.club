datanow=`date "+%Y-%m-%d %H:%M:%S"`
commitStr='[deploy] auto deploy to catdogs at '$datanow  
echo -e '\033[36m Wait for building\n \033[0m' && 
{
yarn build &&
echo -e '\033[35m Build done, wait for pushing to git\n \033[0m' &&
cp -r ./catdogs ../../web-build &&
cd ../../web-build &&
git pull &&
git add . &&
git commit -m "$commitStr" &&
git push &&
curl 'http://118.24.146.34:3001/'
echo -e '\033[42;37m Deploy successfully!\033[0m You can refresh \033[32m http://118.24.146.34 \033[0m now to see your changes\n'
} || {
    echo -e '\033[41;37m Error\033[0m when building or git operation, check it manully please.'
}
