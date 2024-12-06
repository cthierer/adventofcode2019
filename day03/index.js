// node index.js < input.txt
import run from '../util/run.mjs'

const runner = run(process.stdin)

const parseMul = function* parseMul(string) {
    const pattern = /mul\((\d{1,3}),(\d{1,3})\)/g
    let match

    while((match = pattern.exec(string)) !== null) {
        const [, x, y] = match
        yield [x, y]
    }
}

const execMul = ([x, y]) => x * y

let total = 0

runner.onLine((line) => {
    total += Array.from(parseMul(line)).map(execMul).reduce((last, curr) => last += curr, 0)
})

runner.onEnd(() => {
    console.log(`Total:\t${total}`)
})