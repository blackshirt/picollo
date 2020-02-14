package scrapper

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// mockScrapper is an autogenerated mock type for the Scrapper type
type mockScrapper struct {
	mock.Mock
}

// Scrape provides a mock function with given fields: url
func (_m *mockScrapper) Scrape(url string) (*Results, error) {
	ret := _m.Called(url)

	var r0 *Results
	if rf, ok := ret.Get(0).(func(string) *Results); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Results)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func TestNewmockScrapper(t *testing.T) {
	// var ctx context.Context = context.TODO()
	// var w waktu
	// Create a new instance of the mock store
	// c := colly.NewCollector()
	// r := new(model.MockStorage)

	// svc := model.NewService(r)
	m := new(mockScrapper)
	// In the "On" method, we assert that we want the "Get" method
	// to be called with one argument, that is 2
	// In the "Return" method, we define the return values to be 7, and nil (for the result and error values)
	m.On("Scrape", "abc").Return(nil, nil)
	// Next, we create a new instance of our module with the mock store as its "store" dependency
	// s := New(c, svc)
	// The "Get" method call is then made
	_, err := m.Scrape("abc")
	// The expectations that we defined for our mock store earlier are asserted here
	m.AssertExpectations(t)
	// Finally, we assert that we should'nt get any error
	if err != nil {
		t.Errorf("error should be nil, got: %v", err)
	}
}