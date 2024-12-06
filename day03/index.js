// node index.js < input.txt
import run from '../util/run.mjs'

const runner = run(process.stdin)

const findNextOccurrence = (string, pattern, afterIdx) => {
    let nextOccurrence = afterIdx - 1

    while (nextOccurrence < afterIdx) {
        const match = pattern.exec(string)
        if (match === null) {
            return string.length
        }

        nextOccurrence = match.index + 1
    }

    return nextOccurrence
}

const parseBlock = function* parseBlock(string, initialState) {
    const enablePattern = /do\(\)/g
    const disablePattern = /don't\(\)/g
    const { enabled } = initialState
    let startIdx = enabled ? 0 : findNextOccurrence(string, enablePattern, 0)
    let stopIdx

    while (startIdx < string.length) {
        stopIdx = findNextOccurrence(string, disablePattern, startIdx)
        yield string.substring(startIdx, stopIdx)
        
        startIdx = findNextOccurrence(string, enablePattern, stopIdx)
    }

    return { ...initialState, enabled: stopIdx >= startIdx }
}


const parseMul = function* parseMul(string) {
    const pattern = /mul\((\d{1,3}),(\d{1,3})\)/g
    let match

    while((match = pattern.exec(string)) !== null) {
        const [, x, y] = match
        yield [x, y]
    }
}

const execMul = ([x, y]) => x * y

const sum = (last, curr) => last += curr

let totalPt1 = 0
runner.onLine((line) => {
    totalPt1 = Array.from(parseMul(line)).map(execMul).reduce(sum, totalPt1)
})

let totalPt2 = 0
let state = { enabled: true }
runner.onLine((line) => {
    const blocks = parseBlock(line, state)
    let result

    while ((result = blocks.next()) && !result.done) {
        const block = result.value
        for(const args of parseMul(block)) {
            totalPt2 += execMul(args)
        }
    }

    state = result.value
})

runner.onEnd(() => {
    console.log(`Total (part 1):\t${totalPt1}`)
    console.log(`Total (part 2):\t${totalPt2}`)
})