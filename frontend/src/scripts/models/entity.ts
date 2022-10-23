function entity <T extends object> (
    this: Entity<T>, entity: EntityModel<T>
) {
    for (const key of Object.keys(entity)) { // @ts-ignore
        this[key as keyof Entity<T>] = entity[key as keyof EntityModel<T>]
    }

}

const Entity = entity as unknown as {
    new <T extends object> (EntityModel: T): Entity<T>
}

export { Entity }
