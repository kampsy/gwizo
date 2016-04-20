#!/bin/sh
#
#------------------------
# Test all steps
#------------------------

cd ..

go test -v gwizo/gwizo_step_1a_test.go
go test -v gwizo/gwizo_step_1b_test.go
go test -v gwizo/gwizo_step_1c_test.go
go test -v gwizo/gwizo_step_2_test.go
go test -v gwizo/gwizo_step_3_test.go
go test -v gwizo/gwizo_step_4_test.go
go test -v gwizo/gwizo_step_5a_test.go
go test -v gwizo/gwizo_step_5b_test.go
