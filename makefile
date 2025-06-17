install:
	echo "Creating local bin dir..."
	mkdir -p ${HOME}/.local/bin
	echo "Copying local scripts to ${HOME}/.local/bin directory..."
	cp -rfv scripts/bash/* ${HOME}/.local/bin
