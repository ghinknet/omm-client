#!/bin/bash

cd $(dirname $0)

function _build(){
	go generate || exit $?
	go build -o "$1" . || exit $?
}

_build ./target/output-self
