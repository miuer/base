#! /bin/bash

i=1
while((${i}<=100))
do
  ./main >> main.txt
  let i=$i+1
done