#!/bin/bash

RED="$(tput setaf 1)" && readonly RED
GREEN="$(tput setaf 2)" && readonly GREEN
YELLOW="$(tput setaf 3)" && readonly YELLOW
END="$(tput sgr0)" && readonly END

warn() { echo "$YELLOW$*$END"; } >&2
die() { echo; echo "$RED$*$END"; echo; exit 1; } >&2

SOURCE_ROOT="$(cd "$(dirname "$0")" && pwd)" && readonly SOURCE_ROOT
APPLICATION="$(cd "$SOURCE_ROOT/application" && pwd)" && readonly APPLICATION
LIBRARY="$(cd "$SOURCE_ROOT/library" && pwd)" && readonly LIBRARY

COVERAGE_FILE='coverage.out' && readonly COVERAGE_FILE
GODOC_DIR='website/static/api/' && readonly GODOC_DIR
HUGO_DIR='website/public/' && readonly HUGO_DIR

warn 'Testing application...'

echo '(1/3) Running Just commands'
cd "$APPLICATION" || exit 1
just install
just format
just cov
just generate-website
cd "$APPLICATION" || exit 1

echo '(2/3) Checking coverage file'
if [[ ! -e "$COVERAGE_FILE" ]]; then
  die 'Coverage not found.'
fi

echo '(3/3) Checking website directory'
if [[ ! -e "$HUGO_DIR" ]]; then
  die 'Website not built.'
fi

warn 'Testing library...'

echo '(1/4) Running Just commands'
cd "$LIBRARY" || exit 1
just install
just format
just cov
just generate-website
cd "$LIBRARY" || exit 1

echo '(2/4) Checking coverage file'
if [[ ! -e "$COVERAGE_FILE" ]]; then
  die 'Coverage not found.'
fi

echo '(3/4) Checking documentation directory'
if [[ ! -e "$GODOC_DIR" ]]; then
  die 'Documentation not built.'
fi

echo '(4/4) Checking website directory'
if [[ ! -e "$HUGO_DIR" ]]; then
  die 'Website not built.'
fi

echo "${GREEN}Tests complete.$END"
echo
echo 'Goodbye!'
