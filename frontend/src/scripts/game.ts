import { Sprite, sprites } from './models/sprites.js'
import { Soldier, soldier_priorities } from './models/soldier.js'
import { generate_projectile_trail } from './models/projectile.js'

const canvas = document.getElementById('canvas') as HTMLCanvasElement
const ctx = canvas.getContext('2d') as CanvasRenderingContext2D

const framesPerSecond = 30

let world: World | undefined

let entities: {
    projectiles: ProjectileModel[],
    soldiers: {
        [key: string]: Entity<ISoldier> | undefined
    }
} = {
    projectiles: [],
    soldiers: {}
}

function overwrite_entities(ss: SocketValue<'data'>) {
    const soldiers: typeof entities['soldiers'] = {}

    for (const soldier of ss.soldiers) {
        const model = entities.soldiers[soldier.id]

        // If the soldier not exists, create a new.
        if (model === undefined) {
            soldiers[soldier.id] = new Soldier(soldier)

            continue
        }

        // If exists, overwrite the value of properties.
        for (const [key, value] of Object.entries(soldier)) { // @ts-ignore
            soldiers[entity.id]![key] = value

        }

        // In case the entity actions changed, we need to update the sprite.
        const [entity_action] = soldier.actions.sort((a, b) =>
            soldier_priorities[a] - soldier_priorities[b]
        )

        const [model_action] = model.actions.sort((a, b) =>
            soldier_priorities[a] - soldier_priorities[b]
        )

        if (model_action !== entity_action) {
            soldiers[soldier.id]!.sprite = structuredClone(
                sprites.get(`${entity_action}-soldier-${soldier.faction}`)!
            )

        }

    }

    entities.projectiles = ss.projectiles
    entities.soldiers = soldiers
}

function overwrite_map(w: World) {
    world = w
}

async function init() {
    for (const color of ['red', 'blue']) {
        new Sprite(`idle-${color}-soldier`, {
            framesHold: 10,
            framesMax: 5,
            source: `/assets/images/soldiers/${color}/idle.png`
        })

        new Sprite(`move-${color}-soldier`, {
            framesHold: 10,
            framesMax: 6,
            source: `/assets/images/soldiers/${color}/move.png`
        })

        new Sprite(`jump-${color}-soldier`, {
            framesHold: 10,
            framesMax: 2,
            source: `/assets/images/soldiers/${color}/jump.png`
        })

        new Sprite(`death-${color}-soldier`, {
            animationType: 'once',
            framesHold: 10,
            framesMax: 7,
            source: `/assets/images/soldiers/${color}/death.png`
        })

        new Sprite(`crouch-${color}-soldier`, {
            framesHold: 10,
            framesMax: 3,
            source: `/assets/images/soldiers/${color}/crouch.png`
        })

    }

    new Sprite(`small-plataform`, {
        animationType: 'once',
        source: `/assets/images/plataforms/small-plataform.png`
    })

    new Sprite(`plataform`, {
        animationType: 'once',
        source: `/assets/images/plataforms/plataform.png`
    })

    requestAnimationFrame(game)
}

function game() {
    ctx.fillStyle = 'rgba(0, 0, 0, 0.1)'
    ctx.fillRect(0, 0, canvas.width, canvas.height)

    for (const projectile of Object.values(entities.projectiles)) {
        generate_projectile_trail(ctx, projectile.start, projectile.position)
    }

    for (const soldier of Object.values(entities.soldiers)) {
        soldier!.update(ctx)
    }

    for (const definition of world?.defs || []) {
        if (definition.visible === false) continue

        for (const position of definition.positions) {
            sprites.get(definition.id)?.update(ctx, position)
        }
    }

    requestAnimationFrame(game)
}

export {
    entities, canvas, ctx, framesPerSecond, init,
    overwrite_entities,
    overwrite_map
}
