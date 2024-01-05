NAME=zig-cart
ARCHIVE=tic80-${NAME}.zip
GAME_PATH=games/tic80/${NAME}
GAME_URL=https://${HOSTNAME}/${GAME_PATH}
PUBLIC_PATH=~/public_html/${GAME_PATH}
HOSTNAME=peter.tilde.team

all:
	zig build

.PHONY: spy
spy:
	zig build spy

.PHONY: run
run:
	zig build run

.PHONY: clean
clean:
	rm -rf bundle

.PHONY: bundle
bundle: all
	@rm -rf bundle
	@mkdir -p bundle
	@tic80-pro --cli --fs . --cmd 'load cart.wasmp & import binary zig-out/bin/cart.wasm & save & export html bundle/${NAME} alone=1 & export linux bundle/${NAME}.elf alone=1 & export win bundle/${NAME}.exe alone=1 & exit'
	@echo ""
	@unzip -d bundle/web bundle/${NAME}.zip
	@zip -rjuq bundle/${ARCHIVE} bundle/web bundle/${NAME}.elf bundle/${NAME}.exe
	@echo "✔ Updated bundle/${ARCHIVE}"

.PHONY: deploy
deploy: bundle
	@ssh ${HOSTNAME} 'mkdir -p ${PUBLIC_PATH}'
	@scp -q bundle/web/* ${HOSTNAME}:${PUBLIC_PATH}/
	@echo "✔ Updated ${NAME} on ${GAME_URL}"
	@scp -q bundle/${ARCHIVE} ${HOSTNAME}:${PUBLIC_PATH}/${ARCHIVE}
	@echo "✔ Archive ${GAME_URL}/${ARCHIVE}"
