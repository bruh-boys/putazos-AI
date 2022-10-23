import { Entity } from './entity.js'
import { sprites } from './sprites.js'

const soldier_priorities: { [key: string]: number } = {
    'move': 1,
    'crouch': 2,
    'jump': 3,
    'death': 4
}

function update <T extends ISoldier> (
    this: Entity<T>, ctx: CanvasRenderingContext2D
) {
    this.sprite.update(ctx, this.position)
}

function soldier (
    this: Entity<ISoldier>, soldier: EntityModel<SoldierModel>
) { // @ts-ignore
    this = new Entity(soldier)

    this.sprite = structuredClone(
        sprites.get(`idle-${this.faction}-soldier`)!
    )

    this.update = update.bind(this)
}

const Soldier = soldier as unknown as {
    new (soldier: EntityModel<SoldierModel>): Entity<ISoldier>
}

export { Soldier, soldier_priorities }
