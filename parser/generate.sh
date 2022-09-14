#!/bin/sh

export CLASSPATH="./antlr-4.11.1-complete.jar:$CLASSPATH"
java -Xmx500M org.antlr.v4.Tool -Dlanguage=Go -visitor -no-listener -package parser *.g4