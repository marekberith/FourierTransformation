package main

import (
	"github.com/therecipe/qt/widgets"
	"./tr"
)

var (
	mainWindow *widgets.QWidget
	mainLayout *widgets.QVBoxLayout
	resultLayout *widgets.QHBoxLayout
	vectorLine *widgets.QLineEdit
	transformFFT *widgets.QPushButton
	transformIFFT *widgets.QPushButton
	result *widgets.QLabel
)

func setLayout() {
	mainWindow = widgets.NewQWidget(nil, 0)
	mainWindow.SetWindowTitle("Fourier Transformation Calculator")
	mainWindow.SetMinimumWidth(500)
	mainWindow.SetMinimumHeight(400)

	mainLayout = widgets.NewQVBoxLayout()
	vectorLine = widgets.NewQLineEdit(nil)
	vectorLine.SetPlaceholderText("Insert vector to calculate FFT/IFFT(complex numbers in format x + yi)")

	result = widgets.NewQLabel(nil, 0)

	transformFFT = widgets.NewQPushButton(nil)
	transformFFT.SetText("Calculate FFT")
	transformFFT.ConnectClicked(func(checked bool) {
		transformedString := tr.TransformVectorFFT(vectorLine.Text())
		result.SetText(transformedString)
	})

	transformIFFT = widgets.NewQPushButton(nil)
	transformIFFT.SetText("Calculate IFFT")
	transformIFFT.ConnectClicked(func(checked bool) {
		transformedString := tr.TransformVectorIFFT(vectorLine.Text())
		result.SetText(transformedString)
	})

	mainLayout.AddWidget(vectorLine, 0, 0)
	mainLayout.AddWidget(transformFFT, 0, 0)
	mainLayout.AddWidget(transformIFFT, 0, 0)
	resultLayout = widgets.NewQHBoxLayout()
	resultLayout.AddStretch(1)
	resultLayout.AddWidget(result, 0, 0)
	resultLayout.AddStretch(1)
	mainLayout.AddLayout(resultLayout, 0)

	mainWindow.SetLayout(mainLayout)
	mainWindow.Show()
}
