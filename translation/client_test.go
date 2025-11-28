package translation_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fseba/hello-api/translation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestHelloClientSuite(t *testing.T) {
	suite.Run(t, new(HelloClientSuite))
}

type HelloClientSuite struct {
	suite.Suite
	mockServerService *MockService
	server            *httptest.Server
	underTest         translation.HelloClient
}

type MockService struct {
	mock.Mock
}

func (m *MockService) Translate(word, language string) (string, error) {
	args := m.Called(word, language)
	return args.String(0), args.Error(1)
}

func (suite *HelloClientSuite) SetupSuite() {
	suite.mockServerService = new(MockService)
	handler := func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		defer func() {
			_ = r.Body.Close()
		}()

		var m map[string]any
		_ = json.Unmarshal(b, &m)
		word := m["word"].(string)
		language := m["language"].(string)

		resp, err := suite.mockServerService.Translate(word, language)
		fmt.Printf("Mock response: %v, error: %v\n", resp, err)

		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		if resp == "" {
			http.Error(w, "missing", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, resp)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	suite.server = httptest.NewServer(mux)
	suite.underTest = translation.NewHelloClient(suite.server.URL)
}

func (suite *HelloClientSuite) TearDownSuite() {
	suite.server.Close()
}

func (suite *HelloClientSuite) TestCall() {
	// Arrange
	suite.mockServerService.On("Translate", "foo", "bar").Return(`{"translation":"baz"}`, nil).Once()

	// Act
	resp, err := suite.underTest.Translate("foo", "bar")

	// Assert
	suite.NoError(err)
	suite.Equal(resp, "baz")
	suite.mockServerService.AssertExpectations(suite.T())
}

func (suite *HelloClientSuite) TestCall_APIError() {
	// Arrange
	suite.mockServerService.On("Translate", "foo", "bar").Return("", errors.New("this is a test")).Once()

	// Act
	resp, err := suite.underTest.Translate("foo", "bar")

	// Assert
	suite.EqualError(err, "error in api")
	suite.Equal(resp, "")
	suite.mockServerService.AssertExpectations(suite.T())
}

func (suite *HelloClientSuite) TestCall_InvalidJSON() {
	// Arrange
	suite.mockServerService.On("Translate", "foo", "bar").Return(`invalid json`, nil).Once()

	// Act
	resp, err := suite.underTest.Translate("foo", "bar")

	// Assert
	suite.EqualError(err, "unable to decode api response")
	suite.Equal(resp, "")
	suite.mockServerService.AssertExpectations(suite.T())
}
