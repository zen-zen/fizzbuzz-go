#!/usr/bin/env bash

set -e
set -x

## Starting up the server
./fizzbuzz-go &
pid=$!

## Configure URL
baseurl=http://localhost:8080

## (success cases) Test ordinary numbers (not fizz, not buzz)
nums=( 1 4 7 )
expected=""
for num in ${nums[@]}
do
    echo num is "'$num'"
    url=${baseurl}/${num}
    res=$(curl $url)
    if [ "$res" != "$expected" ]
    then
        echo ">>> ERROR"
        echo "curl $url"
        echo "  expected: '$expected'"
        echo "  got:      '$res'"
        exit 1
    fi
done

## (success cases) Test fizz only numbers
nums=(3 6 9)
expected="fizz"
for num in ${nums[@]}
do
    url=${baseurl}/${num}
    res=$(curl $url)
    if [ "$res" != "$expected" ]
    then
        echo ">>> ERROR"
        echo "curl $url"
        echo "  expected: '$expected'"
        echo "  got:      '$res'"
        exit 1
    fi
done

## (success cases) Test buzz only numbers
nums=(5 10 20)
expected="buzz"
for num in ${nums[@]}
do
    url=${baseurl}/${num}
    res=$(curl $url)
    if [ "$res" != "$expected" ]
    then
        echo ">>> ERROR"
        echo "curl $url"
        echo "  expected: '$expected'"
        echo "  got:      '$res'"
        exit 1
    fi
done

## (success cases) Test fuzzbuzz numbers
nums=(15 30 45)
expected="fizzbuzz"
for num in ${nums[@]}
do
    url=${baseurl}/${num}
    res=$(curl $url)
    if [ "$res" != "$expected" ]
    then
        echo ">>> ERROR"
        echo "curl $url"
        echo "  expected: '$expected'"
        echo "  got:      '$res'"
        exit 1
    fi
done

## (error cases) Test non-number strings
strs=(q 8a '%22')
expected='"400"'
set +e
for str in in ${strs[@]}
do
    url=${baseurl}/${str}
    ## cf. https://superuser.com/questions/272265/getting-curl-to-output-http-status-code#442395
    res=$(curl -s -o /dev/null -w \"%{http_code}\" $url)
    if [ "$res" != "$expected" ]
    then
        set -e
        echo ">>> ERROR"
        echo "curl -s -o /dev/null -w \"%{http_code}\" $url"
        echo "  expected: '$expected'"
        echo "  got:      '$res'"
        exit 1
    fi
done

kill $pid
