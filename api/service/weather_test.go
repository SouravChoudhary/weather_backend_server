package service

import (
	"errors"
	"testing"
	"weatherapp/external"
	"weatherapp/pkg/logging"

	"github.com/golang/mock/gomock"
)

/* TEST STRATEGY

If Weather Api Client Call Success (T)
    If Weather History Svc Create Record Call Success (T)
	    then return weather data  ===> Testcase 1

	If Weather History Svc Create Record Call fail (F)
	     then return error  ===> Testcase 2

If Weather Api Client Call Fail (F)
    then return error  ===> Testcase 3

*/

func Test_GetCurrentWeather(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedApiClient := external.NewMockOpenWeatherMapClient(mockCtrl)
	mockedWeatherHistorySvc := NewMockWeatherHistory(mockCtrl)
	mockedLogger := logging.NewMockLogger(mockCtrl)

	mockedWeatherService := NewWeatherService(mockedWeatherHistorySvc, mockedApiClient, mockedLogger)

	type input struct {
		userId   int
		cityName string
	}

	type output struct {
		wd  string
		err error
	}

	type testCase struct {
		result  string
		usecase string
		in      input
		out     output
		mocks   func()
	}

	testcases := []testCase{
		{
			result:  "SUCCESS",
			usecase: "if Weather Api Client Call =  Success ; Weather History Svc Create Record Call= Success; then return weather data",
			in:      input{userId: 10115, cityName: "Berlin"},
			out:     output{err: nil, wd: "json string weather data"},
			mocks: func() {
				mockedApiClient.EXPECT().GetWeatherData(gomock.Any()).Return("json string weather data", nil)
				mockedWeatherHistorySvc.EXPECT().CreateRecord(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			result:  "ERROR",
			usecase: "if Weather Api Client Call =  Success ; Weather History Svc Create Record Call= Fail; then return error",
			in:      input{userId: 10115, cityName: "Berlin"},
			out:     output{err: errors.New("db error"), wd: ""},
			mocks: func() {
				mockedApiClient.EXPECT().GetWeatherData(gomock.Any()).Return("json string weather data", nil)
				mockedWeatherHistorySvc.EXPECT().CreateRecord(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("db error"))
			},
		},
		{
			result:  "ERROR",
			usecase: "if Weather Api Client Call =  Fail ; then return error",
			in:      input{userId: 10115, cityName: "Berlin"},
			out:     output{wd: "", err: errors.New("weather api client call failure")},
			mocks: func() {
				mockedApiClient.EXPECT().GetWeatherData(gomock.Any()).Return("", errors.New("client error"))
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.usecase, func(t *testing.T) {
			if tc.mocks != nil {
				tc.mocks()
			}

			weatherData, err := mockedWeatherService.GetCurrentWeather(tc.in.userId, tc.in.cityName)

			if weatherData != tc.out.wd {
				t.Errorf("Test_GetCurrentWeather: expected weather data:%v;got:%v", tc.out.wd, weatherData)
			}

			if err != nil && tc.out.err != nil && err.Error() != tc.out.err.Error() {
				t.Errorf("Test_GetCurrentWeather: expected error:%v;got:%v", tc.out.err.Error(), err.Error())
			}

		})
	}

}
