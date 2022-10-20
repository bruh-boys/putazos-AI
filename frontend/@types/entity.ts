interface IEntity {
    direction: boolean, // false = left, true = right
    position: Position,
    velocity: Position,
    actions: string[],
    cosmetics: string[]
}

interface ISoldier extends IEntity {
    color: string,
    id: string,
}
