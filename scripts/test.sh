#! /bin/bash

echo "Running test..."
go test -race -p 1 -v -coverpkg=./... -coverprofile=profile.cov ./... | tee test.log
sed -i '' '/^go-template\/cmd\/app\/main.go/ d' profile.cov # Remove '' if you use ubuntu.
go tool cover -func profile.cov
#go tool cover -html profile.cov

echo "------------------"

fail=`grep "FAIL:" test.log` || true
echo "$fail"
if [ -z "$fail" ]; then
  pass=`grep "PASS:" test.log` || true
  echo "$pass"
  echo "--- ALL TESTCASES PASSED"
else
  echo "--- SOME TESTCASES FAILED"
  echo "Exit 1"
  exit 1
fi

echo "Threshold: $TESTCOVERAGE_THRESHOLD%"
totalCoverage=`go tool cover -func=profile.cov | grep total | grep -Eo '[0-9]+\.[0-9]+'`
echo "Current test coverage: $totalCoverage%"
if (( $(echo "$totalCoverage $TESTCOVERAGE_THRESHOLD" | awk '{print ($1 > $2)}') )); then
  echo "OK"
  exit 0
else
  echo "Current test coverage is below threshold. Please add more unit tests."
  echo "Exit 1"
  exit 1
fi
