
export const parseInt = (valueStr) => {
    const valueInt = Number.parseInt(valueStr, 10)
    if (Number.isNaN(valueInt)) {
        return new Error(`invalid number: "${valueStr}"`)
    }

    return valueInt
}