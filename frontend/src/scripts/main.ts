import { canvas, ctx } from './@module.js'

import { entities } from './socket.js'
import('./socket.js')

const scale = 1

const map = await fetch('http://localhost:5500/public/assets/map.json').then((res) => res.json()) as World

canvas.height = map.radius.y * scale
canvas.width = map.radius.x * scale

const images = new Map<string, HTMLImageElement>()
const amount = map.defs.filter((def) => def.visible).length

// It's necessary to record the unnecesary pixels of the image
const soldier_idle = new Image()
soldier_idle.onload = () => {
    images.set("soldier-red-idle", soldier_idle)

}
soldier_idle.src = 'http://localhost:5500/public/assets/soldiers/red/idle2.png'

for (const def of map.defs) {
    if (def.visible === false) continue

    const image = new Image()

    image.onload = () => images.set(def.id, image)
    image.src = "http://localhost:5500/public/assets/plataforms/" + def.id + ".png"

}

function game() {
    //if (images.size !== amount) return

    ctx!.fillStyle = 'rgba(255, 255, 255, 0.3)'
    ctx!.fillRect(0, 0, canvas.width, canvas.height)
    
    for (const def of map.defs) {
        if (def.visible === false) continue

        const image = images.get(def.id)!

        for (const position of def.positions) {
            ctx!.drawImage(image,
                position.x * scale,
                position.y * scale,
                image.width * scale,
                image.height * scale,
            )
        }

    }

    for (const collision of map.defs.filter((def) => def.id === 'collision')) {
        for (const position of collision.positions) {
            ctx!.fillStyle = 'rgba(0, 0, 0, 0.5)'
            ctx!.fillRect(
                position.x * scale,
                position.y * scale,
                16 * scale,
                16 * scale,
            )
        }
    }

    for (const entity of (entities || {}).soldiers || []) {
        const image = images.get("soldier-red-idle")!
        console.log(entity)

        ctx!.drawImage(image,
            (entity.position.x) * scale,
            (entity.position.y) * scale,
            image.width  * scale,
            image.height * scale
        )
    }

}

setInterval(
    () => requestAnimationFrame(game), 1000 / 30
)
