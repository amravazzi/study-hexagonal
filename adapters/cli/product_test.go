package cli_test

import (
	"fmt"
	"testing"

	"github.com/amravazzi/study-hexagonal/adapters/cli"
	mock_application "github.com/amravazzi/study-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "abc"
	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	// create
	expectedResult := fmt.Sprintf(
		"Product ID %s with the name %s has been created with price %f and status %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	// enable
	expectedResult = fmt.Sprintf(
		"Product %s has been enabled",
		productName,
	)
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	// disable
	expectedResult = fmt.Sprintf(
		"Product %s has been disabled",
		productName,
	)
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	// get
	expectedResult = fmt.Sprintf(
		"Product ID: %s\n Name: %s\n Price: %f\n Status: %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}
