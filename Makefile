build:
	mkdir -p schoolday.app/Contents/MacOS
	cp ./mac_bundle/Info.plist schoolday.app/Contents
	go build -o schoolday.app/Contents/MacOS/schoolday