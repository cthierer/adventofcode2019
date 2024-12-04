// node index.js < input.txt

const appendToSorted = (list, value) => {
    if (list.length < 1) {
        return [value]
    }

    const idx = list.findLastIndex((element) => element < value)

    if (idx < 0) {
        return [value, ...list]
    }

    return [
        ...list.slice(0, idx + 1),
        value,
        ...list.slice(idx + 1),
    ]
}

const parseInt = (valueStr) => {
    const valueInt = Number.parseInt(valueStr, 10)
    if (Number.isNaN(valueInt)) {
        return new Error(`invalid number: "${valueStr}"`)
    }

    return valueInt
}

const computeDistances = (list1, list2) => list1.map((value1, idx) => {
    if (idx > list2.length) {
        return value1
    }

    const value2 = list2[idx]
    return Math.abs(value1 - value2)
})

const computeSimularities = (list1, list2) => list1.map((value1) => {
    const numRepeats = list2.filter((value2) => value1 === value2).length
    return value1 * numRepeats
})

const run = (input) => {
    let list1 = []
    let list2 = []

    input.on('readable', () => {
        let chunk
        while ((chunk = input.read()) && chunk !== null) {
            const lines = chunk.toString('utf8').trim().split('\n')
            lines.forEach((element) => {
                const line = element.trim() 
                if (line.length < 1) {
                    return
                }

                const [list1Str, list2Str] = line.split(/\s+/)
                list1 = appendToSorted(list1, parseInt(list1Str))
                list2 = appendToSorted(list2, parseInt(list2Str))    
            })
        }
    })

    input.on('error', (err) => {
        console.error(`An error occurred: ${err.message}`)
    })

    input.on('end', () => {
        const distances = computeDistances(list1, list2)
        const totalDistance = distances.reduce((last, curr) => last += curr, 0)
        console.log('Total distance:\t\t', totalDistance)

        const similarities = computeSimularities(list1, list2)
        const totalSimilarity = similarities.reduce((last, curr) => last += curr, 0)
        console.log('Total similarity:\t', totalSimilarity)

        console.log('Done!')
    })
}

run(process.stdin)