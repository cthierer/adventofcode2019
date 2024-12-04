import EventEmitter from 'node:events'

const Line = 'line'
const End = 'end'

const run = (input) => {
    const runner = new EventEmitter()

    setImmediate(() => {
        input.on('readable', () => {
            let chunk
            while ((chunk = input.read()) && chunk !== null) {
                const lines = chunk.toString('utf8').trim().split('\n')
                lines.forEach((element) => {
                    const line = element.trim() 
                    if (line.length < 1) {
                        return
                    }

                    runner.emit(Line, line)
                })
            }
        })

        input.on('error', (err) => {
            console.error(`An error occurred: ${err.message}`)
        })

        input.on('end', () => {
            runner.emit(End)
            console.log('Done!')
        })
    })

    return {
        onLine: (handler) => {
            runner.on(Line, handler)
            return () => {
                runner.removeListener(Line, handler)
            }
        },
        onEnd: (handler) => {
            runner.on(End, handler)
            return () => {
                runner.removeListener(End, handler)
            }
        },
    }
}

export default run