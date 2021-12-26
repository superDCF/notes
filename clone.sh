#!/bin/bash
## login gitlab cookie
set -Eeuxo pipefail

page=1
count=0
gitlab=$2 # gitlab domain ex:code.hellotalk.com
url="https://$gitlab/?non_archived=true&page=$page&sort=latest_activity_desc"
response="$(curl -LSs -b "$1" $url)"
declare -a dirs
# while (test $page -lt 2) ;do
while  [[ $response =~ class=\"project\" ]];do
    list="$(echo "$response" | awk -F '"' '/class="project"/ {print $4}')"
    echo "$list"
    for str in $list;do
        echo $str
        dirs[$count]="$str"
        ((count++))
        echo $count
    done
    ((page++))
    echo $page
    url="https://$gitlab/?non_archived=true&page=$page&sort=latest_activity_desc"
    echo $url
    response="$(curl -LSs -b "$1" $url)"
done
echo ${dirs[@]}
basepath=~/$gitlab # clone dir
for dir in ${dirs[@]};do
    echo $dir
    name=${dir##/*/}
    echo $name
    subpath=${dir%/*}
    echo $subpath
    path=$basepath$subpath
    echo $path
    mkdir -p $path
    cd $path
    git clone git@$gitlab:$dir.git
done
