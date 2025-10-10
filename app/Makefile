all: acme.app

acme.app:
	mkdir -p acme.app/Contents/{MacOS,Resources}
	cp Info.plist acme.app/Contents/
	cp launcher wrapper acme.app/Contents/MacOS/
	cp spaceglenda.icns acme.app/Contents/Resources/
	chmod +x acme.app/Contents/MacOS/{launcher,wrapper}

clean:
	rm -rf acme.app

.PHONY: all clean
