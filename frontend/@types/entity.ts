interface IEntity {
    direction: boolean, // false = left, true = right
    position: IPosition,
    velocity: IPosition,
    actions: string[],
    cosmetics: string[]
}

interface ISoldier extends IEntity {
    color: string,
    id: string,
}
