// node index.js < input.txt
import run from '../util/run.mjs'
import { parseInt } from '../util/numbers.mjs'

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

const runner = run(process.stdin)
let list1 = []
let list2 = []

runner.onLine((line) => {
    const [list1Str, list2Str] = line.split(/\s+/)
    list1 = appendToSorted(list1, parseInt(list1Str))
    list2 = appendToSorted(list2, parseInt(list2Str))    
})

runner.onEnd(() => {
    const distances = computeDistances(list1, list2)
    const totalDistance = distances.reduce((last, curr) => last += curr, 0)
    console.log('Total distance:\t\t', totalDistance)

    const similarities = computeSimularities(list1, list2)
    const totalSimilarity = similarities.reduce((last, curr) => last += curr, 0)
    console.log('Total similarity:\t', totalSimilarity)
})
