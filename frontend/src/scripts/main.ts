import { canvas, ctx } from './@module.js'

const scale = 1.5

const map = await fetch('http://localhost:5500/public/assets/map.json').then((res) => res.json()) as World

canvas.height = map.radius.y * scale
canvas.width = map.radius.x * scale

const images = new Map<string, HTMLImageElement>()
const amount = map.defs.filter((def) => def.visible).length

for (const def of map.defs) {
    if (def.visible === false) continue

    const image = new Image()

    image.onload = () => images.set(def.id, image)
    image.src = def.image_src

}

function game() {
    if (images.size !== amount) return
    
    for (const def of map.defs) {
        if (def.visible === false) continue

        const image = images.get(def.id)!

        for (const position of def.positions) {
            console.log(position)
            ctx!.drawImage(image,
                position.x * scale,
                position.y * scale,
                image.width * scale,
                image.height * scale,
            )
        }

    }

}

setInterval(
    () => requestAnimationFrame(game), 1000 / 1
)
