// node index.js < input.txt
import run from '../util/run.mjs'
import { parseInt } from '../util/numbers.mjs'

const allIncreasing = (values) => values.every((value, idx, arr) => idx + 1 >= arr.length || value < arr[idx + 1])

const allDecreasing = (values) => values.every((value, idx, arr) => idx + 1 >= arr.length || value > arr[idx + 1])

const allIncreasingOrDecreasing = (values) => allIncreasing(values) || allDecreasing(values)

const differenceWithinRange = (values, min, max) => values.every((value, idx, arr) => {
    const nextIdx = idx + 1
    if (nextIdx >= arr.length) {
        return true
    }

    const nextValue = arr[nextIdx]
    const difference = Math.abs(value - nextValue)
    return difference >= min && difference <= max 
})

const safe = (values, minRange, maxRange) => allIncreasingOrDecreasing(values) && differenceWithinRange(values, minRange, maxRange)

const runner = run(process.stdin)
let reports = []

runner.onLine((line) => {
    const levels = line.trim().split(/\s+/).map(parseInt)
    reports = [...reports, levels]
})

runner.onEnd(() => {
    const safeReports = reports.filter(report => safe(report, 1, 3))
    console.log(`Number of safe reports:\t${safeReports.length}`)
})