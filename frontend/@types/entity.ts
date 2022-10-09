type EntityTypes = Soldier | Projectile

interface Entity<T extends EntityTypes> {
    _id: T extends Soldier ? string : never,

    position: { x: number, y: number, z: number },
    states: {
        [key: string]: (e: Entity<T>, i: number) => void
    },

    prev_states: string[],
    state: string,

    entity: T
}
