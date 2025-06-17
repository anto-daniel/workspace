install:
	echo "Creating local bin dir..."
	mkdir -p ${HOME}/.local/bin
	echo "Copying local BASH scripts to ${HOME}/.local/bin directory..."
	cp -rf scripts/bash/* ${HOME}/.local/bin
	echo "Copying local PYTHON scripts to ${HOME}/.local/bin directory..."
	cp -rf scripts/python/* ${HOME}/.local/bin
