#!/bin/sh
 
set -eu
 
echo "Checking DB MySQL connection ..."
 
i=0
until [ $i -ge 20 ]
do
  nc -z mysql 3306 && break
 
  i=$(( i + 1 ))
 
  echo "$i: Waiting for DB MySQL 1 second ..."
  sleep 1
done
 
if [ $i -eq 20 ]
then
  echo "DB MySQL connection refused, terminating ..."
  exit 1
fi
 
echo "DB MySQL is ready ..."
 
./sample_blog_server