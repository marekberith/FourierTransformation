# FourierTransformation

FourierTransformation is a GoLang app used to calculate the FFT or the IFFT from a vector in O(nlog(n)) time. 

Example of a input:

1, 2, 3, 4

FFT Output:

((10+0i),(-2-2i),(-2+0i),(-2+2i))

IFFT Output:

(10,-2,-2,-1.9999999999999998)

## How to use?

You can either build the app or use one of the binaries - they are in the src folder.

If you want to build the app, you'll need to install GO and the qt binding into your GO root. Please, use the following package: https://github.com/therecipe/qt.
Then clone the repository, open the src folder and build the app using the following command:

`go build`
