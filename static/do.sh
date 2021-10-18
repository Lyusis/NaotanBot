#!/bin/bash
# 30 4 * * * sh /root/workspace/do.sh

pgrep -f ./main20210717

sp_pid=$(pgrep -f ./naotanbot)
if [ -z "$sp_pid" ];
then
  echo "\033[31m [ not find bot pid ] \033[0m"
else
  echo "\033[32m find result: $sp_pid  \033[0m"
kill -15 "$sp_pid"
fi
at -m now+30min << EOF
root/workspace/naotanbot
EOF