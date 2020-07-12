# 异步精讲

JavaScript 的异步编程发展经过了四个阶段：<br />

1. 回调函数、发布订阅
1. Promise
1. co 自执行的 Generator 函数
1. async / await



<a name="9VSjW"></a>
## 😢 Promise 源码

<br />核心代码，参考 [最简实现Promise，支持异步链式调用（20行）](https://juejin.im/post/5e6f4579f265da576429a907)<br />

```javascript
function Promise(fn) {
  this.cbs = [];

  const resolve = (value) => {
    setTimeout(() => {
      this.data = value;
      this.cbs.forEach((cb) => cb(value));
    });
  }

  fn(resolve.bind(this));
}

Promise.prototype.then = function (onResolved) {
  return new Promise((resolve) => { // 实现链式调用
    this.cbs.push(() => {
      const res = onResolved(this.data);
      if (res instanceof Promise) {
        res.then(resolve);
      } else {
        resolve(res);
      }
    });
  });
};
```

<br />添加 reject，抽离 resolvePromise，添加其他操作<br />

```javascript
const isFunction = obj => typeof obj === 'function'
const isObject = obj => !!(obj && typeof obj === 'object') // null 的情况
const isThenable = obj => (isFunction(obj) || isObject(obj)) && 'then' in obj && isFunction(obj.then)
const isPromise = promise => promise instanceof Promise

const PENDING = 'pending'
const FULFILLED = 'fulfilled'
const REJECTED = 'rejected'

class Promise {
  constructor(fn) {
    this.status = PENDING
    this.value = undefined
    this.reason = undefined
    this.onFulfilledCallbacks = []
    this.onRejectedCallbacks = []
    function resolve(value) {
      if (this.status !== PENDING)
        return
      setTimeout(() => {
        this.status = FULFILLED
        this.value = value
        this.onFulfilledCallbacks.forEach(cb => cb(this.value))
      }, 0)
    }
    function reject(reason) {
      if (this.status !== PENDING)
        return
      setTimeout(() => {
        this.status = REJECTED
        this.reason = reason
        this.onRejectedCallbacks.forEach(cb => cb(this.reason))
      })
    }
    try {
      fn(resolve, reject)
    } catch (e) {
      reject(e)
    }
  }

  then(onFulfilled, onRejected) {
    onFulfilled = typeof onFulfilled === 'function' ? onFulfilled : value => value
    onRejected = typeof onRejected === 'function' ? onRejected : reason => { throw reason }
    return bridgePromise = new Promise((resolve, reject) => {
      if (this.status === FULFILLED) {
        setTimeout(() => {
          try {
            let result = onFulfilled(this.value)
            resolvePromise(bridgePromise, result, resolve, reject)
          } catch (e) {
            reject(e)
          }
        }, 0)
      } else if (this.status === REJECTED) {
        setTimeout(() => {
          try {
            let result = onRejected(this.reason)
            resolvePromise(bridgePromise, result, resolve, reject)
          } catch (e) {
            reject(e)
          }
        }, 0)
      } else if (this.status === PENDING) {
        this.onFulfilledCallbacks.push(() => {
          try {
            let result = onFulfilled(this.value)
            resolvePromise(bridgePromise, result, resolve, reject)
          } catch (e) {
            reject(e)
          }
        })
        this.onRejectedCallbacks.push(() => {
          try {
            let result = onRejected(this.reason)
            resolvePromise(bridgePromise, result, resolve, reject)
          } catch (e) {
            reject(e)
          }
        })
      }
    })
  }

  catch(onRejected) {
    return this.then(null, onRejected)
  }

  static resolve(p) {
    if (isPromise(p)) return p // Promise.resolve(p) 与 new Promise(resolve => resolve(p)) 的区别
    return new Promise((resolve, reject) => {
      if (isThenable(p)) p.then(resolve, reject)
      else resolve(p)
    })
  }

  static reject(p) {
    return new Promise((_, reject) => reject(p))
  }

  static all(promises) {
    return new Promise((resolve, reject) => {
      let values = []
      let count = 0
      function handle(value, index) {
        values[index] = value
        if (++count === promises.length) resolve(values)
      }
      // p 可能不是 Promise，所以用 Promise.resolve 包一下
      promises.forEach((p, i) => Promise.resolve(p).then(value => handle(value, i), reject))
    })
  }

  static race(promises) {
    return new Promise((resolve, reject) => {
      promises.forEach(p => Promise.resolve(p).then(resolve, reject))
    })
  }

  static allSettled(promises) {
    return new Promise((resolve) => {
      let results = []
      let count = 0
      function handle(result, index) {
        results[index] = result
        if (++count === promises.length) resolve(results)
      }
      promises.forEach((p, i) => Promise.resolve(p).then(
        value => handle({ status: 'fulfilled', value }, i),
        reason => handle({ status: 'rejected', reason }, i),
      ))
    })
  }
}

// 抽离的 resolve 方法
function resolvePromise(bridgePromise, result, resolve, reject) {
  if (bridgePromise === result) { // 循环
    return reject(new TypeError('Chaining cycle detected for promise #<Promise>'))
  }
  if (isPromise(result)) {
    if (result.status === PENDING) {
      result.then(y => resolvePromise(bridgePromise, y, resolve, reject), reject)
    } else {
      result.then(resolve, reject)
    }
  } else if (isThenable(result)) {
    result.then(y => resolvePromise(bridgePromise, y, resolve, reject), reject)
  } else {
    resolve(result)
  }
}
```


<a name="91489af7"></a>
## 📝 co 源码

<br />co 接收一个 generator 函数，返回一个 promise，generator 函数中 yieldable 对象有：<br />

- promises
- thunks (functions)
- array (parallel execution)
- objects (parallel execution)
- generators (delegation)
- generator functions (delegation)


<br />其中 array 和 objects 是并行执行的，里面的值仍然是 promise 和 thunk 函数，而 generators 和 generator functions 是通过代理执行，内部再次调用 co，所以简单来说都是基于 promise 和 thunk 函数的，而 co 内部对于 thunk 的处理是把 thunk 也转化成 promise，所以直接看对于 yield 一个 promise 的 generator 怎么自动执行
```javascript
function* gen() {
  const foo = yield Promise.resolve(1)
  const bar = yield Promise.resolve(2)
  console.log(foo)
  console.log(bar)
}
```

<br />这里我们写一个 GeneratorFunction，每次都 yield 出一个 promise，我们如何让这段代码以类似同步的执行方式从上到下执行<br />

```javascript
const gen = gen()
g.next()
```

<br />先得到一个 generator，然后 next，此时 generatorFunction 执行到 yield 处<br />

```javascript
const gen = gen()
g.next().value.then((data) => {
  // next
})
```

<br />返回的结果的 value 是 yield 出来的 promise 容器包裹的数值 1，那么 then 方法的 callback 的参数就是 1<br />

```javascript
const gen = gen()
g.next().value.then((data) => {
  g.next(data)
})
```

<br />为了让 yield 左边的变量 foo 得到异步代码的结果，我们只需要把 data 通过 generator 的 next 方法传入就可以了，同时 generatorFunction 的控制权也回到 generatorFunction 手中，generatorFunction 继续执行<br />

```javascript
const gen = gen()
g.next().value.then((data) => {
  g.next(data).value.then((data) => {
    g.next(data)
  })
})
```

<br />之后再次 yield 出 promise 的异步操作，交出控制权，同样的通过 next 返回结果和控制权让 generatorFunction 继续执行，这样就实现了包含异步操作的 generatorFunction 的同步执行<br />在直接套用 co 源码：<br />

```javascript
function co(gen) {
  // ...
  return new Promise((resolve, reject) => {
    const g = gen()

    const gResult = g.next() // 第一次 next
    if (gResult.done) resolve(gResult.value)
    if (gResult.value && isPromise(gResult.value)) {
      value.then((res) => {

        const gResult = g.next(res) // 第二次 next
        if (gResult.done) resolve(gResult.value)
        if (gResult.value && isPromise(gResult.value)) {
          value.then((res) => {

            const gResult = g.next(res) // 第三次 next，done 为 true
            if (gResult.done) resolve(gResult.value) // resolve 掉 generator 中 return 的结果
          })
        }
      })
    }
  })
}
```

<br />在看 co 整体代码：<br />

```javascript
function co(gen) {
  var ctx = this; // 那 this，一般是 co.call 这样调用
  var args = slice.call(arguments, 1) // generator 的参数可以在 gen 后面传入

  return new Promise(function(resolve, reject) {
    // 检查 gen
    if (typeof gen === 'function') gen = gen.apply(ctx, args); // 普通函数就会调用得到返回值，下一行 resolve 返回值
    if (!gen || typeof gen.next !== 'function') return resolve(gen);

    onFulfilled();

    function onFulfilled(res) {
      var ret;
      try {
        ret = gen.next(res);
      } catch (e) { // try / catch 做错误捕获
        return reject(e); // 出错就 return reject 掉，return 是为了防止 reject 后仍然执行 next 函数
      }
      next(ret);
    }

    function onRejected(err) {
      var ret;
      try {
        ret = gen.throw(err);
      } catch (e) {
        return reject(e);
      }
      next(ret);
    }

    function next(ret) {
      if (ret.done) return resolve(ret.value); // new Promise 的 resolve 用来 resolve 最终 done 为 true 时的 value
      var value = toPromise.call(ctx, ret.value); // 把其他的 yieldable 转化成 promise
      if (value && isPromise(value)) return value.then(onFulfilled, onRejected);
      return onRejected(new TypeError('You may only yield a function, promise, generator, array, or object, '
        + 'but the following object was passed: "' + String(ret.value) + '"'));
    }
  });
}
```

<br />其中 toPromise 针对不同的 yieldable 进行 xxxToPromise，arrayToPromise 是通过 Promise.all(value.map(toPromise)) 进行转换，objectToPromise 等待对象的所有的值都 resolve 后，并添加到新的对象中，然后再 resolve，类似于 Promise.all<br />
<br />thunkToPromise 类似于一般 Node.js 的 API 的 promisify，只不过是 thunk 函数已经传入了第一个参数，promisify 时只需要传入另一个参数就可以了，我们也可以看出这里 thunk 是针对 Node.js 的 API 的，与 curry 的不同在于 thunk 是分为两次参数传入的<br />

```javascript
function thunkToPromise(fn) {
  var ctx = this;
  return new Promise(function (resolve, reject) {
    fn.call(ctx, function (err, res) {
      if (err) return reject(err);
      if (arguments.length > 2) res = slice.call(arguments, 1);
      resolve(res);
    });
  });
}
```

<br />isPromise 的判断也是通过查看参数的 then 是不是一个函数，体现了鸭子类型的特点<br />

```javascript
function isPromise(obj) {
  return 'function' == typeof obj.then;
}
```


<a name="a051c3c9"></a>
## ⚙️ 原理

<br />co 的原理其实是通过 generator.next() 得到 generatorResult，由于 yield 出是一个 promise，通过 generatorResult.value.then 再把 promise 的结果通过 generator.next 的参数传给 yield 的左边，让 generator 自动执行，通过 generatorResult.done 判断是否执行结束<br />

<a name="ae289456"></a>
## 🍬 async / await

<br />async / await 是语法糖，我们还原一个 async 函数，使用 TypeScript 跟更体现一些类型本质的东西
```typescript
type ExtractType<T> =
  T extends {
    [Symbol.iterator](): {
      next(): { done: true, value: infer U }
    }
  } ? U :
  T extends {
    [Symbol.iterator](): {
      next(): { done: false }
    }
  } ? never :
  T extends {
    [Symbol.iterator](): {
      next(): { value: infer U }
    }
  } ? U :
  T extends {
    [Symbol.iterator](): any
  } ? unknown :
  never

type Async =
  <F extends (...args: any[]) => Generator<unknown>>(fn: F)
    => (...args: Parameters<F>) => Promise<ExtractType<ReturnType<F>>>
```

<br />先对类型进行编写，async function 返回一个 Promise，Promise 包裹内部 return 的值，由于我们模拟 Async 函数要传入一个 GeneratorFunction，返回的一个函数才相当于 async function，所以通过 ExtractType 拿到 `Generator<unknown>` 最终 done 为 true 时的 value 的类型<br />

```typescript
const getData = async(function * (url: string) {
  const result: Response = yield fetch(url)
  const json = yield result.json()
  return json
})
```

<br />我们实现的用法就想这样，除了 `() * yield` 写法不一致其他与 async function 用法一样，与 co 不同的是 yield 后可以跟任何值，不止是 Promise<br />

```typescript
const async: Async = (genFn) =>
  (...args) => new Promise((resolve, reject) => {
    const gen = genFn(...args)

    function next(nextFn: () => IteratorResult<unknown>) {
      let result = nextF()
      // resolve value and deliver it to gen.next
    }

    next(() => gen.next(undefined))
  })
```

<br />`(...args) => new Promise(...)` 相当于我们实际调用的 async function，通过 thunk 和展开运算符把 genFn 的参数拿到，并在传入 genFn，得到 gen<br />
<br />同时我们要在 next 内部执行 gen.next，通过包裹一个函数 nextFn 传入，在内部得到 result<br />

```typescript
const async: Async = (genFn) =>
  (...args) => new Promise((resolve, reject) => {
    const gen = genFn(...args)

    function next(nextFn: () => IteratorResult<unknown>) {
      let result = nextFn()

      if(result.done) return resolve(result.value)
      Promise.resolve(result.value).then((res) => {
        next(() => gen.next(res))
      })
    }
    
    next(() => gen.next(undefined))
  })
```

<br />通过判断是否 done 进行 new Promise 的 resolve，如果没有完成就继续通过 next 进行传递，注意不同于 co 我们内部用 Promise.resolve 处理 result.value，所以我们 yield 时也可以不是一个 Promise<br />

> 之前的标准是使用 `new Promise(res => res(resule.value))` 进行包裹处理，v8 提出 [Faster async functions and promises](https://v8.dev/blog/fast-async#await-under-the-hood) 并 PR，现在已经修改为 `Promise.resolve`
> 对于这两个的区别在于 resolve 一个 Promise 时的表现不同，`Promise.resolve(p)` 对于 Promise 会直接返回这个 Promise，而 `new Promise(res => res(p))` 在内部调用 `p.then(resolve, reject)` 相当于多出一个微任务来处理 `res(p)`，所以目前新版的更快，有些代码执行顺序也会不同



```typescript
const async: Async = (genFn) =>
  (...args) => new Promise((resolve, reject) => {
    const gen = genFn(...args)

    function next(nextFn: () => IteratorResult<unknown>) {
      let result: ReturnType<typeof nextF>
      try {
        result = nextFn()
      } catch(e) {
        return reject(e)
      }

      if(result.done) return resolve(result.value)
      Promise.resolve(result.value).then((res) => {
        next(() => gen.next(res))
      }, (err) => {
        next(() => gen.throw(err))
      })
    }
    
    next(() => gen.next(undefined))
  })
```

<br />现在我们加上错误处理，当 resolve value 出错时会通过 gen.throw(err) 抛出错误，而 gen.throw 通过 genF 内部的 try / catch 捕捉（所以 async / await 的错误处理一般也是在函数内写 try / catch）然后通过上面的 try / catch 将错误 reject 出来，不同于成功时 async 函数返回一个包裹 value 的 Promise，而是返回出一个包裹 error 的 Promise<br />现在我们完成和 async / await 的函数的模拟，我们看到 async function 实际上返回一个 Promise 包裹的 return 值，await 会自动使用 Promise.resolve 进行包裹，并类似 yield 把 flat 后得到的结果代替那个表达式<br />
<br />这个函数与 co 的不同除了使用 Promise.resolve 自动包裹，不能处理 yield 数组和对象时实现的并行以外，还有将 gen.next 和 gen.throw 抽象成 nextFn，这也导致直观上代码行数不同，但本质实际上没有什么区别<br />

<a name="97a7b02c"></a>
## 🤔 对于 JavaScript 异步的思考


<a name="3c419b3b"></a>
### 3.1 raw callback

<br />我们看最开始最朴素的 raw callback，是将 callback 交给另一个函数执行，也就是说我们把 callback 的控制权交给这个函数，这个函数在进行完异步操作之后调用 callback，以此实现异步<br />

<a name="9aafd3ef"></a>
### 3.2 Promise callback

<br />而之后 promise 也是通过传入 callback 的方式，只不过把之前嵌套式的形式展开成链式，其实通过链表为函数增加 next 属性，也可以使嵌套式展开成链式<br />
<br />promise 通过完成异步操作后进行 resolve 或 reject，来控制 callback 的执行，而且提供了 then 返回一个 promise 的自动进行 flat（flatMap），实现了 then 中继续执行异步的操作，所以提供 callback 参数对于 promise 来说也是一种控制权的转移，只不过是从以前直接的函数调用改成了 resolve、reject 控制 callback 的调用时机<br />
<br />同时是一种标准的实现也相较于原来的 raw callback 保证了内部的可控性与安全性<br />

<a name="5c4753fc"></a>
### 3.3 co + Generator

<br />GeneratorFunction 得到的 Generator 可以通过 next 打断 GeneratorFunction 的执行，由于只能通过 Generator 调用 next 把 GeneratorFunction 的执行权还给 GeneratorFunction，所以称作“半协程”<br />
<br />通过保存 GeneratorFunction 的执行上下文，使 GeneratorFunction 可中断执行，从而把 GeneratorFunction 控制权交给 Generator，Generator 拿到控制权后通过 yield 出来的 promise 完成异步操作，等 resolve 之后再通过 then 中调用 next 把异步的结果和 GeneratorFunction 的控制权交给 GeneratorFunction，以继续执行 yield 后的操作<br />

<a name="18627735"></a>
### 3.4 async / await

<br />async 函数是对 GeneratorFunction + co 的语义化和标准化的语法糖<br />
<br />便捷性提升的同时也意味着灵活性的减少，由于 async / await 是语法，而 promise、callback 是对象，对象可以到处传递，React 也通过 throw 一个 promise 如此 creative and hacking 的模拟了 [Algebraic Effects](https://overreacted.io/algebraic-effects-for-the-rest-of-us/) 实现 Suspense<br />
<br />同时 Promise 和 GeneratorFunction 也相对于 raw callback 约束，Promise 是 onFulfilled、onRejected 的约束，GeneratorFunction 是 next、done 的约束，Node.js APIs 中也限制了 cb 的参数，所以也能被统一的 thunk 化，这种约束类似于语法糖，规范的同时也丧失了些许灵活性<br />
<br />Promise 作为 async function 中的异步最小单位通过 await 进行传递，而 Promise 又是由 callback 组成，所以 co + Generator（async / await）也是一种 callback 的形式，只不过写法更加方便规范<br />

<a name="2d19699c"></a>
### 3.5 RxJS

<br />// TODO: RxJS 与 async 区别，RxJS 理念等<br />

<a name="3e6b60ca"></a>
### 3.6 🔑 the key

<br />异步的关键就在于调用 callback 的时机，因为我们不知道异步操作需要多少时间，我们自然也就不知道何时调用异步之后的操作，所以我们通过 callback 将之后操作的控制权交给异步操作，实现控制反转，在异步操作完成之后自动调用 callback，就完成了在合适的时机进行合适的操作<br />

<a name="ref"></a>
## ref

<br />[Generator 函数的异步应用](https://es6.ruanyifeng.com/#docs/generator-async)<br />
<br />[100 行代码实现 Promises/A+ 规范](https://zhuanlan.zhihu.com/p/83965949)<br />
<br />[JAVASCRIPT GETTER-SETTER PYRAMID](https://staltz.com/javascript-getter-setter-pyramid.html)
