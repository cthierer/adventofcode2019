// node index.js < input.txt
import run from '../util/run.mjs'

const searchMatrix = (deltaRow, deltaCol) => (matrix, search, rowIdx, colIdx) => {
    if (search.length < 1) {
        return 0
    }

    if (rowIdx < 0 || rowIdx >= matrix.length) {
        return 0
    }

    const row = matrix[rowIdx]
    if (colIdx < 0 || colIdx >= row.length) {
        return 0
    }

    const actual = row[colIdx]
    const target = search[0]
    
    if (target === actual) {
        return searchMatrix(deltaRow, deltaCol)(matrix, search.slice(1), rowIdx + deltaRow, colIdx + deltaCol) + 1
    }

    return 0
}

const searches = [
    searchMatrix(-1, 0),
    searchMatrix(1, 0),
    searchMatrix(0, -1),
    searchMatrix(0, 1),
    searchMatrix(-1, -1),
    searchMatrix(-1, 1),
    searchMatrix(1, -1),
    searchMatrix(1, 1),
]

const countMatches = (matrix, search) => (rowIdx, colIdx) => searches.map(check => check(matrix, search, rowIdx, colIdx)).filter(length => length === search.length).length

const runner = run(process.stdin)

let wordsearch = []

runner.onLine((line) => {
    wordsearch = [...wordsearch, Array.from(line)]
})

runner.onEnd(() => {
    let totalMatches = 0
    const countFrom = countMatches(wordsearch, Array.from('XMAS'))
    for (let row = 0; row < wordsearch.length; row += 1) {
        for (let col = 0; col < wordsearch[row].length; col += 1) {
            totalMatches += countFrom(row, col)
        }
    }

    console.log(`Total matches:\t${totalMatches}`)
})