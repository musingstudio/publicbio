#!/bin/bash

CSSDIR=static/css/

lessc less/app.less --clean-css="--s1 --advanced" ${CSSDIR}main.css
