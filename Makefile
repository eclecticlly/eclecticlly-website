dev:
# Windows
# start air -c .air-win.toml && start tailwindcss -i ./tailwind.css -o ./static/css/styles.css --watch && start goconvey
# Unix
	tailwindcss -i ./tailwind.css -o ./static/css/styles.css --watch &! goconvey &! air -c .air-unix.toml
.PHONY: dev