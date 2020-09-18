package sayhello_test

import (
	"testing"

	"sayhello"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSayHello(t *testing.T) {
	tests := []struct {
		name        string
		ip          string
		expectedRes string

		setup func(*MockIPStack, *MockDB, *MockClock)
	}{
		{
			name:        "a UK address must return time in english",
			ip:          "3.8.0.16",
			expectedRes: "Hello, it is 3:48PM\n",

			setup: func(ipstack *MockIPStack, db *MockDB, clock *MockClock) {
				ipstack.EXPECT().GetCountryCode("3.8.0.16").Return("GB", nil)
				db.EXPECT().Hit("3.8.0.16")
				clock.EXPECT().Now().Return("3:48PM")
			},
		},
		{
			name:        "a FR address must return time in french",
			ip:          "3.8.0.16",
			expectedRes: "Salut, il est 3:48PM\n",

			setup: func(ipstack *MockIPStack, db *MockDB, clock *MockClock) {
				ipstack.EXPECT().GetCountryCode("3.8.0.16").Return("FR", nil)
				db.EXPECT().Hit("3.8.0.16")
				clock.EXPECT().Now().Return("3:48PM")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			ipstack := NewMockIPStack(ctrl)
			db := NewMockDB(ctrl)
			clock := NewMockClock(ctrl)
			tt.setup(ipstack, db, clock)
			sh := sayhello.New(db, ipstack, clock)

			// test
			res, err := sh.SayHello(tt.ip)

			// assert
			require.NoError(t, err)
			assert.Equal(t, tt.expectedRes, res)
		})
	}
}
