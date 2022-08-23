#!/bin/bash

cd $(dirname $0)

function _build(){
	cargo tauri build || exit $?
}

_build
