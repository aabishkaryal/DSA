#!/bin/bash
go test github.com/aabishkaryal/DSA/problems -failfast -covermode=atomic -coverprofile cover.out > /dev/null
coverage=$(go tool cover -func cover.out | grep total | awk '{print substr($3, 1, length($3)-1)}')
rm cover.out
if (( $(echo "$coverage < $1" | bc -l) )); 
then
    echo "Code coverage is less than $1%. Please fix the code coverage."
    exit 1
else
    echo "Code coverage is $coverage%. Good job!"
    exit 0
fi