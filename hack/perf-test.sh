#!/bin/bash
rss=`ps -eo pid,tid,class,rtprio,stat,vsz,rss,comm | grep k8s-dqlite | awk '{ print $7}'`
echo "START: $rss $size"
microk8s kubectl apply -f /var/snap/microk8s/current/args/cni-network/cni.yaml
microk8s.kubectl apply -f ./nginx.yaml
for i in `seq 1 1 20`
do
  microk8s.kubectl scale deployment nginx-deployment --replicas=80 > /dev/null
  #rss=`pmap -X $(pgrep k8s-dqlite) | tail -n 1 | awk '{ print $2}'`
  rss=`ps -eo pid,tid,class,rtprio,stat,vsz,rss,comm | grep k8s-dqlite | awk '{ print $7}'`
  size=`sudo ls -l /var/snap/k8s-dqlite/current/var/data/snapshot* | head -n 1 | awk '{print $5}'`
  sleep 5
  microk8s.kubectl scale deployment nginx-deployment --replicas=0 > /dev/null
  if [[ "$i" == *00 ]]
  then
    sleep 10
  fi
  echo $i $rss $size
done
