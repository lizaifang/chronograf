import Deferred from 'src/worker/Deferred'

export function restartable<T extends any[], V>(
  f: (...args: T) => Promise<V>
): ((...args: T) => Promise<V>) {
  let deferred: Deferred
  let id: number = 0

  async function checkResult(promise: Promise<V>, promiseID: number) {
    let result
    let isOk = true

    try {
      result = await promise
    } catch (error) {
      result = error
      isOk = false
    }

    if (promiseID !== id) {
      return
    }

    if (isOk) {
      deferred.resolve(result)
    } else {
      deferred.reject(result)
    }

    deferred = null
  }

  return function(...args: T): Promise<V> {
    if (!deferred) {
      deferred = new Deferred()
    }

    const promise = f(...args)

    id += 1
    checkResult(promise, id)

    return deferred.promise
  }
}
