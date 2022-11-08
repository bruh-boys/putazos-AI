type EntityModel<T extends object> = T & {
    position: Position
    id: string
}

type Entity<T extends object> = EntityModel<T> & {
    sprite?: Sprite,

    update?: <T extends object> (
        this: Entity<T>, ctx: CanvasRenderingContext2D
    ) => void
}

//

interface SoldierModel {
    health?: number,
    direction?: boolean
    position?: Position
    actions?: string[]
    faction?: 'red' | 'blue'
}

interface ISoldier extends SoldierModel {
    sprite: Sprite,

    update: <T extends object> (
        this: Entity<T>, ctx: CanvasRenderingContext2D
    ) => void
}

interface ProjectileModel {
    position: Position
    start: Position
}
