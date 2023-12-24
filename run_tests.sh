#!/bin/bash
# Проверка написанного.
# Если скрипт завершился и ничего не написал, код 0, - значит добро.

BIN=run_tests
DATA=run_tests_data

go build -o $BIN .
printf "E456BK17\nA123AA70\nО001ОО98" > $DATA

./test_suite.sh $BIN $DATA
EXIT_CODE=$?

rm $BIN
rm $DATA

exit $EXIT_CODE
