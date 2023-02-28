
all : linux osx

linux : 
	GOOS=linux GOARCH=amd64 go build -o timeseries_linux_amd64

osx :
	GOOS=darwin GOARCH=arm64 go build -o timeseries_osx
clean :
	rm han-mongodb
