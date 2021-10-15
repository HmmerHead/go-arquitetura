package cli_test

import (
	"fmt"
	"testing"

	"github.com/HmmerHead/go-arquit/adapters/cli"
	mock_app "github.com/HmmerHead/go-arquit/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product"
	productPrice := 25.0
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_app.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()

	service := mock_app.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Create Prod ID %s with the name %s created. Price = %f and status = %s",
		productId, productName, productPrice, productStatus)

	result, err := cli.Run(service, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Prod %s enable", productName)

	result, err = cli.Run(service, "enabled", productId, productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Prod %s disable", productName)

	result, err = cli.Run(service, "disabled", productId, productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Prod ID %s with the name %s created. Price = %f and status = %s",
			productId, productName, productPrice, productStatus)

	result, err = cli.Run(service, "get", productId, productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
