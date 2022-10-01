package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ruancaetano/hexagonal-go/adapters/cli"
	mock_application "github.com/ruancaetano/hexagonal-go/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestProduct_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "123"
	productName := "Product Test"
	productStatus := "enabled"
	productPrice := 25.99

	producktMock := mock_application.NewMockProductInterface(ctrl)
	producktMock.EXPECT().GetID().Return(productId).AnyTimes()
	producktMock.EXPECT().GetName().Return(productName).AnyTimes()
	producktMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	producktMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Get(gomock.Any()).Return(producktMock, nil).AnyTimes()
	service.EXPECT().Create(gomock.Any(), gomock.Any()).Return(producktMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(producktMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(producktMock, nil).AnyTimes()
	service.EXPECT().Get(gomock.Any()).Return(producktMock, nil).AnyTimes()

	// create
	expectedResult := fmt.Sprintf(
		"Product ID %s with name %s has been created with price %f and status %s\n",
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
		"Product ID %s has been enabled\n",
		productId,
	)
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	// disable
	expectedResult = fmt.Sprintf(
		"Product ID %s has been disabled\n",
		productId,
	)
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	// get
	expectedResult = fmt.Sprintf(
		"Produc ID: %s\n Name: %s\n Price: %f\n Status: %s\n",
		productId,
		productName,
		productPrice,
		productStatus,
	)
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}
