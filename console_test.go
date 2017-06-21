package console

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockedWriter struct {
	mock.Mock
}

func (m *MockedWriter) Print(a ...interface{}) (int, error) {
	args := m.Called(a)
	return args.Int(0), args.Error(1)
}

func TestPrintShouldCallFmtPrintWithGivenColor(t *testing.T) {
	var writer = new(MockedWriter)
	var data = "Test"
	var expected = []interface{}{fmt.Sprintf("%s%s%s", colorYellow, data, colorReset)}

	writer.On("Print", expected).Return(0, nil)
	setWriter(writer.Print)

	Print(data, colorYellow)

	writer.AssertCalled(t, "Print", expected)
}

func TestPrintlnShouldCallFmtPrintlnWithGivenColor(t *testing.T) {
	var writer = new(MockedWriter)
	var data = "Test"
	var expected = []interface{}{fmt.Sprintf("%s%s%s\n", colorYellow, data, colorReset)}

	writer.On("Print", expected).Return(0, nil)
	setWriter(writer.Print)

	Println(data, colorYellow)

	writer.AssertCalled(t, "Print", expected)
}
