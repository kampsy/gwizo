package signup

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

// instrumentingMiddleware
type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           Servicer
}

// Singin runtime behaviour
func (mw instrumentingMiddleware) Signup(FirstName,LastName, Phone_number, Password,Pin,Username, User_id,  Email, Balance , User_type_name string , Status_id, Account_type_id, Account_id, User_type_id int ) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"service", "signup", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Signup(FirstName,LastName, Phone_number, Password,Pin,Username, User_id,Email, Balance , User_type_name, Status_id, Account_type_id, Account_id, User_type_id ) // i am having an erro here to 
	return output, nil
}