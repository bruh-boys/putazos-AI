interface Position {
    x: number, y: number, z: number
}

interface Soldier {
    _id: string,
    color: string
    actions: string[],
}

//

type EntityTypes = Soldier

interface EntityMethods {
    set_previous_state: (this: Entity<EntityTypes>) => void,
    change_state: (this: Entity<EntityTypes>, state: string) => void,
    call_state: (this: Entity<EntityTypes>, state: string) => void,
}

interface Entity<T extends EntityTypes> extends EntityMethods, Position {
    ctx: CanvasRenderingContext2D,

    sprites: { [key: string]: {
        image_src: string,
        amount: number,
        hold: number
    } },

    previous_states: [string, number][],

    elapsed_states: number,
    current_state: string,

    states: {
        [key: string]: () => void 
    },

    entity: T
}
