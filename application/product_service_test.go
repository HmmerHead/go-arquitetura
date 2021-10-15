package app_test

import (
	"testing"

	application "github.com/HmmerHead/go-arquit/application"
	mock_app "github.com/HmmerHead/go-arquit/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {

	crtl := gomock.NewController(t)
	defer crtl.Finish()

	product := mock_app.NewMockProductInterface(crtl)
	persistence := mock_app.NewMockProductPersistenceInterface(crtl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {

	crtl := gomock.NewController(t)
	defer crtl.Finish()

	product := mock_app.NewMockProductInterface(crtl)
	persistence := mock_app.NewMockProductPersistenceInterface(crtl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Prod 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Enable(t *testing.T) {

	crtl := gomock.NewController(t)
	defer crtl.Finish()
	product := mock_app.NewMockProductInterface(crtl)
	product.EXPECT().Enable().Return(nil)

	persistence := mock_app.NewMockProductPersistenceInterface(crtl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_Disable(t *testing.T) {

	crtl := gomock.NewController(t)
	defer crtl.Finish()
	product := mock_app.NewMockProductInterface(crtl)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_app.NewMockProductPersistenceInterface(crtl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

}
