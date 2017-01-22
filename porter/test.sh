#!/bin/sh
#
#Package porter implement Porter, M. "An algorithm for suffix stripping."
#Program 14.3 (1980): 130-137.
#Martin Porter, the algorithm's inventor, maintains a web page about the
#algorithm at http://www.tartarus.org/~martin/PorterStemmer/

#------------------------
# Test all steps
#------------------------
cd ..

go test -v porter/step1a_test.go
go test -v porter/step1b_test.go
go test -v porter/step1c_test.go
go test -v porter/step2_test.go
go test -v porter/step3_test.go
go test -v porter/step4_test.go
go test -v porter/step5a_test.go
go test -v porter/step5b_test.go