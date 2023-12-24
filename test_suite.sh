#!/bin/bash
# Это сам набор тестов. Запускать через ./run_tests.sh

# упростители жизни со SO
yell() { echo "$0: $*" >&2; }
die() { yell "$*"; exit 111; }

# с данными запустили
./"$1" -file="$2" -concurrency=2

# начинаются тесты!

# отработало успешно?
if [ $? -ne 0 ]; then
    die "the program has failed us! non zero return code"
fi

# есть папка E456BK17?
if [ ! -d './output/E456BK17' ]; then
    die 'Dir E456BK17 does not exist'
else
    ## там есть непустые info.json и preview.jpg?
    if ! [ -s './output/E456BK17/info.json' ]; then
      die "E456BK17/info.json does not exist or is empty."
    fi
    if ! [ -s './output/E456BK17/preview.jpg' ]; then
      die "E456BK17/preview.jpg does not exist or is empty."
    fi
    ### info.json содержит нужный номер?
    if ! grep -qi '"carplate":"E456BK17",' "./output/E456BK17/info.json"; then
      die 'E456BK17/info.json invalid carplate'
    fi
fi

# есть папка A123AA70?
if [ ! -d './output/A123AA70' ]; then
    die 'Dir A123AA70 does not exist'
else
    ## там есть info.json?
    if ! [ -s './output/A123AA70/info.json' ]; then
      die "A123AA70/info.json does not exist or is empty."
    fi
    ### содержит нужный номер?
    if ! grep -qi '"carplate":"A123AA70",' "./output/A123AA70/info.json"; then
      die 'A123AA70/info.json invalid carplate'
    fi
fi
