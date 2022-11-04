import { Sprite } from './models/sprite.js'
import { Soldier } from './models/soldier.js'
import { generate_projectile_trail } from './models/projectile.js'

const canvas = document.getElementById('canvas') as HTMLCanvasElement
const ctx = canvas.getContext('2d') as CanvasRenderingContext2D

const framesPerSecond = 10
const scale = 1

let world: World | undefined

let entities: {
    projectiles: ProjectileModel[],
    soldiers: Soldier[]
} = {
    projectiles: [],
    soldiers: []
}

let sprites = new Map<string, Sprite>()

function overwrite_entities(e: SocketValue<'update'>) {
    for (const soldier of e.soldiers) {
        const entity = entities.soldiers.find(s => s.id === soldier.id)

        if (entity) {
            entity.update(soldier)

            continue
        }

        entities.soldiers.push(new Soldier(
            soldier.id, soldier.faction!, soldier.direction!, soldier.position, soldier.actions
        ))
    }

}

function overwrite_map(w: World) {
    canvas.height = w.radius.y
    canvas.width = w.radius.x
    world = w
}

async function init() {
    sprites.set('small-plataform', new Sprite({
        source: '/public/assets/images/plataforms/small-plataform.png',
        id: 'small-plataform',
        animationType: 'once',
        maxFrames: 1,
        holdFrames: 0,
    }))

    sprites.set('plataform', new Sprite({
        source: '/public/assets/images/plataforms/plataform.png',
        id: 'plataform',
        animationType: 'once',
        maxFrames: 1,
        holdFrames: 0,
    }))

    requestAnimationFrame(game)
}

async function game() {
    ctx!.fillStyle = 'rgba(255, 255, 255, 1)'
    ctx!.clearRect(0, 0, canvas.width, canvas.height)

    for (const definition of world?.defs || []) {
        if (definition.visible === false) continue

        for (const position of definition.positions) {
            const sprite = sprites.get(definition.id)!
            sprite.draw(ctx, scale, position)

        }

    }

    for (const projectile of Object.values(entities.projectiles)) {
        generate_projectile_trail(ctx, projectile.start, projectile.position)
    }

    for (const soldier of entities.soldiers) {
        soldier!.draw(ctx, scale)
    }

    for (const definition of world!.defs) {
        if (definition.id !== 'collision') continue

        for (const position of definition.positions) {
            ctx!.strokeStyle = 'rgba(255, 0, 0, 1)'
            ctx!.strokeRect(position.x, position.y, 16, 16)
        }
    }

    setTimeout(() => requestAnimationFrame(game), (1000 / framesPerSecond));
}

export {
    entities, canvas, ctx, framesPerSecond, init,
    overwrite_map, overwrite_entities
}
