#!/bin/bash
echo " ---------- Downloading htui ---------- "
curl -sSL https://github.com/PierreKieffer/htui/raw/master/bin/64_osx/htui -o htui
chmod +x htui
sudo mv htui /usr/local/bin
echo " ---------- htui is installed ---------- "
echo ""
echo " ---------- usage ---------- "
echo ""
echo "    start cmd : htui "
echo ""
echo "
     -----------------------------
     -        Move around        -
     -----------------------------
     go up               ▲  or 'k'
     go down             ▼  or 'j'
     go to the top       'gg'
     go to the bottom    'G'
     select item         'enter'
     Quit htui           'q'
"
echo " --------------------------- "

